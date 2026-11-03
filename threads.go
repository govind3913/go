package main

import (
	"fmt"
	"sync"
)

func mainthread() {

	var waitGroupCounter sync.WaitGroup

	waitGroupCounter.Add(2)

	employeeList := []string{"Mayank", "TechnoFunnel"}

	go func() {
		employeeList = append(employeeList, "Data 1")
		fmt.Println("Hello Team, This is Golang Channels 0")
		waitGroupCounter.Done()
	}()

	go func() {
		employeeList = append(employeeList, "Data 2")
		fmt.Println("Hello Team, This is Golang Channels1")
		waitGroupCounter.Done()
	}()

	waitGroupCounter.Wait()

	for _, value := range employeeList {
		fmt.Println(value)
	}
}
