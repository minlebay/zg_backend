package server

import (
	"context"
	"fmt"
	"sync"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Server struct {
	Config *Config
	Logger *zap.Logger
	wg     sync.WaitGroup
	Echo   *echo.Echo
}

func NewServer(logger *zap.Logger, config *Config) *Server {
	return &Server{
		Logger: logger,
		Config: config,
	}
}

func (s *Server) StartServer(ctx context.Context) {
	go func() {
		s.Echo = echo.New()
		s.Echo.Start(fmt.Sprintf(":%s", s.Config.Port))
	}()
}

func (s *Server) StopServer(ctx context.Context) {
	s.wg.Wait()
	s.Echo.Shutdown(context.Background())
	s.Logger.Info("Server stopped")
}
