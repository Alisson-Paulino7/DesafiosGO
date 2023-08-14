package main

import "fmt"


func main() {

	for x := 1; x <= 100; x ++ {
		if x % 3 == 0{
			fmt.Printf("\nO valor %d é divisível por 3", x)

		}
	}

}