package follow

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mrkresnofatih/gotwitter/models"
	"github.com/mrkresnofatih/gotwitter/services"
	"net/http"
)

type GetFollowersEndpoint struct {
	FollowService services.IFollowService
}

func (g *GetFollowersEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		listReq := new(models.FollowGetFollowersRequestModel)
		_ = c.Bind(listReq)

		getFollowersResp, err := g.FollowService.GetFollowers(c.Request().Context(), *listReq)
		if err != nil {
			log.Error("failed to get followers list")
			return models.SendBadResponse(c, "failed to get followers list")
		}

		return models.SendGoodResponse[models.FollowGetFollowersResponseModel](c, getFollowersResp)
	}
}

func (g *GetFollowersEndpoint) GetMethod() string {
	return http.MethodPost
}

func (g *GetFollowersEndpoint) GetPath() string {
	return "/followers"
}

func (g *GetFollowersEndpoint) Register(group *echo.Group) {
	group.Match([]string{g.GetMethod()}, g.GetPath(), g.GetHandler())
}
