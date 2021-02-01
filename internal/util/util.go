package util

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
)

func getMessage(err error) string {
	err1, ok := err.(validation.Errors)
	if !ok {
		err2, ok := err.(validation.ErrorObject)
		if ok {
			return err2.Message()
		}
		return err.Error()
	}
	for _, item := range err1 {
		if item == nil {
			continue
		}
		return getMessage(item)
	}
	return err.Error()
}

func getErrorKey(err error) (key string) {
	errString := err.Error()
	errString = strings.ReplaceAll(errString, " ", "")
	values := strings.Split(errString, ";")
	if len(values) <= 0 {
		return
	}
	values = strings.Split(values[0], ":")
	if len(values) > 1 {
		key = values[1]
		if strings.HasSuffix(key, ".") {
			key = key[:len(values[1])-1]
		}
	}
	return
}

// ValidationError ...
func ValidationError(c echo.Context, err error) error {
	cc := EchoGetCustomCtx(c)
	return cc.Response400(nil, getMessage(err))
}
