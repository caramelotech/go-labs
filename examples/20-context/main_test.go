package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestTimeoutCancelaTrabalho(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	err := Trabalhar(ctx, "teste", 200*time.Millisecond)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("queria DeadlineExceeded, veio: %v", err)
	}
}

func TestCancelPropagaNasCamadas(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // já cancelado antes de começar

	err := Handler(ctx, "x")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("queria Canceled, veio: %v", err)
	}
}

func TestRequestIDNoContext(t *testing.T) {
	ctx := ComRequestID(context.Background(), "req-1")
	id, ok := RequestID(ctx)
	if !ok || id != "req-1" {
		t.Fatalf("request_id=%q ok=%v", id, ok)
	}
}

func TestTxNoContext(t *testing.T) {
	tx := &Tx{ID: "tx-test"}
	ctx := ComTx(context.Background(), tx)

	got, ok := TxFrom(ctx)
	if !ok || got.ID != "tx-test" {
		t.Fatalf("tx=%v ok=%v", got, ok)
	}
}
