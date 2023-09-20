package repository

import (
	"github.com/jmoiron/sqlx"
	"testProject/internal/entity"
)

type RedisRepository struct {
	RedisUsage
}

func NewRedisRepository(db *sqlx.DB) *RedisRepository {
	return &RedisRepository{
		RedisUsage: NewRedisFromDB(db),
	}
}

type RedisUsage interface {
	FetchPermissionsFromDB(isAdmin string) ([]entity.Premissoins, error)
}
