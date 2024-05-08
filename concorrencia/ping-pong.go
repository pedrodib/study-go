package main

import (
	"fmt"
)

func main() {

	ping := make(chan string)
	pong := make(chan string)

	go func() {
		//time.Sleep(time.Second * 1)
		ping <- "Ping"
	}()

	go func() {
		//time.Sleep(time.Second * 2)
		pong <- "Pong"
	}()

	for {
		select {
		case msg1 := <-ping:
			fmt.Println(msg1)
		case msg2 := <-pong:
			fmt.Println(msg2)
		}
	}
}
