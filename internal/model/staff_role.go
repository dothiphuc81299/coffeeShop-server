package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// StaffRoleDAO represent staff role data access object
type StaffRoleDAO interface {
	FindOneByCondition(ctx context.Context, cond interface{}) (StaffRoleRaw, error)
	FindByID(ctx context.Context, id AppID) (StaffRoleRaw, error)
	FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) ([]StaffRoleRaw, error)
	CountByCondition(ctx context.Context, cond interface{}) int64
	InsertOne(ctx context.Context, u *StaffRoleRaw) error
	UpdateByID(ctx context.Context, id AppID, payload interface{}) error
	DeleteByID(ctx context.Context, id AppID) error
}

// StaffRoleAdminService represent staff roles service
type StaffRoleAdminService interface {
	ListStaffRole(ctx context.Context, q CommonQuery) ([]StaffRoleAdminResponse, int64)
	Create(ctx context.Context, body CreateStaffRoleCommand) (StaffRoleAdminResponse, error)
	Update(ctx context.Context, data StaffRoleRaw, body CreateStaffRoleCommand) (StaffRoleAdminResponse, error)
	FindByID(ctx context.Context, id AppID) (StaffRoleRaw, error)
	Delete(ctx context.Context, data StaffRoleRaw) error
}

// StaffRoleRaw ...
type StaffRoleRaw struct {
	ID          AppID     `bson:"_id"`
	Name        string    `bson:"name"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
	Permissions []string  `bson:"permissions"`
}

// GetResponse ...
func (sr *StaffRoleRaw) GetResponse() StaffRoleAdminResponse {
	return StaffRoleAdminResponse{
		ID:          sr.ID,
		Name:        sr.Name,
		CreatedAt:   TimeResponse{sr.CreatedAt},
		Permissions: sr.Permissions,
	}
}
