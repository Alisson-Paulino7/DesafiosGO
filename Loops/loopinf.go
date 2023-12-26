package main

import "fmt"

func main() {
	
	for x := 0; x < 20; x++{
		if x == 3 {
			fmt.Println(x)
			break
		} 
		fmt.Println(x)
	}
}