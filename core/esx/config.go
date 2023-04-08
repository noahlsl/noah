package esx

type Cfg struct {
	Address  []string // 集群连接地址([]string{127.0.0.1:2379})
	Username string   // 账号
	Password string   // 密码
	TLS      int      // 是否开启TLS(HTTPS)
}
