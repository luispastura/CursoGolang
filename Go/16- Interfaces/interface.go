package main

import (
	"fmt"
	"math"
)

// so da pra passar assinaturas de metodo
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A sua área se refere a %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

// assinatura tem que ser igual e precisa retornar um floa64
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10, 20}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
