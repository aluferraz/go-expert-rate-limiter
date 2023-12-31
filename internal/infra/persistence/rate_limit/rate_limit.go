package rate_limit

type RateLimitRepository interface {
	SetRequestCounter(counterKey string, maxRequest int64) error
	GetLastRequestTime(resetTimeKey string) (int64, error)
	DecreaseTokenBucket(counterKey string) error
}
