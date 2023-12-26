//Para consultar http://pokeapi.co/api/v2/pokedex/kanto/
// Usar método GET

package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"os"
	"io/ioutil"
)

type response struct {
	Nome string `json:"nome"`
	Pokemon []pokemon `json:"pokemon_entries"`
}

type pokemon struct {
	Numero int `json:"entry_number"`
	Especie pokemonSpecies `json:"pokemon_species"`
}

type pokemonSpecies struct {
	Nome string `json:"name"`
}

func main() {

	// Leitura de dados em página Web
	resposta, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer resposta.Body.Close()

	// Só ele tem a função pra mapear tudo da resposta, mapeando o que tá no corpo HTTP
	respostaData, err := ioutil.ReadAll(resposta.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Converter de byte pra string
	// Byte impresso fica tudo bugado
	fmt.Println(string(respostaData))

	var responseObject response

	err = json.Unmarshal(respostaData, &responseObject)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(responseObject.Nome)
	fmt.Println(responseObject.Pokemon)

	for _, pokemon := range responseObject.Pokemon{
		fmt.Println(pokemon.Especie.Nome)
	}
}