package main

import (
	"fmt"
	"os"
	"time"
	"errors"
)

func main() {
	args := os.Args;

	hourOfDay := time.Now().Hour()
	greeting, err := getGreeting(hourOfDay)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	if len(args) > 1 {
		fmt.Println(args[1])
	} else {
		fmt.Println(greeting)
	}
}

func getGreeting(hour int) (string, error) {
	var message string

	if hour < 7 {
		err := errors.New("Too early...")
		return message, err
	}

	if hour < 12 {
		message = "Good Morning"
	} else if hour < 18 {
		message = "Good Afternoon"
	} else {
		message = "Good Evening"
	}
	return message, nil
}


