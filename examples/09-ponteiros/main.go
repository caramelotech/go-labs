// Demonstra ponteiros: endereço (&) e desreferência (*).
package main

import "fmt"

func dobraCopia(n int) {
	n = n * 2 // altera só a cópia local
}

func dobraPonteiro(n *int) {
	*n = *n * 2 // altera o valor no endereço apontado
}

func main() {
	x := 10
	p := &x // p guarda o endereço de x

	fmt.Println("x:", x)
	fmt.Println("endereço (&x):", p)
	fmt.Println("valor via ponteiro (*p):", *p)

	*p = 20
	fmt.Println("depois de *p = 20, x:", x)

	n := 5
	dobraCopia(n)
	fmt.Println("após dobraCopia:", n) // continua 5

	dobraPonteiro(&n)
	fmt.Println("após dobraPonteiro:", n) // virou 10

	// new(T) aloca um T com zero value e devolve *T
	pi := new(int)
	fmt.Println("new(int):", *pi)
	*pi = 42
	fmt.Println("após *pi = 42:", *pi)
}
