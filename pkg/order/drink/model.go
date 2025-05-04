package drink

import (
	"errors"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/category"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/format"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrDrinkNameExisted = errors.New("Drink name existed")
)

type DrinkRaw struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Name         string             `bson:"name" json:"name"`
	Category     primitive.ObjectID `bson:"category" json:"category"`
	Price        float64            `bson:"price" json:"price"`
	SearchString string             `bson:"searchString" json:"search_string"`
	CreatedAt    time.Time          `bson:"createdAt" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updated_at"`
	Active       bool               `bson:"active" json:"active"`
	Image        string             `bson:"image" json:"image"`
}

type DrinkBody struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Image    string  `json:"image"`
}

type DrinkAdminResponse struct {
	ID       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name"`
	Category CategoryInfo       `json:"category"`
	Price    float64            `json:"price"`
	Image    string             `json:"image"`
	Active   bool               `json:"active"`
}

type CategoryInfo struct {
	ID   primitive.ObjectID `json:"_id"`
	Name string             `json:"name"`
}

func (d DrinkBody) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name, validation.Required),
		validation.Field(&d.Category, validation.Required,
			is.MongoID.Error("MongoID invalid")),
		validation.Field(&d.Price, validation.Required),
	)
}

func (d DrinkBody) NewDrinkRaw() DrinkRaw {
	categoryID, _ := primitive.ObjectIDFromHex(d.Category)
	now := time.Now()
	return DrinkRaw{
		ID:           primitive.NewObjectID(),
		Name:         d.Name,
		SearchString: format.NonAccentVietnamese(d.Name),
		Price:        d.Price,
		Category:     categoryID,
		CreatedAt:    now,
		UpdatedAt:    now,
		Active:       true,
		Image:        d.Image,
	}
}

func (b DrinkRaw) DrinkGetAdminResponse(c CategoryInfo) DrinkAdminResponse {
	return DrinkAdminResponse{
		ID:       b.ID,
		Name:     b.Name,
		Category: c,
		Price:    b.Price,
		Image:    b.Image,
		Active:   b.Active,
	}
}

func CategoryGetInfo(r category.CategoryRaw) CategoryInfo {
	return CategoryInfo{
		ID:   r.ID,
		Name: r.Name,
	}
}
