// Usar pacote time e fmt
// Utilizar o aprendido com concorência

package main

import (
	"fmt"
 	"time"
)

func pingar(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
		c <- "pong"
	}
}


func imprimir(c chan string) { 
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 2)
	}
}

func main() {

	var c chan string = make(chan string)

	// var p chan string = make(chan string)

	go pingar(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)

}

