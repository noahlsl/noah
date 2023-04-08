package slicex

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// SplitArray 切片拆分
// num 指每个切片里最大容量
func SplitArray[T any](arr interface{}, num int) ([][]T, error) {

	slice, ok := CreateAnyTypeSlice(arr)
	if !ok {
		return nil, errors.New("conversion error")
	}

	var data [][]T
	max := len(slice)
	if max <= num {
		var val []T
		for _, v := range slice {
			if v1, ok := v.(T); ok {
				val = append(val, v1)
			}
		}
		data = append(data, val)
		return data, nil
	}

	var quantity int
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	var start, end, i int
	for i = 1; i <= quantity; i++ {
		end = i * num
		var (
			val  []T
			list []interface{}
		)
		if i != quantity {
			list = slice[start:end]
		} else {
			list = slice[start:]
		}
		for _, v := range list {
			if v1, ok := v.(T); ok {
				val = append(val, v1)
			}
		}
		data = append(data, val)
		start = i * num
	}

	return data, nil
}

// CreateAnyTypeSlice 转interface切片
func CreateAnyTypeSlice(slice interface{}) ([]interface{}, bool) {

	val, ok := IsSlice(slice)
	if !ok {
		return nil, false
	}

	sliceLen := val.Len()
	out := make([]interface{}, sliceLen)
	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}
	return out, true
}

// IsSlice 校验是否是切片
func IsSlice(arg interface{}) (val reflect.Value, ok bool) {

	val = reflect.ValueOf(arg)
	if val.Kind() == reflect.Slice {
		ok = true
	}
	return
}

func AnyJoin(arr interface{}, seps ...string) string {

	slice, ok := CreateAnyTypeSlice(arr)
	if !ok {
		return fmt.Sprintf("%v", arr)
	}

	var sep = ","
	if len(seps) != 0 {
		sep = seps[0]
	}

	var out []string
	for _, d := range slice {
		out = append(out, fmt.Sprintf("%v", d))
	}
	return strings.Join(out, sep)
}
