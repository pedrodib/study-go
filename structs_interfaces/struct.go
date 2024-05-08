// Coleções de Campos -> Tipo uma classe
package main

import "fmt"

type pessoa struct {
	name string
	age  int
}

func main() {
	pedro := pessoa{"Pedro", 27}
	fmt.Println(pedro)

	fmt.Println(pessoa{"Victória", 29})
}
