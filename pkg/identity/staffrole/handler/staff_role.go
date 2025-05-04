 package handler

// import (
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
// 	"github.com/labstack/echo/v4"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// // StaffRoleAdminHandler ...
// type StaffRoleAdminHandler struct {
// 	StaffRoleAdminService model.StaffRoleAdminService
// }

// // Create ...
// func (d *StaffRoleAdminHandler) Create(c echo.Context) error {
// 	cc := util.EchoGetCustomCtx(c)
// 	var (
// 		body = c.Get("body").(model.CreateStaffRoleCommand)
// 	)

// 	data, err := d.StaffRoleAdminService.Create(cc.GetRequestCtx(), body)
// 	if err != nil {
// 		return cc.Response400(nil, err.Error())
// 	}
// 	return cc.Response200(echo.Map{
// 		"role": data,
// 	}, "")
// }

// // Update ...
// func (d *StaffRoleAdminHandler) Update(c echo.Context) error {
// 	cc := util.EchoGetCustomCtx(c)
// 	var (
// 		body = c.Get("body").(model.CreateStaffRoleCommand)
// 		role = c.Get("role").(model.StaffRoleRaw)
// 	)
// 	data, err := d.StaffRoleAdminService.Update(cc.GetRequestCtx(), role, body)
// 	if err != nil {
// 		return cc.Response400(nil, err.Error())
// 	}
// 	return cc.Response200(echo.Map{
// 		"role": data,
// 	}, "")
// }

// func (d *StaffRoleAdminHandler) ListRoleStaff(c echo.Context) error {
// 	var (
// 		cc  = util.EchoGetCustomCtx(c)
// 		ctx = cc.GetRequestCtx()
// 		q   = model.CommonQuery{
// 			Sort: bson.D{
// 				bson.E{Key: "createdAt", Value: -1},
// 			},
// 		}
// 	)

// 	roles, total := d.StaffRoleAdminService.ListStaffRole(ctx, q)

// 	return cc.Response200(echo.Map{
// 		"staffRoles": roles,
// 		"total":      total,
// 		"limit":      q.Limit,
// 	}, "")
// }

// // SearchPermission ...
// func (h *StaffRoleAdminHandler) SearchPermission(c echo.Context) error {
// 	cc := util.EchoGetCustomCtx(c)
// 	permissions := config.Permissions
// 	return cc.Response200(echo.Map{
// 		"permissions": permissions,
// 	}, "")
// }

// func (h *StaffRoleAdminHandler) Delete(c echo.Context) error {
// 	cc := util.EchoGetCustomCtx(c)
// 	var (
// 		role = c.Get("role").(model.StaffRoleRaw)
// 	)
// 	err := h.StaffRoleAdminService.Delete(cc.GetRequestCtx(), role)
// 	if err != nil {
// 		return cc.Response400(nil, err.Error())
// 	}
// 	return cc.Response200(nil, "")
// }

// func (h *StaffRoleAdminHandler) StaffRoleGetByID(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var (
// 			cc = util.EchoGetCustomCtx(c)
// 		)
// 		staffRoleIDString := c.Param("roleID")
// 		staffRoleID, err := primitive.ObjectIDFromHex(staffRoleIDString)
// 		if staffRoleID.IsZero() || err != nil {
// 			return cc.Response404(nil, locale.CommonKeyNotFound)
// 		}

// 		role, err := h.StaffRoleAdminService.FindByID(cc.GetRequestCtx(), staffRoleID)
// 		if role.ID.IsZero() || err != nil {
// 			return cc.Response404(nil, locale.CommonKeyNotFound)
// 		}
// 		c.Set("role", role)
// 		return next(c)
// 	}
// }
