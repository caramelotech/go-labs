package main

import (
	"fmt"
	"sync"
)

// CasoSyncMap: dois cenários em que sync.Map encaixa bem
// (documentação oficial do pacote sync):
//
//  1. chave escrita uma vez, lida muitas vezes (cache que só cresce)
//  2. goroutines mexem em conjuntos de chaves disjuntos
//
// Não substitui map+Mutex no caso geral (tipos any, API diferente).
func CasoSyncMap() {
	var cache sync.Map

	// 1) write-once, read-many
	cache.Store("versão", "1.22")
	if v, ok := cache.Load("versão"); ok {
		fmt.Println("cache versão:", v)
	}

	// 2) cada goroutine escreve só a própria chave
	var wg sync.WaitGroup
	const n = 4
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg.Done()
			chave := fmt.Sprintf("worker-%d", id)
			cache.Store(chave, id*10)
		}(i + 1)
	}
	wg.Wait()

	cache.Range(func(chave, valor any) bool {
		fmt.Printf("  %v = %v\n", chave, valor)
		return true // continua
	})
}
