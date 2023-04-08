package logx

import (
	"github.com/noahlsl/noah/tools/ipx"
	"net/http"
)

type Field struct {
	Key   string
	Value interface{}
}

func GetFields(r *http.Request) []Field {
	var fields []Field
	fields = append(fields, Field{Key: "trace", Value: r.Header.Get("trace")})
	fields = append(fields, Field{Key: "path", Value: r.URL.Path})
	fields = append(fields, Field{Key: "ip", Value: ipx.RemoteIp(r)})
	fields = append(fields, Field{Key: "param", Value: r.Header.Get("param")})
	fields = append(fields, Field{Key: "method", Value: r.Method})
	fields = append(fields, Field{Key: "content_type", Value: r.Header.Get("Content-Type")})
	return fields
}

func GetFieldsByErr(r *http.Request, err error) []Field {
	var fields []Field
	getFields := GetFields(r)
	fields = append(fields, getFields...)
	fields = append(fields, Field{Key: "stack", Value: GetStack(err)})
	return fields
}

func ToMap(fields ...Field) map[string]interface{} {
	var data = make(map[string]interface{})
	for _, datum := range fields {
		data[datum.Key] = datum.Value
	}
	return data
}
