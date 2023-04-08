package lockx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestSetLock(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		select {
		case <-time.After(11 * time.Second):
			cancelFunc()
		}
	}()
	if !SetLock(ctx, nil, 0) {
		fmt.Println("加锁失败")
	}

	time.Sleep(12 * time.Second)
}
