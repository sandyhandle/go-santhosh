// good to go
package main

import (
	"fmt"
	// we can also use "math" module
)

func main() {
	var num1 int = 9
	var num2 int = 5
	var num3 float64 = float64(num1) / float64(num2)
	fmt.Printf("The output will be: %f", num3)

}
