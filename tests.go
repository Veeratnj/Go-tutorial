package main

import "fmt"

// type Speaker interface {
// 	Speak()
// }



type Person struct {
	Name string
	Age  int
}

func saySomething(s Person) {
	s.Speak()
}

func (p Person) Speak() {
	fmt.Println("Hi, I'm", p.Name)
}

func main() {
	println("Hello, World!")
}
