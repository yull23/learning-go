package main

import (
	"fmt"
)

func main() {
	var number int = 35
	var letters string = "Hola"
	fmt.Println("Esta es un string =>", letters, "y este es un numero", number)
	var i int8 = 0
	fmt.Println(i)
	fmt.Println(i + 5)
	fmt.Println(i)

	// Combinando tipos de variables
	// fmt.Println(letters + number) rebota un error

	var text string = "(Este es un texto)"
	var newNumber int = 8
	fmt.Println(text, newNumber)
}
