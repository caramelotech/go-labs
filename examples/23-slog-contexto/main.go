// Demonstra slog com dados da request no context: reutiliza slogx (exemplo 11)
// no middleware HTTP - erro herda attrs; sucesso acumula Append até o log final.
package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/caramelotech/go-labs/examples/22-slog/slogx"
)

func main() {
	slogx.SetDefaultJSON(os.Stdout, slog.LevelInfo)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /pedido", handlePedido)
	handler := MiddlewareLog(mux)

	if len(os.Args) > 1 && os.Args[1] == "serve" {
		addr := ":8080"
		fmt.Fprintf(os.Stderr, "ouvindo em http://localhost%s  (Ctrl+C para parar)\n", addr)
		fmt.Fprintf(os.Stderr, "  curl -H 'X-Request-ID: abc' 'http://localhost%s/pedido?id=42'\n", addr)
		fmt.Fprintf(os.Stderr, "  curl -H 'X-Request-ID: abc' 'http://localhost%s/pedido?id=42&falha=1'\n", addr)
		if err := http.ListenAndServe(addr, handler); err != nil {
			fmt.Fprintf(os.Stderr, "erro: %v\n", err)
			os.Exit(1)
		}
		return
	}

	demonstrar(handler)
}

func demonstrar(handler http.Handler) {
	server := httptest.NewServer(handler)
	defer server.Close()

	fmt.Fprintln(os.Stderr, "=== sucesso: service faz Append; middleware loga tudo no fim ===")
	doGET(server.URL+"/pedido?id=42", "req-ok")

	fmt.Fprintln(os.Stderr, "\n=== erro: service loga Error já com dados do middleware ===")
	doGET(server.URL+"/pedido?id=99&falha=1", "req-erro")
}

func doGET(url, requestID string) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "erro:", err)
		return
	}
	req.Header.Set("X-Request-ID", requestID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "erro:", err)
		return
	}
	defer resp.Body.Close()
	corpo, _ := io.ReadAll(resp.Body)
	fmt.Fprintf(os.Stderr, "HTTP %d %s\n", resp.StatusCode, corpo)
}
