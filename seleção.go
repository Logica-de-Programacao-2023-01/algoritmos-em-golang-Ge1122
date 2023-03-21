package main

import "fmt"

func main() {
	var x int
	fmt.Print("Escreva um número inteiro ")
	fmt.Scan(&x)
	var y int
	fmt.Printf("Escreva outro número inteiro ")
	fmt.Scan(&y)

	if x > y {
		fmt.Print("x é maior que y ")
	} else if x < y {
		fmt.Print("y é maior que x")
	} else {
		fmt.Print("x é igual ao y")
	}
}
