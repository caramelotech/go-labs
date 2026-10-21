package main

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/caramelotech/go-labs/examples/11-slog/slogx"
)

func TestSucessoMiddlewareLogaAttrsDasCamadas(t *testing.T) {
	var buf bytes.Buffer
	slogx.SetDefaultJSON(&buf, slog.LevelInfo)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /pedido", handlePedido)
	srv := httptest.NewServer(MiddlewareLog(mux))
	t.Cleanup(srv.Close)

	req, err := http.NewRequest(http.MethodGet, srv.URL+"/pedido?id=7", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-Request-ID", "test-ok")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d", resp.StatusCode)
	}

	final := acharLog(t, buf.Bytes(), "request finalizada")
	if final["request_id"] != "test-ok" {
		t.Errorf("request_id = %v", final["request_id"])
	}
	if final["pedido_id"] != "7" {
		t.Errorf("pedido_id = %v (queria attrs da service)", final["pedido_id"])
	}
	if final["resultado"] != "aprovado" {
		t.Errorf("resultado = %v", final["resultado"])
	}
	if final["camada"] != "handler" {
		t.Errorf("camada = %v", final["camada"])
	}
}

func TestErroLogTrazDadosDoMiddleware(t *testing.T) {
	var buf bytes.Buffer
	slogx.SetDefaultJSON(&buf, slog.LevelInfo)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /pedido", handlePedido)
	srv := httptest.NewServer(MiddlewareLog(mux))
	t.Cleanup(srv.Close)

	req, err := http.NewRequest(http.MethodGet, srv.URL+"/pedido?id=9&falha=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-Request-ID", "test-erro")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("status = %d", resp.StatusCode)
	}

	erroLog := acharLog(t, buf.Bytes(), "processar pedido")
	if erroLog["request_id"] != "test-erro" {
		t.Errorf("Error sem request_id do middleware: %v", erroLog["request_id"])
	}
	if erroLog["method"] != http.MethodGet {
		t.Errorf("method = %v", erroLog["method"])
	}
	if erroLog["pedido_id"] != "9" {
		t.Errorf("pedido_id = %v", erroLog["pedido_id"])
	}
	if !strings.Contains(str(erroLog["erro"]), "pagamento recusado") {
		t.Errorf("erro = %v", erroLog["erro"])
	}
}

func acharLog(t *testing.T, raw []byte, msg string) map[string]any {
	t.Helper()
	for _, line := range bytes.Split(raw, []byte("\n")) {
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}
		var m map[string]any
		if err := json.Unmarshal(line, &m); err != nil {
			t.Fatalf("json: %v (linha %s)", err, line)
		}
		if m["msg"] == msg {
			return m
		}
	}
	t.Fatalf("não achou log com msg=%q em:\n%s", msg, raw)
	return nil
}

func str(v any) string {
	s, _ := v.(string)
	return s
}
