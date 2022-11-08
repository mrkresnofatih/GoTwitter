package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/gotwitter/endpoints/follow"
	"github.com/mrkresnofatih/gotwitter/models"
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

	getFollowersEndpoint := &follow.GetFollowersEndpoint{
		FollowService: f.FollowService,
	}
	getFollowersWithValidation := &RequireValidationDecorator[models.FollowGetFollowersRequestModel]{
		Endpoint: getFollowersEndpoint,
	}
	getFollowersWithAuth := &RequireAuthorizationDecorator{
		Endpoint: getFollowersWithValidation,
	}
	controllerRouter.AddEndpoint(getFollowersWithAuth)

	getFollowingsEndpoint := &follow.GetFollowingsEndpoint{
		FollowService: f.FollowService,
	}
	getFollowingsWithValidation := &RequireValidationDecorator[models.FollowGetFollowingsRequestModel]{
		Endpoint: getFollowingsEndpoint,
	}
	getFollowingsWithAuth := &RequireAuthorizationDecorator{
		Endpoint: getFollowingsWithValidation,
	}
	controllerRouter.AddEndpoint(getFollowingsWithAuth)

	controllerRouter.Build()
}
