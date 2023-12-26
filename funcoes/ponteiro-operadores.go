package main

import "fmt"

// Armazenar um valor na memória, mas não é o valor propriamente escrito
//

func inicial (a *int) {
	*a = 5
}
func main() {
	
	x := 5
	p := &x
	*p = 30
	fmt.Println(x)

	inicial(&x)
	fmt.Println("Retorno de x: ", x)

}