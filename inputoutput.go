package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type your birth year :")
	scanner.Scan()
	inp, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("your age will be '%d' at the end of the 2021", 2021-inp)

}
