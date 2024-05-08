package main

import "fmt"

func main() {

	// forfunc()
	// while()
	infiniteBreak()
}

func forfunc() {
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		} else if i == 3 {
			break
		}

		fmt.Println(i)
	}
}

func while() {
	i := 0

	for i <= 20 {
		fmt.Println("i é menor que 20.")
		i++
	}
}

func infiniteBreak() {
	i := 0

	for {
		fmt.Println(i)

		if i == 20 {
			break
		}

		i++
	}
}
