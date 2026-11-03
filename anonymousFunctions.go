package main

import (
	"fmt"

)

func main2() {
	inputFunction := declareFunction(0)
	inputFunction()
	inputFunction = declareFunction(10)
	inputFunction()
	executeOperation(10,2, sum)
	executeOperation(10,2, multiply)
	executeOperation(10,2, division)
}

func executeOperation(a,b int, operateFunction func(int,int)){
	operateFunction(a,b)
}
func sum(a,b int){
	fmt.Println(a+b)
}
func multiply(a,b int){
	fmt.Println(a*b)
}
func division(a,b int){
	fmt.Println(a/b)
}

func declareFunction(someInput int) func(){
	var sampleFunction func() 
		if someInput ==0 {
			sampleFunction = func(){
				fmt.Println("Hello from anonymous function")	
			}
		} else{
			sampleFunction= func(){
				fmt.Println("Hello from specific function")	
			}
		}
	
	return sampleFunction
}