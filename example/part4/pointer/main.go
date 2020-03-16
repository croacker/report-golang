package main

import "fmt"

type employee struct {
	name   string
	salary int
}

func (e employee) update() {
	e.name = "unknown"
	e.salary = 0
}

func main() {
	e := employee{"Concrete", 100}
	fmt.Println(e)
	// e.update()
	// fmt.Println("after update", e)
	// e.dismis()
	// fmt.Println("after dismis", e)
	// e.replace()
	// fmt.Println("after replace", e)
	// e.super_replace()
	// fmt.Println("after super_replace", e)
}

func (e *employee) dismis() {
	e.name = "unknown"
	e.salary = 0
}

func (e *employee) replace() {
	e = &employee{"always unknown", -1}
}

func (e *employee) super_replace() {
	*e = employee{"always unknown", -1}
}
