// Demonstra testes em tabela (table-driven): o padrão idiomático
// para cobrir vários casos sem repetir o corpo do teste.
package main

import (
	"fmt"
	"strings"
)

// ContarPalavras conta ocorrências de cada palavra, ignorando maiúsculas.
func ContarPalavras(texto string) map[string]int {
	contagem := make(map[string]int)
	for _, palavra := range strings.Fields(texto) {
		chave := strings.ToLower(palavra)
		contagem[chave]++
	}
	return contagem
}

func main() {
	texto := "Go go GO labs"
	fmt.Printf("texto: %q\n", texto)
	fmt.Printf("contagem: %v\n", ContarPalavras(texto))
}
