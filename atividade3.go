package main

import "fmt"

func main() {

	var peso float64
	fmt.Print("Qual o seu peso? ")
	fmt.Scan(&peso)

	var altura float64
	fmt.Print("Qual a sua altura? ")
	fmt.Scan(&altura)

	var IMC float64 = peso / (altura * altura)
	fmt.Print("Seu índice de massas corporal é  ", IMC)

}
