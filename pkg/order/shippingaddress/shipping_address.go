package shippingaddress

import "context"

type Store interface {
	InsertOne(ctx context.Context, u UserShippingAddressRaw) error
	FindOneByCondition(ctx context.Context, cond interface{}) (u UserShippingAddressRaw, err error)
}
