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
	"app/domains/user"
	"app/sql/db"
	"github.com/go-chi/jwtauth"
	"github.com/google/wire"
)

type reqControllers struct {
	AuthC *auth.Controller
	AccC  *user.UserController
}

func controllers(jwtAuth *jwtauth.JWTAuth, queries *db.Queries) (
	reqControllers, error,
) {
	wire.Build(
		auth.NewAuthController,
		auth.NewAuthService,
		user.NewUserController,
		user.NewUserService,
		user.NewUserRepo,
		newReqControllers,
	)
	return reqControllers{}, nil
}

func newReqControllers(
	authC *auth.Controller,
	accC *user.UserController,
) reqControllers {
	return reqControllers{
		AuthC: authC,
	}
}
