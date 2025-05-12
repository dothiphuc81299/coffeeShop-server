package util

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EchoGetCustomCtx(c echo.Context) *EchoCustomCtx {
	return &EchoCustomCtx{c}
}

type EchoCustomCtx struct {
	echo.Context
}

func (c *EchoCustomCtx) GetHeaderKey(k string) string {
	return c.Request().Header.Get(k)
}

func (c *EchoCustomCtx) GetPageQuery() int64 {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	return int64(page)
}

func (c *EchoCustomCtx) GetLimitQuery() int64 {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit <= 0 {
		limit = 20
	}
	return int64(limit)
}

func (c *EchoCustomCtx) GetCurrentUserID() (id primitive.ObjectID) {
	token := c.Get("user")
	if token == nil {
		return
	}
	data, ok := token.(*jwt.Token)

	if !ok {
		return
	}

	m, ok := data.Claims.(jwt.MapClaims)

	if ok && data.Valid && m["_id"] != "" {
		id = GetObjectIDFromHex(m["_id"].(string))
	}

	return id
}

func (c *EchoCustomCtx) GetRequestCtx() context.Context {
	return c.Request().Context()
}

func (c *EchoCustomCtx) GetUserPlatform() string {
	return strings.ToLower(c.Request().Header.Get("OS-NAME"))
}

func (c *EchoCustomCtx) GetObjectIDFromQuery(key string) primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(c.QueryParam(key))
	return id
}

func (c *EchoCustomCtx) Response200(data interface{}, msgKey string) error {
	if msgKey == "" {
		msgKey = "ok"
	}

	resp := getResponse(data, msgKey, http.StatusOK)
	sendResponse(c, resp)
	return nil
}

func (c *EchoCustomCtx) Response400(data interface{}, msgKey string) error {
	if msgKey == "" {
		msgKey = "bad request"
	}

	resp := getResponse(data, msgKey, http.StatusBadRequest)
	sendResponse(c, resp)
	return nil
}

func (c *EchoCustomCtx) Response401(data interface{}, msgKey string) error {
	if msgKey == "" {
		msgKey = "unauthorized"
	}

	resp := getResponse(data, msgKey, http.StatusUnauthorized)
	sendResponse(c, resp)
	return nil
}

func (c *EchoCustomCtx) Response404(data interface{}, msgKey string) error {
	if msgKey == "" {
		msgKey = "not found"
	}

	resp := getResponse(data, msgKey, http.StatusNotFound)
	sendResponse(c, resp)
	return nil
}

type Response struct {
	HTTPCode int         `json:"-"`
	Data     interface{} `json:"data"`
	Code     int         `json:"code"`
	Message  string      `json:"message"`
}

func sendResponse(c echo.Context, data Response) {
	c.JSON(data.HTTPCode, echo.Map{
		"data":    data.Data,
		"message": data.Message,
		"code":    data.Code,
	})
}

func getResponse(data interface{}, messageKey string, httpCode int) Response {
	if data == nil {
		data = echo.Map{}
	}

	return Response{
		HTTPCode: httpCode,
		Data:     data,
		Message:  messageKey,
		Code:     httpCode,
	}
}
