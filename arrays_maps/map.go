package main

import "fmt"

func main() {
	x := make(map[string]int)

	x["key"] = 10

	fmt.Println(x["key"])

	// --
	x2 := make(map[int]int)

	x2[0] = 20
	x2[1] = 30
	fmt.Println(x2[0], x2[1])

	// --
	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "Lítio"
	fmt.Println(elemento["Li"])

}
