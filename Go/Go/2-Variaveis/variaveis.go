package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var(
		variavel3 string = "lalalala"
		variavel4 string = "lalalala"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "Variavel5", "Variavel6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)
	
	
}