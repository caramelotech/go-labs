package main

import (
	"context"
	"fmt"
	"time"
)

// Handler -> Service -> Repo
// O context desce a pilha. Cada camada só adiciona o que precisa.

func Handler(ctx context.Context, acao string) error {
	fmt.Println("  handler: recebeu request")
	return Service(ctx, acao)
}

func Service(ctx context.Context, acao string) error {
	fmt.Println("  service: regra de negócio")
	return Repo(ctx, acao)
}

func Repo(ctx context.Context, acao string) error {
	if id, ok := RequestID(ctx); ok {
		fmt.Printf("  repo: request_id=%s acao=%s\n", id, acao)
	} else {
		fmt.Printf("  repo: acao=%s\n", acao)
	}
	return Trabalhar(ctx, "repo."+acao, 80*time.Millisecond)
}

// ServiceComTx abre uma Tx, coloca no context e chama o repo.
func ServiceComTx(ctx context.Context) error {
	tx := &Tx{ID: "tx-1"}
	fmt.Printf("  service: abriu %s\n", tx.ID)

	ctx = ComTx(ctx, tx)
	if err := RepoComTx(ctx, "INSERT INTO pedidos"); err != nil {
		return err
	}
	if err := RepoComTx(ctx, "INSERT INTO itens"); err != nil {
		return err
	}

	fmt.Printf("  service: commit %s\n", tx.ID)
	return nil
}

func RepoComTx(ctx context.Context, sql string) error {
	tx, ok := TxFrom(ctx)
	if !ok {
		return fmt.Errorf("repo: transação ausente no context")
	}
	return tx.Exec(ctx, sql)
}
