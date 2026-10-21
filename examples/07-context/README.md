# 07 - Context entre camadas

Como passar e usar `context.Context` em Go: timeout, cancelamento (cliente fechou a conexão), dados da request e transação.

## Arquivos

| Arquivo | O que tem |
| ------- | --------- |
| `trabalho.go` | Trabalho lento que respeita `ctx.Done()` |
| `camadas.go` | Handler → Service → Repo recebendo `ctx` |
| `ctx_valores.go` | `request_id` via `WithValue` (chave tipada) |
| `ctx_transacao.go` | Transação mock propagada no context |
| `caso_*.go` | Um caso de uso por arquivo |
| `main.go` | Só chama os casos |

## Como executar

```bash
go run .
go test -v
```

## O que cada caso mostra

1. **Timeout** - `WithTimeout`: a request morre no prazo
2. **Close connection (simulado)** - `cancel()` = cliente desconectou; as camadas param
3. **Close connection (HTTP)** - `r.Context()` + client cancela a request
4. **Dados** - `request_id` desce handler → repo sem poluir assinaturas
5. **Transação** - service abre `Tx`, repo reutiliza via `TxFrom(ctx)`

## Ideia central

```text
context vai como PRIMEIRO argumento: Func(ctx context.Context, ...)

timeout / cancel  →  ctx.Done() / ctx.Err()
dados da request  →  WithValue + chave tipada privada
transação         →  WithValue (mesmo padrão dos dados)

NÃO use context para passar dependências obrigatórias
(isso é parâmetro normal ou struct de serviço).
```

## Padrão nas camadas

```go
func Handler(ctx context.Context, ...) error {
	return Service(ctx, ...)
}

func Service(ctx context.Context, ...) error {
	return Repo(ctx, ...)
}
```

Quem cria o context (HTTP middleware, `main`, teste) define timeout/cancel.
Quem recebe só propaga e consulta `ctx.Err()` / `<-ctx.Done()`.

## Relacionado

- Anterior: [06 - Worker Pool](../06-worker-pool/)
- Próximo: [11 - slog](../11-slog/) → [12 - slog e contexto](../12-slog-contexto/)
- [Projeto 04 - Verificador de Links](../../notes/pratica/02-projetos.md)
- Roadmap: seção Concorrência (`context`)
