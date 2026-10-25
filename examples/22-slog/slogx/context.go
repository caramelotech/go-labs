package slogx

import (
	"context"
	"log/slog"
)

type chaveLogger struct{}
type chaveDados struct{}

// dados acumula atributos das camadas. Ponteiro no context para o
// middleware ver o que handler/service acrescentaram na volta.
type dados struct {
	attrs []any
}

// ComLogger coloca o slog no context e inicia o saco de atributos da request.
func ComLogger(ctx context.Context, logger *slog.Logger) context.Context {
	ctx = context.WithValue(ctx, chaveLogger{}, logger)
	return context.WithValue(ctx, chaveDados{}, &dados{})
}

// Logger devolve o slog do context já com os attrs acumulados pelas camadas.
func Logger(ctx context.Context) *slog.Logger {
	logger, _ := ctx.Value(chaveLogger{}).(*slog.Logger)
	if logger == nil {
		logger = slog.Default()
	}
	if d, ok := ctx.Value(chaveDados{}).(*dados); ok && len(d.attrs) > 0 {
		return logger.With(d.attrs...)
	}
	return logger
}

// Append registra dados da camada atual para o middleware logar no fim.
// Também entram em Logger(ctx) nas linhas seguintes da mesma request.
func Append(ctx context.Context, args ...any) {
	d, ok := ctx.Value(chaveDados{}).(*dados)
	if !ok || d == nil {
		return
	}
	d.attrs = append(d.attrs, args...)
}

// Attrs devolve cópia dos atributos acumulados (útil nos testes).
func Attrs(ctx context.Context) []any {
	d, ok := ctx.Value(chaveDados{}).(*dados)
	if !ok || d == nil || len(d.attrs) == 0 {
		return nil
	}
	out := make([]any, len(d.attrs))
	copy(out, d.attrs)
	return out
}
