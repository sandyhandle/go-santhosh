package main

import (
	"fmt"
)

func add(x int, y int) {
	fmt.Println(x + y)
}

func test() {
	fmt.Println("This test works")
}

func add1(x, y int) int {
	return x + y
}

func add2(x, y int) (int, int) {
	return x + y, x - y
}

func add3(x, y int) (x3, x4 int) { // can also use (x3 int, x4 int)
	x3 = x * y
	x4 = x % y
	return x3, x4
}

func main() {
	test()
	add(7, 9)
	ans := add1(3, 4)
	fmt.Println(ans)
	ans1, ans2 := add2(3, 4)
	fmt.Println(ans1, ans2)
	ans3, ans4 := add3(7, 9)
	fmt.Println(ans3, ans4)

}
