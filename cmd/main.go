package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(
		os.Stdout,
		"app ",
		log.LstdFlags,
	)

	// Создаем сервер
	srv := server.NewServer(logger)

	// Запускаем сервер
	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		logger.Fatal(err)
	}
}
