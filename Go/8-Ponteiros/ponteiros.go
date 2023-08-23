package main

import "fmt"

func main(){
	var variavel1 int = 66
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// Ponteiros são referencias de memoria
	var variavel3 int 
	var ponteiro *int 
	

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro) // * Usado para ver o valor dentro do endereço de memória(desreferenciaçao)

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
}