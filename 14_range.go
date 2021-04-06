package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5, 1, 2}
	// first method
	for i, element := range a {
		// for j = i +1; j < len(a); j++{
		//   if element == a[j]  the print element}
		for _, element2 := range a[i+1:] {
			if element == element2 {
				fmt.Println(element)
			}
		}
	}
	/* if any doubts
	    https://youtu.be/DYqpu3jF2_4
		here's the youtube link..
	*/
}
