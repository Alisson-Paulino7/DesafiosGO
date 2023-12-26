// Servidor relacionado ao que o usuário precisa
// Nesse código vamos emitir uma solicitação de um cliente a um servidor http

package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// Emitir uma solicitação HTTP GET para um servidor, chamando seu método após
	// Criar um client server objeto

	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Ele usa o objeto que tem configurações padrão úteis, como predefinição
	// Imprimir o status da resposta HTTP.
	fmt.Println("Qual é o status: ", resp.Status)

	// Imprimir as primeiras 5 linhas do corpo da resposta
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err()
		err != nil {
		panic(err)
		}
	
}