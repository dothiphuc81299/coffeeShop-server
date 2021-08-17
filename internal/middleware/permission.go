package middleware

import (
	"strings"

	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
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
			token := util.Base64EncodeToString(strings.Split(cc.GetHeaderKey(config.HeaderAuthorization), " ")[1])
			sessionTotal := d.Session.CountByCondition(ctx, bson.M{"staff": userID, "token": token})
			if sessionTotal <= 0 {
				return cc.Response401(nil, "token het han")
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

func CheckPermission(model string, fieldPermission string, d *model.CommonDAO) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := util.EchoGetCustomCtx(c)
			ctx := cc.GetRequestCtx()
			userID := cc.GetCurrentUserID()

			if userID.IsZero() {
				return cc.Response401(nil, "")
			}
			// Check session
			token := util.Base64EncodeToString(strings.Split(cc.GetHeaderKey(config.HeaderAuthorization), " ")[1])
			sessionTotal := d.Session.CountByCondition(ctx, bson.M{"staff": userID, "token": token})
			if sessionTotal <= 0 {
				return cc.Response401(nil, "token het han")
			}

			cond := bson.M{
				"_id": userID,
				"$or": []bson.M{
					bson.M{
						"permissions": model + "_" + config.PermissionAdmin,
					},
					bson.M{
						"permissions": model + "_" + fieldPermission,
					},
					bson.M{
						"isRoot": true,
					},
				},
			}

			staff, err := d.Staff.FindOneByCondition(ctx, cond)
			if err != nil || staff.ID.IsZero() {
				return cc.Response401(nil, locale.CommonNoPermission)
			}

			if !staff.Active {
				return cc.Response401(nil, locale.CommonKeyStaffDeactive)
			}

			c.Set("staff", staff)
			return next(c)
		}
	}
}

func CheckUser(d *model.CommonDAO) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := util.EchoGetCustomCtx(c)
			ctx := cc.GetRequestCtx()
			userID := cc.GetCurrentUserID()
			if userID.IsZero() {
				return cc.Response401(nil, "")
			}

			cond := bson.M{
				"_id":    userID,
				"active": true,
			}
			user, err := d.User.FindOneByCondition(ctx, cond)

			if err != nil || user.ID.IsZero() {
				return cc.Response401(nil, "tai khoan khong hop le")
			}

			if !user.Active {
				return cc.Response401(nil, locale.CommonKeyStaffDeactive)
			}

			c.Set("user", user)
			return next(c)
		}
	}
}
