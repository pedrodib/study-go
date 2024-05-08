package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func cabecalhos(w http.ResponseWriter, req *http.Request) {
	for nome, cabecalhos := range req.Header {
		for _, c := range cabecalhos {
			fmt.Fprintf(w, "%v: %v\n", nome, c)
		}
	}
}

func main() {
	http.HandleFunc("/ola", hello)
	http.HandleFunc("/cabecalho", cabecalhos)

	http.ListenAndServe(":8090", nil)
}
