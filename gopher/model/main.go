package model

type gopher struct {
	name    string
	age     int
	isAdult bool
}

func (g gopher) Jump() string {
	if g.age < 65 {
		return g.name + " can Jump high"
	}
	return g.name + " can still Jump"
}

type horse struct {
	name   string
	weight int
}

func (h horse) Jump() string {
	if h.weight > 2500 {
		return "I'm too heavy..."
	}
	return "I will Jump..."
}

type jumper interface {
	Jump() string
}

func GetList() []jumper {
	gopher1 := &gopher{name: "Joe", age: 37}
	gopher2 := &gopher{name: "Bill", age: 70}
	horse1 := &horse{name: "Kincsem", weight: 2000}

	list := []jumper{gopher1, gopher2, horse1}
	return list
}
