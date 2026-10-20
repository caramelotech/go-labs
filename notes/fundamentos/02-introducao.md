# Introdução ao Go

Go (também chamada de **Golang**, por causa do antigo domínio golang.org) é uma linguagem de programação de código aberto criada no Google por Robert Griesemer, Rob Pike e Ken Thompson. O projeto começou em 2007, foi anunciado publicamente em 2009 e chegou à versão 1.0 em 2012.

## Por que o Go foi criado

No fim dos anos 2000, o Google tinha bases de código enormes em C++ e Java. Compilar levava muito tempo, as dependências eram difíceis de gerenciar e cada equipe escrevia código de um jeito diferente. O Go nasceu para resolver problemas de **engenharia de software em escala**, não para ser a linguagem com mais recursos.

Os objetivos eram:

- **Compilação rápida**, mesmo em projetos grandes
- **Código fácil de ler**, mesmo escrito por outra pessoa
- **Concorrência simples**, para aproveitar processadores com muitos núcleos
- **Ferramentas padronizadas**, para que todo projeto Go se pareça com qualquer outro

## Características principais

### Compilada e com tipagem estática

O código Go é compilado direto para código de máquina. Os tipos são verificados em tempo de compilação, então muitos erros aparecem antes de o programa rodar. Ao mesmo tempo, a **inferência de tipos** deixa o código enxuto:

```go
nome := "Caramelo" // o compilador sabe que é string
idade := 3         // e que isto é int
```

### Simplicidade deliberada

A linguagem tem poucas palavras-chave (25) e evita recursos que tornam o código difícil de entender. Não existe herança de classes, sobrecarga de métodos ou exceções. Em vez disso, Go usa:

- **Structs** e **composição** no lugar de classes e herança
- **Interfaces implícitas** - um tipo implementa uma interface só por ter os métodos certos, sem declarar `implements`
- **Erros como valores** - funções retornam um `error` que você verifica explicitamente

```go
arquivo, err := os.Open("dados.txt")
if err != nil {
	return err
}
```

### Concorrência nativa

Concorrência faz parte da linguagem. Uma **goroutine** é uma função que roda de forma concorrente, iniciada com a palavra-chave `go`. Goroutines são muito leves - um programa pode ter milhares delas. Para se comunicarem, usam **channels**:

```go
mensagens := make(chan string)

go func() {
	mensagens <- "olá da goroutine"
}()

fmt.Println(<-mensagens)
```

Você vai ver goroutines e channels com calma mais adiante. Por enquanto, guarde a ideia: em Go, fazer várias coisas ao mesmo tempo é algo do dia a dia.

### Garbage collector

O Go gerencia a memória automaticamente com um coletor de lixo projetado para pausas curtas. Você não precisa liberar memória manualmente, mas ainda tem controle sobre ponteiros e sobre como os dados são organizados.

### Um binário, sem dependências

`go build` gera um único executável que já contém tudo que o programa precisa. Basta copiar o arquivo para o servidor e executar - não é preciso instalar uma máquina virtual ou um interpretador. Também é fácil gerar binários para outros sistemas:

```bash
GOOS=linux GOARCH=amd64 go build
```

### Ferramentas incluídas

O comando `go` já traz formatação (`go fmt`), testes (`go test`), benchmarks, análise estática (`go vet`), gerenciamento de dependências (`go mod`) e documentação (`go doc`). Você não precisa escolher entre várias ferramentas concorrentes para começar.

## Onde o Go é usado

Go é muito popular em **infraestrutura, cloud e backend**. Algumas ferramentas conhecidas escritas em Go:

- **Docker** e **Kubernetes** - containers e orquestração
- **Terraform** - infraestrutura como código
- **Prometheus** e **Grafana** (backend) - monitoramento e observabilidade
- **Hugo** - gerador de sites estáticos

Também é comum em APIs e microsserviços, ferramentas de linha de comando (CLIs) e serviços de rede de alta performance.

## Anatomia de um programa Go

```go
package main

import "fmt"

func main() {
	fmt.Println("Olá, Caramelo Tech!")
}
```

- **`package main`** - todo arquivo Go pertence a um pacote. O pacote `main` é especial: ele gera um programa executável.
- **`import "fmt"`** - importa o pacote `fmt` da biblioteca padrão, usado para formatar e imprimir texto.
- **`func main()`** - o ponto de entrada do programa. A execução começa aqui.
- **`fmt.Println`** - chama a função `Println` do pacote `fmt`. Nomes que começam com letra maiúscula são **exportados**, ou seja, visíveis fora do pacote.

Repare que não há ponto e vírgula no fim das linhas: o compilador insere automaticamente. Por isso a chave `{` precisa ficar na mesma linha do `func` - colocá-la na linha de baixo é erro de compilação.

## `go run` vs `go build`

```bash
go run .     # compila em uma pasta temporária e executa
go build     # gera o executável na pasta atual
./ola-go     # executa o binário (ola-go.exe no Windows)
```

Use `go run` enquanto está desenvolvendo e `go build` quando quiser distribuir o programa.

## O jeito Go de escrever código

A comunidade Go valoriza código **claro em vez de esperto**. Alguns princípios que você vai encontrar sempre:

- O código é formatado com `gofmt` - sem discussões sobre estilo
- Variáveis e imports não usados são **erro de compilação**, não aviso
- Nomes curtos para escopos pequenos (`i`, `err`, `r`), nomes descritivos para escopos maiores
- Tratar erros explicitamente, logo depois de cada chamada que pode falhar

Uma frase popular entre os Gophers (como são chamadas as pessoas que programam em Go) resume bem: *"Clear is better than clever"* - claro é melhor que esperto.

## Referências

- [Documentação oficial do Go](https://go.dev/doc/)
- [A Tour of Go](https://go.dev/tour/)
- [Go at Google: Language Design in the Service of Software Engineering - Rob Pike](https://go.dev/talks/2012/splash.article)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go FAQ - Origins](https://go.dev/doc/faq#Origins)
