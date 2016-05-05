package main

import "fmt"

type singleton struct {
	// Some fields
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = func() *singleton {
			fmt.Println("何か複雑な初期化がある")
			return &(singleton{})
		}()
	}
	return instance
}

func main() {
	GetInstance()
}
