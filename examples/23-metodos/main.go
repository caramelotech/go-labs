// Demonstra métodos: funções com receiver ligado a um tipo.
package main

import "fmt"

type Retangulo struct {
	Largura, Altura float64
}

// Método com value receiver - não altera o struct; trabalha numa cópia.
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (r Retangulo) String() string {
	return fmt.Sprintf("%.0fx%.0f", r.Largura, r.Altura)
}

// Método com pointer receiver - altera o original.
func (r *Retangulo) Escalar(fator float64) {
	r.Largura *= fator
	r.Altura *= fator
}

func main() {
	r := Retangulo{Largura: 4, Altura: 3}

	fmt.Println("retângulo:", r.String())
	fmt.Println("área:", r.Area())

	r.Escalar(2)
	fmt.Println("após escalar x2:", r.String())
	fmt.Println("nova área:", r.Area())

	// Sintaxe equivalente a método: função com receiver explícito é só o método
	fmt.Println("Area como função:", Retangulo.Area(r))
}
