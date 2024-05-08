package main

import "fmt"

func main() {
	x := 4

	if x == 2 || x == 3 {
		fmt.Println("Sim, x é 2 ou 3")
	} else {
		fmt.Println("Sim, x não é 2 ou 3")
	}
}
