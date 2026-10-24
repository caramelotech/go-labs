package main

import "errors"

// ErrNaoEncontrado é um erro sentinela.
//
// Sentinela = um valor fixo de erro.
// Use errors.Is quando só importa "foi este erro?".
var ErrNaoEncontrado = errors.New("não encontrado")
