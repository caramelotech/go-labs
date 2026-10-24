// Demonstra goroutines coordenadas com sync.WaitGroup: o ponto de
// partida da concorrência em Go, sem channels ainda.
package main

import (
	"fmt"
	"sync"
	"time"
)

// processar simula trabalho (I/O, CPU, chamada de rede).
func processar(id int, demora time.Duration) {
	fmt.Printf("tarefa %d: iniciou\n", id)
	time.Sleep(demora)
	fmt.Printf("tarefa %d: concluiu em %s\n", id, demora)
}

func main() {
	tarefas := []time.Duration{
		120 * time.Millisecond,
		80 * time.Millisecond,
		150 * time.Millisecond,
		50 * time.Millisecond,
	}

	inicio := time.Now()
	var wg sync.WaitGroup

	for i, demora := range tarefas {
		wg.Add(1)
		// Capture i e demora por valor no parâmetro da closure.
		// Sem isso, todas as goroutines veriam o último índice do loop.
		go func(id int, d time.Duration) {
			defer wg.Done()
			processar(id, d)
		}(i+1, demora)
	}

	wg.Wait()
	fmt.Printf("todas as tarefas concluídas em %s\n", time.Since(inicio).Round(time.Millisecond))
}
