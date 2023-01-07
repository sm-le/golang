// Go supports constants of char, str, bool, num

package main

import (
	"fmt"
)

const a string = "constant"
const b int = 500000
const c bool = true

func main() {
	// str
	fmt.Println(a)

	// num
	fmt.Println(b)

	// bool
	fmt.Println(c)

	// numeric operation
	// return int not float
	fmt.Println(1e6 / b)

}
