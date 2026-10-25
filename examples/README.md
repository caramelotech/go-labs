# Examples

Este diretório reúne exemplos de código para acompanhar o conteúdo do Go Labs, ordenados do mais simples ao mais complexo.

Cada exemplo é um módulo Go independente, com seu próprio `go.mod`.

## Exemplos disponíveis

| Exemplo | Conceito |
| ------- | -------- |
| [01-hello-world](01-hello-world/) | Estrutura básica de um programa Go |
| [02-variaveis-tipos](02-variaveis-tipos/) | `var`, `:=`, `const` e zero values |
| [03-if-switch](03-if-switch/) | `if` com declaração curta e `switch` |
| [04-for](04-for/) | Única estrutura de repetição (`for` / `range`) |
| [05-funcoes](05-funcoes/) | Múltiplos retornos, nomeados e variádicas |
| [06-slices](06-slices/) | Slice, `append`, fatiamento e `copy` |
| [07-maps](07-maps/) | Map: insert, comma ok, `delete`, `range` |
| [08-structs](08-structs/) | Structs e structs aninhadas |
| [09-ponteiros](09-ponteiros/) | `&`, `*` e mutação via ponteiro |
| [10-visibilidade](10-visibilidade/) | Exportado (maiúscula) vs não exportado |
| [11-metodos](11-metodos/) | Métodos com value e pointer receiver |
| [12-nomes-de-pacotes](12-nomes-de-pacotes/) | Nomes de pacote e identificadores (Effective Go) |
| [13-interface-consumidor](13-interface-consumidor/) | Interface no consumidor e injeção de dependência |
| [14-receivers](14-receivers/) | Value vs pointer receiver - preferir `(s Servico)` |
| [15-erros-idiomaticos](15-erros-idiomaticos/) | Sentinela, erro estruturado, `%w`, `errors.Is` / `As` |
| [16-testes-em-tabela](16-testes-em-tabela/) | Table-driven tests com `t.Run` |
| [17-goroutines-waitgroup](17-goroutines-waitgroup/) | Concorrência básica com goroutines e `WaitGroup` |
| [18-channels](18-channels/) | Channels: buffer, close+range, `select` e direção |
| [19-worker-pool](19-worker-pool/) | Workers limitados, channels e `context` |
| [20-context](20-context/) | Context entre camadas: timeout, cancel, valores e Tx |
| [21-sync-mutex-map](21-sync-mutex-map/) | `Mutex`, `RWMutex` e `sync.Map` |
| [22-slog](22-slog/) | `log/slog` básico + pacote reutilizável `slogx` |
| [23-slog-contexto](23-slog-contexto/) | Middleware HTTP com `slogx` (attrs no context) |

## Como usar

1. Escolha um exemplo
2. Leia o README do exemplo
3. Entre na pasta e execute com `go run .`
4. Nos exemplos com teste, rode também `go test` / `go test -race`
5. Para validar todos de uma vez (com race detector): `make test`
6. Modifique, teste e experimente variações

## Trilha sugerida

Siga a numeração em ordem:

1. **Sintaxe e tipos:** 01 → 11
2. **Pacotes, interfaces e erros:** 12 → 16
3. **Concorrência:** 17 → 21
4. **HTTP / observabilidade:** 22 → 23
5. Depois, aplique no [Projeto 04 - Verificador de Links](../notes/pratica/02-projetos.md)

## Prática guiada

- [Exercícios](../notes/pratica/01-exercicios.md)
- [Projetos](../notes/pratica/02-projetos.md)
