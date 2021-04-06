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

	fmt.Printf("%t\n", "Hello" == "hello")
}
