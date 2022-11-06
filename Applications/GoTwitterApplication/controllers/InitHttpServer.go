package controllers

import (
	"github.com/mrkresnofatih/gotwitter/applications"
	"github.com/mrkresnofatih/gotwitter/services"
	"sync"
)

func InitHttpServer(appRunState *sync.WaitGroup) {
	go func() {
		httpServerObj := &ApplicationServer{}

		authController := &AuthController{
			AuthService: &services.AuthService{
				PlayerService: &services.PlayerService{
					MongoDb: applications.GetMongoInstance(),
				},
				OauthService: &services.OauthService{},
			},
		}
		httpServerObj.AddController(authController)

		httpServerObj.Initialize()
		httpServerObj.Router.Logger.Fatal(httpServerObj.Router.Start(":1324"))
		appRunState.Done()
	}()
}
