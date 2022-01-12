package biz

import (
	"context"
	"simple-rest-curd/modules/user/model"
)

type UpdateUserStore interface {
	UpdateData(ctx context.Context, id int, data *model.UserUpdate) error
}

type updateUserBiz struct {
	store UpdateUserStore
}

func NewUpdateUserBiz(store UpdateUserStore) *updateUserBiz {
	return &updateUserBiz{
		store: store,
	}
}

func (biz *updateUserBiz) UpdateUser(ctx context.Context, id int, data *model.UserUpdate) error {
	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}
	return nil
}
