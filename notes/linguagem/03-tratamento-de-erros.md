# Tratamento de Erros

Go não tem exceções. Uma função que pode falhar devolve um valor do tipo `error` como último retorno, e quem chama decide o que fazer com ele, ali mesmo, na linha seguinte.

## Erros são valores

`error` é uma interface com um método só:

```go
type error interface {
	Error() string
}
```

O padrão de uso aparece em todo código Go:

```go
idade, err := strconv.Atoi(texto)
if err != nil {
	return 0, fmt.Errorf("converter idade: %w", err)
}
// daqui para baixo, idade é válida
```

O `if err != nil` repetido incomoda no começo, mas tem uma vantagem: o caminho de falha fica visível no código, e não escondido num `catch` três funções acima. O caminho feliz segue sem indentação, e cada erro é tratado assim que aparece.

## Criando erros

```go
errors.New("saldo insuficiente")
fmt.Errorf("pedido %s não encontrado", id)
```

As mensagens de erro seguem uma convenção: **começam com minúscula e não terminam com pontuação**, porque costumam ser concatenadas em outras mensagens (`buscar pedido: conectar ao banco: timeout`).

## Envolvendo erros com `%w`

Conforme o erro sobe pelas camadas, cada uma acrescenta contexto. O verbo `%w` do `fmt.Errorf` **envolve** (wrap) o erro original, guardando uma referência a ele:

```go
func BuscarUsuario(id string) (string, error) {
	if id != "42" {
		return "", fmt.Errorf("buscar usuario %q: %w", id, ErrNaoEncontrado)
	}
	return "caramelo", nil
}
```

O resultado é uma cadeia de erros. A mensagem mostra o caminho todo, e o erro original continua acessível para inspeção.

Usar `%v` no lugar de `%w` gera a mesma mensagem, mas **quebra a cadeia**: o erro vira só texto, e `errors.Is` e `errors.As` deixam de encontrar o original. Use `%v` de propósito só quando não quiser expor o erro interno como parte da sua API.

## Erros sentinela e `errors.Is`

Um **erro sentinela** é um valor fixo, exportado, que representa uma condição específica. Por convenção, o nome começa com `Err`:

```go
var ErrNaoEncontrado = errors.New("não encontrado")
```

Para perguntar "esse erro é o `ErrNaoEncontrado`?", use `errors.Is`, que percorre toda a cadeia de wrapping:

```go
_, err := BuscarUsuario("99")
if errors.Is(err, ErrNaoEncontrado) {
	// responder 404, por exemplo
}
```

Comparar com `==` só funciona se o erro não tiver sido envolvido, por isso prefira sempre `errors.Is`. A biblioteca padrão tem vários sentinelas: `io.EOF`, `sql.ErrNoRows`, `context.DeadlineExceeded`.

## Erros estruturados e `errors.As`

Quando quem trata o erro precisa de **dados** (qual campo falhou, qual valor veio), crie um tipo de erro próprio:

```go
type ErrValidacao struct {
	Campo string
	Valor string
	Msg   string
	causa error
}

func (e ErrValidacao) Error() string {
	return fmt.Sprintf("%s: %s (%q)", e.Campo, e.Msg, e.Valor)
}

// Unwrap liga este erro à causa, como o %w faz.
func (e ErrValidacao) Unwrap() error {
	return e.causa
}
```

Para extrair esse tipo de dentro da cadeia, use `errors.As`, passando um ponteiro para uma variável do tipo procurado:

```go
var ev *ErrValidacao
if errors.As(err, &ev) {
	fmt.Println("campo inválido:", ev.Campo)
}
```

## Quando usar cada um

| Situação                                              | Ferramenta                        |
| ----------------------------------------------------- | --------------------------------- |
| Só importa a mensagem, ninguém vai inspecionar        | `errors.New` / `fmt.Errorf`       |
| Quem chama precisa reconhecer uma condição específica | sentinela + `errors.Is`           |
| Quem chama precisa de dados sobre a falha             | tipo de erro + `errors.As`        |
| Acrescentar contexto ao subir de camada               | `fmt.Errorf("contexto: %w", err)` |
| Juntar vários erros independentes (Go 1.20+)          | `errors.Join(err1, err2)`         |

## Trate o erro uma vez só

Ou você **trata** o erro (registra no log, responde ao usuário, tenta de novo), ou você o **devolve** com contexto. Fazer as duas coisas gera o mesmo erro logado várias vezes, uma por camada. O padrão comum é devolver até a borda do programa (o handler HTTP, o `main`) e tratar lá.

Nos exemplos do repositório, o `main` delega para uma função `run() error` e trata o erro num lugar só:

```go
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "erro: %v\n", err)
		os.Exit(1)
	}
}
```

## `panic` e `recover`

`panic` interrompe a execução normal e sobe pela pilha rodando os `defer`, até derrubar o programa. Ele é para **bugs e situações impossíveis**, como índice fora do slice ou estado interno corrompido, não para erros esperados como arquivo inexistente ou entrada inválida.

`recover`, chamado dentro de um `defer`, captura o panic e devolve o controle. Servidores HTTP usam isso para que um bug em uma requisição não derrube o processo inteiro. No seu código de aplicação, prefira devolver `error`.

## Exemplo no repositório

O exemplo [15-erros-idiomaticos](https://github.com/caramelotech/go-labs/tree/main/examples/15-erros-idiomaticos) separa cada ideia em um arquivo: o sentinela `ErrNaoEncontrado` com `errors.Is`, o erro estruturado `ErrValidacao` com `errors.As` (inclusive extraindo o `*strconv.NumError` da causa via `Unwrap`) e a comparação entre `%w` e `%v`. Os testes em `main_test.go` provam cada comportamento:

```bash
cd examples/15-erros-idiomaticos
go run .
go test -v
```

Pratique no exercício Divisão Segura, em [Exercícios](/labs/go/pratica/01-exercicios/).

## Referências

- [Error handling and Go](https://go.dev/blog/error-handling-and-go) - Andrew Gerrand, Go Blog, inglês
- [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors) - Damien Neil e Jonathan Amsterdam, Go Blog, inglês
- [Errors are values](https://go.dev/blog/errors-are-values) - Rob Pike, Go Blog, inglês
- [Go Code Review Comments - Error Strings](https://go.dev/wiki/CodeReviewComments#error-strings) - Go Wiki, inglês
- [Pacote errors](https://pkg.go.dev/errors) - documentação oficial, inglês
