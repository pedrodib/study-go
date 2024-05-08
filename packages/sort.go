package main

import (
	"fmt"
	"sort"
)

type Dados struct {
	Name string
	Age  int
}

type NameTo []Dados

func (ps NameTo) Len() int {
	return len(ps)
}

func (ps NameTo) Less(i, j int) bool { // Item na Posição i é menor que o j
	return ps[i].Name < ps[j].Name
}

func (ps NameTo) Swap(i, j int) { // Troca os Itens de Posição
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	kids := []Dados{
		{"Julia", 5},
		{"João", 10},
	}

	sort.Sort(NameTo(kids))
	fmt.Println(kids)
}
