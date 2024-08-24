package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(10, 20)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d, esperado %d", total, 30)
	}
}
