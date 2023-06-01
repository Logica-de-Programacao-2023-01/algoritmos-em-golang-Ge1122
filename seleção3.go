package main

import "fmt"

func main() {
	var x int
	fmt.Print("Digite um número: ")
	fmt.Scan(&x)
	if x%2 == 0 {
		fmt.Print("O número é par")
	} else {
		fmt.Print("O número é impar")
	}

}
