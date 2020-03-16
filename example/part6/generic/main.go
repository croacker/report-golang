package main

import "fmt"

type stringlike interface {
	toString() string
}

type developer struct {
	name string
}

func (d developer) toString() string {
	return d.name
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
