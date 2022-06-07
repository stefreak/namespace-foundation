// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package aws

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go/aws/session"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace/devhost"
)

const identityTokenEnv = "AWS_WEB_IDENTITY_TOKEN_FILE"

func useInclusterConfig() bool {
	// Check if we run inside an AWS cluster with a configured IAM role.
	token := os.Getenv(identityTokenEnv)
	return token != ""
}

func ConfiguredSessionV1(ctx context.Context, devHost *schema.DevHost, selector devhost.Selector) (*session.Session, string, error) {
	conf := &Conf{}

	if !selector.Select(devHost).Get(conf) {
		if useInclusterConfig() {
			sess, err := session.NewSession()
			// TODO remove profile?
			return sess, "default", err
		}

		return nil, "", fnerrors.UsageError("Run `fn prepare`.", "Foundation has not been configured to access AWS.")
	}

	profile := conf.Profile
	if profile == "" {
		profile = "default"
	}

	sess, err := session.NewSessionWithOptions(session.Options{Profile: profile})
	return sess, profile, err
}

func ConfiguredSession(ctx context.Context, devHost *schema.DevHost, selector devhost.Selector) (aws.Config, string, error) {
	conf := &Conf{}
	if !selector.Select(devHost).Get(conf) {
		if useInclusterConfig() {
			cfg, err := config.LoadDefaultConfig(ctx)
			// TODO remove profile?
			return cfg, "default", err
		}

		return aws.Config{}, "", fnerrors.UsageError("Run `fn prepare`.", "Foundation has not been configured to access AWS.")
	}

	profile := conf.Profile
	if profile == "" {
		profile = "default"
	}

	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(profile))
	return cfg, profile, err
}
