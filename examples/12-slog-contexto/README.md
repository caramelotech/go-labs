# 12 - slog e contexto

Propagar dados de log na request HTTP: reutiliza o pacote `slogx` do [11 - slog](../11-slog/).

## Fluxo

```text
request
  │
  ▼
MiddlewareLog
  • logger = slog.Default().With(request_id, method, path)
  • slogx.ComLogger(ctx, logger)
  │
  ▼
handlePedido
  • slogx.Append(ctx, "camada", "handler")
  │
  ▼
Service
  • sucesso → slogx.Append(ctx, pedido_id, total, ...)
  • erro    → slogx.Logger(ctx).Error(...)  // herda attrs do middleware
  │
  ▼
MiddlewareLog (volta)
  • slogx.Logger(ctx).Info("request finalizada", status)
```

## Arquivos

| Arquivo | Papel |
| ------- | ----- |
| `middleware.go` | injeta logger via `slogx` e fecha a request |
| `camadas.go` | handler HTTP + service |
| `main.go` | demo com `httptest` ou `serve` na :8080 |
| `main_test.go` | sucesso (attrs no fim) e erro (dados do middleware) |
| `go.mod` | `replace` → `../11-slog` (pacote `slogx`) |

## Como executar

```bash
go run .          # demo automática (sucesso + erro)
go test ./...
go run . serve    # servidor em :8080
```

## Dependência do 11

```go
import "github.com/caramelotech/go-labs/examples/11-slog/slogx"

slogx.SetDefaultJSON(os.Stdout, slog.LevelInfo)
ctx = slogx.ComLogger(ctx, logger)
slogx.Append(ctx, "pedido_id", id)
slogx.Logger(ctx).Error("...", "erro", err.Error())
```

O `go.mod` usa `replace` para apontar ao módulo irmão local - nos labs não precisamos publicar o 11.

## O que observar nos logs (JSON)

**Sucesso** - `Append` nas camadas; middleware emite o fechamento com tudo:

```json
{"msg":"request finalizada","request_id":"req-ok","method":"GET","path":"/pedido","status":200,"camada":"handler","pedido_id":"42","total_centavos":1990,"resultado":"aprovado"}
```

**Erro** - `Error` na service já traz o que o middleware colocou:

```json
{"level":"ERROR","msg":"processar pedido","request_id":"req-erro","method":"GET","path":"/pedido","camada":"handler","pedido_id":"99","erro":"pagamento recusado"}
```

## Experimente

1. Rode antes o [11 - slog](../11-slog/) para ver `NewJSON` / `With` / `Group`
2. Tire um `Append` e compare o log final
3. No erro, troque `slogx.Logger(ctx)` por `slog.Default()` - some o `request_id`

## Relacionado

- Anterior: [11 - slog](../11-slog/)
- [07 - Context](../07-context/) - timeout, cancel e valores no context
- Docs: [log/slog](https://pkg.go.dev/log/slog)
