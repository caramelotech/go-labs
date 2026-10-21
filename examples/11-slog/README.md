# 11 - slog

O básico de `log/slog` (biblioteca padrão desde Go 1.21): níveis, attrs, `With`, grupos e handlers JSON/texto.

O pacote reutilizável fica em `slogx/` - o exemplo [12 - slog e contexto](../12-slog-contexto/) importa o mesmo pacote.

## Arquivos

| Arquivo | Papel |
| ------- | ----- |
| `slogx/handler.go` | `NewJSON`, `NewText`, `SetDefaultJSON` |
| `slogx/context.go` | `ComLogger`, `Logger`, `Append`, `Attrs` (usado no 12) |
| `main.go` | Demo dos níveis, `With` e `Group` |

## Como executar

```bash
go run .
```

## Conceitos

### Logger + Handler

```go
logger := slogx.NewJSON(os.Stdout, slog.LevelInfo)
logger.Info("pedido criado", "pedido_id", "42")
```

Cada linha é um evento com nível, mensagem e pares chave/valor. JSON é o formato usual em produção; texto ajuda no terminal local.

### Níveis

`Debug` < `Info` < `Warn` < `Error`. O `Level` do handler filtra o que aparece - com `LevelInfo`, `Debug` some.

### `With`

```go
reqLog := logger.With("request_id", "req-1")
reqLog.Info("cobrando") // sempre inclui request_id
```

Attrs comuns da request ficam no logger filho - sem repetir em cada chamada.

### `Group`

```go
slog.Group("user", "id", "u-9", "plano", "pro")
```

No JSON vira um objeto aninhado `"user":{"id":"u-9","plano":"pro"}`.

### Context (próximo exemplo)

`slogx.ComLogger` / `Logger` / `Append` colocam o logger na request e acumulam attrs entre camadas. Veja [12 - slog e contexto](../12-slog-contexto/).

## Experimente

1. Mude `NewJSON(..., slog.LevelDebug)` para `LevelWarn` - só Warn e Error saem
2. Troque `NewJSON` por `NewText` no bloco JSON
3. Adicione mais um `With("versão", "1")` em cima do `reqLog`

## Relacionado

- Próximo: [12 - slog e contexto](../12-slog-contexto/)
- Docs: [log/slog](https://pkg.go.dev/log/slog)
