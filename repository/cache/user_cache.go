package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ryanbaskara/learning-go/entity"
)

const (
	TTL = 1 * time.Minute
)

type UserCacheRepo struct {
	redisClient *redis.Client
}

func NewUserCacheRepo(rdb *redis.Client) *UserCacheRepo {
	return &UserCacheRepo{redisClient: rdb}
}

func (r *UserCacheRepo) SetUser(ctx context.Context, user *entity.User) error {
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}
	r.redisClient.Set(ctx, generateUserKey(user.ID), b, TTL)

	return nil
}

func (r *UserCacheRepo) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	val, err := r.redisClient.Get(ctx, generateUserKey(id)).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var user entity.User
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func generateUserKey(id int64) string {
	return fmt.Sprintf("user:%d", id)
}
