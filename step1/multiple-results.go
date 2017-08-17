package main

import "fmt"

func swap(x string, y int) (string, int) {
	return x, y
}

func main() {
	a, b := swap("hello", 8)
	fmt.Println(a, b)
}
