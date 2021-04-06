package main

import "fmt"

func main() {
	var name string = fmt.Sprintf("number %2.f is cool", 45.333)
	fmt.Println(name)
}
