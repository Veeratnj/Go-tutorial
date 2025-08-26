package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	x:=5
	fmt.Fprintf(w, "Hello, World! You requested: ", x)
}

func main() {
	// Register route
	http.HandleFunc("/", helloHandler)

	// Start server on port 8080
	fmt.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
