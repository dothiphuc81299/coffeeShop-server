package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/game/handler"
	validation "github.com/dothiphuc81299/coffeeShop-server/game/validation"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/labstack/echo/v4"
)

// InitQuestionAdmin ...
func InitQuestionAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.QuestionAdminHandler{
		QuestionAdminService: cs.Question,
	}

	g := e.Group("/game/questions")
	g.POST("", h.Create, validation.QuestionBodyValidation)

	g.PUT("/:questionID", h.Update, h.QuestionGetByID, validation.QuestionValidationBodyUpdate)

	// change status
	g.PATCH("/:questionID/status", h.ChangeStatus, h.QuestionGetByID)
	g.GET("", h.ListAll)

	// get detail quesition
	g.GET("/:questionID", h.GetDetail, h.QuestionGetByID)
}
