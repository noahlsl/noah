package prometheusx

//Name: order-rpc
//Endpoint:  http://jaeger:14268/api/traces
//Sampler: 1.0
//Batcher: jaeger

type Cfg struct {
	Name     string
	Endpoint string
	Sampler  string
	Batcher  string
}
