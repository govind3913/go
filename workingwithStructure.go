package main

import ("fmt")


func withstructs() {
	employeeOne := EmployeeDetails{}
	employeeOne.employeeName = "Mayank"
	employeeOne.employeeAddress = Address{
		city:    "delhi",
		street:  "delhi",
		country: "India",
	}

	employeeTwo := new(EmployeeDetails)
	employeeTwo.employeeName = "Mayank"
	employeeTwo.employeeAddress = Address{
		city:    "delhi",
		street:  "delhi",
		country: "India",
	}

	isEqual := (employeeOne == *employeeTwo)
	fmt.Println(isEqual)
}
