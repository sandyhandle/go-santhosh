package main

import "fmt"

// any doubets visi https://youtu.be/a4HcEsJ1hIE
func main() {
	x := 49
	y := &x               // & symbol will give the memory location fo the variable that we are using after the &
	fmt.Println(x, y, *y) // * symbol will give yout the value that stored in the memory location of that varibable
	name := "Popu"
	var str *string = &name
	fmt.Println(str, *str)
}
