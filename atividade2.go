package main

import "fmt"

func main() {

	var numero int
	fmt.Print("Escreva o numero que deseja descobrir o dobro, triplo e quadruplo")
	fmt.Scan(&numero)

	var dobro int = numero * 2
	var triplo int = numero * 3
	var quadruplo int = numero * 4

	fmt.Print("Dobro desse numero é  ", dobro, ", triplo desse numero é  ", triplo, ", quadruplo desse numero é  ", quadruplo)

}
