// Modo de duas goroutines se comunicarem e sincronizagem sua execução

package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" //Usado para enviar e receber mensagem pelo canal. Neste caso, c recebendo "ping" como canal
	}
}

func imprimir(c chan string) { //chan = Palavra reservada
	for {
		msg := <- c //Neste caso, msg recebendo "c" que pos sua vez recebe "ping"
		fmt.Println(msg)
		time.Sleep(time.Second * 1) //Definido soneca de 1 seg entre pings
	}
}

func main() {

	var c chan string = make(chan string) // Faz assim porque o C vai receber uma string, que no caso é "ping". No momento, tá vazio

	go pingar(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)

}