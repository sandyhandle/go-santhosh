package main

import "fmt"

func main() {
	//for loop first type
	x := 0
	for x <= 5 {
		fmt.Println(x)
		x++
	}
	// for loop second type
	for y := 1; y <= 5; y++ {
		fmt.Println(y)
	}
	// This is the third type
	// First create a infinite for loop
	// and then use break and continue to come our of the loop
	z := 1
	for { // can also write in for true {}
		if z == 5 {
			break
		} else {
			fmt.Println(z)
		}
		z += 1
	}

}
