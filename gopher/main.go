package main

import (
	"fmt"
	"playing-with-golang/gopher/model"
)

func main() {
	jumperList := model.GetList()

	for _, jumper := range jumperList {
		fmt.Println(jumper.Jump())
	}
}
