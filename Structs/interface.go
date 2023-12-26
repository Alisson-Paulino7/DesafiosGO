// Interfaces são coleções
// de métodos.

package main

import (
	"fmt"
	"math" //Biblioteca de matemática
)

// Aqui temos uma interface usada para formas geométricas.
type geometria interface {
	area() float64 //Método abstrato
	perimetro() float64 //Método abstrato
}

// área de um quadrado: lado² ou lado*lado
// perímetro = soma dos lados

type quadrado struct {
	lado float64
}
type círculo struct { //área círculo: (Pi)*raio² perímetro do círculo: 2*r*(pi)
	raio float64
}

type retang struct {
	comprimento float64
	altura float64
}

// Para usar uma interface em Go, só precisamos
// usar todos os métodos da interface. Aqui nós
// usaremos `geometria` em `quadrado`s.
func (q quadrado) area() float64 {
	return q.lado * q.lado
}
func (q quadrado) perimetro() float64 {
	return 4 * q.lado
}

// A implementação dos `círculo`s.
func (c círculo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c círculo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func (r retang) area() float64 {
	return r.comprimento * r.altura
}

func (r retang) perimetro() float64 {
	return 2*r.comprimento + 2*r.altura
}

// Se a variável tem um tipo interface, então nós podemos chamar
// métodos que estão na interface nomeada. Aqui temos uma
// função genérica `medida` tomando vantagem desse
// trabalho em qualquer `geometria`.
func medir(g geometria) {
	fmt.Printf("\nDado informado: %.2f", g) // Valor do lado do quadrado
	fmt.Printf("\nValor da área: %.2f", g.area()) // Valor da área do quadrado
	fmt.Printf("\nValor do perimetro: %.2f\n", g.perimetro()) // Perímetro do quadrado
}

func main() {
	q := quadrado{lado: 25}
	c := círculo{raio: 100}
	r := retang{comprimento: 50, altura: 35}

	// Os tipos de estrutura `círculo`, `retangulo` e `quadrado`
	// implementam a interface `geometria` então nós podemos usar
	// instâncias dessas estruturas como argumentos para `medir`.

	medir(q)
	medir(c)
	medir(r)
}