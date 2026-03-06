package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New()

	// Создаем сервер
	srv := server.CreateServer(logger)

	// Запускаем сервер
	log.Fatal(srv.ListenAndServe())
}
