// O main é o único lugar que conhece provedor + consumidor:
// instancia o concreto e injeta na interface do consumidor.
package main

import (
	"fmt"
	"os"

	"github.com/caramelotech/go-labs/examples/09-interface-consumidor/console"
	"github.com/caramelotech/go-labs/examples/09-interface-consumidor/email"
	"github.com/caramelotech/go-labs/examples/09-interface-consumidor/pedido"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "erro: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("=== injeção com email.Cliente ===")
	mail := email.New("noreply@caramelo.dev")
	svcMail, err := pedido.New(mail)
	if err != nil {
		return err
	}
	if err := svcMail.Confirmar("42", "aluno@exemplo.com"); err != nil {
		return err
	}

	fmt.Println("\n=== mesma regra, outro provedor (console.Saida) ===")
	svcConsole, err := pedido.New(console.Saida{})
	if err != nil {
		return err
	}
	return svcConsole.Confirmar("43", "aluno@exemplo.com")
}
