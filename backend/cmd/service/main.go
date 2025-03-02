package main

import (
	"homework_ipl/internal/config"
	"homework_ipl/internal/http-server/server"
	"homework_ipl/router"
	"homework_ipl/utils/logger"
)

func main() {
	logger := logger.Logger()
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed to load config", "error", err)
		return
	}

	logger.Info("Start config", "config", cfg)

	router := router.SetupRouter(cfg)

	if err := server.StartServer(router, cfg); err != nil {
		logger.Error("Failed to start server", "error", err)
		return
	}
}
