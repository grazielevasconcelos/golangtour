package main

import "fmt"

var javaB = true

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
	teste()
}

func teste() {
	fmt.Println(javaB)
}