package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printMessage("Goroutine") // runs concurrently
	printMessage("Main")         // main goroutine
}
