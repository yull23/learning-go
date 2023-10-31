package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Comenzando !!!!!!!!!!!!!!")

	fmt.Println("♫ \nDifference between declaration of variables\n")

	fmt.Println("♫ \nGlobal statement\n")

	a := 3
	fmt.Println(a) // Imprime 3

	fmt.Println("♫ \nLocal declaration\n")

	var b int
	b = 3
	fmt.Println(b) // Imprime 3

	fmt.Println("♫ \nDeclaring variables and assigning them values\n")

	var number int = 5
	var number1 int8 = 5
	var number2 int16 = 5
	var numberFloat float32 = 43.25
	var numberFloat1 float64 = 43.25
	fmt.Println(number)
	fmt.Println(number1)
	fmt.Println(number2)
	fmt.Println(numberFloat)
	fmt.Println(numberFloat1)
	var myString string = "Hello world!!"
	fmt.Println(myString)

	fmt.Println("♫ \nDeclaring variables without assignmet our direct assignment\n")
	var numberA int
	var numberB float32
	var stringA string
	numberA = 5
	numberB = 50.25
	stringA = "New String"
	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(stringA)

	fmt.Println("♫ \nViewing the type of variable \n")

	fmt.Println(reflect.TypeOf(number))
	fmt.Println(reflect.TypeOf(number1))
	fmt.Println(reflect.TypeOf(number2))
	fmt.Println(reflect.TypeOf(numberFloat))
	fmt.Println(reflect.TypeOf(numberFloat1))
	fmt.Println(reflect.TypeOf(myString))

	fmt.Println("♫ \nAssigning of variables\n")
	number = 5
	fmt.Println(number)
	numberFloat = 45.55
	fmt.Println(numberFloat)
	myString = myString + "--" + myString
	fmt.Println(myString)

}
