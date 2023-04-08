package rocketmqx

import (
	"os"
	"time"

	client "github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
)

func (c *Cfg) SetConsumerGroup(s string) *Cfg {
	c.ConsumerGroup = s
	return c
}

func (c *Cfg) SetProducerGroup(s string) *Cfg {
	c.ProducerGroup = s
	return c
}

func init() {
	// 设置日志环境变量
	_ = os.Setenv(client.ENABLE_CONSOLE_APPENDER, "true")
	_ = os.Setenv(client.CLIENT_LOG_LEVEL, "error")
}

func (c *Cfg) NewConsumer() client.SimpleConsumer {

	client.ResetLogger()
	if c.AwaitDuration == 0 {
		c.AwaitDuration = 5
	}
	// new simpleConsumer instance
	consumer, err := client.NewSimpleConsumer(&client.Config{
		Endpoint:      c.Endpoint,
		ConsumerGroup: c.ConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    c.AccessKey,
			AccessSecret: c.SecretKey,
		},
	},
		client.WithAwaitDuration(time.Duration(c.AwaitDuration)*time.Second))
	if err != nil {
		panic(err)
	}

	return consumer
}

func (c *Cfg) NewProducer() client.Producer {

	producer, err := client.NewProducer(&client.Config{
		Endpoint: c.Endpoint,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    c.AccessKey,
			AccessSecret: c.SecretKey,
		},
	})
	if err != nil {
		panic(err)
	}

	// start producer
	err = producer.Start()
	if err != nil {
		panic(err)
	}

	return producer
}
