package main

import (
	"github.com/Rt-00/gopportunities.git/config"
	"github.com/Rt-00/gopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	// Initialize Logger
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
