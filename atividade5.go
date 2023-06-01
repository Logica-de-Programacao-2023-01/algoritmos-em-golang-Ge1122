 Faça um algoritmo que leia a idade de uma pessoa em anos e mostre a idade em dias.

import "fmt"

func main() {
	var x int
	fmt.Print("Quantos anos você tem? ")
	fmt.Scan(&x)

	var y int
	fmt.Print("Quantos meses se passaram desde seu ultimo aniversário? ")
	fmt.Scan(&y)

	var z int = x*365 + y*30
	fmt.Print("Você tem ", z, " dias de vida")

}
