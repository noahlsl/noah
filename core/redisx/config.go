package redisx

type Cfg struct {
	Address    []string // 集群连接地址
	Username   string   // 账号
	Password   string   // 密码
	DB         int      // 数据库建议使用默认0
	PoolSize   int      // 连接池大小，连接池中的连接的最大数量
	ClientName string   // 客户端名称
}
