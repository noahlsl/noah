package serverx

import (
	"context"
	"fmt"

	"github.com/goccy/go-json"
	"github.com/noahlsl/noah/consts"
	"go.etcd.io/etcd/client/v3"
)

func AnyLoad[T any](cli *clientv3.Client, project, env string) T {

	var c T
	key := fmt.Sprintf(consts.CfgServer, env, project)
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
