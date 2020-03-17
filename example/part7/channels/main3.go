package main

import (
	"fmt"
	"time"
)

//Без каналов

var msg = "start"

func producer() {
	i := 0
	for true {
		i++
		msg = "from producer " + string(i)
		time.Sleep(time.Second * 2)
	}
}

func consumer() {
	for {
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go producer()
	go consumer()

	var input string
	fmt.Scanln(&input)
}
