package main

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

type Client struct {
	Users []Usuario
}

type Usuario struct {
	Id       string       `json:"id"`
	Nome     string       `json:"nome"`
	Email    string       `json:"email"`
	Telefone string       `json:"telefone"`
	Endereço *Localização `json:"endereço"`
}

type Localização struct {
	Logradouro string `json:"logradouro"`
	Cidade     string `json:"cidade"`
}

var client Client

func CreateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var cliente Usuario

	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	cliente.Id = params["id"]
	client.Users = append(client.Users, cliente)
	json.NewEncoder(w).Encode(client.Users)
}

func ReadAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(client.Users)
}

func ReadOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range client.Users {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// response := []string{"Usuário não encontrado"}
	// json.NewEncoder(w).Encode(response)
	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range client.Users {
		if item.Id == params["id"] {
			client.Users = append(client.Users[:i], client.Users[i+1:]...)
			break
		}
	}
	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var newUser Usuario
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}
	for i, item := range client.Users {
		if item.Id == params["id"] {
			// item.Nome = newUser.Nome
			// item.Email = newUser.Email
			// item.Telefone = newUser.Telefone
			// item.Endereço.Logradouro = newUser.Endereço.Logradouro
			// item.Endereço.Cidade = newUser.Endereço.Cidade
			
			client.Users[i] = newUser

			json.NewEncoder(w).Encode(newUser)
			w.Write([]byte("Usuário atualizado com sucesso"))
			return
		}
	}
	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func main() {
	
	client.Users = append(client.Users, Usuario{Id: "1", Nome: "John", Email: "Doe", Telefone: "88888888", Endereço: &Localização{Logradouro: "Rua das Malvas", Cidade: "Barbalha"}})
	client.Users = append(client.Users, Usuario{Id: "2", Nome: "Koko", Email: "Doe", Telefone: "99999999", Endereço: &Localização{Logradouro: "Rua das Fezes", Cidade: "Juazeiro"}})

	CarregarRouter()
}

func CarregarRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/cliente/{id}", CreateUser).Methods("POST")
	router.HandleFunc("/cliente", ReadAll).Methods("GET")
	router.HandleFunc("/cliente/{id}", ReadOne).Methods("GET")
	router.HandleFunc("/cliente/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
