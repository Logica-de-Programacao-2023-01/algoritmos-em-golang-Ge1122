package main

import (
	"fmt"
)

func main() {
	var x int
	var y int
	fmt.Print("Digite um numero: ")
	fmt.Scan(&x)
	fmt.Print("Digite outro numero: ")
	fmt.Scan(&y)
	if x > 0 && y > 0 {
		fmt.Print(x * y)
	} else {
		fmt.Print(x + y)
	}

}
