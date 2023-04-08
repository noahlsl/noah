package pulsarx

import (
	"fmt"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func (c *Cfg) NewClient() pulsar.Client {

	cli, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               fmt.Sprintf("pulsar://%s:%d", c.Host, c.Port),
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	return cli
}
