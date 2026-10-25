package main

import (
	"maps"
	"testing"
)

func TestContarPalavras(t *testing.T) {
	testes := []struct {
		nome  string
		texto string
		quer  map[string]int
	}{
		{
			nome:  "vazio",
			texto: "",
			quer:  map[string]int{},
		},
		{
			nome:  "uma palavra",
			texto: "Go",
			quer:  map[string]int{"go": 1},
		},
		{
			nome:  "ignora maiúsculas",
			texto: "Go go GO",
			quer:  map[string]int{"go": 3},
		},
		{
			nome:  "varias palavras",
			texto: "a b a c b a",
			quer:  map[string]int{"a": 3, "b": 2, "c": 1},
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			obtido := ContarPalavras(tt.texto)
			if !maps.Equal(obtido, tt.quer) {
				t.Fatalf("ContarPalavras(%q) = %v; quer %v", tt.texto, obtido, tt.quer)
			}
		})
	}
}
