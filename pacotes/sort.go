package main

import (
	"fmt"
	"sort"
)

type Dados struct {
	Nome string
	Idade int
}

//Interface abaixo
type ParaNome []Dados //Chamando a estrutura "Dados"

func (p ParaNome) Len() int { 
	return len(p)

}

func (p ParaNome) Less(i, j int) bool { //Verificar se item na posição i é menor que o item na posição j.
	return p[i].Nome < p[j].Nome
}

func (p ParaNome) Swap(i, j int) { //Troca os itens
	p[i], p[j] = p[j], p[i]
}

func main() {

	crianças := []Dados{
		{"Julia", 5},
		{"Joao", 10},
	} // Atribuindo vários objetos na interface

	sort.Sort(ParaNome(crianças))

	fmt.Println(crianças)

}