package main

import (
	"github.com/hcncc/gopportunities-api/config"
	"github.com/hcncc/gopportunities-api/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")

	// Initialize Configurations
	err := config.InitializeConfigurations()

	if err != nil {
		logger.Errorf("Initialize Configurations Error: %v", err)

		return
	}

	//initialize Server application
	router.InitializeServer()
}
