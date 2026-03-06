package main

import (
	"log"

	"6sprint/6/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New()

	// Создаем сервер
	srv := server.CreateServer(logger)

	// Запускаем сервер
	log.Fatal(srv.ListenAndServe())
}
