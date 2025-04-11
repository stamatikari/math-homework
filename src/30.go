package main

import (
    "fmt"
    "math/rand"
)

func main() {
    // Generate a random number between 0 and 1
    var r rand.Rand = rand.New(rand.NewSource(42))
    num := float64(r.Int()) / 1.0

    if num <= 0.5 {
        fmt.Println("The number is less than or equal to 0.5")
    } else {
        fmt.Println("The number is greater than 0.5")
    }
}
