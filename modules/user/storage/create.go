package storage

import (
	"context"
	"simple-rest-curd/modules/user/model"
)

func (s *sqlStore) Create(ctx context.Context, data *model.UserCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
