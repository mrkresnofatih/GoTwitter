package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/mrkresnofatih/gotwitter/models"
	"io"
	"net/http"
	"time"
)

type IOauthService interface {
	AuthenticateGrant(model models.OauthAuthenticateGrantRequestModel) (models.OauthAuthenticateGrantResponseModel, error)
	GetOpenId(model models.OauthOpenIdRequestModel) (models.OauthOpenIdResponseModel, error)
}

type OauthService struct{}

func (o *OauthService) AuthenticateGrant(model models.OauthAuthenticateGrantRequestModel) (models.OauthAuthenticateGrantResponseModel, error) {
	log.Info("Start AuthenticateGrant")
	serialized, err := json.Marshal(model)
	if err != nil {
		log.Error("failed to marshal authenticate grant request model")
		return *new(models.OauthAuthenticateGrantResponseModel), errors.New("failed to marshal authenticate grant request model")
	}
	bodyReader := bytes.NewReader(serialized)
	authGrantUrl := "http://localhost:1323/oauth2/authenticate-grant"
	req, err := http.NewRequest(http.MethodPost, authGrantUrl, bodyReader)
	if err != nil {
		log.Info("failed to create request")
		return *new(models.OauthAuthenticateGrantResponseModel), errors.New("failed to marshal authenticate grant request model")
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Do(req)
	if err != nil {
		log.Error("failed when executing request")
		return *new(models.OauthAuthenticateGrantResponseModel), errors.New("failed to execute request")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("failed to read request body")
		return *new(models.OauthAuthenticateGrantResponseModel), errors.New("failed to read request body")
	}

	oauthGrantResponse := models.BaseResponseModel[models.OauthAuthenticateGrantResponseModel]{}
	err = json.Unmarshal(body, &oauthGrantResponse)
	if err != nil {
		log.Error("failed to unmarshal response body")
		return *new(models.OauthAuthenticateGrantResponseModel), errors.New("failed to unmarshal response body")
	}

	log.Info("successful auth grant")
	return oauthGrantResponse.Data, nil
}

func (o *OauthService) GetOpenId(model models.OauthOpenIdRequestModel) (models.OauthOpenIdResponseModel, error) {
	log.Info("Start GetOpenId")
	getOpenIdUrl := "http://localhost:1323/player/get-my-profile"
	req, err := http.NewRequest(http.MethodGet, getOpenIdUrl, nil)
	if err != nil {
		log.Error("failed to create new req")
		return *new(models.OauthOpenIdResponseModel), errors.New("failed to create req")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", model.Token))

	client := http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Do(req)
	if err != nil {
		log.Error("failed to exec req")
		return *new(models.OauthOpenIdResponseModel), errors.New("failed to exec req")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("failed to read response")
		return *new(models.OauthOpenIdResponseModel), errors.New("failed to read response")
	}

	var result models.BaseResponseModel[models.OauthOpenIdResponseModel]
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Error("failed to unmarshal response")
		return *new(models.OauthOpenIdResponseModel), errors.New("failed to unmarshal response")
	}

	log.Info("successfully obtained openid")
	return result.Data, nil
}
