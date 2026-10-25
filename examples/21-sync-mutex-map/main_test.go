package main

import (
	"sync"
	"testing"
)

func TestContadorSemRace(t *testing.T) {
	var c Contador
	var wg sync.WaitGroup

	const goroutines = 16
	const porGoroutine = 500
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

	want := goroutines * porGoroutine
	if got := c.Valor(); got != want {
		t.Fatalf("Valor() = %d, queria %d", got, want)
	}
}

func TestCatalogoGetSet(t *testing.T) {
	cat := NovoCatalogo()
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		cat.Set("a", "1")
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			cat.Get("a")
		}
	}()
	wg.Wait()

	if v, ok := cat.Get("a"); !ok || v != "1" {
		t.Fatalf("Get(a) = %q, %v", v, ok)
	}
}

func TestSyncMapLoadOrStore(t *testing.T) {
	var m sync.Map

	atual, carregou := m.LoadOrStore("cfg", "prod")
	if carregou {
		t.Fatal("primeira vez não deveria ter carregado valor antigo")
	}
	if atual != "prod" {
		t.Fatalf("atual = %v", atual)
	}

	atual, carregou = m.LoadOrStore("cfg", "dev")
	if !carregou {
		t.Fatal("segunda vez deveria carregar o valor já guardado")
	}
	if atual != "prod" {
		t.Fatalf("não deveria sobrescrever: %v", atual)
	}
}
