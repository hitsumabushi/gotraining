package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	// Some fields
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = func() *singleton {
			fmt.Println("何か複雑な初期化がある")
			return &(singleton{})
		}()
	})
	return instance
}

func main() {
	go GetInstance()
	go GetInstance()
	go GetInstance()
	GetInstance()
}
