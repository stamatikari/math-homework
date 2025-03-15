package main

import (
    "fmt"
    "math/rand"
)

func generateRandomNumber() int {
    return rand.Intn(100)
}

func main() {
    fmt.Println("The random number is:", generateRandomNumber())
}
