package main

import "fmt"

type gopher struct {
	name string
	age  int
	isAdult bool
}

func (g gopher) jump() string {
	if g.age < 65 {
		return g.name + " can jump high"
	}
	return g.name + " can still jump"
}

func validateAge(g *gopher) {
	g.isAdult = g.age >= 21
}

func main() {
	gopher1 := &gopher{name: "Joe", age: 37}
	gopher2 := gopher{name: "Bill", age: 70}

	validateAge(gopher1)

	fmt.Println(gopher1)
	fmt.Println(gopher2)
}
