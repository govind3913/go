package main

import (
	"fmt"
)

type EmployeeDetails struct {
	employeeName    string
	employeeAge     int
	employeeSalary  float64
	employeeAddress Address
}
type Address struct {
	city    string
	street  string
	country string
}

func structure() {
	employeeOther := EmployeeDetails{}
	fmt.Println(employeeOther.employeeName)

	employeeDetailsObject := EmployeeDetails{
		employeeName:   "Mayank",
		employeeAge:    30,
		employeeSalary: 300,
	}

	employeePointer := new(EmployeeDetails)
	employeeDataClone := employeePointer
	employeeDataClone.employeeSalary = 232434253

	employeePointer.employeeName = "TechnoFunnel"
	(*employeePointer).employeeAge = 200
	(&(*employeePointer)).employeeSalary = 100000
	changeNamePointer(employeePointer)

	fmt.Println(employeePointer)

	employeeDetailsObject.employeeAge = 100

	// It is copied by Value
	employeeClone := employeeDetailsObject
	employeeClone.employeeAge = 30076

	fmt.Print(employeeDetailsObject)
	fmt.Println(employeeClone)

	// Copied by Value
	changeName(employeeDetailsObject)
	changeNamePointer(&employeeDetailsObject)

	fmt.Print(employeeDetailsObject)
}

func changeName(emp EmployeeDetails) {
	emp.employeeName = "Random Name"
	fmt.Println(emp)
}

func changeNamePointer(emp *EmployeeDetails) {
	emp.employeeName = "Random Name Pointer"
	fmt.Println(emp)
}
