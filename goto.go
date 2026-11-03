package main

import "fmt"

func gotokeyword() {

	userNumerator := 0
	divideData := 0

cleanup:

	fmt.Scanln(&userNumerator)

	if userNumerator == 0 {
		goto cleanup
	}

success:

	divideData = 10 / userNumerator
	fmt.Println(divideData)

	userDetails := "Mayank"
	fmt.Println(userDetails)

	goto success

}
