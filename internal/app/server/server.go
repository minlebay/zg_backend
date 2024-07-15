package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"sync"
)

type Router interface {
	RegisterRoutes(e *echo.Echo)
}

type Server struct {
	Config *Config
	Logger *zap.Logger
	wg     sync.WaitGroup
	Echo   *echo.Echo
	Router Router
}

func NewServer(logger *zap.Logger, config *Config, router Router) *Server {
	return &Server{
		Logger: logger,
		Config: config,
		Router: router,
	}
}

func (s *Server) StartServer(ctx context.Context) {
	go func() {
		s.Echo = echo.New()
		s.Router.RegisterRoutes(s.Echo)
		s.Echo.Start(fmt.Sprintf(":%s", s.Config.Port))
	}()
}

func (s *Server) StopServer(ctx context.Context) {
	s.wg.Wait()
	s.Echo.Shutdown(context.Background())
	s.Logger.Info("Server stopped")
}
