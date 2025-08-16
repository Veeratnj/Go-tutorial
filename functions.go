package main

import (
	"fmt"
	"reflect"
)

func main() {
	x :=5
	y:=&x
	fmt.Println("This is the functions.go file.")
	fmt.Println("Calling add function with 5 and 10:", add(5, 10))
	fmt.Println(reflect.TypeOf(add(1,2))) 
	fmt.Println(y)
	fmt.Printf("Address: %p\n", &x)


	// t := reflect.TypeOf(x)
	// for i := 0; i < t.NumField(); i++ {
	// 	fmt.Println("Field:", t.Field(i).Name)
	// }
}


func add(a int,b int) int{
	return a+b
}