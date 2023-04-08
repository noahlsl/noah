package hash

import (
	"strconv"
	"sync"

	"stathat.com/c/consistent"
)

var (
	c   *consistent.Consistent
	one sync.Once
)

func InitConsistent() {
	c = consistent.New()
	for i := 1; i < 1001; i++ {
		v := strconv.Itoa(i)
		if len(v) == 1 {
			v = "000" + v
		} else if len(v) == 2 {
			v = "00" + v
		} else if len(v) == 3 {
			v = "0" + v
		}
		c.Add("piece:" + v)
	}
}

func GetOneHash(data string) string {
	one.Do(InitConsistent)
	f, _ := c.Get(data)
	return f
}
