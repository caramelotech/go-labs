# Channels e Select

O `WaitGroup` resolve "espere todo mundo terminar", mas não deixa as goroutines trocarem dados. Para isso existem os **channels**: filas tipadas por onde uma goroutine envia valores e outra recebe, com a sincronização embutida.

Um dos provérbios de Go resume a filosofia: _"Don't communicate by sharing memory, share memory by communicating"_. Em vez de várias goroutines mexendo na mesma variável protegida por trava, os dados passam de mão em mão pelo channel.

## Criando e usando um channel

```go
ch := make(chan string) // channel de strings

go func() {
	ch <- "ping" // envia
}()

msg := <-ch // recebe
fmt.Println(msg)
```

A seta indica a direção do dado: `ch <- v` coloca no channel, `<-ch` tira dele. O zero value de um channel é `nil`, então ele sempre precisa de `make`.

## Sem buffer: um aperto de mão

Um channel criado com `make(chan T)` não tem espaço interno. O envio **bloqueia** até alguém receber, e o recebimento bloqueia até alguém enviar. As duas goroutines se encontram naquele ponto, como num aperto de mão. Por isso um channel sem buffer serve tanto para passar dados quanto para sincronizar.

## Com buffer: uma fila limitada

`make(chan T, n)` cria um channel com espaço para `n` valores. Os envios só bloqueiam quando o buffer está cheio, e os recebimentos quando está vazio:

```go
ch := make(chan int, 3)
ch <- 10
ch <- 20
ch <- 30
// ch <- 40 bloquearia: buffer cheio e ninguém lendo

fmt.Println(len(ch), cap(ch)) // 3 3
fmt.Println(<-ch, <-ch, <-ch) // 10 20 30
```

Se todas as goroutines ficarem bloqueadas esperando um channel, o runtime detecta e encerra o programa com `fatal error: all goroutines are asleep - deadlock!`.

Buffer não é solução para problema de sincronização. Comece sem buffer e só adicione quando houver um motivo concreto, como limitar quantas tarefas ficam na fila.

## Fechando channels e `range`

O produtor chama `close(ch)` para avisar que não vai mais enviar nada. O consumidor percorre com `for range`, que termina sozinho quando o channel fecha e esvazia:

```go
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
```

Regras de fechamento:

- **Quem envia fecha**, nunca quem recebe
- Enviar num channel fechado causa panic, assim como fechar duas vezes
- Receber de um channel fechado devolve o zero value na hora; use `v, ok := <-ch` para saber se `ok` é `false` (fechado)
- Fechar não é obrigatório: só feche quando o receptor precisa saber que acabou, como no `range`

## Direção nos parâmetros

Na assinatura de uma função, dá para restringir o channel a só envio (`chan<- T`) ou só recebimento (`<-chan T`):

```go
func enviar(saida chan<- string, msg string) {
	saida <- msg
}

func receber(entrada <-chan string) string {
	return <-entrada
}
```

Um `chan string` comum é convertido automaticamente na chamada. A assinatura documenta o papel de cada função, e o compilador barra o uso invertido, como tentar receber de um `chan<-`.

## `select`

`select` espera várias operações de channel ao mesmo tempo e executa a **primeira que ficar pronta**. Se várias estiverem prontas, ele escolhe uma ao acaso:

```go
select {
case msg := <-rapido:
	fmt.Println("ganhou:", msg)
case msg := <-lento:
	fmt.Println("ganhou:", msg)
case <-time.After(100 * time.Millisecond):
	fmt.Println("timeout")
}
```

`time.After` devolve um channel que recebe um valor depois do tempo indicado, e é a forma idiomática de colocar timeout numa espera.

Com um `case default`, o `select` não bloqueia: se nenhum channel estiver pronto, executa o `default` na hora. É útil para tentar enviar sem travar, descartando o valor se a fila estiver cheia.

O `select` é a base do cancelamento: na nota de [context](/labs/go/concorrencia/04-context/), um dos `case` será `<-ctx.Done()`.

## Channel ou WaitGroup?

| Preciso de...                                        | Use                                                                    |
| ---------------------------------------------------- | ---------------------------------------------------------------------- |
| Esperar N tarefas sem coletar resultado              | `sync.WaitGroup`                                                       |
| Passar dados de uma goroutine para outra             | channel                                                                |
| Sinalizar "acabou" para vários receptores de uma vez | `close` num channel                                                    |
| Esperar o primeiro de vários eventos, ou um timeout  | `select`                                                               |
| Proteger uma variável compartilhada                  | [`sync.Mutex`](/labs/go/concorrencia/05-mutex-e-estado-compartilhado/) |

## Exemplo no repositório

O exemplo [18-channels](https://github.com/caramelotech/go-labs/tree/main/examples/18-channels) tem um caso por arquivo: sem buffer (`caso_sem_buffer.go`), com buffer (`caso_com_buffer.go`), `close` com `range` (`caso_close_range.go`), `select` com timeout (`caso_select.go`) e direção (`caso_direcao.go`).

```bash
cd examples/18-channels
go run -race .
```

Descomente o quarto envio em `caso_com_buffer.go` para ver o deadlock, e remova o `close` em `caso_close_range.go` para ver o `range` nunca terminar.

## Referências

- [Effective Go - Channels](https://go.dev/doc/effective_go#channels) - Go Team, inglês
- [Share Memory By Communicating](https://go.dev/blog/codelab-share) - Andrew Gerrand, Go Blog, inglês
- [A Tour of Go - Channels](https://go.dev/tour/concurrency/2) - Go Team, inglês
- [Go Concurrency Patterns](https://go.dev/talks/2012/concurrency.slide) - Rob Pike, Google I/O 2012, inglês
