// Demonstra declaração de variáveis, constantes e zero values em Go.
package main

import "fmt"

func main() {
	// Forma longa: tipo explícito
	var nome string = "Caramelo"
	var idade int = 3

	// Inferência com var (o compilador descobre o tipo)
	var ativo = true

	// Forma curta (só dentro de funções): :=
	peso := 12.5
	inicial := 'C' // rune (alias de int32) - um caractere Unicode

	// Constantes não mudam depois de declaradas
	const lab = "Go Labs"

	fmt.Println("nome:", nome)
	fmt.Println("idade:", idade)
	fmt.Println("ativo:", ativo)
	fmt.Println("peso:", peso)
	fmt.Printf("inicial: %c código: %d\n", inicial, inicial)
	fmt.Println("lab:", lab)

	// Zero values: valor padrão de variáveis declaradas sem inicialização
	var i int
	var s string
	var b bool
	var f float64
	fmt.Println("--- zero values ---")
	fmt.Printf("int=%d string=%q bool=%v float64=%v\n", i, s, b, f)
}
