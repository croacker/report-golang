package main

import "fmt"

type worker interface {
	do() string
}

type developer struct {
	name string
}

func (d developer) do() string {
	fmt.Println("public static void main", name)
}

type project struct {
	name string
}

func (p project) toString() string {
	return p.name
}

type List []stringlike

func main() {
	list := List{developer{"dev1"},
		project{"proj1"},
		developer{"dev2"},
		project{"proj2"},
		developer{"dev3"},
	}
	for _, item := range list {
		fmt.Println(item.toString())
	}
}
