package main

import (
    "fmt"
    "math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
    for i := 0; i < 5; i++ {
        // Use rand.Perm to generate a random array of numbers.
        numbers := rand.Perm(5)
        fmt.Println(numbers)
    }
}
