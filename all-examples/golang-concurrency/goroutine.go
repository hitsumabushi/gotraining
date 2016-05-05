package main

import (
	"fmt"
	"time"
)

func echo(s string, n int) {
	for n > 0 {
		fmt.Println(s)
		n--
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("--- start")
	go echo("○", 10)
	//go echo("●", 10)
	echo("●", 10)
	fmt.Println("--- end")
}
