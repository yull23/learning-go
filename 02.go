package main

import "fmt"

func main() {
	fmt.Println("Conditionals")
	a := 3
	b := 6
	if a == b {
		fmt.Println("a is equal to y")
	} else {
		fmt.Println("a is diferent to y")
	}
	fmt.Println("Loops")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}

}
