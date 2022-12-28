package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	maxRetries      = 5
	minRetryBackoff = 300 * time.Millisecond
	maxRetryBackoff = 500 * time.Millisecond
	dialTimeout     = 5 * time.Second
	readTimeout     = 5 * time.Second
	writeTimeout    = 3 * time.Second
	minIdleConn     = 20
	poolTimeout     = 6 * time.Second
	idleTimeout     = 12 * time.Second
	poolSize        = 300
)

// NewUniversalClient creates a new redis universal client.
func NewUniversalClient(url string, port int, db int, password string) redis.UniversalClient {
	addr := fmt.Sprintf("%s:%d", url, port)
	universalClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:           []string{addr},
		Password:        password,
		DB:              db,
		MaxRetries:      maxRetries,
		MinRetryBackoff: minRetryBackoff,
		MaxRetryBackoff: maxRetryBackoff,
		DialTimeout:     dialTimeout,
		ReadTimeout:     readTimeout,
		WriteTimeout:    writeTimeout,
		PoolSize:        poolSize,
		MinIdleConns:    minIdleConn,
		PoolTimeout:     poolTimeout,
		IdleTimeout:     idleTimeout,
	})

	return universalClient
}
