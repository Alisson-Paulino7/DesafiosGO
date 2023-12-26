package main

import "fmt"

func fatorial (x int) int {

	if x == 0 {
		return 1
	}
	return x * fatorial(x-1)
	
}


func contagem (y int) int {
	if y == 0 {
		return 1
	}
	return y + contagem(y-1)
}

func main() {

	//fmt.Println(fatorial(5))
	fmt.Println(contagem(7))

}