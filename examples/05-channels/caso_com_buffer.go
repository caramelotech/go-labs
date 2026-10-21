package main

import "fmt"

// CasoComBuffer: channel com capacidade - envios cabem no buffer
// sem precisar de um receiver imediato (até esgotar a capacidade).
func CasoComBuffer() {
	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30
	// ch <- 40  // descomente: deadlock - buffer cheio e ninguém lendo

	fmt.Println("len:", len(ch), "cap:", cap(ch))
	fmt.Println("saiu:", <-ch, <-ch, <-ch)
}
