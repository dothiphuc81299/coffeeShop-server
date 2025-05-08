package rest

import (
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/util"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/role"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/user"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"github.com/labstack/echo/v4"
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
		body user.CreateUserCommand
	)

	if err := c.Bind(&body); err != nil {
		return cc.Response400(nil, err.Error())
	}

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
		body user.CreateLoginUserCommand
	)

	if err := c.Bind(&body); err != nil {
		return cc.Response400(nil, err.Error())
	}

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
		cc = util.EchoGetCustomCtx(c)
	)

	cmd := query.CommonQuery{
		Page:  cc.GetPageQuery(),
		Limit: cc.GetLimitQuery(),
		Sort: bson.D{
			bson.E{Key: "createdAt", Value: -1},
		},
		Active:   c.QueryParam("active"),
		Username: c.QueryParam("username"),
	}

	if err := c.Bind(&cmd); err != nil {
		return cc.Response400(nil, err.Error())
	}

	data, total := s.Dependences.UserSrv.Search(cc.GetRequestCtx(), &cmd)
	result := query.ResponseAppListData{
		Data:         data,
		Total:        total,
		LimitPerPage: cmd.Limit,
	}
	return cc.Response200(result, "")
}

func (s *Server) SendEmail(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		body user.SendUserEmailCommand
	)

	if err := c.Bind(&body); err != nil {
		return cc.Response400(nil, err.Error())
	}

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
	var cmd user.VerifyEmailCommand

	if err := c.Bind(&cmd); err != nil {
		return cc.Response400(nil, err.Error())
	}

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
		cmd user.UpdateUserCommand
	)

	if err := c.Bind(&cmd); err != nil {
		return cc.Response400(nil, err.Error())
	}

	err := s.Dependences.UserSrv.UpdateUser(cc.GetRequestCtx(), &cmd)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(nil, "")
}

func (s *Server) getDetailUser(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	account, ok := cc.GetRequestCtx().Value("current_account").(*token.AccountData)
	if !ok || account.AccountType != token.User {
		return cc.Response400(nil, "account is invalid")
	}

	data := s.Dependences.UserSrv.GetDetailUser(cc.GetRequestCtx(), account.ID)
	return cc.Response200(echo.Map{
		"data": data,
	}, "")
}

func (s *Server) ChangePassword(c echo.Context) error {
	cc := util.EchoGetCustomCtx(c)
	var (
		cmd user.ChangePasswordUserCommand
	)

	if err := c.Bind(&cmd); err != nil {
		return cc.Response400(nil, err.Error())
	}

	err := s.Dependences.UserSrv.ChangePassword(cc.GetRequestCtx(), &cmd)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{}, "")
}
