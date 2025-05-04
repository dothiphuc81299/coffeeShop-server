package rest

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) NewStaffHandler(e *echo.Echo) {

	staffRoute := e.Group("/api/staff")
	staffAdmin := e.Group("/api/admin/staff")

	staffRoute.PUT("/update", s.updateStaff, middleware.CheckStaff(token.Staff))
	staffRoute.PUT("/me/password", s.updateStaffPassword, middleware.CheckStaff(token.Staff))
	staffRoute.GET("/me", s.getDetailStaff, middleware.CheckStaff(token.Staff))
	staffRoute.POST("/log-in", s.LoginStaff)

	staffAdmin.POST("", s.createStaff, middleware.CheckPermissionRoot(token.Root))
	staffAdmin.GET("", s.ListStaff, middleware.CheckPermission(config.ModelFieldStaff, config.PermissionView, token.Staff))
	staffAdmin.GET("/:staffID", s.GetStaffByID, middleware.CheckPermissionRoot(token.Root))
	staffAdmin.PUT("/:staffID", s.UpdateRole, middleware.CheckPermissionRoot(token.Root))
}

func (s *Server) updateStaff(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body   = c.Get("bodyUpdate").(staff.UpdateStaffCommand)
		enitty = c.Get("staff").(staff.Staff)
	)

	if err := body.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	err := s.Dependences.StaffSrv.Update(cc.GetRequestCtx(), body, enitty)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}

	return cc.Response200(echo.Map{}, "")
}

func (s *Server) updateStaffPassword(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body   = c.Get("body").(staff.PasswordBody)
		entity = c.Get("staff").(staff.Staff)
	)

	if err := body.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	err := s.Dependences.StaffSrv.ChangePassword(cc.GetRequestCtx(), entity, body)

	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{}, "")
}

func (s *Server) getDetailStaff(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		staff = c.Get("staff").(staff.Staff)
	)

	data, err := s.Dependences.StaffSrv.GetStaffByID(cc.GetRequestCtx(), staff.ID)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}

	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}

func (s *Server) ListStaff(c echo.Context) error {
	var (
		cc  = util.EchoGetCustomCtx(c)
		ctx = cc.GetRequestCtx()
	)

	q := query.CommonQuery{
		Page:  cc.GetPageQuery(),
		Limit: cc.GetLimitQuery(),
		Sort: bson.D{
			bson.E{Key: "createdAt", Value: -1},
		},
		Active:   c.QueryParam("active"),
		Username: c.QueryParam("username"),
	}

	staffs, total := s.Dependences.StaffSrv.ListStaff(ctx, q)

	return cc.Response200(echo.Map{
		"staffs": staffs,
		"total":  total,
		"limit":  q.Limit,
	}, "")
}

func (s *Server) createStaff(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body = c.Get("body").(staff.CreateStaffCommand)
	)

	if err := body.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	err := s.Dependences.StaffSrv.Create(cc.GetRequestCtx(), body)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}

	return cc.Response200(echo.Map{}, "")
}

func (s *Server) UpdateRole(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body   = c.Get("body").(staff.UpdateStaffRoleCommand)
		enitty = c.Get("staff").(staff.Staff)
	)

	err := s.Dependences.StaffSrv.UpdateRole(cc.GetRequestCtx(), body, enitty)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}

	return cc.Response200(nil, "")
}

func (s *Server) GetStaffByID(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		staffIDString = cc.Param("staffID")
	)
	staffID, err := primitive.ObjectIDFromHex(staffIDString)
	if staffID.IsZero() || err != nil {
		return cc.Response404(nil, "staff not found")
	}

	data, err := s.Dependences.StaffSrv.GetStaffByID(cc.GetRequestCtx(), staffID)

	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}

func (s *Server) LoginStaff(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body = c.Get("body").(staff.LoginStaffCommand)
	)

	if err := body.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	data, err := s.Dependences.StaffSrv.LoginStaff(cc.GetRequestCtx(), body)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}
