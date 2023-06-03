package main

import (
	"github.com/viza/banking/app"
	"github.com/viza/banking/app/logger"
)

func main() {

	logger.Info("Starting the application...")
	app.Start()
}
