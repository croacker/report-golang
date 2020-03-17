package main

import (
	"fmt"
	"time"
)

//Запись чтение из канала

func producer(c chan string) {
	i := 0
	for true {
		i++
		c <- "from producer " + string(i)
		time.Sleep(time.Second * 2)
	}
}

func consumer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c := make(chan string)
	go producer(c)
	go consumer(c)

	var input string
	fmt.Scanln(&input)
}
