package main

import (
	"github.com/eduardylopes/gopportunities/config"
	"github.com/eduardylopes/gopportunities/router"
)

func main() {
	logger := config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
