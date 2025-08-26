package main

import (
    "fmt"
    "github.com/veeratnj/Go-tutorial/myMath"  // import your custom package
)

func main() {
    sum := myMath.Add(3, 4)
    product := myMath.Multiply(3, 4)

    fmt.Println("Sum:", sum)
    fmt.Println("Product:", product)
}
