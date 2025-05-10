package rest

import (
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/role"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/query"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	reqRoleView = role.ResourceRole + "_" + role.PermissionView
)

func (s *Server) NewStaffRoleHandler(e *echo.Echo) {
	admin := e.Group("/api/admin/staff-role")

	admin.POST("/", s.createStaffRole, middleware.AuthMiddleware(token.Root, ""))
	admin.PUT("/detail/:roleID", s.updateRole, middleware.AuthMiddleware(token.Root, ""))
	admin.GET("/", s.ListRoleStaff, middleware.AuthMiddleware(token.Staff, ""))
	admin.GET("/permissions", s.SearchPermission, middleware.AuthMiddleware(token.Staff, reqRoleView))
	admin.GET("/detail/:roleID", s.GetRoleByID, middleware.AuthMiddleware(token.Staff, reqRoleView))
	admin.DELETE("/detail/:roleID", s.deleteRole, middleware.AuthMiddleware(token.Root, ""))
}

func (s *Server) createStaffRole(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		cmd role.CreateStaffRoleCommand
	)

	if err := c.Bind(&cmd); err != nil {
		return cc.Response400(nil, err.Error())
	}

	if err := cmd.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	err := s.Dependences.StaffRoleSrv.Create(cc.GetRequestCtx(), cmd)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{}, "")
}

func (s *Server) updateRole(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		cmd role.UpdateStaffRoleCommand
	)

	if err := c.Bind(&cmd); err != nil {
		return cc.Response400(nil, err.Error())
	}

	staffRoleIDString := c.Param("roleID")
	staffRoleID, err := primitive.ObjectIDFromHex(staffRoleIDString)
	if staffRoleID.IsZero() || err != nil {
		return cc.Response400(nil, "")
	}

	cmd.ID = staffRoleID
	if err := cmd.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	err = s.Dependences.StaffRoleSrv.Update(cc.GetRequestCtx(), cmd)
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

	if err := c.Bind(&q); err != nil {
		return cc.Response400(nil, err.Error())
	}

	roles, total := s.Dependences.StaffRoleSrv.ListStaffRole(ctx, &q)
	return cc.Response200(echo.Map{
		"staffRoles": roles,
		"total":      total,
		"limit":      q.Limit,
	}, "")
}

func (s *Server) SearchPermission(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	permissions := role.Permissions
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
