package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo!") // goroutine => EXECUTA A LINHA 9 MAS NÃO ESPERA PRA SEGUIR O PROGRAMA(logo esta escrevendo os dois ao mesmo tempo)
	escrever("Programando em Go!") // ou seja esse comando Go ele não espera o print de cima acabar para começar a rodar o outro
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}