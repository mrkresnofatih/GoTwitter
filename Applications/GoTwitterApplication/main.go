package main

import (
	"github.com/mrkresnofatih/gotwitter/applications"
	"github.com/mrkresnofatih/gotwitter/controllers"
	"sync"
)

func main() {
	var appRunState sync.WaitGroup
	appRunState.Add(1)
	applications.LoadEnvVars()
	controllers.InitHttpServer(&appRunState)
	appRunState.Wait()
}
