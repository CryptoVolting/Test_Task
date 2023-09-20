package usecase

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"testProject/internal/entity"
	"testProject/internal/usecase/repository"
	"time"
)

var ctx = context.Background()

type RedisUsecase struct {
	repositoryCache repository.RedisUsage
	client          *redis.Client
}

func NewRedisUsecase(repositoryCache repository.RedisUsage, client *redis.Client) *RedisUsecase {
	return &RedisUsecase{
		repositoryCache: repositoryCache,
		client:          client,
	}
}

func (c *RedisUsecase) GetPermissionsByRole(isAdmin string) ([]entity.Premissoins, error) {
	cached, err := c.client.Get(ctx, isAdmin).Result()
	if err == redis.Nil {
		permissions, err := c.repositoryCache.FetchPermissionsFromDB(isAdmin)
		if err != nil {
			return nil, err
		}

		serialized, errr := json.Marshal(permissions)
		if errr != nil {
			return nil, errr
		}
		c.client.Set(ctx, isAdmin, serialized, 1*time.Minute)
		return permissions, nil

	} else if err != nil {
		return nil, err
	}

	var permissions []entity.Premissoins
	if err := json.Unmarshal([]byte(cached), &permissions); err != nil {
		return nil, err
	}
	return permissions, nil
}
