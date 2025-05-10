package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(accountType token.AccountType, requiredPermissions string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing token"})
			}

			authToken := strings.Split(authHeader, " ")
			if len(authToken) != 2 || authToken[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token format"})
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := token.ValidateJWT(tokenStr, accountType)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
			}

			expirationTime, err := claims.GetExpirationTime()
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token expiration"})
			}

			if float64(time.Now().UTC().Unix()) > float64(expirationTime.Unix()) {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Token expired"})
			}

			if claims.AccountType == token.Staff && requiredPermissions != "" {
				if !hasPermission(claims.Permissions, requiredPermissions) {
					return c.JSON(http.StatusForbidden, echo.Map{"error": "Insufficient permissions"})
				}
			}

			account := &token.AccountData{
				LoginName:   claims.LoginName,
				ID:          claims.UserID,
				AccountType: accountType,
			}

			ctx := context.WithValue(c.Request().Context(), "current_account", account)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

func hasPermission(userPermissions []string, requiredPermission string) bool {
	for _, p := range userPermissions {
		if p == requiredPermission {
			return true
		}
	}
	return false
}
