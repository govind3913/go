package main

import "fmt"

func sliceData() {
	sliceData := []int{1, 2, 3}

	anotherSlice := sliceData
	anotherSlice[0] = 10000

	fmt.Println(anotherSlice)

	//. Append data, existing array array increase its size
	sliceData = append(sliceData, 10)
	fmt.Println(sliceData)

	fmt.Printf("Memory address of the array: %p\n", &sliceData)
	fmt.Printf("Memory address of the array: %p\n", &anotherSlice)

	fmt.Printf("Memory address of the array: %p\n", &sliceData[0])
	fmt.Printf("Memory address of the array: %p\n", &anotherSlice[0])

}
