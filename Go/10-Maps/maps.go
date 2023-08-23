package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Guilherme",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus1",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)
}
