package model

// import (
// 	"time"

// 	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
// 	"github.com/go-ozzo/ozzo-validation/is"
// 	validation "github.com/go-ozzo/ozzo-validation/v4"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type PackageGroupBody struct {
// 	PackageID string `json:"packageId"`
// 	GroupID   string `json:"groupId"`
// }
// type PackageGroupAdminResponse struct {
// 	ID        primitive.ObjectID           `json:"_id"`
// 	PackageID string          `json:"packageId"`
// 	Group     QuizGroupCommon `json:"group"`
// 	CreatedAt TimeResponse    `json:"createdAt"`
// }

// func (p PackageGroupBody) Validate() error {
// 	return validation.ValidateStruct(&p,
// 		validation.Field(&p.PackageID, validation.Required.Error("packageID is required"),
// 			is.MongoID.Error(locale.DrinkKeyCategoryInvalid)),

// 		validation.Field(&p.GroupID, validation.Required.Error(locale.DrinkKeyCategoryIDIsRequired),
// 			is.MongoID.Error(locale.DrinkKeyCategoryInvalid)),
// 	)
// }

// func (p *PackageGroupRaw) GetPackageGroupAdminResponse(group QuizGroupCommon) PackageGroupAdminResponse {
// 	return PackageGroupAdminResponse{
// 		ID:        p.ID,
// 		PackageID: p.PackageID.Hex(),
// 		Group:     group,
// 		CreatedAt: TimeResponse{Time: p.CreatedAt},
// 	}
// }

// func (p *PackageGroupBody) PackageGroupNewBson() PackageGroupRaw {
// 	packageID, _ := primitive.ObjectIDFromHex(p.PackageID)
// 	groupID, _ := primitive.ObjectIDFromHex(p.GroupID)

// 	return PackageGroupRaw{
// 		ID:        primitive.NewObjectID(),
// 		PackageID: packageID,
// 		GroupID:   groupID,
// 		CreatedAt: time.Now(),
// 	}
// }
