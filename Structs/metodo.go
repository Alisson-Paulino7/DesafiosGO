package main

import "fmt"


type retangulo struct {
	comprimento, altura int
}

// Este método `área` possui um tipo de receptor de retangulo

func (r retangulo) area () int {
	return r.comprimento * r.altura
}

func (r retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
}

func (r retangulo) getComprimento() int{
	return r.comprimento
}

func (r retangulo) getAltura() int {
	return r.altura
}


func main() {

	r := retangulo{comprimento: 50, altura: 25}


	fmt.Println("Área do retângulo: ", r.area())
	fmt.Println("Perímetro do retângulo: ", r.perimetro())
	fmt.Printf("Comprimento: %d , Altura: %d ", r.getComprimento(), r.getAltura())

}