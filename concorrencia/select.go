// Funciona como um Switch para canais
// Permite que você aguarde operações de vários camais
// Combinar goroutines e canais com select é um recurso poderoso do GO
// Para nosso exemplo, selecionaremos em dois canais

package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func () { // Recebemos os valores "um" e depois "dois" conforme esperado
		time.Sleep(1 * time.Second) // Tempo de execução de ~2 segundos
		//pois o 1 e o 2 segundos são Sleeps executados simultaneamente

		c1 <- "um"
}()
	go func ()  {
		time.Sleep(2 * time.Second) 
		c2 <- "dois"
	}()

	for i := 0; i < 2; i ++ {
		select { //Usaremos select para aguardar esses 
			//dois valores simultaneamente, imprimindo cada um à medida que chegam

		case msg1 := <-c1: 
			fmt.Println("Receba", msg1)
		case msg2 := <-c2: 
			fmt.Println("Receba", msg2)
	}

}

}