// Demonstra funções: múltiplos retornos, retornos nomeados e variádicas.
package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

// Múltiplos retornos - padrão comum em Go (valor + erro, quociente + resto, etc.)
func dividir(a, b int) (int, int) {
	return a / b, a % b
}

// Retornos nomeados: já existem como variáveis locais; bare return devolve elas
func coordenadas() (x, y int) {
	x = 10
	y = 20
	return
}

// Variádica: recebe zero ou mais argumentos do mesmo tipo
func media(nums ...int) float64 {
	if len(nums) == 0 {
		return 0
	}
	total := 0
	for _, n := range nums {
		total += n
	}
	return float64(total) / float64(len(nums))
}

func main() {
	fmt.Println("soma:", soma(2, 3))

	q, r := dividir(17, 5)
	fmt.Println("dividir 17/5 → quociente:", q, "resto:", r)

	x, y := coordenadas()
	fmt.Println("coordenadas:", x, y)

	fmt.Println("média:", media(2, 4, 6, 8))
	fmt.Println("média vazia:", media())
}
