package hello

import (
	"fmt"
)

func a(x float64) string {
	fmt.Println("aqui", x)
	if x < 0 {
		return a(-x) + "i"
	}
	
	return fmt.Sprint(x)
}

func main() {
	fmt.Println(a(2), a(-4))
}
