package main

import (
	"fmt"
	"gopher/model"
)

func main() {
	jumperList := model.GetList()

	for _, jumper := range jumperList {
		fmt.Println(jumper.Jump())
	}
}
