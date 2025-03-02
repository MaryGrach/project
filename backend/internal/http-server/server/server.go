package server

import (
	"net/http"

	"homework_ipl/internal/config"

	"github.com/go-chi/chi/v5"
	"homework_ipl/utils/logger"
)

func StartServer(router *chi.Mux, cfg *config.Config) error {
	logger.Logger().Info("Server is starting:", "address", cfg.HTTPServer.Address)

	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
