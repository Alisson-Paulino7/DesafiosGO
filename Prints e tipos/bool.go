package main

import "fmt"

/*
func main() {

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)


}
*/

func main () {

	fmt.Println(len("Hello World!")) // Usando a biblioteca fmt pra imprimir o tamanho de caracteres
	fmt.Println("Hello World!"[2]) // Qual a letra no índice 2
	fmt.Println("Hello" + "World!") // Concatenação
	var nome = "Alisson" // Declarando variável
	var idade = 27
	var versao = 7.7
	fmt.Println("Meu nome é:", nome, "e minha idade é:", idade)
	fmt.Println("Atualmente estou na versão:", versao)

}