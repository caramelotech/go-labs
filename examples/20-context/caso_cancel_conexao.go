package main

import (
	"context"
	"fmt"
	"time"
)

// CasoCancelConexao simula o cliente fechando a conexão no meio da request.
//
// No net/http real: r.Context() é cancelado quando o client desconecta.
// Aqui o cancel() representa esse "connection close".
func CasoCancelConexao() {
	ctx, cancel := context.WithCancel(context.Background())

	// Cliente abandona a request após 40ms.
	go func() {
		time.Sleep(40 * time.Millisecond)
		fmt.Println("  cliente: fechou a conexão")
		cancel()
	}()

	err := Handler(ctx, "gerar-relatorio")
	fmt.Println("  resultado:", err)
}
