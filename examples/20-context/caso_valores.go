package main

import (
	"context"
	"fmt"
)

// CasoValores - propaga dados de escopo da request (request_id)
// do handler até o repo, sem poluir assinaturas com parâmetros extras.
//
// Use WithValue só para dados transversais da request.
// Não use para passar dependências obrigatórias (isso vai em argumento).
func CasoValores() {
	ctx := context.Background()
	ctx = ComRequestID(ctx, "req-abc-123")

	err := Handler(ctx, "listar")
	fmt.Println("  resultado:", err)
}
