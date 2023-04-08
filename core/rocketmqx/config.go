package rocketmqx

// Cfg RocketMQ5.0配置,请使用最新的5.0
type Cfg struct {
	Endpoint      string `json:"endpoint"`       // 连接地址
	AccessKey     string `json:"accessKey"`      // 账号
	SecretKey     string `json:"secretKey"`      // 密码
	ProducerGroup string `json:"producer_group"` // 生产者群组
	ConsumerGroup string `json:"consumer_group"` // 消费者群组
	Level         string `json:"level"`          // 日志等级
	AwaitDuration int    `json:"await_duration"` // 接受消息最大等待时间 默认5秒
}
