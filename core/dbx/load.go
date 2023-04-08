package dbx

import (
	"context"
	"fmt"

	"github.com/goccy/go-json"
	"github.com/noahlsl/noah/consts"
	"go.etcd.io/etcd/client/v3"
)

func Load(cli *clientv3.Client, env string) *Cfg {

	c := &Cfg{}
	key := fmt.Sprintf(consts.CfgMySQL, env)
	res, err := cli.Get(context.Background(), key)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(res.Kvs[0].Value, &c)
	if err != nil {
		panic(err)
	}

	return c
}
