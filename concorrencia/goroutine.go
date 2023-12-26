// Função capaz de executar de modo concorrente com outras funções
// Palavra reserva para goroutine: go func() {

package main

import "fmt"

func f (n int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d : %d\n", n, i)
	}
}

func g (s int)  {
	println(s + 15)
}

func main() {
	
	go f(0)
	go g(1)
	var escreva string
	fmt.Scanln(&escreva) //Gera mais uma linha abaixo do código
}