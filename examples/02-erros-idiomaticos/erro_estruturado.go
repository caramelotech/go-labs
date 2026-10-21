package main

import "fmt"

// ErrValidacao é um erro estruturado (tem campos).
//
// Estruturado = um tipo próprio com dados extras.
// Use errors.As quando precisa ler Campo, Valor ou Msg.
type ErrValidacao struct {
	Campo string
	Valor string
	Msg   string
	causa error // opcional; Is/As passam por aqui via Unwrap
}

// Error e Unwrap usam value receiver: só leem campos, não mutam o erro.
// Continuamos devolvendo *ErrValidacao nas funções - *T recebe métodos de T.
func (e ErrValidacao) Error() string {
	if e.causa != nil {
		return fmt.Sprintf("%s: %s (%q): %v", e.Campo, e.Msg, e.Valor, e.causa)
	}
	return fmt.Sprintf("%s: %s (%q)", e.Campo, e.Msg, e.Valor)
}

// Unwrap liga este erro à causa (mesmo papel do %w).
func (e ErrValidacao) Unwrap() error {
	return e.causa
}
