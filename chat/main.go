package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0

	for {
		if i>= 5 {
			break
		}

		fmt.Println(i)
		i++
	}
}
