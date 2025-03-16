package main

import (
    "fmt"
)

func main() {
    // Function to calculate the area of a circle
    func circleArea(radius float64) float64 {
        return 3.14 * radius * radius
    }

    // Test the function with different inputs
    fmt.Println("The area of a circle with radius = 5 is", circleArea(5))
    fmt.Println("The area of a circle with radius = 8 is", circleArea(8))
    fmt.Println("The area of a circle with radius = 12 is", circleArea(12))
}
