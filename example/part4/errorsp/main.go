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
		return "", errors.New("too much")
	} else {
		return "good", nil
	}
}

func catch(e error) {
	fmt.Println(e)
}

func finaly() {
	fmt.Println("Recover")
}

func main() {
	e := employee{name: "name", salary: 100}

	defer finaly()
	msg, err := try(e.salary)
	if err != nil {
		catch(err)
	} else {
		fmt.Println(msg)
	}
}
