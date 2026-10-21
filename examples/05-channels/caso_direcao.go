package main

import "fmt"

// enviar só escreve; receber só lê. A assinatura documenta o papel
// e o compilador impede uso invertido.
func enviar(saida chan<- string, msg string) {
	saida <- msg
}

func receber(entrada <-chan string) string {
	return <-entrada
}

// CasoDirecao: channels tipados como send-only ou receive-only.
func CasoDirecao() {
	ch := make(chan string, 1)

	enviar(ch, "olá")
	fmt.Println("recebido:", receber(ch))
}
