package main

import "fmt"

func main() {
	names := []string{"First", "Second", "Third"}
	for _, name := range names {
		printName(name)
	}
}

func printName(n string) {
	fmt.Println("Name: ", n)
}
