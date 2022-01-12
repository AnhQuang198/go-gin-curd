package storage

import (
	"context"
	"simple-rest-curd/modules/user/model"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *model.UserUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
