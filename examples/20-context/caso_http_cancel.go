package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

// CasoHTTPCancel - versão próxima do mundo real:
// o context da request HTTP cancela quando o client desiste.
func CasoHTTPCancel() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context() // já vem cancelável com a conexão

		err := Trabalhar(ctx, "http.handler", 300*time.Millisecond)
		if err != nil {
			fmt.Println("  servidor: parou porque o client cancelou")
			return
		}
		fmt.Fprint(w, "ok")
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	ctx, cancel := context.WithCancel(context.Background())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL, nil)
	if err != nil {
		cancel()
		fmt.Println("  erro:", err)
		return
	}
	defer cancel()

	// Client desiste cedo (equivale a fechar a conexão / AbortController).
	go func() {
		time.Sleep(40 * time.Millisecond)
		fmt.Println("  client HTTP: cancelou a request")
		cancel()
	}()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("  client HTTP: resultado:", err)
		return
	}
	defer resp.Body.Close()
	corpo, _ := io.ReadAll(resp.Body)
	fmt.Printf("  client HTTP: status=%d corpo=%q\n", resp.StatusCode, corpo)
}
