package main

import "fmt"

var x = 0
var numero = func () int {
	x++
	return x
}

func newteste (a, b int) float64 {
	return float64(a) / float64(b)
}

func main() {

/*	
	z := 0
	numero1 := func() int {
		z++
		return z
	}
*/
	teste := func(a, b int) float64 {
		fmt.Println("oxente")
		return float64(a / b)
	}
		
/*	for y := 0; y < 5; y++ { 
	fmt.Println(numero(), "---")
	fmt.Println(numero1(), ">>>")
	}
*/

num1 := 3

num2 := 5
	fmt.Printf("Retorno: %.2f", teste(3, 5))
	fmt.Printf("\nRetorno: %.2f", newteste(num1, num2))

}