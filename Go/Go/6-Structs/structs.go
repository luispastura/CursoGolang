package main

import "fmt"

type usuario struct{
	nome string
	idade uint8
	endereco endereco
}

type endereco struct{
	endereco string
	numero uint8
}

func main(){
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Luis"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua os Macacos", 0}

	usuario2 := usuario{"Caio", 20, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Joao"}
	fmt.Println(usuario3)

}