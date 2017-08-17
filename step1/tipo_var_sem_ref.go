package main

import "fmt"

func main() {

	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128

	fmt.Printf("i type %T\n", i)
	fmt.Printf("f type %T\n", f)
	fmt.Printf("g type %T\n", g)

}
