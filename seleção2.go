package main

import "fmt"

func main() {
	var x int
	fmt.Print("Escreva o primeiro número ")
	fmt.Scan(&x)

	var y int
	fmt.Print("Escreva o segundo número ")
	fmt.Scan(&y)

	var z int
	fmt.Print("Escreva o terceiro número ")
	fmt.Scan(&z)

	if x > y {
		fmt.Print("O menor número é ", y)
	} else if y > z {
		fmt.Print("O menor número é ", z)
	} else {
		fmt.Print("O menor número é ", x)
	}
	if x == y {
		fmt.Print(" e os números são iguais ")
	} else if x == z {
		fmt.Print(" e os números são iguais")
	}

}
