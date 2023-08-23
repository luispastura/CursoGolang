package main

import "fmt"

func diaSemana(numero int) string {
	switch numero {
	case 1:
		return "segunda"
	case 2:
		return "terça"
	case 3:
		return "quarta"
	case 4:
		return "quinta"
	case 5:
		return "sexta"
	default:
		return "NADA"
	}
}

func main() {
	dia := diaSemana(1)
	fmt.Println(dia)
}
