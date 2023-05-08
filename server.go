package course_work

import (
	"context"
	"net/http"
	"time"
)

// Server - Структура сервера
type Server struct {
	httpServer *http.Server
}

// Run - метод для инициализации сервера
func (s *Server) Run(port string, handler http.Handler) error {
	// Создаем экземпляр (синглтон) для сервера и конфигурируем его
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        handler,
	}

	// Возвращаем метод для запуска сервера
	// внутри запускается горутина с бесконечным циклом for
	return s.httpServer.ListenAndServe()
}

// Shutdown - метод для завершения работы сервера
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
