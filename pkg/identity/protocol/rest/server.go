package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/config"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/role"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Dependences *Dependences
	Cfg         *config.Config
	HTTPServer  *http.Server
}

type Dependences struct {
	UserSrv      user.Service
	StaffSrv     staff.Service
	StaffRoleSrv role.Service
}

func NewServer(deps *Dependences, cfg *config.Config) *Server {
	return &Server{
		Cfg:         cfg,
		Dependences: deps,
	}
}

func (s *Server) Run(ctx context.Context) error {
	stopCh := ctx.Done()
	server := echo.New()

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		MaxAge:           600,
	}))

	s.NewStaffHandler(server)
	s.NewUserHandler(server)
	s.NewStaffRoleHandler(server)

	s.HTTPServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Cfg.Server.HTTPPort),
		Handler: server,
	}

	go func() {
		<-stopCh
		log.Println("Shutting down HTTP server...")

		if err := s.HTTPServer.Shutdown(context.Background()); err != nil {
			log.Printf("❌ Server forced to shutdown: %v\n", err)
		} else {
			log.Println("✅ Server exited properly")
		}
	}()

	log.Printf("🚀 Starting HTTP server on port %s...\n", s.Cfg.Server.HTTPPort)
	if err := s.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.HTTPServer != nil {
		return s.HTTPServer.Shutdown(ctx)
	}
	return nil
}
