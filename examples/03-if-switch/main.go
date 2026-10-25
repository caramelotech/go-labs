// Demonstra if (incluindo declaração curta) e switch em Go.
package main

import "fmt"

func main() {
	idade := 17

	// if / else clássico
	if idade >= 18 {
		fmt.Println("maior de idade")
	} else {
		fmt.Println("menor de idade")
	}

	// if com declaração curta: a variável vive só no if/else
	if n := len("caramelo"); n > 5 {
		fmt.Println("palavra longa, letras:", n)
	} else {
		fmt.Println("palavra curta, letras:", n)
	}

	// switch com expressão
	nota := 8
	switch nota {
	case 10:
		fmt.Println("nota: excelente")
	case 8, 9:
		fmt.Println("nota: muito bom")
	case 6, 7:
		fmt.Println("nota: ok")
	default:
		fmt.Println("nota: precisa melhorar")
	}

	// switch sem expressão: equivalente a if/else encadeado
	switch {
	case idade < 13:
		fmt.Println("faixa: criança")
	case idade < 18:
		fmt.Println("faixa: adolescente")
	default:
		fmt.Println("faixa: adulto")
	}
}
