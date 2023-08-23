package main

import "fmt"

func main() {

	numero := -5

	if numero > 15 {
		fmt.Println("Você é maior de idade")
	} else {
		fmt.Println("Você é cuzao de idade")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Você é maior de idade")
	} else if numero < -10 {
		fmt.Println("Você é cuzao de idade -10")
	} else {
		fmt.Println("Você é menor de idade")
	}
}
