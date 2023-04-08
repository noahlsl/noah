package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/noahlsl/noah/tools/bytex"
	"github.com/noahlsl/noah/tools/ipx"
	"github.com/noahlsl/noah/tools/strx"
	"github.com/redis/go-redis/v9"
)

type BaseDataMiddleware struct {
	key   string
	flags []string
	r     *redis.ClusterClient
}

func NewBaseDataMiddleware(r *redis.ClusterClient, key string, flags ...string) *BaseDataMiddleware {
	return &BaseDataMiddleware{
		key:   key,
		r:     r,
		flags: flags,
	}
}

func (m *BaseDataMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		aid := r.Header.Get("aid")
		data := map[string]interface{}{}
		if aid != "" {
			get, err := m.r.Get(context.Background(),
				fmt.Sprintf(m.key, aid)).Result()
			if err != nil {
				return
			}

			data = bytex.ToMap(strx.S2b(get))
		}
		for _, flag := range m.flags {
			data[flag] = r.Header.Get(flag)
		}
		data["ip"] = ipx.RemoteIp(r)
		data["path"] = r.URL.Path
		data["token"] = r.Header.Get("token")
		data["language"] = r.Header.Get("language")
		marshal, _ := json.Marshal(data)
		r.Header.Set("base", string(marshal))
		next(w, r)
	}
}

func (m *BaseDataMiddleware) OriginalHandle(w http.ResponseWriter, r *http.Request) error {

	aid := r.Header.Get("aid")
	data := map[string]interface{}{}
	if aid != "" {
		get, err := m.r.Get(context.Background(),
			fmt.Sprintf(m.key, aid)).Result()
		if err != nil {
			return err
		}

		data = bytex.ToMap(strx.S2b(get))
	}

	data["ip"] = ipx.RemoteIp(r)
	data["path"] = r.URL.Path
	data["token"] = r.Header.Get("token")
	data["language"] = r.Header.Get("language")
	for _, flag := range m.flags {
		data[flag] = r.Header.Get(flag)
	}
	marshal, _ := json.Marshal(data)
	r.Header.Set("base", string(marshal))
	return nil
}
