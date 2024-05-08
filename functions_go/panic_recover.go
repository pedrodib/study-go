//panic: Erro do programador/erro execução tempo -> Tipo Exception
//recover: Ela interrompe o panic -> Tipo try catch

package main

import "fmt"

func main() {
	defer func() {
		x := recover()
		fmt.Println("Try Catch:", x)
	}()

	panic("Pânico")
}
