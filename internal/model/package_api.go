package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Level string
type PackageBody struct {
	Title      string  `json:"title"`
	Level      string  `json:"level"` // easy , medium, hard
	Reward     float64 `json:"reward"`
	MinusPoint float64 `json:"minusPoint"`
}

type PackageAdminResponse struct {
	ID         AppID        `json:"_id"`
	Title      string       `json:"title"`
	Level      string       `json:"level"` // easy , medium, hard
	Reward     float64      `json:"reward"`
	MinusPoint float64      `json:"minusPoint"`
	CreatedAt  TimeResponse `json:"createdAt"`
}

func (p PackageBody) Validate() error {
	err := validation.ValidateStruct(&p,
		validation.Field(&p.Title, validation.Required.Error("GroupNameIsRequired")),
		validation.Field(&p.Level, validation.Required.Error("QuizIDIsRequired")),
		validation.Field(&p.Reward, validation.Required.Error("QuizIDIsRequired")),
		validation.Field(&p.MinusPoint, validation.Required.Error("QuizIDIsRequired")),
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *PackageBody) PackageNewBSON() PackageRaw {
	return PackageRaw{
		ID:         primitive.NewObjectID(),
		Title:      p.Title,
		Reward:     p.Reward,
		MinusPoint: p.MinusPoint,
		Level:      Level(p.Level),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func PackageAdminGetResponse(raw PackageRaw) PackageAdminResponse {
	return PackageAdminResponse{
		ID:         raw.ID,
		Reward:     raw.Reward,
		Title:      raw.Title,
		MinusPoint: raw.MinusPoint,
		Level:      string(raw.Level),
		CreatedAt:  TimeResponse{Time: raw.CreatedAt},
	}
}
