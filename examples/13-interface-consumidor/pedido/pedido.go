// Package pedido é o consumidor: define a interface do que precisa
// e recebe a implementação pronta via New (injeção de dependência).
package pedido

import "fmt"

// Notificador é o contrato que o pedido precisa.
// Fica aqui (consumidor), não no pacote que envia e-mail ou imprime no console.
type Notificador interface {
	Notificar(destino, mensagem string) error
}

// Servico confirma pedidos e avisa o cliente pelo Notificador injetado.
type Servico struct {
	notificador Notificador
}

// New valida a dependência obrigatória na construção - sem checagem de nil
// espalhada nos métodos de negócio. Devolve valor: Confirmar não muta Servico.
func New(n Notificador) (Servico, error) {
	if n == nil {
		return Servico{}, fmt.Errorf("notificador é obrigatório")
	}
	return Servico{notificador: n}, nil
}

// Confirmar registra a confirmação e notifica o destino.
// Value receiver: só lê s.notificador, não altera o struct.
func (s Servico) Confirmar(id, destino string) error {
	msg := fmt.Sprintf("pedido %s confirmado", id)
	if err := s.notificador.Notificar(destino, msg); err != nil {
		return fmt.Errorf("confirmar pedido %s: %w", id, err)
	}
	return nil
}
