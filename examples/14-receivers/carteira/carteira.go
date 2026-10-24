// Package carteira usa pointer receiver: Depositar e Sacar mudam o saldo.
// Se um método do tipo precisa de ponteiro, os outros do mesmo tipo
// também usam ponteiro - method set consistente.
package carteira

import "fmt"

// Carteira guarda um saldo que muda com o tempo.
type Carteira struct {
	saldo int
}

// New cria uma carteira com saldo inicial.
func New(saldoInicial int) *Carteira {
	return &Carteira{saldo: saldoInicial}
}

// Saldo lê o saldo. Ponteiro por consistência com Depositar/Sacar.
func (c *Carteira) Saldo() int {
	return c.saldo
}

// Depositar altera o estado - precisa de (c *Carteira).
func (c *Carteira) Depositar(valor int) error {
	if valor <= 0 {
		return fmt.Errorf("depósito deve ser positivo")
	}
	c.saldo += valor
	return nil
}

// Sacar altera o estado - também (c *Carteira).
func (c *Carteira) Sacar(valor int) error {
	if valor <= 0 {
		return fmt.Errorf("saque deve ser positivo")
	}
	if valor > c.saldo {
		return fmt.Errorf("saldo insuficiente")
	}
	c.saldo -= valor
	return nil
}
