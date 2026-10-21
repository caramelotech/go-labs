// Package slogx facilita o uso de log/slog nos exemplos do Go Labs:
// handlers prontos e (em context.go) propagação de logger via context.
package slogx

import (
	"io"
	"log/slog"
)

// NewJSON cria um logger que escreve uma linha JSON por evento.
func NewJSON(w io.Writer, level slog.Level) *slog.Logger {
	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: level}))
}

// NewText cria um logger em formato texto (key=value), bom para terminal.
func NewText(w io.Writer, level slog.Level) *slog.Logger {
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{Level: level}))
}

// SetDefaultJSON configura o logger padrão do processo em JSON (stdout típico).
func SetDefaultJSON(w io.Writer, level slog.Level) {
	slog.SetDefault(NewJSON(w, level))
}
