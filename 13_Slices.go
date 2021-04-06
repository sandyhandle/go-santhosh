package main

import "fmt"

func main() {
	/* bacially two thing is created in slicing
	len() refers as length of the array
	cap() refers as a capacity of the array cap = index of the current current array - length of the previous array
	*/

	// creating slices: method 1
	var x [9]int = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s []int = x[3:5]
	fmt.Println(s[0:4], len(s), cap(s))

	// creating slices: method 2
	var y []int = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(cap(y[2:4]), y[2:4])

	var a []int = []int{1, 2, 3, 4}
	a = append(a, 6)
	fmt.Println(a)

	// creating array or slices
	b := make([]int, 5)
	fmt.Println(b)
}
