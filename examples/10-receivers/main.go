// Demonstra value receiver (padrão) vs pointer receiver (mutação).
package main

import (
	"fmt"
	"os"

	"github.com/caramelotech/go-labs/examples/10-receivers/carteira"
	"github.com/caramelotech/go-labs/examples/10-receivers/perfil"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "erro: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("=== value receiver (perfil) ===")
	p := perfil.New("Caramelo", 3)
	fmt.Println(p.Apresentacao())
	fmt.Println("nome:", p.Nome(), "idade:", p.Idade())

	fmt.Println("\n=== pointer receiver (carteira) ===")
	c := carteira.New(100)
	fmt.Println("saldo inicial:", c.Saldo())
	if err := c.Depositar(50); err != nil {
		return err
	}
	fmt.Println("após depósito:", c.Saldo())
	if err := c.Sacar(30); err != nil {
		return err
	}
	fmt.Println("após saque:", c.Saldo())
	return nil
}
