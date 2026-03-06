package main

import (
	"log"

	"6sprint/internal/server"
)

func main() {
    // Создаем логгер
    logger := log.New(os.Stdout, "", log.LstdFlags)

    // Создаем сервер
    srv := server.CreateServer(logger)

    // Запускаем сервер
    log.Fatal(srv.ListenAndServe())
}
