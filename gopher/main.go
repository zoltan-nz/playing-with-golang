package main

import (
	"fmt"
)

type gopher struct {
	name    string
	age     int
	isAdult bool
}

type horse struct {
	name   string
	weight int
}

type jumper interface {
	jump() string
}

func (g gopher) jump() string {
	if g.age < 65 {
		return g.name + " can jump high"
	}
	return g.name + " can still jump"
}

func (h horse) jump() string {
	if h.weight > 2500 {
		return "I'm too heavy..."
	}
	return "I will jump..."
}

func main() {
	jumperList := getList()

	for _, jumper := range jumperList {
		fmt.Println(jumper.jump())
	}
}

func validateAge(g *gopher) {
	g.isAdult = g.age >= 21
}

func getList() []jumper {
	gopher1 := &gopher{name: "Joe", age: 37}
	gopher2 := &gopher{name: "Bill", age: 70}
	horse1 := &horse{name: "Kincsem", weight: 2000}

	list := []jumper{gopher1, gopher2, horse1}
	return list
}
