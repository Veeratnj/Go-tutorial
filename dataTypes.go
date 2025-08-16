package main

import "fmt"

// basic data types
var myInt int;
var myInt16 int16;
var myInt32 int32;
var myInt64 int64;
var myUint uint;
var myFloat32 float32;
var myFloat64 float64;
var myString string;
var myBool bool
var myByte byte    // alias for uint8
var myRune rune    // alias for int32 (Unicode code point)

// aggregate data types
var myArray [3]int                // fixed length array
var mySlice []string              // slice (dynamic array)
var myStruct struct { name string } // anonymous struct

// reference data types
var myMap map[string]int          // map (key-value pairs)
var myPointer *int                 // pointer to int
var myChan chan int                // channel
var myFunc func(a, b int) int      // function type
var myInterface interface{}        // empty interface

func main() {
	// assigning some values
	myInt = 42
	myFloat64 = 3.14
	myString = "Hello Go"
	myBool = true
	mySlice = []string{"apple", "banana"}
	myMap = map[string]int{"age": 25}
	value := 100
	myPointer = &value
	myStruct.name= "John Doe"

	fmt.Println("myInt:", myInt)
	fmt.Println("myFloat64:", myFloat64)
	fmt.Println("myString:", myString)
	fmt.Println("myBool:", myBool)
	fmt.Println("mySlice:", mySlice)
	fmt.Println("myMap:", myMap)
	fmt.Println("myPointer value:", *myPointer)
}
