package main

import (
	"context"
	"fmt"
)

type chaveTx struct{}

// Tx é uma "transação" de mentira, só para mostrar o padrão.
type Tx struct {
	ID string
}

// Exec só lê tx.ID - value receiver basta (ponteiro no context é outra história).
func (tx Tx) Exec(ctx context.Context, sql string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	fmt.Printf("  tx %s: exec %q\n", tx.ID, sql)
	return nil
}

// ComTx propaga a transação pelas camadas via context.
func ComTx(ctx context.Context, tx *Tx) context.Context {
	return context.WithValue(ctx, chaveTx{}, tx)
}

// TxFrom recupera a transação (se existir).
func TxFrom(ctx context.Context) (*Tx, bool) {
	tx, ok := ctx.Value(chaveTx{}).(*Tx)
	return tx, ok
}
