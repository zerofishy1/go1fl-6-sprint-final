package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// 1. Создаём логгер с префиксом для удобства
	logger := log.New(os.Stdout, "[MORSE-CONVERTER] ", log.LstdFlags)

	// 2. Создаём сервер с нашим логгером
	srv := server.NewServer(logger)

	// 3. Запускаем сервер
	logger.Println("Сервер конвертера азбуки Морзе запускается...")
	logger.Println("Откройте http://localhost:8080 в браузере")
	logger.Println("Используйте тестовые файлы test и test12 для проверки")

	// 4. Запускаем сервер и проверяем ошибки
	if err := srv.Run(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
