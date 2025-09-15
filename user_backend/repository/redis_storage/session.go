package redis_storage

import (
	"context"
	"time"
	"user_backend/domain"

	"github.com/redis/go-redis/v9"
)

type Session struct {
	client *redis.Client
	ttl    time.Duration
}

func NewSession(connStr string, ttl time.Duration) *Session {
	rdb := redis.NewClient(&redis.Options{
		Addr:     connStr,
		Password: "",
		DB:       0,
	})
	return &Session{
		client: rdb,
		ttl:    ttl,
	}
}

func (rds *Session) GetUserIDBySessionID(sessionID string) (string, error) {
	userID, err := rds.client.Get(context.Background(),
		"session:"+sessionID).Result()
	if err == redis.Nil {
		return "", domain.ErrNotFound("user not found")
	}
	if err != nil {
		return "", err
	}
	return userID, nil
}

func (rds *Session) SetUserIDBySessionID(sessionID string, userID string) error {
	return rds.client.Set(
		context.Background(),
		"session:"+sessionID,
		userID,
		rds.ttl,
	).Err()
}
