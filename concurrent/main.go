package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	names := []string{"First", "Second", "Third"}
	var wg sync.WaitGroup
	wg.Add(len(names))
	for _, name := range names {
		go printName(name, &wg)
	}
	wg.Wait()
}

func printName(n string, wg *sync.WaitGroup) {
	result := 0.0
	for i := 0; i < 1000000000; i++ {
		result += math.Pi * math.Sin(float64(len(n)))
	}
	fmt.Println("Name: ", n)
	wg.Done()
}
