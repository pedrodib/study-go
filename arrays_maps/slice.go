package main

import "fmt"

func main() {
	// Jeito de declarar
	fatia := make([]float64, 5)

	fmt.Println(fatia)

	// Outro jeito de declarar usando um array
	arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	fatia2 := arr[2:5]

	fmt.Println(fatia2)

	// Append to Slice
	fatia3 := append(fatia2, 6, 7, 3)
	fmt.Println(fatia3, fatia2)

	// Copy Slice
	fatia4 := make([]float64, 0)
	copy(fatia4, fatia3)
	fmt.Println(fatia4)

}
