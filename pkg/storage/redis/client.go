package redis

import (
	"strings"
	"time"

	rediscleint "github.com/go-redis/redis/v8"
)

const (
	ClusterMode = "cluster"
)

type PoolOptions struct {
	URL         string
	Password    string
	MinIdle     int
	IdleTimeout int
	PoolSize    int
	Mode        string
}

func NewRedisPool(cfg PoolOptions) rediscleint.UniversalClient {
	redisUrls := strings.Split(cfg.URL, ",")
	if cfg.Mode == ClusterMode {
		if len(redisUrls) == 1 {
			return rediscleint.NewClient(&rediscleint.Options{
				Addr:         cfg.URL,
				Password:     cfg.Password,
				PoolSize:     cfg.PoolSize,
				MinIdleConns: cfg.MinIdle,
				IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
			})
		} else {
			return rediscleint.NewClusterClient(&rediscleint.ClusterOptions{
				Addrs:        redisUrls,
				Password:     cfg.Password,
				PoolSize:     cfg.PoolSize,
				MinIdleConns: cfg.MinIdle,
				IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
			})
		}
	} else {
		return rediscleint.NewFailoverClient(&rediscleint.FailoverOptions{
			SentinelAddrs: redisUrls,
			Password:      cfg.Password,
			PoolSize:      cfg.PoolSize,
			MinIdleConns:  cfg.MinIdle,
			IdleTimeout:   time.Duration(cfg.IdleTimeout) * time.Second,
		})
	}
}
