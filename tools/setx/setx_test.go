package setx

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNewSetByInt(t *testing.T) {
	s := NewSet[int]()
	for i := 0; i < 5; i++ {
		s.Add(i)
	}
	for i := 0; i < 5; i++ {
		s.Add(i)
	}
	for i := 0; i < 5; i++ {
		s.Add(i)
	}
	fmt.Println(s.List())
}

func TestNewSetByStr(t *testing.T) {
	s := NewSet[string]()
	for i := 0; i < 5; i++ {
		s.Add(strconv.Itoa(i))
	}
	for i := 0; i < 5; i++ {
		s.Add(strconv.Itoa(i))
	}
	for i := 0; i < 5; i++ {
		s.Add(strconv.Itoa(i))
	}
	fmt.Println(s.List())
}
func TestNewSetByFloat32(t *testing.T) {
	s := NewSet[float32]()
	for i := 0; i < 5; i++ {
		s.Add(float32(i))
	}
	for i := 0; i < 5; i++ {
		s.Add(float32(i))
	}
	for i := 0; i < 5; i++ {
		s.Add(float32(i))
	}
	fmt.Println(s.List())
}
