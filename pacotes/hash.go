
// Aceitar um conjunto de dados que reduz a um tamanho fixo menor
// Hashes são usados principalmente para buscar DADOS e DETECTAR ALTERAÇÕES
// Em GO são divididas em CRIPTOGRAFADAS e NÃO CRIPTOGRAFADAS
// Lista das NÃO CRIP.: adler32. crc32, crc64
// Vamos criar uma HASH e escrever dados nele pra converter em números

package main

import (
	"fmt"
	"hash/crc32"
)

func main() {

	//Criar a hash

	h := crc32.NewIEEE()
	h.Write([]byte("O código com pacote Hash"))
	// calcular o hash
	v := h.Sum32()
	fmt.Println(v)

}