package main

import (
	"fmt"
	"hash/crc32"
)

func main() {

	hash := crc32.NewIEEE()

	hash.Write([]byte("código com pacote hash"))

	hashed := hash.Sum32()

	fmt.Println(hashed)
}
