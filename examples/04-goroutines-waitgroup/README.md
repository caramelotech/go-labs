# 04 - Goroutines e WaitGroup

Primeiro contato com concorrência em Go: lançar trabalho em paralelo e esperar todos terminarem.

## Sobre

Goroutines são baratas. O erro clássico é lançá-las e esquecer de sincronizar - o `main` termina e o processo acaba antes do trabalho.

`sync.WaitGroup` resolve o caso "espere N tarefas":

1. `Add(1)` antes de cada `go`
2. `Done()` (em geral via `defer`) quando a tarefa termina
3. `Wait()` no coordenador

## Como executar

```bash
go run .
```

Saída típica (ordem das linhas intermediárias pode variar):

```
tarefa 1: iniciou
tarefa 2: iniciou
tarefa 3: iniciou
tarefa 4: iniciou
tarefa 2: concluiu em 80ms
tarefa 4: concluiu em 50ms
tarefa 1: concluiu em 120ms
tarefa 3: concluiu em 150ms
todas as tarefas concluídas em 150ms
```

O tempo total fica perto da tarefa mais lenta, não da soma - esse é o ganho.

## Conceitos

### Captura de variáveis no loop

```go
go func(id int, d time.Duration) {
	defer wg.Done()
	processar(id, d)
}(i+1, demora)
```

Passar `i` e `demora` como parâmetros evita o bug clássico de todas as goroutines compartilharem a mesma variável do `for`.

### `defer wg.Done()`

Garante o `Done` mesmo se `processar` der panic. Sem `WaitGroup` (ou channel), o programa termina cedo demais.

## Experimente

1. Comente `wg.Wait()` e rode de novo - quase nada aparece
2. Remova os parâmetros da closure e use `i`/`demora` direto (no Go recente o loop var é por iteração, mas o hábito de passar por parâmetro continua seguro e claro)
3. Rode com `go run -race .` - não deve haver data race neste exemplo

## Relacionado

- Próximo: [05 - Channels](../05-channels/)
- Memória compartilhada: [13 - Mutex, RWMutex e sync.Map](../13-sync-mutex-map/)
- Roadmap: seção Concorrência
