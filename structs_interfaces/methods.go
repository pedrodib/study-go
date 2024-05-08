package main

import "fmt"

func main() {
	retanguloVar := retangulo{comprimento: 10, altura: 5}

	fmt.Println("Area:", retanguloVar.area())
	fmt.Println("Area:", retanguloVar.perimetro())
}

type retangulo struct {
	comprimento, altura int
}

func (r *retangulo) area() int {
	return r.comprimento * r.altura
}

func (r retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
}
