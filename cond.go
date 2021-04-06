package main

import "fmt"

func main() {
	fmt.Println("Numerical Comparisation")

	x := 1
	var y int = 2

	fmt.Printf("%t \n", x < y)
	fmt.Printf("%t \n", x <= y)
	fmt.Printf("%t \n", x > y)
	fmt.Printf("%t \n", x >= y)
	fmt.Printf("%t\n", x == y)
	fmt.Printf("%t\n", x != y)

	fmt.Println("String Comparisation")
	var s string = "hello"
	fmt.Printf("%t\n", "Hello" == s)
	fmt.Printf("%t\t,%t\t,%t\t,%t", 'B' < 'a', 'B' <= 'C', 'G' > 'L', 'Z' >= 'Y')
}
