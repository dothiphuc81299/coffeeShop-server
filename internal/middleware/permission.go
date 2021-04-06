package middleware

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckPermissionRoot ...
func CheckPermissionRoot(d *model.CommonDAO) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := util.EchoGetCustomCtx(c)
			ctx := cc.GetRequestCtx()
			userID := cc.GetCurrentUserID()
			if userID.IsZero() {
				return cc.Response401(nil, "")
			}

			// check session
			cond := bson.M{
				"_id":    userID,
				"isRoot": true,
			}
			staff, err := d.Staff.FindOneByCondition(ctx, cond)
			if staff.ID.IsZero() || err != nil {
				return cc.Response401(nil, locale.CommonNoPermission)
			}

			if !staff.Active {
				return cc.Response401(nil, locale.CommonKeyStaffDeactive)
			}
			return next(c)
		}
	}
}
