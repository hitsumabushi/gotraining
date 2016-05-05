package main

import (
	"fmt"
)

func inputToChan(x chan int) {
	x <- 1
}

var size = 1

func main() {
	foo := make(chan int, size)
	bar := make(chan int, size)

LOOP:
	for {
		select {
		case input := <-foo:
			fmt.Printf("Foo gives %d\n", input)
			break LOOP
		case input := <-bar:
			fmt.Printf("Bar gives %d\n", input)
			break LOOP
		default:
			fmt.Println("no input.")
			go inputToChan(foo)
		}
	}

	fmt.Print("Exit.")
}
