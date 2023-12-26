package main

import "fmt"


//O nome da fatia dentro da função não precisa especificamente ser igual ao nome da lista
// 

func media (notas [] float64) float64 { 
	total := 0.0

	for _, valor := range notas {
		total += valor
	}
	
	return total / float64(len(notas))
}

//O nome dos valores que representam os valores a receber não precisam ser iguais
//Importante sempre definir o tipo de arquivo que irá ser recebido e o tipo de arquivo que deseja como saída

func soma (num10, num20 float64) float64 {
	return float64(num10 / num20) //Como desejo que o retorno seja em float preciso converter o resultado para Float
}

func soma1 (num10, num20 int) int {
	return num10 + num20 //Nessa situação, não há necessidade de conversão, pois tanto o valor recebido quanto de saída são Int
}



func main() {


	notas := [] float64{6, 9, 4, 10, 5, 7}

	notas = append(notas, 3, 2, 10)

	num1 := 5
	num2 := 8

	res := soma1 (num1, num2)

	fmt.Println(notas)

	fmt.Printf("Valor da média é: %.2f", media(notas))
	fmt.Printf("\nValor da soma é: %.2f", soma(num1,num2))
	fmt.Printf("\nValor da soma é: %v\n", soma1(num1,num2))
	fmt.Printf("\nValor da soma é: %v\n", res)
	fmt.Printf("\nValor da soma é: %v\n", soma1(8,8))


	fatia := make([][] int, 3)
	fmt.Println(fatia)

}