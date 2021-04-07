package main

import "fmt"

func changefirst(x []int) {
	x[0] = 100
}
func main() {
	// mutable..
	/*
		x := 10
		y := x
		y = 9
	*/
	// fmt.Println(x, y)
	// slices array and maps are immutable...
	x1 := []int{2, 3, 4}
	fmt.Println(x1)
	y1 := x1
	y1[0] = 93
	fmt.Println(x1, y1)
	changefirst(x1)
	fmt.Println(x1, y1)

}
