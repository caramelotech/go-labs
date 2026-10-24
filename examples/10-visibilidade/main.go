// Demonstra visibilidade: identificador com maiúscula = exportado entre pacotes.
package main

import (
	"fmt"

	"github.com/caramelotech/go-labs/examples/10-visibilidade/conta"
)

func main() {
	c := conta.New("Caramelo", 100)

	// Campo exportado: ok
	fmt.Println("titular:", c.Titular)

	// Campo saldo é minúsculo - não compila se descomentar:
	// fmt.Println(c.saldo)

	// Acesso ao saldo só via método exportado
	fmt.Println("saldo:", c.Saldo())

	c.Depositar(50)
	fmt.Println("após depósito:", c.Saldo())

	// extrato é privada - use Resumo, que o pacote conta exporta
	fmt.Println("resumo:", conta.Resumo(c))

	// Isto também não compilaria:
	// fmt.Println(conta.extrato(c))
}
