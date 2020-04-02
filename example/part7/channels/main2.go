package main

import (
	"fmt"
	"time"
)

//Запись без чтения канала

func producer(c chan string) {
	i := 0
	for {
		i++
		fmt.Println("in producer", i)
		c <- "from producer " + string(i)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	c := make(chan string, 10)
	go producer(c)

	var input string
	fmt.Scanln(&input)
}
