// Demonstra structs: agrupar campos relacionados num tipo.
package main

import "fmt"

type Endereco struct {
	Cidade string
	UF     string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco // struct aninhada
}

func main() {
	// Literal com nomes de campo (preferível)
	p := Pessoa{
		Nome:  "Caramelo",
		Idade: 3,
		Endereco: Endereco{
			Cidade: "São Paulo",
			UF:     "SP",
		},
	}
	fmt.Printf("%s tem %d anos\n", p.Nome, p.Idade)
	fmt.Printf("mora em %s/%s\n", p.Endereco.Cidade, p.Endereco.UF)

	// Campos começam no zero value se omitidos
	outra := Pessoa{Nome: "Ana"}
	fmt.Printf("outra: nome=%s idade=%d cidade=%q\n", outra.Nome, outra.Idade, outra.Endereco.Cidade)

	// Alterar campo
	p.Idade++
	fmt.Println("idade depois do aniversário:", p.Idade)

	// Struct anônima (útil em testes ou dados locais)
	ponto := struct {
		X, Y int
	}{X: 2, Y: 5}
	fmt.Println("ponto:", ponto)
}
