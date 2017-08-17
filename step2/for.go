package main

import "fmt"

func main() {
	sum := 0
	var o int64
	for i := 0; i < 4; i++ {
		fmt.Println(sum)
		sum += i
		o = int64(i)
	}

	fmt.Println("aqui: ", sum, " Outra: ", o)
}
