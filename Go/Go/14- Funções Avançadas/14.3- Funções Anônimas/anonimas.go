package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando aqui pra falar que você é FODA IRMÃO!!!!!!")

	fmt.Println(retorno)

	func(numero int){
		fmt.Println(numero)
	}(15)
}