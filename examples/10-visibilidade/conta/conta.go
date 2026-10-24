// Package conta mostra o que é exportado (público) e o que não é (privado ao pacote).
package conta

import "fmt"

// Conta é exportada: outros pacotes podem usar o tipo conta.Conta.
type Conta struct {
	Titular string // exportado - acessível fora do pacote
	saldo   int    // não exportado - só visível dentro de conta
}

// New cria uma Conta. Função exportada (N maiúsculo).
func New(titular string, saldoInicial int) Conta {
	return Conta{Titular: titular, saldo: saldoInicial}
}

// Saldo devolve o saldo. Getter sem prefixo Get - idiomático em Go.
func (c Conta) Saldo() int {
	return c.saldo
}

// Depositar altera o saldo. Precisa de ponteiro para mutar.
func (c *Conta) Depositar(valor int) {
	if valor > 0 {
		c.saldo += valor
	}
}

// extrato é privada ao pacote - main não consegue chamar conta.extrato.
func extrato(c Conta) string {
	return fmt.Sprintf("%s: R$ %d", c.Titular, c.saldo)
}

// Resumo usa a função privada internamente e exporta o resultado.
func Resumo(c Conta) string {
	return extrato(c)
}
