package controllers

import (
	"github.com/google/wire"
	"mono-base/cmd/http/controllers/app_version"
	"mono-base/cmd/http/controllers/user"
)

var ControllerProviders = wire.NewSet(
	user.NewUserControllerV1,
	user.NewUserControllerV2,
	app_version.NewStoryControllerV1,
)
