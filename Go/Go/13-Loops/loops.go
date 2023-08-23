package main

import "fmt"

func main() {
	//Enquanto essa condição de código for verdadeiro irá repetir um bloco de código
	i := 0
	for i < 10 {
		i++
		fmt.Println("BOM DIA BRUNO i")
	}

	fmt.Println("-----------------------")

	//Enquanto ela for menor que 10 vai incrementando ela e vai escrevendo na tela também
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
	}

	fmt.Println("-----------------------")

	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando j", j)
	}

	fmt.Println("-----------------------")
	
	//Usando for com o RANGE, que significa que você quer iterar em alguma coisa(array por exemplo)
	nomes := [3]string{"Luis", "Ana", "Julia"}

	// i = posição no iteravel("Joao" posição 0, "Ana"...)
	// nome = String(valor)
	// nomes = nome da variável criada
	for i, nome := range nomes {
		fmt.Println(i, nome)
	}

	fmt.Println("-----------------------")

	for _, letra := range "PALAVRA" {
		fmt.Println(string(letra))
	}
}
