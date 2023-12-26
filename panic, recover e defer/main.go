package main

import "fmt"

func divide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Erro recuperado: ", r)
		}
	}()
		result := a / b
		return result
}

func main() {
	fmt.Println(divide(1, 0))
	fmt.Println(divide(1, 1))
	fmt.Println("Fim da execução")
}