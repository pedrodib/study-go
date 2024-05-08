package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	hash := sha1.New()

	hash.Write([]byte("código com pacote hash"))

	bs := hash.Sum([]byte{})

	fmt.Println(bs)
}
