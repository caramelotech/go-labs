package main

import "fmt"

func main() {
	fmt.Println("=== sucesso ===")
	CasoSucesso()

	fmt.Println("\n=== erro sentinela ===")
	fmt.Println("1) errors.Is")
	CasoSentinelaComIs()
	fmt.Println("2) wrap com percent-w (preserva)")
	CasoWrapComW()
	fmt.Println("3) wrap com percent-v (quebra)")
	CasoWrapComV()

	fmt.Println("\n=== erro estruturado ===")
	fmt.Println("1) errors.As")
	CasoEstruturadoComAs()
}
