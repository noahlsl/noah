package sortx

import (
	"reflect"
	"sort"
	"strconv"
)

// SortStruct
// s 是倒序标识。填入任意值，排序变倒序
func SortStruct[T any](data []T, filed string, s ...any) []T {
	var (
		sortSlice    []int
		sortStrSlice []string
		val          = make(map[string]T)
		out          []T
		t            reflect.Type
		v            reflect.Value
	)

	for _, datum := range data {
		// 获取结构体的类型信息
		t = reflect.TypeOf(datum)
		// 获取结构体的值信息
		v = reflect.ValueOf(datum)
		if reflect.ValueOf(datum).Kind() == reflect.Ptr {
			t = reflect.TypeOf(datum).Elem()
			v = reflect.ValueOf(datum).Elem()
		}

		// 遍历结构体的字段
		for i := 0; i < t.NumField(); i++ {
			// 获取字段的名称和类型
			fieldName := t.Field(i).Name
			if fieldName != filed {
				continue
			}

			fieldType := t.Field(i).Type.Name()
			if fieldType == "string" {
				v1 := v.Field(i).String()
				sortStrSlice = append(sortStrSlice, v1)
				val[v1] = datum
			} else {
				v1 := int(v.Field(i).Int())
				sortSlice = append(sortSlice, v1)
				val[strconv.Itoa(v1)] = datum
			}
			break
		}
	}

	if len(sortSlice) != 0 {
		sort.Ints(sortSlice)
		if len(s) != 0 {
			for i := len(sortSlice); i > 0; i-- {
				out = append(out, val[strconv.Itoa(sortSlice[i-1])])
			}
		} else {
			for _, k := range sortSlice {
				out = append(out, val[strconv.Itoa(k)])
			}
		}
		return out
	}

	sort.Strings(sortStrSlice)
	if len(s) != 0 {
		for i := len(sortStrSlice); i > 0; i-- {
			out = append(out, val[sortStrSlice[i-1]])
		}
	} else {
		for _, k := range sortStrSlice {
			out = append(out, val[k])
		}
	}
	return out

}
