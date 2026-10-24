package main

import (
	"fmt"
	"sync"
)

// Contador protege um int com Mutex exclusivo.
// Use Mutex quando leituras e escritas se misturam sem vantagem clara
// de vários leitores ao mesmo tempo.
//
// mu nunca pode ser copiado - por isso o receiver é sempre ponteiro.
type Contador struct {
	mu sync.Mutex
	n  int
}

func (c *Contador) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *Contador) Valor() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

// CasoMutex: N goroutines incrementam o mesmo contador com segurança.
func CasoMutex() {
	var c Contador
	var wg sync.WaitGroup

	const goroutines = 8
	const porGoroutine = 1000

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < porGoroutine; j++ {
				c.Inc()
			}
		}()
	}
	wg.Wait()

	fmt.Printf("contador=%d (esperado %d)\n", c.Valor(), goroutines*porGoroutine)
}
