package main

import (
	"fmt"
)

func main(){
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alternada"
	fmt.Println(slice2)



	// ARRAYS INTERNOS(Mais título de curiosidade)
	//Slice NÃO TEM TAMANHO LIMITE,á o array SIM
	fmt.Println("----------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	//Estourando capacidade do slice(quando ele ve que vai ultrapassar o seu limite ele pega o valor e dobra)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade


	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4)) 
	fmt.Println(cap(slice4)) 

}