package main

import (
	"fmt"
)

func chanelmain() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
		}

		close(ch)
	}()

	for inputData := range ch {
		fmt.Println(inputData)
	}
}
