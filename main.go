package main

import (
	"github.com/PedroBSanchez/gobooks.git/config"
	"github.com/PedroBSanchez/gobooks.git/router"
)


var (
	logger *config.Logger
)


func main () {
	
	
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Init()

	if err != nil {
		logger.ErrorF("CONFIG initialization error: %v", err)
		return
	}

	//Initialize Router
	router.InitializeRouter()


	
}