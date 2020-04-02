package main

import (
	"fmt"
	"time"
)

func asyncFunc() {
	i := 0
	for true {
		i++
		fmt.Println("from asyncFunc", i)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	go asyncFunc()
	fmt.Println("program started.....")
	var input string
	fmt.Scanln(&input)
}
