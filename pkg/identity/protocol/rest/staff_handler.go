package rest

import (
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff"
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
	reqStaffView = role.ResourceStaff + "_" + role.PermissionView
)

func (s *Server) NewStaffHandler(e *echo.Echo) {
	staffRoute := e.Group("/api/staff")
	staffAdmin := e.Group("/api/admin/staff")

	staffRoute.PUT("/update", s.updateStaff, middleware.AuthMiddleware(token.Staff, ""))
	staffRoute.PUT("/me/password", s.updateStaffPassword, middleware.AuthMiddleware(token.Staff, ""))
	staffRoute.GET("/me", s.getDetailStaff, middleware.AuthMiddleware(token.Staff, ""))
	staffRoute.POST("/log-in", s.LoginStaff)

	staffAdmin.POST("/", s.createStaff, middleware.AuthMiddleware(token.Root, ""))
	staffAdmin.GET("/", s.ListStaff, middleware.AuthMiddleware(token.Staff, reqStaffView))
	staffAdmin.GET("/detail/:staffID", s.GetStaffByID, middleware.AuthMiddleware(token.Root, ""))
	staffAdmin.PUT("/detail/:staffID", s.UpdateRole, middleware.AuthMiddleware(token.Root, ""))
}

func (s *Server) updateStaff(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		cmd staff.UpdateStaffCommand
	)

	if err := c.Bind(&cmd); err != nil {
		return cc.Response400(nil, err.Error())
	}

	if err := cmd.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	err := s.Dependences.StaffSrv.Update(cc.GetRequestCtx(), &cmd)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}

	return cc.Response200(echo.Map{}, "")
}

func (s *Server) updateStaffPassword(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		cmd staff.PasswordBody
	)

	if err := c.Bind(&cmd); err != nil {
		return cc.Response400(nil, err.Error())
	}

	if err := cmd.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	err := s.Dependences.StaffSrv.ChangePassword(cc.GetRequestCtx(), &cmd)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{}, "")
}

func (s *Server) getDetailStaff(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	account, ok := cc.GetRequestCtx().Value("current_account").(*token.AccountData)
	if !ok || account.AccountType != token.Staff {
		return cc.Response400(nil, "account is invalid")
	}

	data, err := s.Dependences.StaffSrv.GetStaffByID(cc.GetRequestCtx(), account.ID)
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

	if err := c.Bind(&q); err != nil {
		return cc.Response400(nil, err.Error())
	}

	staffs, total := s.Dependences.StaffSrv.ListStaff(ctx, &q)

	return cc.Response200(echo.Map{
		"staffs": staffs,
		"total":  total,
		"limit":  q.Limit,
	}, "")
}

func (s *Server) createStaff(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body staff.CreateStaffCommand
	)

	if err := c.Bind(&body); err != nil {
		return cc.Response400(nil, err.Error())
	}

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
		staffIDString = cc.Param("staffID")
	)
	staffID, err := primitive.ObjectIDFromHex(staffIDString)
	if staffID.IsZero() || err != nil {
		return cc.Response404(nil, "staff not found")
	}

	var cmd staff.UpdateStaffRoleCommand
	if err := c.Bind(&cmd); err != nil {
		return cc.Response400(nil, err.Error())
	}

	cmd.ID = staffID

	if err := cmd.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	err = s.Dependences.StaffSrv.UpdateRole(cc.GetRequestCtx(), &cmd)
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
	if err != nil {
		return cc.Response400(nil, err.Error())
	}

	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}

func (s *Server) LoginStaff(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		cmd staff.LoginStaffCommand
	)

	if err := c.Bind(&cmd); err != nil {
		return cc.Response400(nil, err.Error())
	}

	if err := cmd.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	data, err := s.Dependences.StaffSrv.LoginStaff(cc.GetRequestCtx(), cmd)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}
