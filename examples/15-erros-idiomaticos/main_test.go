package main

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestSentinelaComIs(t *testing.T) {
	_, err := BuscarUsuario("99")
	if !errors.Is(err, ErrNaoEncontrado) {
		t.Fatalf("queria ErrNaoEncontrado, veio: %v", err)
	}
}

func TestEstruturadoComAs(t *testing.T) {
	_, err := ValidarIdade("abc")

	var ev *ErrValidacao
	if !errors.As(err, &ev) {
		t.Fatalf("queria *ErrValidacao, veio: %v", err)
	}
	if ev.Campo != "idade" {
		t.Fatalf("campo=%s, queria idade", ev.Campo)
	}

	var numErr *strconv.NumError
	if !errors.As(err, &numErr) {
		t.Fatalf("queria *strconv.NumError via Unwrap, veio: %v", err)
	}
}

func TestWrapPercentWVersusV(t *testing.T) {
	base := fmt.Errorf("buscar: %w", ErrNaoEncontrado)

	comW := fmt.Errorf("servico: %w", base)
	if !errors.Is(comW, ErrNaoEncontrado) {
		t.Fatal("percent-w deveria preservar a cadeia")
	}

	comV := fmt.Errorf("servico: %v", base)
	if errors.Is(comV, ErrNaoEncontrado) {
		t.Fatal("percent-v não deveria preservar a cadeia")
	}
}
