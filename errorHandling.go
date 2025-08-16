package main

import (
	"errors"
	"fmt"
)

func errorExample(i int, j int) (result int, errorMsg error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			errorMsg = errors.New("an error occurred")
		}else{
			fmt.Println("No panic occurred")
		}
	}()

	if i == 5 {
		result = 0
		errorMsg = errors.New("i cannot be 5")
		panic("i cannot be 5")
		return
	}else{
		result = i + j
		errorMsg = nil
		return
	}

}

func main() {
	result , err := errorExample(5,10)
	fmt.Println("Result:", result)
	fmt.Println("Error:", err)
}
