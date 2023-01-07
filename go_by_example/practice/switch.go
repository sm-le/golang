// switch statement express conditionals

package main

import (
	"fmt"
	"time"
)

func main() {

	// basic switch
	i := 2
	fmt.Println(i, " is expressed as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// multiple expression allowed and a default setting
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("This is weekend")
	default:
		fmt.Println("This is weekday")
	}

	// If/else operation alternative
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Before Noon aye")
	default:
		fmt.Println("Affffternoon")
	}

}
