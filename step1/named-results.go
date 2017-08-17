package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	w := y
	fmt.Print(w)
	return
}

func main() {
	fmt.Println(split(10))
}
