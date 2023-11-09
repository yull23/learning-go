package main

import "fmt"

func main() {
	fmt.Println("Creating maps")
	myMap := make(map[string]int)
	fmt.Println(myMap)
	fmt.Println("Assigning values to the array")
	myMap["1"] = 3
	myMap["1"] = 3
	myMap["1"] = 1
	myMap["2"] = 3
	fmt.Println(myMap)

	fmt.Println("Operations with maps\n")

	fmt.Println("Accesing maps elements")
	fmt.Println("Existing element")
	a := myMap["1"]
	fmt.Println(a)
	fmt.Println("Non-Existing element")
	a = myMap["3"]
	fmt.Println(a)
	// Si el elemento no existe, entonces retorna cero.

	fmt.Println("Obteniendo dos variables de la asignación")

	x, y := myMap["1"]
	w, z := myMap["3"]
	fmt.Println(x, y)
	fmt.Println(w, z)

	// La segunda variable, retorna verdadero si existe, y falso sino.
	fmt.Println("Removing items")
	delete(myMap, "2")
	fmt.Println(myMap)

	fmt.Println("Iterating on a map")

	myMap["2"] = 2
	myMap["3"] = 3
	myMap["4"] = 4

	// Cuando se recorre en un for, se obtiene el par clave valor.
	for a, b := range myMap {
		fmt.Println("Getting key-value pair", a, b)
	}

}
