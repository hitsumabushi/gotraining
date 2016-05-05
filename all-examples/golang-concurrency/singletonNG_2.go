package main

import (
	"fmt"
	"math/rand"
	"time"
)

type singleton struct {
	label int
}

func (s singleton) printLabel() {
	fmt.Println(s.label)
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = func() *singleton {
			fmt.Println("何か複雑な初期化がある")
			time.Sleep(2 * time.Second)
			return &(singleton{label: rand.Intn(100)})
		}()
	}
	return instance
}

// 試験用関数
func getInstance() {
	a := GetInstance()
	a.printLabel()
}

func main() {
	// わかりやすさのための乱数
	rand.Seed(31)
	go getInstance()
	go getInstance()
	getInstance()
}
