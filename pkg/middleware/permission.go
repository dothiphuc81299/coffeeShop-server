package middleware

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/labstack/echo/v4"
)

func CheckPermissionRoot(accountType token.AccountType) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := util.EchoGetCustomCtx(c)

			account, exists := c.Get("current_account").(*token.AccountData)
			if !exists {
				return cc.Response401(nil, "Account not found")
			}

			if account.AccountType != token.Route {
				return cc.Response401(nil, locale.CommonNoPermission)
			}

			return next(c)
		}
	}
}

func CheckPermission(model string, fieldPermission string, accountType token.AccountType) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := util.EchoGetCustomCtx(c)

			account, exists := c.Get("current_account").(*token.AccountData)
			if !exists {
				return cc.Response401(nil, "Account not found")
			}

			if !checkUserPermission(account, model, fieldPermission) {
				return cc.Response401(nil, locale.CommonNoPermission)
			}

			return next(c)
		}
	}
}

func checkUserPermission(account *token.AccountData, model string, fieldPermission string) bool {
	if account.AccountType == token.Route {
		return true
	}

	if account.LoginName == model+"_"+fieldPermission {
		return true
	}

	return false
}

func CheckStaff(accountType token.AccountType) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := util.EchoGetCustomCtx(c)

			account, exists := c.Get("current_account").(*token.AccountData)
			if !exists {
				return cc.Response401(nil, "Account not found")
			}

			if account.AccountType != token.Staff {
				return cc.Response401(nil, locale.CommonNoPermission)
			}

			return next(c)
		}
	}
}

func CheckUser(accountType token.AccountType) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := util.EchoGetCustomCtx(c)

			account, exists := c.Get("current_account").(*token.AccountData)
			if !exists {
				return cc.Response401(nil, "Account not found")
			}

			if account.AccountType != token.User {
				return cc.Response401(nil, "Tai khoan khong hop le")
			}

			// // Kiểm tra xem tài khoản có đang hoạt động hay không
			// if !account.Active {
			// 	return cc.Response401(nil, locale.CommonKeyStaffDeactive)
			// }

			//c.Set("user", account)

			return next(c)
		}
	}
}
