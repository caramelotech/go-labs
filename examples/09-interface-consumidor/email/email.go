// Package email é o provedor: tipo concreto, sem importar pedido e sem
// declarar "implements". Basta ter o método com a assinatura certa.
package email

import "fmt"

// Cliente envia notificações por e-mail (simulado com fmt).
type Cliente struct {
	de string
}

// New cria um Cliente que envia a partir do endereço informado.
func New(de string) Cliente {
	return Cliente{de: de}
}

// Notificar simula o envio. Value receiver: não altera Cliente.
// Satisfaz pedido.Notificador sem conhecer pedido.
func (c Cliente) Notificar(destino, mensagem string) error {
	fmt.Printf("[email %s -> %s] %s\n", c.de, destino, mensagem)
	return nil
}
