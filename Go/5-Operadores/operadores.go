package main

import "fmt"

func main() {
	//Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 5
	multiplicacao := 10 * 6
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	/* Esta operação não se realiza pois o Go não deixa com que dois inteiros de classificações diferentes realizam a operação
	var numero1 int16 = 10
	var numero2 int32 = 15
	soma := numero1 + numero2
	fmt.Println(soma)
	*/

	var numero1 int16 = 10
	var numero2 int16 = 15
	soma1 := numero1 + numero2
	subtracao1 := numero2 - numero1
	fmt.Println(soma1, subtracao1)

	//Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2) // diferente

	//Operadores Lógicos
	fmt.Println("-----------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//Operadores Unarios
	fmt.Println("-----------------")
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 // numero = numero - 20

	numero *= 3 // numero = numero vezes 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)

	//Operador Ternário
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)

}
