package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(10, 10)

	if total != 20 {
		t.Errorf("Resuldado da soma é inválido: Resuldado %d. Esperado %d", total, 20)
	}
}
