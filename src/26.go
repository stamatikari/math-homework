package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    primes := make([]int, 0)
    for i := 2; i <= 100; i++ {
        if isPrime(i) {
            primes = append(primes, i)
        }
    }

    fmt.Println("Primes:", primes)

    var questions []interface{}
    questions = append(questions, 576 + 39 * 42 * 83 + math.Sqrt(float64(10000)))
    questions = append(questions, 576 + 39 * 42 * 83 + math.Sqrt(float64(10000)))

    for _, q := range questions {
        fmt.Println("Question:", q)
    }
}

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    if n == 2 || n == 3 {
        return true
    }
    if n%2 == 0 || n%3 == 0 {
        return false
    }

    for i := 5; i*i <= n; i += 6 {
        if n%i == 0 || n%(i+2) == 0 {
            return false
        }
    }
    return true
}
