package main

import (
	"fmt"
)

func main() {

	// Standard var set and for loop scale
	i := 0
	for i <= 5 {
		fmt.Println("Loop1", i)
		i += 1
	}

	// Classic initial, condition, after
	for j := 1; j <= 3; j++ {
		fmt.Println("Loop2", j)
	}

	// infinite loop until break
	for {
		fmt.Println("Loop3")
		break
	}

	// continue in the iteration
	for n := 1; n <= 3; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("Loop4", n)
	}
}
