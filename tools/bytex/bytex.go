package bytex

import "github.com/goccy/go-json"

func ToBytes(data any) []byte {
	marshal, _ := json.Marshal(data)
	return marshal
}

func ToMap(data []byte) map[string]interface{} {
	out := map[string]interface{}{}
	_ = json.Unmarshal(data, &out)
	return out
}

func ToAny[T any](data []byte) T {
	var out T
	_ = json.Unmarshal(data, &out)
	return out
}
