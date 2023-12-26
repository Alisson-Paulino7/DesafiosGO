package main

import "fmt"

func main () {


	for i := 1; i <= 10; i++ {
		fmt.Println(i)

	}
	for x := 0; x <= 10; x++ {
		if x % 2 == 0 {
			fmt.Printf("\nO numero %d é par!", x)
		} else {
			fmt.Printf("\nO numero %d é ímpar!", x)
		}
	}
	for u := 0; u <= 7; u++ {
		switch u {
			case 0: fmt.Println("\nZero")
			case 1: fmt.Println("Um")
			case 2: fmt.Println("Dois")
			case 3: fmt.Println("Três")
			case 4: fmt.Println("Quatro")
			case 5: fmt.Println("Cinco")
			case 6: fmt.Println("Seis")
			case 7: fmt.Println("Sete")
		}
	}
	
	z := 4

	if (z == 2 || z == 3) {
		fmt.Println("Sim, Z é 2 ou 3")
	} else {
		fmt.Println("Não, Z não é 2 nem 3")
	}

}


