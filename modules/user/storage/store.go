package storage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

//config storage
func NewSqlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

func RedisStore() {

}
