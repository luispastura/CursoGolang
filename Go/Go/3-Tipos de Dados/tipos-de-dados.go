package main

import (
	"errors"
	"fmt"
)

func main() {
	//Números inteiros
	var numero int = 5551587575757852
	fmt.Println(numero)

	numero2 := 10057689823951
	fmt.Println(numero2)

	var numero3 uint32 = 10000  // uint = int sem sinal(unsigned)
	fmt.Println(numero3)


	// alias(numeros que representam caracteres)
	// int32 = rune
	var numero4 rune = 100000
	fmt.Println(numero4)


	// uint8 = byte(byte = 8bits)
	var numero5 byte = 123
	fmt.Println(numero5)


	//Números reais(suportam pontos ou números quebrados),além do que O FLOAT não funciona que nem o int ou uint que o valor vai de acordo com a sua máquina!!!!!
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 678.90
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)


	//Strings(Texto)
	var str string = "Adoro Go"
	fmt.Println(str)

	str2 := "Curso Udemy"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)


	// Valor 0(É o valor que vais er atribuído a uma variável quando vc não inicializa ela com valor)
	var texto int16
	fmt.Println(texto) // No Go todo tipo de dado tem seu valor "0"||Para STRING é vazio e para INT é valor 0

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro) //<nil> = 0(ponteiros,interface....)

	//Atribuindo valor ao erro
	var erro1 error = errors.New("GAY PRIDE FLAG")
	fmt.Println(erro1)




	
}