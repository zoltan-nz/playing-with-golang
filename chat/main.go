package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0

	for {
		if i >= 5 {
			break
		}

		fmt.Println(i)
		i++
	}

	var langs [3]string

	langs[0] = "Go"
	langs[1] = "Ruby"

	var langs2 []string

	langs2 = append(langs2, "Go")
	langs2 = append(langs2, "Ruby")

	fmt.Println(langs2)

	langs3 := []string{"Go", "Ruby", "JavaScript"}
	for _, element := range langs3 {
		fmt.Println(element)
	}

}
