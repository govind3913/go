
package main

import "fmt"

func main3() {

// Creating basic Array 
	var arrayData [3]int
	fmt.Println(arrayData)
	arrayData[0] = 10

	//ShortHand Notation for creating array
	salaryArray := [5]float64{10.1, 23.4, 56.6, 68.5, 123}
	fmt.Println(salaryArray)

	// That the operation copies values of salaryArray in otherDetails (not by reference)
	otherDetails := salaryArray
	otherDetails[0] = 56.87
	fmt.Println(otherDetails)

	// Copy array from one reference and create other reference
	cloneArray := salaryArray
	fmt.Println(salaryArray == cloneArray)

	
	updateArrayData(&arrayData)
	fmt.Println(arrayData)

}

func updateArrayData(arrayInput *[3]int) {
	(*arrayInput)[2] = 1000
	arrayInput[1]=450

}
