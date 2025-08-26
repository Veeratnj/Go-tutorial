package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("This is the main function of the FileHandling package.")
	data, err := os.ReadFile("example.txt")
	if err != nil {
		// panic(err)
		fmt.Println("Error reading file:", err)
	}
	fmt.Println(string(data))
	fmt.Println(data)
}
