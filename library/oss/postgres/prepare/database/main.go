// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"namespacelabs.dev/foundation/framework/resources/provider"
	postgresclass "namespacelabs.dev/foundation/library/database/postgres"
	"namespacelabs.dev/foundation/library/oss/postgres"
	universepg "namespacelabs.dev/foundation/universe/db/postgres"
)

const (
	providerPkg = "namespacelabs.dev/foundation/library/oss/postgres"
	connBackoff = 1500 * time.Millisecond

	caCertPath = "/tmp/ca.pem"
)

func main() {
	ctx, p := provider.MustPrepare[*postgres.DatabaseIntent]()

	cluster := &postgresclass.ClusterInstance{}
	if err := p.Resources.Unmarshal(fmt.Sprintf("%s:cluster", providerPkg), cluster); err != nil {
		log.Fatalf("unable to read required resource \"cluster\": %v", err)
	}

	// TODO inject file as secret ref and propagate secret ref to server, too.
	if cluster.CaCert != "" {
		if err := os.WriteFile(caCertPath, []byte(cluster.CaCert), 0644); err != nil {
			log.Fatalf("failed to write %q: %v", caCertPath, err)
		}

		if err := os.Setenv("PGSSLROOTCERT", caCertPath); err != nil {
			log.Fatalf("failed to set PGSSLROOTCERT: %v", err)
		}

	}

	exists, err := ensureDatabase(ctx, cluster, p.Intent.Name)
	if err != nil {
		log.Fatalf("unable to create database %q: %v", p.Intent.Name, err)
	}

	instance := &postgresclass.DatabaseInstance{
		ConnectionUri:  postgres.ConnectionUri(cluster, p.Intent.Name),
		Name:           p.Intent.Name,
		User:           postgres.UserOrDefault(cluster.User),
		Password:       cluster.Password,
		ClusterAddress: cluster.Address,
		ClusterHost:    cluster.Host,
		ClusterPort:    cluster.Port,
		SslMode:        cluster.SslMode,
	}

	if !exists || !p.Intent.SkipSchemaInitializationIfExists {
		db, err := universepg.NewDatabaseFromConnectionUri(ctx, instance, instance.ConnectionUri, nil)
		if err != nil {
			log.Fatalf("unable to open connection: %v", err)
		}
		defer db.Close()

		for _, schema := range p.Intent.Schema {
			if _, err := universepg.ReturnFromReadWriteTx(ctx, db, backOff{
				interval: 100 * time.Millisecond,
				deadline: time.Now().Add(5 * time.Second),
				jitter:   100 * time.Millisecond,
			}, func(ctx context.Context, tx pgx.Tx) (pgconn.CommandTag, error) {
				return tx.Exec(ctx, string(schema.Contents))
			}); err != nil {
				log.Fatalf("unable to apply schema %q: %v", schema.Path, err)
			}
		}
	}

	p.EmitResult(instance)
}

func ensureDatabase(ctx context.Context, cluster *postgresclass.ClusterInstance, name string) (bool, error) {
	// Postgres needs a database to connect to so we pin one that is guaranteed to exist.
	postgresConn, err := connect(ctx, cluster, "postgres")
	if err != nil {
		return false, err
	}
	defer postgresConn.Close(ctx)

	exists, err := existsDatabase(ctx, postgresConn, name)
	if err != nil {
		return false, err
	}

	if !exists {
		// SQL arguments can only be values, not identifiers.
		// https://www.postgresql.org/docs/9.5/xfunc-sql.html
		// `existsDb` already uses the database name as an SQL argument, so we already passed its validation.
		// Still, let's do some basic sanity checking (whitespaces are forbidden), as we need to use Sprintf here.
		// Valid database names are defined at https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS
		if len(strings.Fields(name)) > 1 || strings.Contains(name, "-") {
			return false, fmt.Errorf("invalid database name: %s", name)
		}

		if _, err := postgresConn.Exec(ctx, fmt.Sprintf("CREATE DATABASE \"%s\";", name)); err != nil {
			return false, fmt.Errorf("failed to create database %q: %w", name, err)
		}
	}

	return exists, err
}

func existsDatabase(ctx context.Context, conn *pgx.Conn, name string) (bool, error) {
	rows, err := conn.Query(ctx, "SELECT FROM pg_database WHERE datname = $1;", name)
	if err != nil {
		return false, fmt.Errorf("failed to check for database %q: %w", name, err)
	}
	defer rows.Close()

	return rows.Next(), nil
}

func connect(ctx context.Context, cluster *postgresclass.ClusterInstance, db string) (*pgx.Conn, error) {
	cfg, err := pgx.ParseConfig(postgres.ConnectionUri(cluster, db))
	if err != nil {
		return nil, err
	}
	cfg.ConnectTimeout = connBackoff

	// Retry until backend is ready.
	return backoff.RetryWithData(func() (*pgx.Conn, error) {
		conn, err := pgx.ConnectConfig(ctx, cfg)
		if err == nil {
			return conn, nil
		}

		log.Printf("failed to connect to postgres: %v\n", err)
		return nil, err
	}, backoff.WithContext(backoff.NewConstantBackOff(connBackoff), ctx))
}

type backOff struct {
	interval time.Duration
	deadline time.Time
	jitter   time.Duration
}

func (b backOff) Reset() {}
func (b backOff) NextBackOff() time.Duration {
	if time.Now().After(b.deadline) {
		return backoff.Stop
	}
	return b.interval - b.jitter/2 + time.Duration(rand.Int63n(int64(b.jitter)))
}
