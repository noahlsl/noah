package etcdx

type Cfg struct {
	Endpoints []string `json:"endpoints"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Timeout   int      `json:"timeout"`
}
