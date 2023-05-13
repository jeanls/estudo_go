package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

func (ret retangulo) area() float64 {
	return ret.altura * ret.largura
}

func (cir circulo) area() float64 {
	return (cir.raio * cir.raio) * math.Pi
}

type circulo struct {
	raio float64
}

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Println("A area é ", f.area())
}

func main() {
	ret := retangulo{10, 15}
	escreverArea(ret)

	cir := circulo{10}
	escreverArea(cir)
}
