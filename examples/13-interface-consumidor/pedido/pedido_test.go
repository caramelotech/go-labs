package pedido_test

import (
	"errors"
	"testing"

	"github.com/caramelotech/go-labs/examples/13-interface-consumidor/pedido"
)

// fakeNotificador vive no teste do consumidor - sem importar email/console.
type fakeNotificador struct {
	destino  string
	mensagem string
	err      error
}

func (f *fakeNotificador) Notificar(destino, mensagem string) error {
	f.destino = destino
	f.mensagem = mensagem
	return f.err
}

func TestConfirmarNotificaDestino(t *testing.T) {
	fake := &fakeNotificador{}
	svc, err := pedido.New(fake)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	if err := svc.Confirmar("7", "a@b.com"); err != nil {
		t.Fatalf("Confirmar: %v", err)
	}

	if fake.destino != "a@b.com" {
		t.Errorf("destino = %q, queria a@b.com", fake.destino)
	}
	if fake.mensagem != "pedido 7 confirmado" {
		t.Errorf("mensagem = %q", fake.mensagem)
	}
}

func TestConfirmarPropagaErro(t *testing.T) {
	fake := &fakeNotificador{err: errors.New("falha de rede")}
	svc, err := pedido.New(fake)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	err = svc.Confirmar("7", "a@b.com")
	if err == nil {
		t.Fatal("queria erro")
	}
}

func TestNewRejeitaNil(t *testing.T) {
	if _, err := pedido.New(nil); err == nil {
		t.Fatal("queria erro com notificador nil")
	}
}
