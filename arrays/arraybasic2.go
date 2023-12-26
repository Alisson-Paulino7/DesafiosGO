package main

import "fmt"


func main () {

	x := []string {0:"Oi", 1:"Olá"}
/*
	x[0] = 5.3
	x[1] = 8
	x[2] = 4.2
	x[3] = 2.1
	x[4] = 7.8

	total  := 0.0
	//media := total / 5

	for y := 0; y < len(x); y++{
		total += x[y]
	}
	fmt.Println("Média dos alunos é: ", total/float64(len(x))) 

	for z := 0; z <len(x); z++ {
		fmt.Println(x[z])

	}
	*/
	fmt.Println(x)
	for chave, valor := range x {
		fmt.Printf("\n%d -> %s", chave, valor)
	}
	for d := 0; d < len(x); d++ {
		fmt.Printf("\nValor de %d: %s", x[d], x[d])
	}
}