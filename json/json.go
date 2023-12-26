//Lendo arquivos Json
// 1. json.Marshal
// 2. json.Unmarshal
// 3. json.NewEncoder
// 4. json.NewDecoder

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	Users []User `json:"usuarios"`
}

// Estrutura para representar os dados do usuário
type User struct {
	Nome     string   `json:"nome"`
	Tipo     string   `json:"Tipo"`
	Idade    int      `json:"idade"`
	Social   Social   `json:"Social"`
	Endereco Endereco `json:"endereco"`
}

// Estrutura para representar os dados de mídia social
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

// Estrutura para representar os dados de endereço
type Endereco struct {
	Rua    string `json:"rua"`
	Cidade string `json:"cidade"`
	Estado string `json:"estado"`
	Cep    string `json:"cep"`
}

func main() {

	//Abrindo arquivo presente em mesma pasta para leitura
	jsonFile, err := os.Open("usuarios.json")
	if err != nil {
		fmt.Println("Erro encontrado na abertura do arquivo:", err)
	}
	defer jsonFile.Close()
	fmt.Println("Arquivo Aberto com sucesso")

	// Decodificando o conteúdo do arquivo JSON em uma estrutura de dados
	// var data map[string][]Usuario
	// decoder := json.NewDecoder(jsonFile)
	// if err := decoder.Decode(&data); err != nil {
	// 	fmt.Println("Erro ao decodificar o JSON:", err)
	// 	return
	// }

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Erro encontrado na leitura do arquivo:", err)
	}

	var usuarios Users

	// Desempacotou o arquivo Json
	// Pegou os dados lidos através da linha 62 com o ReadAll
	// Adicionou essas informações diretamente na variável usuarios
	err = json.Unmarshal(byteValue, &usuarios)
	if err != nil {
		fmt.Println("Erro encontrado no desempacotamento:", err)
	}

	// Exibindo cada usuário
	for i := 0; i < len(usuarios.Users); i++ {
		fmt.Println("Usuário nº", i+1)
		fmt.Println("\tNome:", usuarios.Users[i].Nome)
		fmt.Println("\tTipo:", usuarios.Users[i].Tipo)
		fmt.Println("\tIdade:", usuarios.Users[i].Idade)
		fmt.Println("\tFacebook:", usuarios.Users[i].Social.Facebook)
		fmt.Println("\tTwitter:", usuarios.Users[i].Social.Twitter)
		fmt.Println("\tRua:", usuarios.Users[i].Endereco.Rua)
		fmt.Println("\tCidade:", usuarios.Users[i].Endereco.Cidade)
		fmt.Println("\tEstado:", usuarios.Users[i].Endereco.Estado)
		fmt.Println("\tCep:", usuarios.Users[i].Endereco.Cep)
	}
}
