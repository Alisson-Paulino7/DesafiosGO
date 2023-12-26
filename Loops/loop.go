package main

import "fmt"

func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Printf("\nHoras: %d", horas)
		
		for min := 0; min < 60; min++ {
			fmt.Printf("\nMinutos: %d atuais", min)
		}
	}
}