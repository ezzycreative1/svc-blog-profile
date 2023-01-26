package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	expiredMinnute = 10
	timeOut        = time.Second * 6
)

type queueUserRepo struct {
	redis *redis.Client
}

func NewQueueUserRepo(redis *redis.Client) *queueUserRepo {
	return &queueUserRepo{
		redis: redis,
	}
}

func (qr *queueUserRepo) AddUserToQueue(ctx context.Context, userID int64, token, transactionDate string) error {
	key := fmt.Sprintf("%d-%s", userID, token)

	ctxWT, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	return qr.redis.Set(ctxWT, key, transactionDate, time.Duration(time.Minute*expiredMinnute)).Err()
}

// CheckQueue return true if value exist in cache
func (qr *queueUserRepo) CheckUserExist(ctx context.Context, userID int64, token string) (bool, error) {
	key := fmt.Sprintf("%d-%s", userID, token)
	_, err := qr.redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (qr *queueUserRepo) GetUserValue(ctx context.Context, userID int64, token string) (string, error) {
	key := fmt.Sprintf("%d-%s", userID, token)
	return qr.redis.Get(ctx, key).Result()
}

func (qr *queueUserRepo) RemoveUserFromQueue(ctx context.Context, userID int64, token string) error {
	key := fmt.Sprintf("%d-%s", userID, token)

	ctxWT, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	return qr.redis.Del(ctxWT, key).Err()
}
