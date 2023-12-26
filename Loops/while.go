package main

import "fmt"

func main() {

	i := 0

	for i < 20 {
		fmt.Printf("%d é menor que 20\n", i)
		i ++
	} 
	if i == 20 {
		fmt.Printf("%d agora é igual a 20", i)
	}
}