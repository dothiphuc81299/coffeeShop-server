package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/feedback/handler"
	"github.com/dothiphuc81299/coffeeShop-server/feedback/validation"
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/labstack/echo/v4"
)

// InitFeedbackApp ...
func InitFeedbackApp(e *echo.Echo, cs *model.AppService, d *model.CommonDAO) {
	h := &handler.FeedbackAppHandler{
		FeedbackAppService: cs.Feedback,
	}

	g := e.Group("/feedback")
	g.POST("", h.Create, middleware.CheckUser(d), validation.FeedbackBodyValidation)

	g.GET("", h.GetList)

	g.GET("/:feedbackID", h.GetDetail, h.FeedbackGetByID)

	g.PUT("/:feedbackID", h.Update, middleware.CheckUser(d), validation.FeedbackBodyValidation, h.FeedbackGetByID)

}
