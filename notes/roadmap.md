# Go Roadmap

Um guia prático e progressivo para aprender Go e construir software de backend e infraestrutura.

## 🧱 1. Fundamentos de Programação

### Lógica de Programação

- Algoritmos e resolução de problemas
- Pseudocódigo
- Linguagens compiladas vs interpretadas

### Conceitos básicos

- Variáveis e constantes
- Tipos de dados
- Comentários
- Entrada e saída no terminal

### Controle de fluxo

- Condicionais
- Estruturas de repetição
- Funções

## 🐹 2. Primeiros Passos com Go

### Ambiente

- Instalação do Go
- `go run`, `go build`, `go fmt`, `go vet`
- Módulos com `go mod init` e `go.mod`
- Editor com `gopls`

### Sintaxe

- Pacotes e imports
- Declaração de variáveis (`var`, `:=`, `const`)
- Tipos básicos: `int`, `float64`, `string`, `bool`, `rune`, `byte`
- Zero values
- Conversão de tipos
- `iota` e constantes enumeradas

### Controle de fluxo em Go

- `if` com declaração curta
- `for` (a única estrutura de repetição)
- `switch` com e sem expressão
- `defer`

### Funções

- Múltiplos retornos
- Retornos nomeados
- Funções variádicas
- Funções como valores e closures

## 📦 3. Tipos Compostos

- Arrays
- Slices (`len`, `cap`, `append`, fatiamento)
- Maps
- Strings, bytes e runes (UTF-8)
- Structs e structs aninhadas
- Ponteiros
- Pacotes `slices` e `maps` da biblioteca padrão

## 🧩 4. Métodos, Interfaces e Composição

- Métodos com receptor por valor e por ponteiro
- Interfaces implícitas
- Interfaces pequenas (`io.Reader`, `io.Writer`, `fmt.Stringer`)
- Interface vazia (`any`) e type assertions
- Type switches
- Composição com embedding
- Visibilidade: exportado vs não exportado

## ⚠️ 5. Tratamento de Erros

- Erros como valores (`error`)
- `errors.New` e `fmt.Errorf`
- Wrapping com `%w`
- `errors.Is` e `errors.As`
- Erros sentinela e tipos de erro customizados
- `panic` e `recover` (e por que evitar)

## 🧬 6. Generics

- Funções genéricas
- Tipos genéricos
- Constraints e `comparable`
- Quando usar e quando não usar generics

## ⚡ 7. Concorrência

- Goroutines
- Channels com e sem buffer
- `select`
- `sync.WaitGroup` e `sync.Mutex`
- Padrões: worker pool, pipeline, fan-in e fan-out
- `context` para cancelamento e timeouts
- Detector de corrida (`-race`)

## 📚 8. Biblioteca Padrão

- `fmt`, `strings`, `strconv`
- `os`, `io`, `bufio`
- `encoding/json`
- `time`
- `log/slog` (logging estruturado)
- `net/http` (cliente e servidor)
- `flag` para CLIs

## 🧪 9. Testes e Qualidade

- `go test` e arquivos `_test.go`
- Testes em tabela (table-driven tests)
- Subtestes com `t.Run`
- Benchmarks e fuzzing
- Cobertura com `-cover`
- `net/http/httptest`
- Linters: `go vet`, `staticcheck`, `golangci-lint`

## 🗂️ 10. Organização de Projetos

- Pacotes e responsabilidades
- Pasta `internal/`
- Layout comum: `cmd/`, `internal/`, `pkg/` (e quando não usar)
- Versionamento semântico de módulos
- Workspaces (`go work`)

## 🌐 11. Desenvolvimento Web e APIs

### HTTP com a biblioteca padrão

- `net/http`, handlers e `ServeMux`
- Roteamento por método e caminho
- Middlewares
- JSON na entrada e na saída
- Validação de entrada

### Frameworks e roteadores

- chi
- Gin
- Echo
- Fiber

### Outras formas de comunicação

- gRPC e Protocol Buffers
- WebSockets
- GraphQL (gqlgen)

## 🗄️ 12. Banco de Dados

- `database/sql`
- Drivers (pgx para PostgreSQL)
- sqlc (código gerado a partir de SQL)
- GORM (ORM)
- Migrações com golang-migrate ou goose
- Redis

## 🔐 13. Segurança

- Autenticação vs autorização
- Hash de senhas com bcrypt
- JWT
- OAuth 2 e OIDC
- `govulncheck` para vulnerabilidades em dependências
- OWASP Top 10

## 📨 14. Mensageria

- Kafka
- RabbitMQ
- NATS

## 📊 15. Observabilidade

- Logging estruturado com `slog`
- Métricas com Prometheus
- Rastreamento distribuído com OpenTelemetry
- Profiling com `pprof`

## ☁️ 16. Cloud, Containers e Deploy

- Build de binários para várias plataformas (`GOOS`, `GOARCH`)
- Imagens Docker multi-stage e imagens mínimas (distroless, scratch)
- Kubernetes
- CI/CD com GitHub Actions
- Graceful shutdown

## 🛠️ 17. Go no Ecossistema de Infraestrutura

- CLIs com Cobra
- Ferramentas de DevOps escritas em Go (Docker, Kubernetes, Terraform)
- Operadores e controllers do Kubernetes
- Plugins do Terraform

## 🎯 18. Tópicos Avançados

- Modelo de memória do Go
- Escape analysis e alocação
- Garbage collector e ajustes (`GOGC`, `GOMEMLIMIT`)
- Pacote `unsafe` e cgo
- Reflexão (`reflect`)
- Performance tuning com benchmarks e `pprof`

## 🚀 Próximos Passos - Da Sintaxe à Produção

1. Domine a base: slices, maps, structs, interfaces e tratamento de erros
2. Escreva testes desde o início com `go test` e testes em tabela
3. Construa uma API com `net/http`, banco de dados e logging com `slog`
4. Aprenda concorrência com goroutines, channels e `context`
5. Empacote com Docker, publique com CI/CD e observe com Prometheus e OpenTelemetry

## 🧠 Dica

> Foque em **projetos reais** ao invés de só teoria.

Sugestões:

- CLI útil para o seu dia a dia
- API REST completa com autenticação e PostgreSQL
- Verificador de links ou crawler concorrente
- Serviço gRPC com observabilidade

🔥 **Regra de ouro:**
Aprenda - Aplique - Ensine - Repita
