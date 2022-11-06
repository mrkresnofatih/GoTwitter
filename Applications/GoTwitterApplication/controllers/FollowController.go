package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/gotwitter/endpoints/follow"
	"github.com/mrkresnofatih/gotwitter/services"
)

type FollowController struct {
	FollowService services.IFollowService
}

func (f *FollowController) Register(echo *echo.Echo) {
	controllerRouter := ControllerRouter{
		MainRouter: echo,
		PathPrefix: "/api/v2/follow",
	}

	followEndpoint := &follow.FollowEndpoint{
		FollowService: f.FollowService,
	}
	followEndpointWithAuth := &RequireAuthorizationDecorator{
		Endpoint: followEndpoint,
	}
	controllerRouter.AddEndpoint(followEndpointWithAuth)

	unfollowEndpoint := &follow.UnfollowEndpoint{
		FollowService: f.FollowService,
	}
	unfollowEdWithAuth := &RequireAuthorizationDecorator{
		Endpoint: unfollowEndpoint,
	}
	controllerRouter.AddEndpoint(unfollowEdWithAuth)

	controllerRouter.Build()
}
