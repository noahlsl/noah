package copyx

import (
	"fmt"
	"strconv"
	"testing"
)

type A struct {
	Age  int64
	Name string
}
type B struct {
	Age  int64
	Name string
}

func TestCopyOne(t *testing.T) {
	var one *B
	a := A{Name: "test"}
	one = CopyOne[*B](a)
	fmt.Println(one)
}

// 1.2秒处理500W个
func TestCopySlice(t *testing.T) {
	var (
		aSlice []A
		bSlice []B
	)
	for i := 0; i < 50; i++ {
		aSlice = append(aSlice, A{Age: int64(i + 1), Name: strconv.Itoa(i)})
	}

	bSlice = CopySlice[B](aSlice)
	_ = bSlice
}
