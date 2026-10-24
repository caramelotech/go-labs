package main

import (
	"errors"
	"fmt"
	"strconv"
)

// CasoEstruturadoComAs - errors.As extrai *ErrValidacao e seus campos.
func CasoEstruturadoComAs() {
	_, err := ValidarIdade("abc")
	if err == nil {
		fmt.Println("  erro inesperado: queria falha")
		return
	}

	// Uma camada a mais de %w (como faria um serviço).
	err = fmt.Errorf("criar perfil: %w", err)

	fmt.Println("  texto:", err)

	var ev *ErrValidacao
	if errors.As(err, &ev) {
		fmt.Printf("  errors.As: campo=%s valor=%q msg=%s\n", ev.Campo, ev.Valor, ev.Msg)
	}

	var numErr *strconv.NumError
	if errors.As(err, &numErr) {
		fmt.Printf("  errors.As: strconv.NumError num=%q\n", numErr.Num)
	}
}
