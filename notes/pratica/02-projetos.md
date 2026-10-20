# Projetos

Mini projetos para aplicar o que foi aprendido de forma integrada.

## Como usar

1. Leia o enunciado do projeto
2. Planeje antes de codificar - defina pacotes, tipos, funções e fluxo
3. Implemente incrementalmente
4. Escreva testes com `go test` para as partes principais antes de avançar

## Projeto 01 - Calculadora de Linha de Comando

**Nível:** iniciante

**Objetivo:** Construir uma calculadora interativa via terminal.

**Descrição:**
Crie um programa que leia dois números e uma operação digitados pelo usuário e exiba o resultado.

**Requisitos:**

- [ ] Suporta adição (`+`), subtração (`-`), multiplicação (`*`) e divisão (`/`)
- [ ] Lê a entrada com `bufio.Scanner` ou `fmt.Scan`
- [ ] Trata divisão por zero e operador inválido com mensagens claras
- [ ] A lógica de cálculo fica em uma função separada e testada com `go test`

**Exemplo de uso:**

```
Informe o primeiro número: 10
Informe o operador (+, -, *, /): *
Informe o segundo número: 4
Resultado: 40
```

## Projeto 02 - Lista de Tarefas (CLI)

**Nível:** intermediário

**Objetivo:** Criar uma ferramenta de linha de comando que persiste dados em arquivo.

**Descrição:**
Construa um gerenciador de tarefas que recebe subcomandos pela linha de comando e salva as tarefas em um arquivo JSON.

**Requisitos:**

- [ ] Subcomandos `add`, `list`, `done` e `remove`
- [ ] Tarefas representadas por uma `struct` com `ID`, `Titulo` e `Concluida`
- [ ] Persistência em `tarefas.json` usando o pacote `encoding/json`
- [ ] Argumentos lidos com `os.Args` ou com o pacote `flag`
- [ ] Erros de leitura e escrita do arquivo tratados e exibidos ao usuário

**Exemplo de uso:**

```
$ tarefas add "Estudar goroutines"
Tarefa 1 adicionada.

$ tarefas list
[ ] 1 - Estudar goroutines

$ tarefas done 1
Tarefa 1 concluída.
```

## Projeto 03 - API REST de Livros

**Nível:** intermediário

**Objetivo:** Construir uma API HTTP usando apenas a biblioteca padrão.

**Descrição:**
Crie uma API que permite cadastrar, listar, buscar e remover livros, guardando os dados em memória.

**Requisitos:**

- [ ] Servidor com `net/http` e roteamento por método e caminho (`mux.HandleFunc("GET /livros/{id}", ...)`)
- [ ] Rotas `GET /livros`, `GET /livros/{id}`, `POST /livros` e `DELETE /livros/{id}`
- [ ] Entrada e saída em JSON
- [ ] Acesso ao armazenamento em memória protegido com `sync.Mutex`
- [ ] Códigos de status corretos (`200`, `201`, `204`, `400`, `404`)
- [ ] Testes dos handlers com o pacote `net/http/httptest`

**Exemplo de entregável:**

```
$ curl -X POST localhost:8080/livros -d '{"titulo":"A Linguagem Go","autor":"Donovan e Kernighan"}'
{"id":1,"titulo":"A Linguagem Go","autor":"Donovan e Kernighan"}
```

## Projeto 04 - Verificador de Links Concorrente

**Nível:** avançado

**Objetivo:** Aplicar goroutines, channels e controle de concorrência.

**Descrição:**
Crie um programa que recebe uma lista de URLs (de um arquivo ou da linha de comando) e verifica, de forma concorrente, quais estão no ar.

**Requisitos:**

- [ ] Cada URL é verificada em uma goroutine
- [ ] Limita o número de verificações simultâneas (worker pool ou semáforo com channel)
- [ ] Usa `context.WithTimeout` para não ficar preso em URLs lentas
- [ ] Coleta os resultados por um channel e imprime um resumo no final
- [ ] Passa sem alertas em `go test -race`

**Exemplo de entregável:**

```
OK    200  https://go.dev          (120ms)
ERRO  404  https://exemplo.com/x   (85ms)
FALHA  -   https://lento.exemplo   (timeout)

3 verificadas - 1 ok, 2 com problema
```
