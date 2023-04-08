package slicex

import (
	"fmt"
	"strconv"
	"testing"
)

type xxs struct {
	Xx string
}

func TestSplitArray(t *testing.T) {
	var data []int
	for i := 0; i < 100; i++ {
		data = append(data, i)
	}
	array, err := SplitArray[int](data, 10)
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range array {
		fmt.Println(v)
	}
}
func TestSplitArrayStruct(t *testing.T) {
	var data []xxs
	for i := 0; i < 100; i++ {
		data = append(data, xxs{Xx: strconv.Itoa(i)})
	}
	array, err := SplitArray[xxs](data, 10)
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range array {
		for _, v2 := range v {
			fmt.Println(v2.Xx)
		}
	}
}

func TestAnyJoin(t *testing.T) {
	s := []int{1, 2, 3}
	join := AnyJoin(s, "&")
	fmt.Println(join)
}
