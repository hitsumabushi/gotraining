package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var size = 25

func main() {
	var num int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer (*wg).Done()
			// atomic
			atomic.AddInt32(&num, 1)
		}(wg)
	}
	wg.Wait()
	fmt.Printf("Shoud be %d.\n", size)
	fmt.Println(num)
}
