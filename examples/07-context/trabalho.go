package main

import (
	"context"
	"fmt"
	"time"
)

// Trabalhar simula I/O lento (query, HTTP, disco).
// Sempre respeita o context: se cancelar/timeout, para na hora.
func Trabalhar(ctx context.Context, nome string, demora time.Duration) error {
	fmt.Printf("  %s: começou\n", nome)

	timer := time.NewTimer(demora)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		fmt.Printf("  %s: cancelado (%v)\n", nome, ctx.Err())
		return fmt.Errorf("%s: %w", nome, ctx.Err())
	case <-timer.C:
		fmt.Printf("  %s: concluiu\n", nome)
		return nil
	}
}
