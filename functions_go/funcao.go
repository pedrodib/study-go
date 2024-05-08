// function / procedure / sub routine

package main

import "fmt"

func calculaMedia(lista []float64) float64 {
	total := 0.0

	for _, valor := range lista { // Foreach
		total += valor
	}

	return total / float64(len(lista))
}

func main() {
	lista := []float64{98, 93, 77, 82, 83}

	fmt.Println(calculaMedia(lista))
}
