package main

import (
	"fmt"
	"sync"
	"time"
)

func echo(wg *sync.WaitGroup, s string, n int) {
	defer wg.Done()
	defer fmt.Println("Exit:" + s)
	for n > 0 {
		fmt.Println(s)
		n--
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("--- start")
	wg.Add(1)
	go echo(&wg, "○", 10)
	wg.Add(1)
	go echo(&wg, "●", 10)
	//echo("●", 10)

	wg.Wait()
	fmt.Println("--- end")
}
