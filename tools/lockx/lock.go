package lockx

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	key        = "public:lock:%v"
	expiration = 10 * time.Second
)

func SetLock(ctx context.Context, r *redis.ClusterClient, val interface{}) bool {

	k := fmt.Sprintf(key, val)
	err := r.SetNX(ctx, key, 0, expiration).Err()
	if err != nil {
		return false
	}

	go func(k string) {
		defer r.Del(context.Background(), k).Err()
		for {
			select {
			case <-ctx.Done():
				return

			case <-time.After(expiration):
				return
			}
		}
	}(k)

	return true
}
