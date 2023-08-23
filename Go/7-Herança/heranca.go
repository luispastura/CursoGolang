package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

// O termo pessoa dentro da struct "estudante" esta puxando os dados feitos na struct "pessoa" acima
type estudante struct {
	pessoa
	faculdade string
	curso     string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Luis", "Pastura", 21, 182}
	fmt.Println(p1)

	e1 := estudante{p1, "Ciencia de dados", "Ibmec"}
	fmt.Println(e1)
}
