// Demonstra as formas básicas de usar channels em Go: sem buffer,
// com buffer, close+range, select e direção (send/receive).
package main

import "fmt"

func main() {
	fmt.Println("=== 1) sem buffer (handshake) ===")
	CasoSemBuffer()

	fmt.Println("\n=== 2) com buffer ===")
	CasoComBuffer()

	fmt.Println("\n=== 3) close + range ===")
	CasoCloseRange()

	fmt.Println("\n=== 4) select ===")
	CasoSelect()

	fmt.Println("\n=== 5) direção (chan<- / <-chan) ===")
	CasoDirecao()
}
