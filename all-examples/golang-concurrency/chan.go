package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- string, s string) {
	defer fmt.Println("exit sender")
	fmt.Println("Sender: " + s)
	// チャンネルに入れる
	ch <- s
}

func reciever(ch <-chan string) {
	defer fmt.Println("exit reciever")

	fmt.Println("Recieve from cahnnel")
	// chにsenderから値を入れられるまでブロックされる
	fmt.Println("Reciever: " + <-ch)
	fmt.Println("Done.")
}

func main() {
	ch := make(chan string)
	fmt.Println("--- start")

	// reciever
	go reciever(ch)
	time.Sleep(5 * time.Second)

	// sender
	sender(ch, "○")

	fmt.Println("--- end")
}
