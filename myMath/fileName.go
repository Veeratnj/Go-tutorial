package myMath

func Add(a int, b int) int {
    return a + b
}

// Multiply two numbers
func Multiply(a int, b int) int {
    return a * b
}

func localFunction() {
    // This function is not exported and can only be used within this package
    // It can be used for internal calculations or helper functions
}


