package main 

import "fmt"


func main() {

	for y := 1; y <= 100; y ++{
		if y % 3 == 0 && y % 5 == 0 {
			fmt.Println("PIN e PAN")
		} else if y % 3 == 0 {
			fmt.Println("PIN")
		} else if y % 5 == 0 {
			fmt.Println("PAN")
		}

	} 

}