package handler

import (
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AccountAdminHandler ...
type AccountAdminHandler struct {
	AccountAdminService model.AccountAdminService
}

// Update ...
func (h *AccountAdminHandler) Update(c echo.Context) error {
	var (
		cc      = util.EchoGetCustomCtx(c)
		body    = c.Get("body").(model.AccountBody)
		account = c.Get("account").(model.AccountRaw)
	)
	data, err := h.AccountAdminService.Update(cc.GetRequestCtx(), body, account)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"data": data,
	}, "")	
}

// ChangeStatus ...
func (h *AccountAdminHandler) ChangeStatus(c echo.Context) error {
	var (
		cc      = util.EchoGetCustomCtx(c)
		account = c.Get("account").(model.AccountRaw)
	)
	active, err := h.AccountAdminService.ChangeStatus(cc.GetRequestCtx(), account)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"active": active,
	}, "")
}

// GenerateToken ...
func (h *AccountAdminHandler) GenerateToken(c echo.Context) error {
	var (
		cc      = util.EchoGetCustomCtx(c)
		account = c.Get("account").(model.AccountRaw)
		userID  = cc.GetCurrentUserID()
	)
	token, err := h.AccountAdminService.GenerateToken(cc.GetRequestCtx(), account, userID)
	if err != nil {
		return cc.Response400(nil, err.Error())
	}
	return cc.Response200(echo.Map{
		"token": token,
	}, "")
}

// GetByID ...
func (h *AccountAdminHandler) GetByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			cc = util.EchoGetCustomCtx(c)
		)
		accountIDString := c.Param("accountID")
		accountID, err := primitive.ObjectIDFromHex(accountIDString)
		if accountID.IsZero() || err != nil {
			return cc.Response404(nil, locale.CommonKeyNotFound)
		}

		account, err := h.AccountAdminService.FindByID(cc.GetRequestCtx(), accountID)
		if account.ID.IsZero() || err != nil {
			return cc.Response404(nil, locale.CommonKeyNotFound)
		}
		c.Set("account", account)
		return next(c)
	}
}
