package rate_limit

import (
	"context"
	"errors"
	"github.com/aluferraz/go-expert-rate-limiter/internal/entity/web_session"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type RateLimitRepositoryRedis struct {
	Client *redis.Client // Unique to redis, can't use sql.DB abstraction here.
}

func NewRateLimitRepositoryRedis(client *redis.Client) *RateLimitRepositoryRedis {
	return &RateLimitRepositoryRedis{Client: client}
}

func (rlim *RateLimitRepositoryRedis) SetRequestCounter(session *web_session.WebSession) error {
	ctx := context.Background()
	counterKey, maxRequest := session.GetSessionId(), session.GetRequestsLimit()
	return rlim.Client.Set(ctx, counterKey, maxRequest, 0).Err()
}

func (rlim *RateLimitRepositoryRedis) GetLastRequestTime(resetTimeKey string) (int64, error) {
	ctx := context.Background()

	lastResetTimeStr, err := rlim.Client.Get(ctx, resetTimeKey).Result()
	if err != redis.Nil {
		return -1, err
	}
	lastResetTime, err := strconv.ParseInt(lastResetTimeStr, 10, 64)

	return lastResetTime, err
}

func (rlim *RateLimitRepositoryRedis) DecreaseTokenBucket(counterKey string) error {
	ctx := context.Background()
	// Transactional function, optimistic lock.
	txf := func(tx *redis.Tx) error {
		// Get the current value or zero.
		remaingRequests, err := tx.Get(ctx, counterKey).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		if remaingRequests <= 0 {
			return errors.New("you have reached the maximum number of requests or actions allowed within a certain time frame")
		}

		// Actual operation (local in optimistic lock).
		remaingRequests--
		// Operation is commited only if the watched keys remain unchanged.
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, counterKey, remaingRequests, 0)
			return nil
		})
		return err
	}
	return rlim.Client.Watch(ctx, txf, counterKey) // Will return error if not possible.
}
