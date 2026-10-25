package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/caramelotech/go-labs/examples/22-slog/slogx"
)

// handlePedido: HTTP -> Service. O context (com logger) já veio do middleware.
func handlePedido(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")
	if id == "" {
		id = "desconhecido"
	}
	falha := r.URL.Query().Get("falha") == "1"

	slogx.Append(ctx, "camada", "handler")

	if err := Service(ctx, id, falha); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "pedido ok")
}

// Service é a "camada seguinte":
//   - erro: loga com slogx.Logger(ctx) → traz request_id/method/path do middleware
//   - sucesso: slogx.Append no context → middleware inclui no log final
func Service(ctx context.Context, pedidoID string, falha bool) error {
	if falha {
		err := fmt.Errorf("pagamento recusado")
		slogx.Logger(ctx).Error("processar pedido",
			"pedido_id", pedidoID,
			"erro", err.Error(),
		)
		return err
	}

	slogx.Append(ctx,
		"pedido_id", pedidoID,
		"total_centavos", 1990,
		"resultado", "aprovado",
	)
	return nil
}
