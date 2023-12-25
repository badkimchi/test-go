//go:build wireinject
// +build wireinject

// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"app/domains/auth"
	"github.com/google/wire"
)

type apiHandlers struct {
	AuthC *auth.AuthController
}

func controllers() (
	apiHandlers, error,
) {
	wire.Build(
		auth.NewAuthController,
		newApiHandlers,
	)
	return apiHandlers{}, nil
}

func newApiHandlers(
	authC *auth.AuthController,
) apiHandlers {
	return apiHandlers{
		AuthC: authC,
	}
}
