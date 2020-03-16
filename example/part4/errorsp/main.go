package main

import (
	"errors"
	"fmt"
)

type employee struct {
	name   string
	salary int
}

func try(v int) (string, error) {
	if v > 10 {
		return "baaaaad", errors.New("too much")
	} else {
		return "good", nil
	}
}

func main() {
	msg, err := try(11)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
}
