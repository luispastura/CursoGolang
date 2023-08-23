package main

import "fmt"

//	DEFER É ADIAR

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Recalculando, aguarde.")
	fmt.Println("entrando na sua função para verificar se o aluno está aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 7))
}
