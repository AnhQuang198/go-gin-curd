package storage

import (
	"context"
	"simple-rest-curd/common"
	"simple-rest-curd/modules/user/model"
)

func (s *sqlStore) ListDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.User, error) {
	var results []model.User

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(model.User{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if len(v.Username) > 0 {
			db = db.Where("username = ?", v.Username)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit). //so dong se bo qua khi lay record
		Limit(paging.Limit).                      //gioi han so dong lay ra
		Order("id desc").
		Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}
