package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var value int32

func setNX(newValue int32) bool {
	return atomic.CompareAndSwapInt32(&value, 0, newValue)
}

func main() {
	numWorkers := 100

	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			if setNX(int32(id + 1)) {
				fmt.Printf("Worker %d successfully set the value.\n", id)
			} else {
				fmt.Printf("Worker %d failed to set the value.\n", id)
			}
		}(i)
	}

	time.Sleep(time.Second) // 等待 goroutines 完成
	fmt.Printf("Final value: %d\n", value)
}
