// Demonstra as formas do for: a única estrutura de repetição em Go.
package main

import "fmt"

func main() {
	// 1) for clássico (estilo C)
	fmt.Println("--- clássico ---")
	for i := 0; i < 3; i++ {
		fmt.Println("i =", i)
	}

	// 2) for como while: só a condição
	fmt.Println("--- condição ---")
	n := 3
	for n > 0 {
		fmt.Println("n =", n)
		n--
	}

	// 3) for range em slice
	fmt.Println("--- range slice ---")
	frutas := []string{"maçã", "banana", "uva"}
	for i, fruta := range frutas {
		fmt.Printf("%d: %s\n", i, fruta)
	}

	// 4) só o valor (ignore o índice com _)
	fmt.Println("--- só valores ---")
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}

	// 5) break e continue
	fmt.Println("--- break / continue ---")
	for i := 1; i <= 5; i++ {
		if i == 2 {
			continue // pula o 2
		}
		if i == 4 {
			break // para no 4
		}
		fmt.Println("número:", i)
	}
}
