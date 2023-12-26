package main

import (
	"fmt"
	"crypto/sha1"
)

func main() {


	p := "Código crypton com Go"

	h := sha1.New()
	h.Write([]byte(p))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}

