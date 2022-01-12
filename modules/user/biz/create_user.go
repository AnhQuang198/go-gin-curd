package biz

import (
	"context"
	"simple-rest-curd/modules/user/model"
)

type CreateUserStore interface {
	Create(ctx context.Context, data *model.UserCreate) error
}

type createUserBiz struct {
	store CreateUserStore
}

func NewCreateUserBiz(store CreateUserStore) *createUserBiz {
	return &createUserBiz{store: store}
}

func (biz *createUserBiz) CreateUser(ctx context.Context, data *model.UserCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}
	//xu ly logic
	err := biz.store.Create(ctx, data)
	return err
}
