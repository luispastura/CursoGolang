package main

import "fmt"

func soma(numeros ...int) int{
	fmt.Println(numeros)
	total := 0
	for _, numero := range numeros{
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int){
	for _, numero := range numeros{
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma( 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(totalDaSoma)

	escrever("Olá, mundo", 12, 38, 45, 76, 109)

}