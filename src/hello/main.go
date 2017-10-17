package main

import (
	"fmt"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	} else {
		fmt.Println("Hello, I'm Gopher")
	}
}
