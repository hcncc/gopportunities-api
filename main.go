package main

import (
	"fmt"

	"github.com/hcncc/gopportunities-api/config"
	"github.com/hcncc/gopportunities-api/router"
)

func main() {
	// Initialize Configurations
	err := config.Init()

	if err != nil {
		fmt.Println("Error in initialize configurations")

		return
	}

	//initialize Server application
	router.InitializeServer()
}
