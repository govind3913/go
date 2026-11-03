package main

import "fmt"

// Using goto Statement
// Panic Creation and the Recovery Machanism
// Sending error from a function
// Handling Errors coming from internal Function
// Deffered statements (Expressions)

func error() {
	someFunctionCall()
}

// LIFO: 
func someFunctionCall() {
	defer func() {
		fmt.Println("Function 1")
	}()

	userName := "Mayank"
	fmt.Println(userName)

	defer func() {
		fmt.Println("Function 3")
	}()

	fmt.Println("This is End of the Function")

	defer func() {
		fmt.Println("Function 2")
	}()
}
