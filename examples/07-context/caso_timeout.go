package main

import (
	"context"
	"fmt"
	"time"
)

// CasoTimeout - a request morre quando estoura o prazo.
// Equivale a context.WithTimeout no início do handler HTTP.
func CasoTimeout() {
	// Prazo curto demais para o trabalho de 200ms.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel() // sempre liberar o timer do timeout

	err := Trabalhar(ctx, "consulta-lenta", 200*time.Millisecond)
	fmt.Println("  resultado:", err)
}
