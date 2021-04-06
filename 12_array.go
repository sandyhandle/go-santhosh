package main

import "fmt"

func main() {
	// creating array: method 1
	var arr [4]int
	arr[0] = 2
	arr[3] = 32
	fmt.Println(arr)
	// creating array: method 2
	num := [5]int{1, 2, 3, 4, 5}
	var sum int = 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	fmt.Println(sum)
	arrr := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(arrr)
	fmt.Println(arrr[1][3])
}
