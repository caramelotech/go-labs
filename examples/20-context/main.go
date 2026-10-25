package main

import "fmt"

func main() {
	fmt.Println("=== 1) cancel por timeout ===")
	CasoTimeout()

	fmt.Println("\n=== 2) cancel por close connection (simulado) ===")
	CasoCancelConexao()

	fmt.Println("\n=== 3) cancel por close connection (HTTP) ===")
	CasoHTTPCancel()

	fmt.Println("\n=== 4) propagar dados (request_id) ===")
	CasoValores()

	fmt.Println("\n=== 5) propagar transação ===")
	CasoTransacao()
}
