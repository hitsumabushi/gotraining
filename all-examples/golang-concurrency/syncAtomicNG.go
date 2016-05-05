package main

import (
	"fmt"
	"sync"
	"time"
)

var size = 25

func main() {
	num := 0
	wg := new(sync.WaitGroup)

	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer (*wg).Done()
			// atomicでない
			tmp := num
			time.Sleep(time.Millisecond * 100)
			num = tmp + 1
		}(wg)
	}
	wg.Wait()
	fmt.Printf("Shoud be %d.\n", size)
	fmt.Println(num)
}
