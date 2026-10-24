// Package console é outro provedor: mesma assinatura de método,
// implementação diferente - main troca na injeção sem mudar pedido.
package console

import "fmt"

// Saida escreve a notificação no terminal.
type Saida struct{}

// Notificar imprime no console. Também satisfaz pedido.Notificador.
func (Saida) Notificar(destino, mensagem string) error {
	fmt.Printf("[console -> %s] %s\n", destino, mensagem)
	return nil
}
