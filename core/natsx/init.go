package natsx

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

var (
	conn *nats.Conn
)

func (c *Cfg) NewConn() *nats.Conn {

	var err error
	addr := fmt.Sprintf("nats://%s:%d", c.Host, c.Port)

	var ops []nats.Option
	ops = append(ops, nats.UserInfo(c.Username, c.Password))

	conn, err = nats.Connect(addr, ops...)
	if err != nil {
		panic(err)
	}

	return conn
}

func (c *Cfg) NewJetStream() nats.JetStreamContext {

	if conn == nil {
		conn = c.NewConn()
	}

	jetStream, err := conn.JetStream()
	if err != nil {
		panic(err)
	}

	return jetStream
}
