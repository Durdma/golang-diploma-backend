package server

import (
	"context"
	"net/http"
	"sas/internal/config"
)

// @title University Platform API
// @version 1.0
// @description API Server for University Platform

// @host localhost:8080
// @BasePath /

// @securityDefinition.apikey AdminAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey EditorsAuth
// @in header
// @name Authorization

// Server - Структура для работы с сервером
type Server struct {
	httpServer *http.Server
}

// NewServer - Создание нового сервера
func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         ":" + cfg.HTTP.Port,
			Handler:      handler,
			ReadTimeout:  cfg.HTTP.ReadTimeout,
			WriteTimeout: cfg.HTTP.WriteTimeout,
		},
	}
}

// Run - запуск сервера
func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

// Stop - остановка сервера
func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
