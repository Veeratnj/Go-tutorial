package main

import "fmt"

// Define an interface
type Shape interface {
    Area() float64
}

// Structs that will implement Shape
type Rectangle struct {
    Width, Height float64
}

type Circle struct {
    Radius float64
}

// Rectangle implements Shape
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Circle implements Shape
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    var s Shape // Interface variable

    s = Rectangle{Width: 5, Height: 3}
    fmt.Println("Rectangle area:", s.Area())

    s = Circle{Radius: 4}
    fmt.Println("Circle area:", s.Area())
}
