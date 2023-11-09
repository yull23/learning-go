package main

import "fmt"

func main() {
	fmt.Println("\nCreating arrays\n")
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	fmt.Println(myArray)

	fmt.Println("Directly create the arrays ")

	myArray1 := [4]int{4, 5, 6} //AUTOCMPLETA EL VALOR FALTANTE CON 0
	myArray2 := [4]int{4, 5, 6, 9}
	myArray3 := [4]int{}
	fmt.Println(myArray1)
	fmt.Println(myArray2)
	fmt.Println(myArray3)
	fmt.Println("Accessing the element")
	fmt.Println(myArray1[0])
	fmt.Println(myArray2[1])
	fmt.Println(myArray3[2])
	fmt.Println("\nCreating slices\n")

	var mySlice = []int{1}
	var mySlice1 = []int{1, 3, 4}
	var mySlice2 = []int{}
	fmt.Println(mySlice)
	fmt.Println(mySlice1)
	fmt.Println(mySlice2)

	mySliceA := []string{"a", "b"}
	fmt.Println(mySliceA)
	mySliceA = append(mySliceA, "c")
	fmt.Println(mySliceA)
	fmt.Println(mySliceA[0])

}
