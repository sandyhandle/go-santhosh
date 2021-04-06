// good to go
package main

import (
	"fmt"
)

func main() {
	var mp map[string]int = map[string]int{
		"Ranjith":    5,
		"Santhosh":   6,
		"Sowmiya":    7,
		"Anagumathi": 8,
	}
	fmt.Println(mp)
	mp["Ranjith"] = 7 // to update
	fmt.Println(mp)
	mp["Priya"] = 10 // to insert a new element
	fmt.Println(mp)
	delete(mp, "Santhosh")
	fmt.Println(mp)

	// other uses of maps
	val, ok := mp["Preetha"]
	fmt.Println(val, ok)
	value, okey := mp["Priya"]
	fmt.Println(value, okey)

	fmt.Println(len(mp))
}
