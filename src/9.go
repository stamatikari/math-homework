
package main
import "fmt"
func randInt(min int, max int) int {
    return min + (rand.Int() % (max - min))
}
func main() {
    fmt.Println("Random number between 1 and 10:", randInt(1, 10))
}