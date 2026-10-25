package main

import (
	"context"
	"fmt"
)

// CasoTransacao - a service abre a Tx, coloca no context,
// e os repos reutilizam a mesma transação sem receber *Tx em todo método.
func CasoTransacao() {
	ctx := context.Background()
	ctx = ComRequestID(ctx, "req-tx-9")

	err := ServiceComTx(ctx)
	fmt.Println("  resultado:", err)
}
