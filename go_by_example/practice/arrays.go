// An array is a numbered sequence of elements
// of a specified length

package main

import (
	"fmt"
)

func main() {

	// Set up an empty array
	var a [5]int
	fmt.Println("zero array:", a)

	// 0-based array replace
	a[4] = 500
	fmt.Println("non-zero array:", a)

	// Retrieve changed element in array
	fmt.Println("fifth element:", a[4])

	// Length of an array
	fmt.Println("length:", len(a))

	// Declare array in one line
	b := [5]int{6, 2, 3, 4, 1}
	fmt.Println("custom array:", b)

	// MultiDimensional Array, 2D
	var tD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			tD[i][j] = i + j
		}
	}
	fmt.Println("2D:", tD)
}
