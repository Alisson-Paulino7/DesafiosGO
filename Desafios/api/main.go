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
	_ = json.NewDecoder(r.Body).Decode(&cliente)
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
	response := []string{"Usuário não encontrado"}
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range client.Users {
		if item.Id == params["id"] {
			client.Users = append(client.Users[:i], client.Users[i+1:]...)
			break
		}
	}
}

func main() {
	
	client.Users = append(client.Users, Usuario{Id: "1", Nome: "John", Email: "Doe", Telefone: "88888888", Endereço: &Localização{Logradouro: "Rua das Malvas", Cidade: "Barbalha"}})
	client.Users = append(client.Users, Usuario{Id: "2", Nome: "Koko", Email: "Doe", Telefone: "99999999", Endereço: &Localização{Logradouro: "Rua das Fezes", Cidade: "Juazeiro"}})

	Carregarrouter()
}

func Carregarrouter() {
	router := mux.NewRouter()

	router.HandleFunc("/cliente/{id}", CreateUser).Methods("POST")
	router.HandleFunc("/cliente", ReadAll).Methods("GET")
	router.HandleFunc("/cliente/{id}", ReadOne).Methods("GET")
	router.HandleFunc("/cliente/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
