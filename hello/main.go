package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args;

	hourOfDay := time.Now().Hour()
	greeting:= getGreeting(hourOfDay)

	if len(args) > 1 {
		fmt.Println(args[1])
	} else {
		fmt.Println(greeting)
	}
}

func getGreeting(hour int) string {
	if hour < 12 {
		return "Good Morning"
	} else if hour < 18 {
		return "Good Afternoon"
	} else {
		return "Good Evening"
	}
}


