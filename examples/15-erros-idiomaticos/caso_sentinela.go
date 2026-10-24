package main

import (
	"errors"
	"fmt"
)

// CasoSentinelaComIs - errors.Is encontra a sentinela sob o %w.
func CasoSentinelaComIs() {
	_, err := BuscarUsuario("99")
	if err == nil {
		fmt.Println("  erro inesperado: queria falha")
		return
	}

	fmt.Println("  texto:", err)

	if errors.Is(err, ErrNaoEncontrado) {
		fmt.Println("  errors.Is: é ErrNaoEncontrado")
	}
}

// CasoWrapComW - %w preserva a cadeia (Is continua funcionando).
func CasoWrapComW() {
	base := fmt.Errorf("buscar: %w", ErrNaoEncontrado)
	err := fmt.Errorf("servico: %w", base)

	fmt.Println("  texto:", err)
	fmt.Println("  errors.Is(ErrNaoEncontrado)?", errors.Is(err, ErrNaoEncontrado))
}

// CasoWrapComV - %v NÃO preserva a cadeia (Is deixa de funcionar).
func CasoWrapComV() {
	base := fmt.Errorf("buscar: %w", ErrNaoEncontrado)
	err := fmt.Errorf("servico: %v", base) // anti-padrão

	fmt.Println("  texto:", err)
	fmt.Println("  errors.Is(ErrNaoEncontrado)?", errors.Is(err, ErrNaoEncontrado), "(esperava false)")
}
