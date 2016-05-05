// うぅ...汚い
// (´・ω・｀)

package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var size = 10

// data source
type data struct {
	slice *([]int)
}

func (d *data) init() {
	init := make([]int, size)
	for i := 0; i < size; i++ {
		init[i] = i
	}
	d.slice = &init
	fmt.Printf("#%v\n", (*d.slice))
}

// 先頭を削除して、削除した値を返す
func (d *data) pop(mutex *sync.Mutex) (int, error) {
	// lock を取る
	mutex.Lock()
	// 最後にunlockする
	defer mutex.Unlock()

	if len(*d.slice) < 1 {
		return 0, errors.New("slice size = 0")
	}

	result := (*d.slice)[0]

	fmt.Printf("Before Wait: #%v\n", (*d.slice))
	// ここにwaitを入れると、不整合が怒る
	time.Sleep(time.Millisecond * 1)
	fmt.Printf("After Wait : #%v\n", (*d.slice))
	if len(*d.slice) == 1 {
		*d.slice = (*d.slice)[:0]

	} else {
		*d.slice = (*d.slice)[1:]
	}
	return result, nil

}

// data source から取れる限り値を取って和を返す
func fold(ch chan int, d data, sum int, mutex *sync.Mutex) int {
	n, err := d.pop(mutex)
	if err != nil {
		if ch == nil {
			return sum
		} else {
			ch <- sum
			return 0
		}
	}

	if ch == nil {
		return fold(nil, d, n+sum, mutex)
	} else {
		ch <- fold(nil, d, n+sum, mutex)
		return 0
	}
}

func main() {
	ch := make(chan int, 2)
	d := data{}
	d.init()
	// mutex
	mutex := new(sync.Mutex)

	fmt.Println("--- start")

	// s の中の数字の和を求めたい
	// size = 100のとき、4950
	go fold(ch, d, 0, mutex)

	//time.Sleep(time.Millisecond * 1000)

	go fold(ch, d, 0, mutex)
	n := <-ch
	m := <-ch

	fmt.Printf("--- result shoud be %d.\n", size*(size-1)/2)
	fmt.Println(n + m) // == 4950

	fmt.Println("--- end")
}
