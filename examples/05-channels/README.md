# 05 - Channels

Formas de usar channels em Go: comunicação entre goroutines sem compartilhar memória diretamente.

## Sobre

Channel é a fila tipada da linguagem. Este exemplo cobre o essencial, um caso por arquivo - sem worker pool ainda (isso fica no próximo).

## Arquivos

| Arquivo | O que tem |
| ------- | --------- |
| `caso_sem_buffer.go` | Handshake: send bloqueia até o receive |
| `caso_com_buffer.go` | Buffer: envia sem receiver imediato |
| `caso_close_range.go` | Produtor fecha; consumidor faz `range` |
| `caso_select.go` | Primeiro channel pronto + timeout |
| `caso_direcao.go` | `chan<-` (só envia) e `<-chan` (só recebe) |
| `main.go` | Só chama os casos |

## Como executar

```bash
go run .
go run -race .
```

Saída esperada:

```
=== 1) sem buffer (handshake) ===
recebido: ping

=== 2) com buffer ===
len: 3 cap: 3
saiu: 10 20 30

=== 3) close + range ===
item: a
item: b
item: c
channel fechado - range terminou

=== 4) select ===
ganhou: rápido

=== 5) direção (chan<- / <-chan) ===
recebido: olá
```

## Conceitos

### Sem buffer vs com buffer

```go
ch := make(chan string)    // capacidade 0 - sincroniza
ch := make(chan int, 3)    // cabe 3 valores sem receiver
```

Sem buffer = handshake. Com buffer = fila limitada. Buffer cheio sem leitor = bloqueio (deadlock se ninguém ler).

### Quem fecha o channel

Só o **produtor** (quem envia) deve fechar. Fechar duas vezes ou enviar em channel fechado = panic. `for range` termina quando o channel fecha.

### `select`

Espera vários casos; o primeiro pronto executa. `time.After` no `select` é o timeout idiomático.

### Direção na assinatura

```go
func enviar(saida chan<- string, msg string)
func receber(entrada <-chan string) string
```

O tipo comunica a intenção; o compilador barra o uso invertido.

## Experimente

1. Em `caso_com_buffer.go`, descomente o 4º envio - veja o deadlock
2. Em `caso_close_range.go`, remova o `close` - o `range` nunca termina
3. Em `caso_select.go`, tire o valor de `rapido` antes do `select` - deve cair no timeout
4. Rode `go run -race .` - deve passar limpo

## Relacionado

- Anterior: [04 - Goroutines e WaitGroup](../04-goroutines-waitgroup/)
- Próximo: [06 - Worker Pool](../06-worker-pool/)
- Roadmap: seção Concorrência
