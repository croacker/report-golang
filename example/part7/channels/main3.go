package main

import (
	"fmt"
	"time"
)

//Без каналов

var msg1 = "start"
var msg2 = "start"

func producer() {
	i := 0
	for true {
		i++
		msg1 = "from producer " + string(i)
		time.Sleep(time.Second * 2)
	}
}

func consumer() {
	for {
		fmt.Println(msg2)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go producer()
	go consumer()

	var input string
	fmt.Scanln(&input)
}
