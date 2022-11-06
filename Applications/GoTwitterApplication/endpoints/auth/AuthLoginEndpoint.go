package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mrkresnofatih/gotwitter/models"
	"github.com/mrkresnofatih/gotwitter/services"
	"net/http"
)

type AuthLoginEndpoint struct {
	AuthService services.IAuthService
}

func (a *AuthLoginEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		grantId := c.QueryParam("grantid")
		if grantId == "" {
			log.Error("failed to retrieve grantId")
			return models.SendBadResponse(c, "failed to retrieve grantId")
		}

		accessTokenObj, err := a.AuthService.Login(c.Request().Context(), models.AuthLoginRequestModel{
			GrantId: grantId,
		})
		if err != nil {
			log.Error("failed to authenticate grant")
			return models.SendBadResponse(c, "failed to authenticate grant")
		}

		log.Info("auth grant successfully logged in")
		return models.SendGoodResponse[models.AuthLoginResponseModel](c, accessTokenObj)
	}
}

func (a *AuthLoginEndpoint) GetMethod() string {
	return http.MethodGet
}

func (a *AuthLoginEndpoint) GetPath() string {
	return "/login"
}

func (a *AuthLoginEndpoint) Register(group *echo.Group) {
	group.Match([]string{a.GetMethod()}, a.GetPath(), a.GetHandler())
}
