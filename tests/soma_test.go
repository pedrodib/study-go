package main

import "testing"

// Padrão Triple A - AAA
// A - Arrange (Preparar)
// A - Act (Rodar)
// A - Assert (Verificar as asserções)

func TestShoulSumCorrect(t *testing.T) {
	teste := soma(3, 2, 1) // Arrange

	resultado := 6 // Act

	// Assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado", teste)
	}
}

func TestShoulMultiplyCorrect(t *testing.T) {
	teste := multiplica(10, 10)

	resultado := 100

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado", teste)
	}
}
