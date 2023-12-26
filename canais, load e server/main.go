package main

import (
	"fmt"
//	"encoding/json"
//	"net/http"
	"time"
)

type Course struct {
	Name string
	Description string
	Price int
}

func (c Course) GetfullInfo() string {
	return fmt.Sprintf("Name: %s, Description: %s, Price: %d", c.Name, c.Description, c.Price)
}

func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	// rodar só 1 deles. Mas se botar o GO e mais vezes ele, rodará em paralelo d

	//Rodar a função 3x ao mesmo tempo.
	go counter()
	go counter()
	counter()

	// A maneira que a classe da API irá dar retorno
//	http.HandleFunc("/struct", home)
	// A porta utilizada será a 8080
//	http.ListenAndServe(":8080", nil)

}

/*
// Uma api simples que retorna a estrutura e seus atributos
// w é referência para o response e o r para o request
func home(w http.ResponseWriter, r *http.Request) {
	course := Course {
				 Name: "Golang", 
				 Description: "Golang Course", 
				 Price: 100}
	
	// Transformando os dados da classe como um json que pode ser consumida
	json.NewEncoder(w).Encode(course)
}
*/