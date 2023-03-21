package main

import "fmt"

func main() {
	var x int
	fmt.Print("Escreva o número desejado ")
	fmt.Scan(&x)

	for i := 1; i <= 10; i++ {
		fmt.Println(i * x)
	}

}
