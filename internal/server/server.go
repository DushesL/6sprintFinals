package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Структура сервера
type Server struct {
	logger *log.Logger
	server *http.Server
}

// Функция создания сервера
func NewServer(logger *log.Logger) *Server {
	// Создаем новый HTTP-роутер
	router := http.NewServeMux()

	// Регистрируем хендлеры
	router.HandleFunc("/", handlers.HtmlHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	// Создаем экземпляр http.Server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		server: server,
	}
}

// Метод для запуска сервера
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

// Метод для остановки сервера
