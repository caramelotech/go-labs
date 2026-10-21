package main

import "fmt"

// BuscarUsuario simula uma busca.
// Em falha, devolve erro sentinela envolvido com %w.
func BuscarUsuario(id string) (string, error) {
	if id != "42" {
		return "", fmt.Errorf("buscar usuario %q: %w", id, ErrNaoEncontrado)
	}
	return "caramelo", nil
}
