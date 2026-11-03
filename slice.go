package main

import "fmt"

func slice() {

	// Defining a array
	arrayData := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sliceData := arrayData[2:6] // Starts from 3, 4, 5, 6
	sliceData[0] = 3000
	fmt.Println(sliceData)
	fmt.Println(arrayData)
	sliceData1 := arrayData[:7] // Starts from 3, 4, 5, 6
	sliceData[0] = 3000
	fmt.Println(sliceData1)
	fmt.Println(arrayData)

}
