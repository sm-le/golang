// Slices are a key data type in Go
// differing from arrays

package main

import (
	"fmt"
)

func main() {

	// initialize empty slice
	s := make([]string, 3)
	fmt.Println("empty array", s)

	// add elements
	s[0], s[1], s[2] = "a", "b", "c"

	fmt.Println("Slice", s)

	// append elements
	s = append(s, "d")
	fmt.Println("1 apd", s)

	// multiple append
	s = append(s, "e", "f")
	fmt.Println("multi apd", s)

	// copy slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied", c)

	// slice operator
	l := s[:5]
	fmt.Println("sl1:", l)
}
