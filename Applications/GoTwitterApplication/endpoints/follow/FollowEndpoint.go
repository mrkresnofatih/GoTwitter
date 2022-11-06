package follow

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mrkresnofatih/gotwitter/models"
	"github.com/mrkresnofatih/gotwitter/services"
	"github.com/mrkresnofatih/gotwitter/tools/jwt"
	"net/http"
)

type FollowEndpoint struct {
	FollowService services.IFollowService
}

func (f *FollowEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		if username == "" {
			log.Error("failed to retrieve username from url param")
			return models.SendBadResponse(c, "failed to retrieve username")
		}

		authHeader := c.Request().Header.Get("Authorization")
		accessToken := authHeader[7:]
		usernameFromToken, _ := jwt.GetClaimFromToken[string](accessToken, jwt.ApplicationJwtClaimsKeyUsername)

		followReq := models.FollowCreateRequestModel{
			Username: username,
			Follower: usernameFromToken,
		}
		followResp, err := f.FollowService.Follow(c.Request().Context(), followReq)
		if err != nil {
			log.Error("failed to follow user")
			return models.SendBadResponse(c, "failed to follow user")
		}

		return models.SendGoodResponse[models.FollowCreateResponseModel](c, followResp)
	}
}

func (f *FollowEndpoint) GetMethod() string {
	return http.MethodGet
}

func (f *FollowEndpoint) GetPath() string {
	return "/follow/:username"
}

func (f *FollowEndpoint) Register(group *echo.Group) {
	group.Match([]string{f.GetMethod()}, f.GetPath(), f.GetHandler())
}
