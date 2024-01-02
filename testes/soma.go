package main

import (
	"fmt"
)

func main() {

	x := Soma(1, 2, 3, 4, 5, 6, 7)
	y := Multiplica(10, 10)
	fmt.Println(x, y)

}

func Soma(i ...int) int {

	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func Multiplica(i ...int) int {
	
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}