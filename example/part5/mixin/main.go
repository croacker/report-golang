package main

import "fmt"

type developer struct {
	firstName string
	lastName  string
}

func (d developer) fullName() string {
	return fmt.Sprintf("%s %s", d.firstName, d.lastName)
}

type project struct {
	name  string
	prize int
	developer
}

func (p project) info() {
	fmt.Println("name", p.name)
	fmt.Println("prize", p.name)
	fmt.Println("developer", p.fullName())
}

func main() {
	d := developer{"developer1", "developer1"}
	p := project{"superproject", 100, d}
	p.info()
}
