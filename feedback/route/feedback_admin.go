package route

import (
	"github.com/dothiphuc81299/coffeeShop-server/feedback/handler"
	"github.com/dothiphuc81299/coffeeShop-server/internal/middleware"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/labstack/echo/v4"
)

// InitFeedbackAdmin ...
func InitFeedbackAdmin(e *echo.Echo, cs *model.AdminService, d *model.CommonDAO) {
	h := &handler.FeedbackAdminHandler{
		FeedbackAdminService: cs.Feedback,
	}

	g := e.Group("/feedback")

	g.PATCH("/:feedbackID/status", h.ChangeStatus, middleware.CheckPermissionRoot(d), h.FeedbackGetByID)

}
