package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// metodo esta obrigatoriamente em alguma estrutura
// u => variável
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do %s no banco de dados ", u.nome)
}

func (u usuario) maiordeIdade() bool {
	return u.idade >= 18
}

// quando for fazer alguma alteração vc terá de usar o * antes do campo denominado(usuario ue no caso é o nome da estrutura do struct acima)
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Luis", 21}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Vicentina", 61}
	usuario2.salvar()

	maiordeIdade := usuario2.maiordeIdade()
	fmt.Println(maiordeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
