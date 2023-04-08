package redisx

import (
	"context"
	"runtime"
	"time"

	"github.com/redis/go-redis/v9"
)

func (c *Cfg) NewClusterClient() *redis.ClusterClient {

	if c.PoolSize == 0 {
		c.PoolSize = 4 * runtime.NumCPU()
	}
	cli := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:          c.Address,
		Username:       c.Username,
		Password:       c.Password,
		RouteByLatency: true,
		PoolSize:       c.PoolSize,
		ClientName:     c.ClientName,
	})
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := cli.Ping(ct).Err()
	if err != nil {
		panic(err)
	}

	return cli
}

func (c *Cfg) NewClient() *redis.Client {

	if c.PoolSize == 0 {
		c.PoolSize = 4 * runtime.NumCPU()
	}
	cli := redis.NewClient(&redis.Options{
		Addr:       c.Address[0],
		Username:   c.Username,
		Password:   c.Password,
		PoolSize:   c.PoolSize,
		ClientName: c.ClientName,
	})
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := cli.Ping(ct).Err()
	if err != nil {
		panic(err)
	}

	return cli
}
