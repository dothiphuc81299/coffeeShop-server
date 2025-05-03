package validation

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
)

// QuestionBodyValidation ...
func QuestionBodyValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			body model.QuestionBody
		)

		c.Bind(&body)
		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}

		c.Set("questionBody", body)
		return next(c)
	}
}

// QuestionValidationBodyUpdate ...
func QuestionValidationBodyUpdate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			body model.QuestionBodyUpdate
		)

		c.Bind(&body)
		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}

		c.Set("body", body)
		return next(c)
	}
}

func QuestionAnswerValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			body model.QuestionAnswersBody
		)

		c.Bind(&body)
		if err := body.Validate(); err != nil {
			return util.ValidationError(c, err)
		}

		c.Set("body", body)
		return next(c)
	}
}
