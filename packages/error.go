package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string { // Por ter o metodo Error, ele ja diz que é uma implementação de error
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"O arquivo do sistema desapareceu",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
