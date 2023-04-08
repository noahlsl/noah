package miniox

type Cfg struct {
	Address     string // 地址-必填(例 127.0.0.1:9000)
	Username    string // 账号-必填
	Password    string // 密码-必填
	Bucket      string // 桶名-必填
	Token       string // 默认无
	Location    string // 地点/服务
	Timeout     int    // 超时
	ContentType string // 文件类型
	TLS         bool   // 是否开启TLS(HTTPS)
}
