package rest

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/role"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) NewStaffRoleHandler(e *echo.Echo) {
	admin := e.Group("/api/admin/staff-role")

	admin.POST("", s.createStaffRole, middleware.CheckPermissionRoot(token.Root))
	admin.PUT("/:roleID", s.Update, middleware.CheckPermissionRoot(token.Root))
	admin.GET("", s.ListRoleStaff, middleware.CheckPermission(config.ModelFieldRole, config.PermissionView, token.Staff))
	admin.GET("/permissions", s.SearchPermission)
	admin.GET("/:roleID", s.GetRoleByID, middleware.CheckPermission(config.ModelFieldRole, config.PermissionView, token.Staff))
	admin.DELETE("/:roleID", s.deleteRole, middleware.CheckPermissionRoot(token.Root))
}

func (s *Server) createStaffRole(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body = c.Get("body").(role.CreateStaffRoleCommand)
	)

	if err := body.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	err := s.Dependences.StaffRoleSrv.Create(cc.GetRequestCtx(), body)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{}, "")
}

func (s *Server) Update(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body = c.Get("body").(role.CreateStaffRoleCommand)
	)

	staffRoleIDString := c.Param("roleID")
	staffRoleID, err := primitive.ObjectIDFromHex(staffRoleIDString)
	if staffRoleID.IsZero() || err != nil {
		return cc.Response400(nil, "")
	}

	err = s.Dependences.StaffRoleSrv.Update(cc.GetRequestCtx(), staffRoleID, body)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{}, "")
}

func (s *Server) ListRoleStaff(c echo.Context) error {
	var (
		cc  = util.EchoGetCustomCtx(c)
		ctx = cc.GetRequestCtx()
		q   = query.CommonQuery{
			Sort: bson.D{
				bson.E{Key: "createdAt", Value: -1},
			},
		}
	)

	roles, total := s.Dependences.StaffRoleSrv.ListStaffRole(ctx, q)
	return cc.Response200(echo.Map{
		"staffRoles": roles,
		"total":      total,
		"limit":      q.Limit,
	}, "")
}

func (s *Server) SearchPermission(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	permissions := config.Permissions
	return cc.Response200(echo.Map{
		"permissions": permissions,
	}, "")
}

func (s *Server) deleteRole(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	staffRoleIDString := c.Param("roleID")
	staffRoleID, err := primitive.ObjectIDFromHex(staffRoleIDString)
	if staffRoleID.IsZero() || err != nil {
		return cc.Response400(nil, "")
	}

	err = s.Dependences.StaffRoleSrv.Delete(cc.GetRequestCtx(), staffRoleID)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(nil, "")
}

func (s *Server) GetRoleByID(c echo.Context) error {
	var (
		cc = util.EchoGetCustomCtx(c)
	)

	staffRoleIDString := c.Param("roleID")
	staffRoleID, err := primitive.ObjectIDFromHex(staffRoleIDString)
	if staffRoleID.IsZero() || err != nil {
		return cc.Response400(nil, "")
	}

	role, err := s.Dependences.StaffRoleSrv.FindByID(cc.GetRequestCtx(), staffRoleID)
	if role.ID.IsZero() || err != nil {
		return cc.Response404(nil, "not found")
	}

	return cc.Response200(echo.Map{
		"role": role,
	}, "")
}
