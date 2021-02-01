package util

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
)

// EchoGetCustomCtx ...
func EchoGetCustomCtx(c echo.Context) *EchoCustomCtx {
	return &EchoCustomCtx{c}
}

// EchoCustomCtx custom echo context
type EchoCustomCtx struct {
	echo.Context
}

// GetHeaderKey ...
func (c *EchoCustomCtx) GetHeaderKey(k string) string {
	return c.Request().Header.Get(k)
}

// GetString ...
func (c *EchoCustomCtx) GetString(key string) string {
	v := c.Get(key)
	res, _ := v.(string)
	return res
}

// GetPageQuery ...
func (c *EchoCustomCtx) GetPageQuery() int64 {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	return int64(page)
}

// GetLimitQuery ...
func (c *EchoCustomCtx) GetLimitQuery() int64 {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit <= 0 {
		limit = 20
	}
	return int64(limit)
}

// GetLang ...
func (c *EchoCustomCtx) GetLang() string {
	lang := c.GetHeaderKey(config.HeaderAcceptLanguage)
	if lang != config.LangEn {
		lang = config.LangVi
	}
	return lang
}

// GetCurrentUserID ...
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
		id = GetAppIDFromHex(m["_id"].(string))
	}

	return id
}

// GetRequestCtx get request context
func (c *EchoCustomCtx) GetRequestCtx() context.Context {
	return c.Request().Context()
}

// GetUserPlatform ...
func (c *EchoCustomCtx) GetUserPlatform() string {
	return strings.ToLower(c.Request().Header.Get("OS-NAME"))
}

// GetAppIDFromQuery ...
func (c *EchoCustomCtx) GetAppIDFromQuery(key string) primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(c.QueryParam(key))
	return id
}

// Response200 response success
func (c *EchoCustomCtx) Response200(data interface{}, msgKey string) error {
	if msgKey == "" {
		msgKey = locale.Default200
	}

	resp := getResponse(c.GetLang(), data, msgKey, http.StatusOK)
	sendResponse(c, resp)
	return nil
}

// Response400 bad request
func (c *EchoCustomCtx) Response400(data interface{}, msgKey string) error {
	if msgKey == "" {
		msgKey = locale.Default400
	}

	resp := getResponse(c.GetLang(), data, msgKey, http.StatusBadRequest)
	sendResponse(c, resp)
	return nil
}

// Response401 Unauthorized
func (c *EchoCustomCtx) Response401(data interface{}, msgKey string) error {
	if msgKey == "" {
		msgKey = locale.Default401
	}

	resp := getResponse(c.GetLang(), data, msgKey, http.StatusUnauthorized)
	sendResponse(c, resp)
	return nil
}

// Response404 not found
func (c *EchoCustomCtx) Response404(data interface{}, msgKey string) error {
	if msgKey == "" {
		msgKey = locale.Default404
	}

	resp := getResponse(c.GetLang(), data, msgKey, http.StatusNotFound)
	sendResponse(c, resp)
	return nil
}

// Response ...
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

func getResponse(lang string, data interface{}, messageKey string, httpCode int) Response {
	if data == nil {
		data = echo.Map{}
	}

	var respInfo locale.Locale
	respInfo = locale.GetByKey(lang, messageKey)
	respInfo.Message.GetDisplay(lang)

	return Response{
		HTTPCode: httpCode,
		Data:     data,
		Message:  respInfo.Message.Display,
		Code:     respInfo.Code,
	}
}
