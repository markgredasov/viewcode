package main

import (
	"fmt"
	"os"

	"github.com/markgredasov/viewcode/internal/logger"
)

func main() {
	loggerConfig := logger.NewConfigMust()
	logger, err := logger.NewLogger(loggerConfig)
	if err != nil {
		//nolint:forbidigo // logger was not initialized to write here smth
		fmt.Println("failed to initialize logger:", err)
		os.Exit(1)
	}

	logger.Debug("starting application")
}
