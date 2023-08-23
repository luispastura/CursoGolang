package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("string")
	generica(1)
	generica(true)

	fmt.Println(1, 2, "string", false, true, float64(123456))
}