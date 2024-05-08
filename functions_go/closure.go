// Closure: Capacidade de uma função chamar uma variavel que está em outra função

package main

import "fmt"

func main() {
	x := 0

	numero := func() int {
		x++
		return x
	}

	fmt.Println(numero(), numero())
}
