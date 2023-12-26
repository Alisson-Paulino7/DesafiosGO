package main

import "fmt"

func main() {

	x := make(map[string]int)

	x["Chave"] = 10
	x["Nome"] = 20
	x["5"] = 30


	y :=make(map[int]int)
	y[1] = 50

	z := map[string]string {
	"H" : "Hidrogênio",
	"HE": "Hélio",
	"O" : "Oxigênio",
}

	for chave, valor := range z {
		fmt.Printf("\nElemento: %s -> %s", chave, valor)
	}
	fmt.Println("\n",y)
	fmt.Println("\n",x)
	fmt.Println("\n",z)

	d := map[int]int {
		0 : 20,
		6 : 30,
		4 : 10,
		2 : 5,
		1 : 2,
		9 : 7,
	}
	fmt.Println(d)

	for chave, valor := range d{
		fmt.Printf("\n O índice é: %d, de valor %d", chave, valor)
	}


}