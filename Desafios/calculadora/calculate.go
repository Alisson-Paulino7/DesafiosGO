package main

import "fmt"

func Sum(i ...int) int {
	soma := 0
	for _, v := range i {
		soma += v
	}
	return soma
}

func Multiply(i ...int) int {
	mult := 1
	for _, v := range i {
		mult *= v
	}
	return mult
}

func Subtract(i ...int) int {
	sub := i[0]
	for _, v := range i[1:] {
		sub -= v
	}	
	return sub
}

func Divide(i ...int) float64 {
	div := float64(i[0])
	for _, v := range i[1:] {
		if v != 0 {
			div /= float64(v)
		} else {
			fmt.Println("Não é possível dividir por 0")
			return 0
		}
	}
	return float64(div)
}

func main() {
	sum := Sum(1, 2, 3, 4, 5)
	mult := Multiply(0, 1, 2, 3)
	sub := Subtract(1, 2, 3, 4, 5)
	div := Divide(1, 2, 3, 4, 5)


	fmt.Println("\n*** Calculadora ***")
	fmt.Printf("Soma: %d\n", sum)
	fmt.Printf("Multiplicação: %d\n", mult)
	fmt.Printf("Subtração: %d\n", sub)
	fmt.Printf("Divisão: %.4f", div)

}
