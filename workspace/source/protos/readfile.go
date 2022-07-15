// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package protos

import (
	"io/ioutil"
	"reflect"

	"google.golang.org/protobuf/proto"
	"namespacelabs.dev/foundation/internal/fnerrors"
)

func ReadFileAndBytes[V proto.Message](path string) (V, []byte, error) {
	var empty V

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return empty, nil, fnerrors.New("%s: failed to load: %w", path, err)
	}

	s := reflect.New(reflect.TypeOf(empty).Elem()).Interface().(V)

	if err := proto.Unmarshal(bytes, s); err != nil {
		return empty, nil, fnerrors.New("%s: unmarshal failed: %w", path, err)
	}

	return s, bytes, nil
}

func ReadFile[V proto.Message](path string) (V, error) {
	v, _, err := ReadFileAndBytes[V](path)
	return v, err
}
