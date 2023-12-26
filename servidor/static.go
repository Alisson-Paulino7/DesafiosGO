//Servidor de serviço - Prestar algo ao usuário

package main

import (
	"fmt"
	"net/http"
	
)

func ola (w http.ResponseWriter, r *http.Request) {
	// As funções que servem como manipuladoras recebem a http.ResponseWriter 
	// e a http.Request como argumentos. O gravador de resposta é usado para preencher
	// a resposta HTTP. Aqui só vamos responder "Olá"
	fmt.Fprintf(w, "Olá\n")
}

func cabecalhos(w http.ResponseWriter, r *http.Request) {
	// Essas funções fazem algo um pouco mais sofisticado lendo todos os cabeçalhos de
	// solicitação HTTP e ecoando-os no corpo da resposta.
	for nome, cabecalhos := range r.Header {
		for _, c := range cabecalhos {
			fmt.Fprintf(w, "%v: %v\n", nome, c)

		}
	}
}

func main() {
	// Um manipulador é um objeto que implementa http.Handler.
	// Uma maneira comum de escrever ele é usar o http.HandlerFunc
	// Adaptado as funções com a assinatura própria.
	// Registramos nossos manipuladores nas rotas do servidor usando a 
	// http.HandlerFunc, rota de função que ele deve chamar de "/Olá" e "/cabecalho"
	// Ele configura o roteador padrão no pacote net/http e recebe uma função
	// como argumento. ("".ola) e ("", cabecalhos)
	http.HandleFunc("/ola", ola)
	http.HandleFunc("/cabecalho", cabecalhos)
	// Finalmente, chamamos o ListenAndServe com a porta ":8090"
	// E um manipulador nil para que seja usado o roteador padrão que acabamos de configurar
	// Execute o servidor em segundo plano
	// Acesse: http://localhost:8090/ola ou 
	// http://localhost:8090/cabecalho
	http.ListenAndServe(":8090", nil)
}