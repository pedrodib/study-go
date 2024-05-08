package main

import "fmt"

func main() {
	for1()
	fmt.Println("")
	for2()
}

func for1() {

	i := 1

	for i <= 10 {
		fmt.Print(i, ",")

		i = i + 1
	}
}

func for2() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i, ",")
	}
}
