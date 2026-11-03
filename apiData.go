package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Employees struct {
	Id             string
	Name           string
	CreatedAt      string
	Avatar         string
	SomeOtherProps int
}

var employeeArray []Employees

func getEmloyeeDataFromApi(appUrl string) []Employees {
	response, err := http.Get(appUrl)

	if err == nil {
		byteData, _ := io.ReadAll(response.Body)
		marshalErr := json.Unmarshal(byteData, &employeeArray)
		if marshalErr != nil {
			fmt.Println(employeeArray)
		}
		return employeeArray
	}

	return []Employees{}
}
