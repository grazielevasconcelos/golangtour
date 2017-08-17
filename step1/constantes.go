package main

import "fmt"

const Pi = 3.14

func main() {
	const world = "Earth"
	fmt.Println("Hello", world)
	fmt.Println("Hi", Pi, "day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
}
