package main

import (
	"log"
	"os"
	"6sprint/internal/server" // Путь к пакету server
)

func main() {
	// Создание логгера
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	// Создание сервера
	srv := server.CreateServer(logger)

	// Запуск сервера и обработка ошибок
	log.Println("Starting server on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Error starting server: ", err) // Логируем ошибку с использованием Fatal
	}
}
