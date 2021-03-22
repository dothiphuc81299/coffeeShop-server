package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type RoleDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (RoleRaw, error)
	InsertOne(ctx context.Context, u RoleRaw) error
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]RoleRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
}

// RoleService ...
type RoleService interface {
	Create(ctx context.Context, body RoleBody) (RoleAdminResponse, error)
	List(ctx context.Context, q CommonQuery) ([]RoleAdminResponse, int64)
	FindByID(ctx context.Context, id AppID) (RoleRaw, error)
	Update(ctx context.Context, body RoleBody, raw RoleRaw) (RoleAdminResponse, error)
	GetDetail(ctx context.Context, role RoleRaw) RoleAdminResponse
}

// RoleRaw ...
type RoleRaw struct {
	ID          AppID     `bson:"_id"`
	Name        string    `bson:"name"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
	Permissions []string  `bson:"permissions"`
}

// GetResponse ...
func (r *RoleRaw) GetResponse() RoleAdminResponse {
	return RoleAdminResponse{
		ID:          r.ID,
		Name:        r.Name,
		CreatedAt:   r.CreatedAt,
		Permissions: r.Permissions,
	}
}
