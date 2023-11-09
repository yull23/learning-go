package main

import "fmt"

func main() {
	fmt.Println("Conditionals")
	a := 3
	b := 6
	if a == b {
		fmt.Println("a is equal to y")
	} else if a <= b {
		fmt.Println("a is less than y")
	} else {
		fmt.Println("a is greater than y")
	}

	fmt.Println("Switch")

	value := 3

	switch value {
	case 1:
		fmt.Println("Values is :", value)
	case 2:
		fmt.Println("Values is :", value)
	case 3:
		fmt.Println("Values is :", value)
	case 4:
		fmt.Println("Values is :", value)
	}

}
