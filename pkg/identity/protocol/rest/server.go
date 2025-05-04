package rest

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/role"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/user"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Dependences *Dependences
}

type Dependences struct {
	UserSrv      user.Service
	StaffSrv     staff.Service
	StaffRoleSrv role.Service
}

func NewServer(deps *Dependences) *Server {
	return &Server{
		Dependences: deps,
	}
}

func (s *Server) New() {
	server := echo.New()

	server.Use(middleware.AppMiddleware()...)
	s.NewStaffHandler(server)
	s.NewUserHandler(server)
	s.NewStaffRoleHandler(server)
}
