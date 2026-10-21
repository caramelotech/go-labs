// Demonstra o essencial de log/slog via o pacote slogx (reutilizado em 12-slog-contexto).
package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/caramelotech/go-labs/examples/11-slog/slogx"
)

func main() {
	fmt.Println("=== JSON (slogx.NewJSON) ===")
	jsonLog := slogx.NewJSON(os.Stdout, slog.LevelDebug)
	jsonLog.Debug("detalhe de desenvolvimento", "passo", 1)
	jsonLog.Info("pedido criado", "pedido_id", "42", "total_centavos", 1990)
	jsonLog.Warn("estoque baixo", "sku", "CAM-1", "qtd", 2)
	jsonLog.Error("pagamento recusado", "pedido_id", "42", "erro", "cartão inválido")

	fmt.Println("\n=== With: attrs fixos em todas as linhas ===")
	reqLog := jsonLog.With("request_id", "req-1", "serviço", "checkout")
	reqLog.Info("validando carrinho")
	reqLog.Info("cobrando cartão", "meio", "pix")

	fmt.Println("\n=== Group: agrupa chaves sob um prefixo ===")
	jsonLog.Info("cliente",
		slog.Group("user",
			"id", "u-9",
			"plano", "pro",
		),
	)

	fmt.Println("\n=== Text (slogx.NewText) - leitura humana no terminal ===")
	textLog := slogx.NewText(os.Stdout, slog.LevelInfo)
	textLog.Info("servidor no ar", "addr", ":8080")
}
