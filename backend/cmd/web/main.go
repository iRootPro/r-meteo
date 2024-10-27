package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/irootpro/r-meteo/internal/config"
	"github.com/irootpro/r-meteo/pkg/logger"
)

func main() {
	logger := logger.NewLogger()
	config := config.NewConfig(logger)

	logger.Info("Starting server", "port", config.HTTP_PORT)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Получен запрос", "path", r.URL.Path)
		fmt.Fprintln(w, "Добро пожаловать в MeteoWeb API!")
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.HTTP_PORT), nil); err != nil {
		logger.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
