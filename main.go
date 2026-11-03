package main

import (
	"fmt"
	"reflect"
)

var username string = "Govind"
var userCompany = "JPMC"
var userSalary = 40
var (
	empName, empDesignation string
	empSalary, empBonus     int
)

func main1() {
	// name := ""
	// company := ""
	// count := 10
	// fmt.Scanln(&name)
	// fmt.Scanln(&company)
	// fmt.Scanln(&count)
	userBonus := 20
	fmt.Println(username)
	fmt.Println(userCompany)
	fmt.Println(userSalary)
	fmt.Println(userBonus)
	fmt.Println(empName)
	fmt.Println(empSalary)
	fmt.Println(empBonus)
	//printValue := fmt.Sprintf("Hello %s welcome to %s %d", name, company, userBonus)
	//fmt.Println(printValue)
	fmt.Println("The variable  user Bonus is of type %s", reflect.TypeOf((userBonus)))
	userAge := 10
	updateAage(&userAge)
	fmt.Println(userAge)
	userSalary := new(int)
	*userSalary =100;
	updateAage(userSalary)
	fmt.Println(*userSalary)

	resultantStrValue,resultint,resultbool := someFunctionImplementation("user")
	fmt.Println(resultantStrValue)
	fmt.Println(resultint)
	fmt.Println(resultbool)
	myVeriadicFunction(32.32,"Govind", 2,4,6,8,10,12,14,26,18,20,22,24,26,28)
	myVeriadicFunction(32.32,"Govind", 2,4,6,8,10,12,14,26)
	myVeriadicFunction(32.32,"Govind", 2,4,6,8,10,12)

}
func updateAage(userAgeRef *int) {
	*userAgeRef = 50
}
func someFunctionImplementation(inputStr string)(string,int,bool){
updateData:="Hello "+inputStr
return updateData,10,false
}
func myVeriadicFunction(f float64,str string, inputParams ... int){
	fmt.Println(inputParams)
	fmt.Println(f)
	fmt.Println(str)
}
