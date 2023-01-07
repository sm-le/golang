// This is a test
package main

import "fmt"

func main() {
	// declare first var with var
	var a = "initial var"
	fmt.Println(a)

	// declare two vars at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// declare zero value var
	var d int
	fmt.Println(d)

	// := to declare var
	e := "shortened version"
	fmt.Println(e)
}
