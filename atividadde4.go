4: Faça um algoritmo que leia três números reais e mostre a média ponderada entre eles, com pesos 2, 3 e 5, respectivamente.

import "fmt"

func main() {
	var x float64
	fmt.Print("Escreva o primeiro número")
	fmt.Scan(&x)

	var y float64
	fmt.Print("Escreva o segundo número")
	fmt.Scan(&y)

	var z float64
	fmt.Print("Escreva o terceiro número")
	fmt.Scan(&z)

	var media float64 = (x*2 + y*3 + z*5) / 10
	fmt.Print("A média ponderada destes três números é  ", media)

}
