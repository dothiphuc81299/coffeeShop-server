package model

import (
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserPackageBody struct {
	PackageID string `json:"packageId"`
}

func (u UserPackageBody) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.PackageID, validation.Required.Error("PackageId is required"),
			is.MongoID.Error(locale.DrinkKeyCategoryInvalid)),
	)
}

func (u *UserPackageBody) UserPackageGroupNewBSON(userID AppID, group QuizGroupRaw) UserPackageGroupRaw {
	packageID, _ := primitive.ObjectIDFromHex(u.PackageID)
	return UserPackageGroupRaw{
		ID:             primitive.NewObjectID(),
		GroupID:        group.ID,
		PackageID:      packageID,
		TotalPoint:     0,
		TotalQuiz:      group.TotalQuestion,
		SubmissionTime: 5 * time.Minute,
		UserID:         userID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
