package main

import "fmt"

const ebulitionF float64 = 212.0

func main() {

	var tempF float64 = ebulitionF
	tempC := (tempF - 32) * 5 / 9 // Gopher ou Operador curto, serve para criar variavel já atribuindo. Só funciona dentro de Blocks

	fmt.Printf("A Temperatura de Ebulição da Água em ºF = %g \nA Temperatura de Ebulição da Água em ºC = %g", tempF, tempC)
	fmt.Printf("\n%T - %T", tempF, tempC) // Get variable Type

}
