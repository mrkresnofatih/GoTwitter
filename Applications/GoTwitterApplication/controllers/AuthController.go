package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/gotwitter/endpoints"
	"github.com/mrkresnofatih/gotwitter/services"
)

type AuthController struct {
	AuthService services.IAuthService
}

func (a AuthController) Register(echo *echo.Echo) {
	controllerRouter := ControllerRouter{
		MainRouter: echo,
		PathPrefix: "/api/v2/auth",
	}

	authLoginEndpoint := &endpoints.AuthLoginEndpoint{
		AuthService: a.AuthService,
	}
	controllerRouter.AddEndpoint(authLoginEndpoint)

	controllerRouter.Build()
}
