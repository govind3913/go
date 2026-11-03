package main

import "fmt"
type cabeAddedData interface{
	int|float32|float64
}

func mainInterface() {
	sumOfData(23.3,44.4)
	sumOfData(3,4)
	sumOfData(23,44)
}
func sumOfData[T cabeAddedData ](a,b T) {
	fmt.Println(a+b)
	
}