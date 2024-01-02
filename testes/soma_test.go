package main

import (
	"testing"
)

//Padrão triple A - AAA
//Arrange - Organizar (Preparar)
//Action - Acão (rodar)
//Assert - Verificar (Asserções)

func Test_If_Should_Sum_Correct(t *testing.T) {
	//Arrange
	teste := Soma(1, 2, 3, 4, 5, 6, 7)
	//Action
	Resultado := 28
	//Assert
	if teste != Resultado {
		t.Errorf("Resultado da soma é inválido: Resultado: %d | Esperado: %d", teste, Resultado)
	}
}

func Test_If_Should_Sum_Incorrect(t *testing.T) {
	teste := Soma(1, 2, 3, 4, 5, 6, 7)
	Resultado := 27
	if teste != Resultado {
		t.Errorf("Resultado da soma é inválido: Resultado: %d | Esperado: %d", teste, Resultado)
	}
}

func Test_If_Should_Mult_Correct(t *testing.T) {
	teste := Multiplica(10, 10)
	Resultado := 100
	if teste != Resultado {
		t.Errorf("Resultado da multiplicação é inválido: Resultado: %d | Esperado: %d", teste, Resultado)
	}
}

func Test_If_Should_Mult_Incorrect(t *testing.T) {
	teste := Multiplica(10, 10)
	Resultado := 101
	if teste != Resultado {
		t.Errorf("Resultado da multiplicação é inválido: Resultado: %d | Esperado: %d", teste, Resultado)
	}
}