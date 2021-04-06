// conditional statements
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// if class
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	age, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	if age >= 18 {
		fmt.Println("you can ride!, just enjoy")
	} else if age >= 14 && age < 18 {
		fmt.Println("you can ride with your parents!")
	} else {
		fmt.Printf("Please come after %d years", 18-age)
	}
}
