// Demonstra o padrão worker pool: goroutines limitadas, channels para
// jobs/resultados e context para timeout - o combo em que Go se destaca.
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	numWorkers = 3
	timeout    = 400 * time.Millisecond
)

type job struct {
	URL    string
	Demora time.Duration
	OK     bool
}

type resultado struct {
	URL    string
	OK     bool
	Demora time.Duration
	Err    error
}

func verificar(ctx context.Context, j job) resultado {
	timer := time.NewTimer(j.Demora)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return resultado{URL: j.URL, Err: ctx.Err()}
	case <-timer.C:
		return resultado{URL: j.URL, OK: j.OK, Demora: j.Demora}
	}
}

func worker(ctx context.Context, id int, jobs <-chan job, saida chan<- resultado) {
	for j := range jobs {
		fmt.Printf("worker %d: verificando %s\n", id, j.URL)
		saida <- verificar(ctx, j)
	}
}

func main() {
	jobsLista := []job{
		{URL: "https://go.dev", Demora: 80 * time.Millisecond, OK: true},
		{URL: "https://pkg.go.dev", Demora: 120 * time.Millisecond, OK: true},
		{URL: "https://lento.exemplo", Demora: 800 * time.Millisecond, OK: true},
		{URL: "https://fora.exemplo", Demora: 60 * time.Millisecond, OK: false},
		{URL: "https://blog.go.dev", Demora: 90 * time.Millisecond, OK: true},
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	jobs := make(chan job)
	saida := make(chan resultado)

	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(ctx, id, jobs, saida)
		}(w)
	}

	go func() {
		defer close(jobs)
		for _, j := range jobsLista {
			select {
			case <-ctx.Done():
				return
			case jobs <- j:
			}
		}
	}()

	go func() {
		wg.Wait()
		close(saida)
	}()

	var ok, falha, cancelados int
	for r := range saida {
		switch {
		case r.Err != nil:
			cancelados++
			fmt.Printf("CANCELADO  %s  (%v)\n", r.URL, r.Err)
		case r.OK:
			ok++
			fmt.Printf("OK         %s  (%s)\n", r.URL, r.Demora)
		default:
			falha++
			fmt.Printf("FALHA      %s  (%s)\n", r.URL, r.Demora)
		}
	}

	fmt.Printf("\n%d resultados - %d ok, %d falha, %d cancelados (timeout %s, %d workers)\n",
		ok+falha+cancelados, ok, falha, cancelados, timeout, numWorkers)
}
