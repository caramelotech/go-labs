package main

import "fmt"

// CasoCloseRange: o produtor fecha o channel quando termina.
// O consumidor usa for range até o close - sem busy-wait nem len.
func CasoCloseRange() {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for _, item := range []string{"a", "b", "c"} {
			ch <- item
		}
	}()

	for item := range ch {
		fmt.Println("item:", item)
	}
	fmt.Println("channel fechado - range terminou")
}
