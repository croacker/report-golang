package main

import "fmt"

type animal struct {
	name string
}

func main() {
	a := &animal{name: "lamina"}
	if a != nil 
	{
		fmt.Println("animal:", *a)
	}
}
