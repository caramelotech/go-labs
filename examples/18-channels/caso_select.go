package main

import (
	"fmt"
	"time"
)

// CasoSelect: espera o primeiro caso pronto entre vários channels.
// Útil para timeout, cancelamento e multiplexação.
func CasoSelect() {
	rapido := make(chan string, 1)
	lento := make(chan string, 1)

	rapido <- "rápido"
	go func() {
		time.Sleep(50 * time.Millisecond)
		lento <- "lento"
	}()

	select {
	case msg := <-rapido:
		fmt.Println("ganhou:", msg)
	case msg := <-lento:
		fmt.Println("ganhou:", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout")
	}
}
