package rest

import (
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/util"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/role"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/user"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	reqUserView = role.ResourceUser + "_" + role.PermissionView
)

func (s *Server) NewUserHandler(e *echo.Echo) {
	admin := e.Group("/api/admin/users")
	user := e.Group("/api/users")

	user.POST("/sign-up", s.createUser)
	user.POST("/log-in", s.loginUser)
	user.GET("/me", s.getDetailUser, middleware.AuthMiddleware(token.User, ""))
	user.PUT("/me/update", s.UpdateUser, middleware.AuthMiddleware(token.User, ""))
	user.PUT("/me/password", s.ChangePassword, middleware.AuthMiddleware(token.User, ""))
	user.POST("/send-email", s.SendEmail)
	user.POST("/verify-email", s.VerifyEmail)

	admin.GET("/list", s.search, middleware.AuthMiddleware(token.Staff, reqUserView))
}

func (s *Server) createUser(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body = c.Get("body").(user.CreateUserCommand)
	)

	err := body.Validate()
	if err != nil {
		return cc.Response400(nil, err.Error())
	}

	email, err := s.Dependences.UserSrv.CreateUser(cc.GetRequestCtx(), body)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"data": email,
	}, "")
}

func (s *Server) loginUser(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body = c.Get("body").(user.CreateLoginUserCommand)
	)

	err := body.Validate()
	if err != nil {
		return cc.Response400(nil, err.Error())
	}

	data, err := s.Dependences.UserSrv.LoginUser(cc.GetRequestCtx(), body)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}

func (s *Server) search(c echo.Context) error {
	var (
		cc  = util.EchoGetCustomCtx(c)
		cmd = query.CommonQuery{
			Limit:   cc.GetLimitQuery(),
			Page:    cc.GetPageQuery(),
			Active:  c.QueryParam("active"),
			Keyword: c.QueryParam("keyword"),
			Sort:    bson.D{{"createdAt", -1}},
		}
	)

	data, total := s.Dependences.UserSrv.Search(cc.GetRequestCtx(), cmd)
	result := query.ResponseAppListData{
		Data:         data,
		Total:        total,
		LimitPerPage: cmd.Limit,
	}
	return cc.Response200(result, "")
}

func (s *Server) GetDetailUser(c echo.Context) error {
	var (
		cc   = util.EchoGetCustomCtx(c)
		user = c.Get("user").(user.UserRaw)
	)

	result := s.Dependences.UserSrv.GetDetailUser(cc.GetRequestCtx(), user)
	return cc.Response200(echo.Map{
		"data": result,
	}, "")
}

func (s *Server) UserGetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customCtx := util.EchoGetCustomCtx(c)
		id := c.Param("userID")
		if id == "" {
			return next(c)
		}
		userID := util.GetObjectIDFromHex(id)
		if userID.IsZero() {
			return customCtx.Response400(nil, "bad request")
		}
		user, err := s.Dependences.UserSrv.FindByID(customCtx.GetRequestCtx(), userID)
		if err != nil {
			return customCtx.Response404(nil, "not found")
		}

		c.Set("user", user)
		return next(c)
	}
}

func (s *Server) SendEmail(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body = c.Get("body").(user.SendUserEmailCommand)
	)

	if err := body.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	err := s.Dependences.UserSrv.SendEmail(cc.GetRequestCtx(), body)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(nil, "")
}

func (s *Server) VerifyEmail(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		cmd = c.Get("body").(user.VerifyEmailCommand)
	)

	if err := cmd.Validate(); err != nil {
		return cc.Response400(nil, err.Error())
	}

	err := s.Dependences.UserSrv.VerifyEmail(cc.GetRequestCtx(), cmd)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(nil, "")
}

func (s *Server) UpdateUser(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body = c.Get("body").(user.UpdateUserCommand)
		user = c.Get("user").(user.UserRaw)
	)

	err := s.Dependences.UserSrv.UpdateUser(cc.GetRequestCtx(), user, body)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(nil, "")
}

func (s *Server) getDetailUser(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		user = c.Get("user").(user.UserRaw)
	)

	data := s.Dependences.UserSrv.GetDetailUser(cc.GetRequestCtx(), user)
	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}

func (s *Server) ChangePassword(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		entity = c.Get("user").(user.UserRaw)
		body   = c.Get("body").(user.ChangePasswordUserCommand)
	)

	err := s.Dependences.UserSrv.ChangePassword(cc.GetRequestCtx(), entity, body)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{}, "")
}
