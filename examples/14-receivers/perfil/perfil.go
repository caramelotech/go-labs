// Package perfil usa só value receiver: o tipo é pequeno e os métodos
// não alteram o struct - cópia é barata e suficiente.
package perfil

import "fmt"

// Perfil é um valor imutável do ponto de vista dos métodos.
type Perfil struct {
	nome  string
	idade int
}

// New monta um Perfil. Devolve valor, não ponteiro.
func New(nome string, idade int) Perfil {
	return Perfil{nome: nome, idade: idade}
}

// Nome devolve o nome. (p Perfil) - não precisa de ponteiro.
func (p Perfil) Nome() string {
	return p.nome
}

// Idade devolve a idade. Mesmo padrão: value receiver.
func (p Perfil) Idade() int {
	return p.idade
}

// Apresentacao formata um texto. Continua (p Perfil).
func (p Perfil) Apresentacao() string {
	return fmt.Sprintf("%s (%d anos)", p.nome, p.idade)
}
