//Pacote que escreve na tela uma frase sem a necessidade de utilizar o Fmt
//Além disso, permite tratamentodo do erro na própria declaração

package main

import (
	"io"
	"os"
	"log"
)

func main() {
	if _, err := io.WriteString(os.Stdout, "Hello, world!"); err != nil {
		log.Fatal(err)
	}
}