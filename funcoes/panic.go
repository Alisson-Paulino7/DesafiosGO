package main

import ("fmt"
		"time"
)
/*
func divisao(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Erro recuperado: ", r)
		}
	}()
	result := a / b
	return result
}

func main()	{

	fmt.Println("Divisão: ", divisao(1, 0))
	fmt.Println("Divisão: ", divisao(5, 2))

	panic("Pânico")
	
}

*/


func Test (senha string) {
	senha1 := "teste"
	cont := 3

	for i := 0;;i++{
	fmt.Println("Digite a senha: ")
	fmt.Scanln(&senha)
		if senha == senha1 {
			fmt.Println("Senha correta")
			break
		} else {
			cont -= 1
			fmt.Println("Senha incorreta")
		}
		if cont == 0 {
			fmt.Println("Você errou a senha 3 vezes, tente novamente após 60 segundos")
			for x := 60; x >= 0; x-- {
				fmt.Printf("Tente novamente em: %d segundos\n", x)
			time.Sleep(time.Second)
		}
		
	}
}
}

func main() {

	Test("")

}


