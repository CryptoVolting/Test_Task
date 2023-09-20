package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"testProject/internal/entity"
)

const adminPermissons = " admin_permissons"

type RedisFromDB struct {
	db *sqlx.DB
}

func NewRedisFromDB(db *sqlx.DB) *RedisFromDB {
	return &RedisFromDB{db: db}
}

func (c *RedisFromDB) FetchPermissionsFromDB(isAdmin string) ([]entity.Premissoins, error) {
	var permissions []entity.Premissoins

	if isAdmin == "admin" {
		query := fmt.Sprintf("SELECT * FROM%s;", adminPermissons)
		if err := c.db.Select(&permissions, query); err != nil {
			return nil, err
		}
	}

	return permissions, nil
}
