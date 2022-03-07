package main

import (
	"banking/app"
	"banking/logger"
)

func main() {
	logger.Info("start application...")
	app.Start()
}
