package main

import (
	"fmt"
)

func test() {
	fmt.Println("This test works")
}
func test1(u int) {
	fmt.Println("This test works", u)
}

func testing(tst func(x int) int) {
	fmt.Println(tst(7))
}

func returnfunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

// any doubts please visit..
// https://www.youtube.com/watch?v=vdm04bVzkLg&list=PLzMcBGfZo4-mtY_SE3HuzQJzuj4VlUG0q&index=17
func main() {
	returnfunc("Modera")()
	ans := returnfunc("Uchiha")
	ans()
	/*
		// calling the function type 1
		x := test // we can assign it to a variable
		x()
		u := test1
		u(8)
		y := func() {
			fmt.Println("This also works !")
		}
		y()
		// another type of implementing the above function with variable.
		y1 := func(x int) {
			fmt.Println("This also works !", x)
		}
		y1(8)
		// another type of implementing the above function with variable.
		y2 := func(x int) int {
			return (-x)
		}
		fmt.Println(y2)
		testing(y2)*/

}
