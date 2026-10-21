# Examples

Este diretório reúne exemplos de código e projetos já implementados para acompanhar o conteúdo do Go Labs.

## Exemplos disponíveis

| Exemplo | Conceito |
| ------- | -------- |
| [01-hello-world](01-hello-world/) | Estrutura básica de um programa Go |
| [08-nomes-de-pacotes](08-nomes-de-pacotes/) | Nomes de pacote e identificadores (Effective Go) |
| [09-interface-consumidor](09-interface-consumidor/) | Interface no consumidor e injeção de dependência no `main` |
| [10-receivers](10-receivers/) | Value vs pointer receiver - preferir `(s Servico)` |
| [02-erros-idiomaticos](02-erros-idiomaticos/) | Sentinela, erro estruturado, `%w`, `errors.Is` e `errors.As` |
| [03-testes-em-tabela](03-testes-em-tabela/) | Table-driven tests com `t.Run` |
| [04-goroutines-waitgroup](04-goroutines-waitgroup/) | Concorrência básica com goroutines e `WaitGroup` |
| [05-channels](05-channels/) | Channels: buffer, close+range, `select` e direção |
| [06-worker-pool](06-worker-pool/) | Workers limitados, channels e `context` |
| [07-context](07-context/) | Context entre camadas: timeout, cancel, valores e Tx |
| [13-sync-mutex-map](13-sync-mutex-map/) | `Mutex`, `RWMutex` e `sync.Map` - quando usar cada um |
| [11-slog](11-slog/) | `log/slog` básico + pacote reutilizável `slogx` |
| [12-slog-contexto](12-slog-contexto/) | Middleware HTTP com `slogx` (attrs no context) |

## Como usar

1. Escolha um exemplo implementado
2. Leia o README do exemplo
3. Entre na pasta do exemplo e execute com `go run .`
4. Nos exemplos com teste, rode também `go test` / `go test -race`
5. Para validar todos de uma vez (com race detector): `make test`
6. Modifique, teste e experimente variações

Cada exemplo é um módulo Go independente, com seu próprio `go.mod`.

## Trilha sugerida

1. **Fundamentos:** 01 → 08 → 09 → 10 → 02 → 03
2. **Concorrência (onde Go se destaca):** 04 → 05 → 06 → 07 → 13
3. **HTTP / observabilidade:** 11 → 12
4. Depois, aplique no [Projeto 04 - Verificador de Links](../notes/pratica/02-projetos.md)

## Prática guiada

As sugestões de exercícios e projetos ficam publicadas no hub:

- [Exercícios](../notes/pratica/01-exercicios.md)
- [Projetos](../notes/pratica/02-projetos.md)
