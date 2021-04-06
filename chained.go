package main

import "fmt"

func main() {
	//  && - and
	//  || - or
	//  !  - not
	fmt.Printf("%t\n", true && false)
	fmt.Printf("%t\n", true || false)
	fmt.Printf("%t\n", !true)
	fmt.Printf("%t", ((true || false) && (!false)))

}
