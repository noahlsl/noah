package poolx

import "github.com/panjf2000/ants/v2"

func NewFnPool(fn func(in []byte), size ...int) *ants.PoolWithFunc {

	var max = 1000
	if len(size) != 0 {
		max = size[0]
	}

	fnPool, _ := ants.NewPoolWithFunc(max, func(payload interface{}) {
		if in, ok := payload.([]byte); ok {
			fn(in)
		}
	})

	return fnPool
}
