package main

import "fmt"

func main() {

	fmt.Println("Hi")
	arrray := [3]int{1, 2, 3}
	fmt.Println("Array:", arrray)
	slice := []string{"apple", "banana", "cherry"}	
	fmt.Println("Slice:", slice)
	myMap := map[string]int{"age": 30, "height": 175}
	fmt.Println("Map:", myMap)
}
