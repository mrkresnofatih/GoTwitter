package follow

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mrkresnofatih/gotwitter/models"
	"github.com/mrkresnofatih/gotwitter/services"
	"github.com/mrkresnofatih/gotwitter/tools/jwt"
	"net/http"
)

type UnfollowEndpoint struct {
	FollowService services.IFollowService
}

func (u UnfollowEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		if username == "" {
			log.Error("failed to retrieve username from url param")
			return models.SendBadResponse(c, "failed to retrieve username")
		}

		authHeader := c.Request().Header.Get("Authorization")
		accessToken := authHeader[7:]
		usernameFromToken, _ := jwt.GetClaimFromToken[string](accessToken, jwt.ApplicationJwtClaimsKeyUsername)

		unfollowReq := models.FollowDeleteRequestModel{
			Username: username,
			Follower: usernameFromToken,
		}
		unfollowResp, err := u.FollowService.Unfollow(c.Request().Context(), unfollowReq)
		if err != nil {
			log.Error("failed to unfollow user")
			return models.SendBadResponse(c, "failed to unfollow user")
		}

		return models.SendGoodResponse[models.FollowDeleteResponseModel](c, unfollowResp)
	}
}

func (u UnfollowEndpoint) GetMethod() string {
	return http.MethodGet
}

func (u UnfollowEndpoint) GetPath() string {
	return "/unfollow/:username"
}

func (u UnfollowEndpoint) Register(group *echo.Group) {
	group.Match([]string{u.GetMethod()}, u.GetPath(), u.GetHandler())
}
