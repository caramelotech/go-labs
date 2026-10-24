package main

import "fmt"

// CasoSemBuffer: channel sem capacidade - send e receive sincronizam.
// O envio bloqueia até alguém receber (e vice-versa).
func CasoSemBuffer() {
	ch := make(chan string)

	go func() {
		ch <- "ping"
	}()

	msg := <-ch
	fmt.Println("recebido:", msg)
}
