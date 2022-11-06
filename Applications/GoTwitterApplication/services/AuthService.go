package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/mrkresnofatih/gotwitter/models"
	"github.com/mrkresnofatih/gotwitter/tools/jwt"
	"time"
)

type IAuthService interface {
	Login(ctx context.Context, model models.AuthLoginRequestModel) (models.AuthLoginResponseModel, error)
}

type AuthService struct {
	PlayerService IPlayerService
	OauthService  IOauthService
}

func (a *AuthService) Login(ctx context.Context, model models.AuthLoginRequestModel) (models.AuthLoginResponseModel, error) {
	oauthResponse, err := a.OauthService.AuthenticateGrant(models.OauthAuthenticateGrantRequestModel{
		GrantId:           model.GrantId,
		ApplicationId:     "be244f27-b46b-4ed9-8fce-ca766b225c33",
		ApplicationSecret: "5c7a511f-6dbe-44c1-bb5d-438dd638d23b",
	})
	if err != nil {
		log.Error("failed to authenticate grant id")
		return *new(models.AuthLoginResponseModel), errors.New("failed to authenticate grant id")
	}

	token := oauthResponse.GrantToken
	oauthOpenIdResponse, err := a.OauthService.GetOpenId(models.OauthOpenIdRequestModel{
		Token: token,
	})
	if err != nil {
		log.Error("failed to get open id")
		return *new(models.AuthLoginResponseModel), errors.New("failed to get open id")
	}

	foundPlayer, err := a.PlayerService.GetPlayer(ctx, models.PlayerGetRequestModel{
		Username: oauthOpenIdResponse.Username,
	})
	if err != nil {
		log.Error("failed to find player with username provided")
		createPlayerReq := models.PlayerCreateRequestModel{
			Username:  oauthOpenIdResponse.Username,
			FullName:  oauthOpenIdResponse.FullName,
			ImageUrl:  oauthOpenIdResponse.ImageUrl,
			Bio:       fmt.Sprintf("Hi! I am %s", oauthOpenIdResponse.Username),
			CreatedAt: time.Now().Format(time.RFC3339),
		}

		createdPlayer, err := a.PlayerService.CreatePlayer(ctx, createPlayerReq)
		if err != nil {
			log.Error("failed to create player from open id")
			return *new(models.AuthLoginResponseModel), errors.New("failed to authenticate grant id")
		}

		basicJwt := &jwt.BasicJwtTokenBuilder{
			ExpiresAfter: time.Hour * 1,
		}
		usernameJwt := &jwt.UsernameJwtTokenBuilder{
			JwtTokenBuilder: basicJwt,
			Username:        createdPlayer.Username,
		}
		accessToken, err := usernameJwt.Build()
		if err != nil {
			log.Error("failed to create token")
			return *new(models.AuthLoginResponseModel), errors.New("failed to create token")
		}

		return models.AuthLoginResponseModel{
			Token: accessToken,
		}, nil
	}

	basicJwt := &jwt.BasicJwtTokenBuilder{
		ExpiresAfter: time.Hour * 1,
	}
	usernameJwt := &jwt.UsernameJwtTokenBuilder{
		JwtTokenBuilder: basicJwt,
		Username:        foundPlayer.Username,
	}
	accessToken, err := usernameJwt.Build()
	if err != nil {
		log.Error("failed to create token")
		return *new(models.AuthLoginResponseModel), errors.New("failed to create token")
	}

	return models.AuthLoginResponseModel{
		Token: accessToken,
	}, nil
}
