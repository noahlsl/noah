package structx

import (
	"errors"
	"reflect"

	"github.com/goccy/go-json"
	"github.com/noahlsl/noah/tools/strx"
)

func StructToMap(in interface{}) (map[string]interface{}, error) {

	out := map[string]interface{}{}
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, errors.New("data is not struct")
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		tagValue := fi.Tag.Get("json")
		if tagValue == "" {
			tagValue = fi.Tag.Get("db")
		}
		if tagValue == "" {
			tagValue = fi.Tag.Get("form")
		}
		if tagValue != "" && tagValue != "-" {
			out[tagValue] = v.Field(i).Interface()
		}
	}

	return out, nil
}

func StructToBytes(in interface{}) []byte {

	marshal, _ := json.Marshal(in)

	return marshal
}

func StructToStr(in interface{}) string {

	marshal, _ := json.Marshal(in)

	return strx.B2s(marshal)
}
