package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("My favorite number is", rand.Intn(50))

    fmt.Println(rand.Intn(100))
    fmt.Println(rand.Intn(100))

	value := rand.Int()
	fmt.Println(value)
	valueb := rand.Int()
	fmt.Println(valueb)

	    // Call Seed, using current nanoseconds.
    rand.Seed(int64(time.Now().Nanosecond()))
    // Random int will be different each program execution.
    valuec := rand.Intn(2)
    fmt.Println(valuec)

}