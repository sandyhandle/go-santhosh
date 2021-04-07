package main

import "fmt"

type pointer struct { // it is like creating a class..
	x int16
	y int16
}

func change(pont *pointer) {
	pont.x = 99
}

func main() {
	var num pointer = pointer{1, 2}
	fmt.Println(num)
	num.x = 98
	fmt.Println(num)

	num1 := pointer{2, 4}
	fmt.Println(num1)
	change(&num1)
	fmt.Println(num1)
}
