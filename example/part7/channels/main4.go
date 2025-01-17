package main

import (
	"fmt"
	"time"
)

//Буферизованный канал и чтение в цикле

func producer(c chan string) {
	i := 0
	for {
		i++
		fmt.Println("in producer", i)
		c <- "from producer " + string(i)
		time.Sleep(time.Second * 1)
	}
}

func rangeConsumer(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {
	c := make(chan string, 5)
	go producer(c)
	go rangeConsumer(c)

	var input string
	fmt.Scanln(&input)
}
