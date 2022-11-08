package follow

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mrkresnofatih/gotwitter/models"
	"github.com/mrkresnofatih/gotwitter/services"
	"net/http"
)

type GetFollowingsEndpoint struct {
	FollowService services.IFollowService
}

func (g *GetFollowingsEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		listReq := new(models.FollowGetFollowingsRequestModel)
		_ = c.Bind(listReq)

		getFollowingsResp, err := g.FollowService.GetFollowings(c.Request().Context(), *listReq)
		if err != nil {
			log.Error("failed to get followings list")
			return models.SendBadResponse(c, "failed to get followings list")
		}

		return models.SendGoodResponse[models.FollowGetFollowingsResponseModel](c, getFollowingsResp)
	}
}

func (g *GetFollowingsEndpoint) GetMethod() string {
	return http.MethodPost
}

func (g *GetFollowingsEndpoint) GetPath() string {
	return "/followings"
}

func (g *GetFollowingsEndpoint) Register(group *echo.Group) {
	group.Match([]string{g.GetMethod()}, g.GetPath(), g.GetHandler())
}
