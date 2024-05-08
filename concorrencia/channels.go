/*
Podemos pensar em channels como sendo uma espécie de tunel de comunidação entre goroutines, onde uma goroutine consegue enviar
informações para outra antes mesmo de terminar sua execução. Nesse mesmo cenário, a goroutine que recebe a informação,
ficaria pausada até as informações chegarem.
*/
package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) { // palavra reservada para Canal: chan
	for i := 0; ; i++ {
		c <- "ping" // usado para enviar e receber mensagem pelo canal
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pingar(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)
}
