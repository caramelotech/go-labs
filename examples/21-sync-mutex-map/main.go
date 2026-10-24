// Compara sync.Mutex, sync.RWMutex e sync.Map: quando usar cada um
// para proteger estado compartilhado entre goroutines.
package main

import "fmt"

func main() {
	fmt.Println("=== 1) Mutex - exclusão mútua geral ===")
	CasoMutex()

	fmt.Println("\n=== 2) RWMutex - muitas leituras, poucas escritas ===")
	CasoRWMutex()

	fmt.Println("\n=== 3) sync.Map - casos especiais de mapa concorrente ===")
	CasoSyncMap()
}
