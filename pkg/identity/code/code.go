package code

import "context"

type CodedRegisterDAO interface {
	InsertOne(ctx context.Context, u CodedRegisterRaw) error
	DeleteOne(ctx context.Context, u string) error
	FindOneByCondition(ctx context.Context, cond interface{}) (CodedRegisterRaw, error)
}
