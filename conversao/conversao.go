package main

import "fmt"

var a int = 20
var b string = "Nome"


func main () {

	var c = float64(a)

	fmt.Printf("O valor de C é: %v de tipo %T. E o valor de B é: '%s' do tipo %T", c, c, b, b)



}