package dbx

type Cfg struct {
	Host     string `json:"host"`     // IP
	Port     int    `json:"port"`     // 端口
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Db       string `json:"db"`       // 数据库
	MaxOpen  int    `json:"max_open"` // 最大打开连接数
	MaxIdle  int    `json:"max_idle"` // 空闲状态下的最大连接数
}
