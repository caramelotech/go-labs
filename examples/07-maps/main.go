// Demonstra maps: dicionário chave → valor.
package main

import "fmt"

func main() {
	// Literal
	idades := map[string]int{
		"Ana":      28,
		"Bruno":    34,
		"Caramelo": 3,
	}
	fmt.Println("mapa:", idades)

	// Inserir / atualizar
	idades["Diana"] = 41
	idades["Ana"] = 29
	fmt.Println("após insert/update:", idades)

	// Leitura com “comma ok”: distingue zero value de chave ausente
	idade, ok := idades["Ana"]
	fmt.Println("Ana:", idade, "existe?", ok)

	idade, ok = idades["Zé"]
	fmt.Println("Zé:", idade, "existe?", ok)

	// delete
	delete(idades, "Bruno")
	fmt.Println("após delete Bruno:", idades)

	// range (ordem das chaves não é garantida)
	fmt.Println("--- range ---")
	for nome, idade := range idades {
		fmt.Printf("%s tem %d anos\n", nome, idade)
	}

	// make para mapa vazio
	estoque := make(map[string]int)
	estoque["café"] = 10
	fmt.Println("estoque:", estoque, "len:", len(estoque))
}
