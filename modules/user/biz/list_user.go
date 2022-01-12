package biz

import (
	"context"
	"simple-rest-curd/common"
	"simple-rest-curd/modules/user/model"
)

type ListUserStore interface {
	ListDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.User, error)
}

type listUserBiz struct {
	store ListUserStore
}

func NewListUserBiz(store ListUserStore) *listUserBiz {
	return &listUserBiz{store: store}
}

func (biz *listUserBiz) ListUser(ctx context.Context, paging *common.Paging, filter *model.Filter) ([]model.User, error) {
	results, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)
	return results, err
}
