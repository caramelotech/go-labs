# 06 - Worker Pool

O padrão em que Go brilha: limitar concorrência com workers, distribuir trabalho por channels e cancelar com `context`.

## Sobre

Lançar uma goroutine por item escala mal (milhares de conexões, rate limits, CPU). O worker pool:

1. Cria **N workers** fixos
2. Enfileira jobs em um channel
3. Coleta resultados em outro channel
4. Usa `context.WithTimeout` para não ficar preso

É a base do [Projeto 04 - Verificador de Links Concorrente](../../notes/pratica/02-projetos.md), sem depender de rede real (as "verificações" são timers).

## Como executar

```bash
go run .
go run -race .
```

Saída típica (ordem pode variar):

```
worker 1: verificando https://go.dev
worker 2: verificando https://pkg.go.dev
worker 3: verificando https://lento.exemplo
OK         https://go.dev  (80ms)
OK         https://pkg.go.dev  (120ms)
worker 1: verificando https://fora.exemplo
worker 2: verificando https://blog.go.dev
FALHA      https://fora.exemplo  (60ms)
OK         https://blog.go.dev  (90ms)
CANCELADO  https://lento.exemplo  (context deadline exceeded)

5 resultados - 3 ok, 1 falha, 1 cancelados (timeout 400ms, 3 workers)
```

## Conceitos

### Channel como fila

```go
jobs := make(chan job)
saida := make(chan resultado)
```

Workers leem de `jobs` até o channel fechar (`for j := range jobs`). O coordenador fecha `saida` só depois do `WaitGroup` - evita send em channel fechado.

### `select` + `context`

```go
select {
case <-ctx.Done():
	return resultado{URL: j.URL, Err: ctx.Err()}
case <-selecionar.C:
	return resultado{URL: j.URL, OK: j.OK, Demora: j.Demora}
}
```

O worker reage ao timeout sem busy-wait. Sempre faça `defer timer.Stop()` para não vazar timers.

### Limite de paralelismo

`numWorkers = 3` com 5 jobs: no máximo 3 verificações ao mesmo tempo. Ajuste o número e o `timeout` para ver o efeito.

## Experimente

1. Suba `timeout` para `1s` - o link lento deixa de cancelar
2. Mude `numWorkers` para `1` e compare o tempo total
3. Rode `go run -race .` - deve passar limpo
4. Implemente a versão real com `http.Client` e `Head` (Projeto 04)

## Relacionado

- Anterior: [05 - Channels](../05-channels/)
- Próximo: [07 - Context](../07-context/)
- [Projeto 04 - Verificador de Links](../../notes/pratica/02-projetos.md)
- Roadmap: seção Concorrência
