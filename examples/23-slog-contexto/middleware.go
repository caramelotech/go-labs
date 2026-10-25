package main

import (
	"log/slog"
	"net/http"

	"github.com/caramelotech/go-labs/examples/22-slog/slogx"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// MiddlewareLog:
//  1. cria logger com dados da request e coloca no context (slogx)
//  2. chama o próximo handler
//  3. no retorno, loga status + tudo que as camadas fizeram Append
func MiddlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = "req-demo"
		}

		logger := slog.Default().With(
			"request_id", requestID,
			"method", r.Method,
			"path", r.URL.Path,
		)
		ctx := slogx.ComLogger(r.Context(), logger)

		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(sw, r.WithContext(ctx))

		slogx.Logger(ctx).Info("request finalizada", "status", sw.status)
	})
}
