package main

import (
	"fmt"
)

func main() {
	// Switch case type : 1
	x := 3
	switch x {
	case 1:
		fmt.Println("First case", x)
	case 2:
		fmt.Println("second case", x)
	default:
		fmt.Println("your input is not matching any case")
	}

	// Switch case type : 2
	y := 5
	switch {
	case y < 0:
		fmt.Println("the value is lesser than 0")
	case y > 0:
		fmt.Println("The value is grether than 0")
	default:
		fmt.Println("zero")
	}

}
