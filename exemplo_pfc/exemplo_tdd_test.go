package main

import "testing"

func TestExemplo(t *testing.T) {
	resultado := Exemplo()
	esperado := "String correta"

	if resultado != esperado {
		t.Errorf("O resultado é %q mas o esperado era %q", resultado, esperado)
	}
}