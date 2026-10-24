package main

import "fmt"

// CasoSucesso - caminho feliz, sem erro.
func CasoSucesso() {
	nome, err := BuscarUsuario("42")
	if err != nil {
		fmt.Println("  erro inesperado:", err)
		return
	}

	idade, err := ValidarIdade("25")
	if err != nil {
		fmt.Println("  erro inesperado:", err)
		return
	}

	fmt.Printf("  olá, %s (%d anos)\n", nome, idade)
}
