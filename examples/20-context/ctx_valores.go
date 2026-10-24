package main

import "context"

// chave privada: evita colisão com outros pacotes.
// Nunca use string pura como chave de context.WithValue.
type chaveRequestID struct{}

// ComRequestID guarda um dado de escopo da requisição no context.
func ComRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, chaveRequestID{}, id)
}

// RequestID lê o valor. Ok=false se não existir.
func RequestID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(chaveRequestID{}).(string)
	return id, ok
}
