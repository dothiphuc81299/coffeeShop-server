package category

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrCategoryExisted  = errors.New("Category existed")
	ErrCategoryNotFound = errors.New("Category not found")
)

type CategoryRaw struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Name         string             `bson:"name" json:"name"`
	SearchString string             `bson:"searchString" json:"search_string,omitempty"`
	CreatedAt    time.Time          `bson:"createdAt" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updated_at"`
}

type CategoryBody struct {
	Name string `json:"name"`
}

func (c CategoryBody) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
	)
}

type ResponseAdminListData struct {
	Data         interface{} `json:"data"`
	Total        int64       `json:"total"`
	LimitPerPage int64       `json:"limitPerPage"`
}