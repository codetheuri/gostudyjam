package main

import "fmt"

// Simple function
func greet(name string) {
    fmt.Println("Hello,", name)
}

// Function with return value
func add(a, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    greet("Alice")
    
    sum := add(5, 3)
    fmt.Println("Sum:", sum)
    
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}