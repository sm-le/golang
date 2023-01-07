package main

import (
	"fmt"
)

func main() {

	// basic operation
	if 1%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// multi conditions
	if num := 5; num < 0 {
		fmt.Println("negative")
	} else if num < 10 {
		fmt.Println("single digit")
	} else {
		fmt.Println("multi digits")
	}
}
