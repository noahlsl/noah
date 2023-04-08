package etcdx

import (
	"time"

	"go.etcd.io/etcd/client/v3"
)

func (c *Cfg) NewClient() *clientv3.Client {

	if c.Timeout == 0 {
		c.Timeout = 5
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Endpoints,
		Username:    c.Username,
		Password:    c.Password,
		DialTimeout: time.Duration(c.Timeout) * time.Second,
	})
	if err != nil {
		panic(err)
	}

	return cli
}
