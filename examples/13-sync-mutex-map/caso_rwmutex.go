package main

import (
	"fmt"
	"sync"
)

// Catalogo é um mapa com muitas leituras e poucas escritas.
// RWMutex permite vários Get em paralelo; Set continua exclusivo.
type Catalogo struct {
	mu    sync.RWMutex
	itens map[string]string
}

func NovoCatalogo() *Catalogo {
	return &Catalogo{itens: make(map[string]string)}
}

func (c *Catalogo) Set(chave, valor string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.itens[chave] = valor
}

func (c *Catalogo) Get(chave string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.itens[chave]
	return v, ok
}

// CasoRWMutex: 1 escritor e vários leitores concorrentes.
func CasoRWMutex() {
	cat := NovoCatalogo()
	cat.Set("curso", "go-labs")

	var wg sync.WaitGroup
	const leitores = 6

	wg.Add(leitores + 1)
	go func() {
		defer wg.Done()
		cat.Set("modulo", "sync")
	}()

	for i := 0; i < leitores; i++ {
		go func(id int) {
			defer wg.Done()
			curso, _ := cat.Get("curso")
			fmt.Printf("leitor %d: curso=%s\n", id, curso)
		}(i + 1)
	}
	wg.Wait()

	mod, ok := cat.Get("modulo")
	fmt.Printf("depois do Set: modulo=%q ok=%v\n", mod, ok)
}
