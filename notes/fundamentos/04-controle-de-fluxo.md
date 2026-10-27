# Controle de Fluxo

Go tem poucas estruturas de controle: `if`, `switch`, `for` e `defer`. Não existe `while`, `do-while` nem operador ternário (`cond ? a : b`). A ideia é ter um jeito só de fazer cada coisa, o que deixa código de pessoas diferentes parecido.

## `if`

O `if` não usa parênteses em volta da condição, e as chaves são obrigatórias, mesmo para uma linha só:

```go
if idade >= 18 {
	fmt.Println("maior de idade")
} else {
	fmt.Println("menor de idade")
}
```

### `if` com declaração curta

Antes da condição dá para declarar uma variável, separada por `;`. Ela só existe dentro do `if` e dos seus `else`:

```go
if n := len("caramelo"); n > 5 {
	fmt.Println("palavra longa, letras:", n)
} else {
	fmt.Println("palavra curta, letras:", n)
}
// n não existe aqui
```

Esse formato aparece o tempo todo no tratamento de erros:

```go
if err := salvar(dados); err != nil {
	return err
}
```

A condição precisa ser um `bool` de verdade. `if contador {` não compila quando `contador` é um `int`.

## `switch`

O `switch` de Go é mais flexível que o de C ou Java:

- Não precisa de `break`: só o `case` que casou é executado
- Um `case` aceita vários valores separados por vírgula
- Os valores não precisam ser constantes nem inteiros

```go
switch nota {
case 10:
	fmt.Println("excelente")
case 8, 9:
	fmt.Println("muito bom")
default:
	fmt.Println("precisa melhorar")
}
```

### `switch` sem expressão

Sem nada depois de `switch`, cada `case` é uma condição booleana. É a forma idiomática de substituir uma cadeia de `if/else if`:

```go
switch {
case idade < 13:
	fmt.Println("criança")
case idade < 18:
	fmt.Println("adolescente")
default:
	fmt.Println("adulto")
}
```

Se você realmente precisar continuar no próximo `case`, existe `fallthrough`, mas ele é raro em código Go. Há ainda o **type switch**, que decide pelo tipo de um valor; ele aparece na nota de [interfaces](/labs/go/linguagem/02-interfaces/).

## `for`: a única estrutura de repetição

O `for` cobre todos os casos de repetição:

```go
// 1. Clássico, estilo C
for i := 0; i < 3; i++ {
	fmt.Println(i)
}

// 2. Só condição: o "while" de Go
for n > 0 {
	n--
}

// 3. Infinito, até um break ou return
for {
	if terminou() {
		break
	}
}

// 4. Contando até N (Go 1.22+)
for i := range 3 {
	fmt.Println(i) // 0, 1, 2
}
```

`break` encerra o laço e `continue` pula para a próxima iteração.

### `range`

`range` percorre coleções e devolve dois valores por iteração. O significado deles muda conforme o que está sendo percorrido:

| Percorrendo    | Primeiro valor | Segundo valor     |
| -------------- | -------------- | ----------------- |
| slice ou array | índice         | cópia do elemento |
| string         | índice do byte | rune (caractere)  |
| map            | chave          | valor             |
| channel        | valor recebido | -                 |
| inteiro `n`    | 0 até n-1      | -                 |

Use `_` para ignorar o que não precisa:

```go
frutas := []string{"maçã", "banana", "uva"}

for i, fruta := range frutas {
	fmt.Printf("%d: %s\n", i, fruta)
}

for _, fruta := range frutas {
	fmt.Println(fruta)
}
```

Desde o Go 1.22, a variável do laço é **nova a cada iteração**. Antes disso, todas as iterações compartilhavam a mesma variável, o que causava bugs com goroutines e closures. A regra nova vale para módulos com `go 1.22` ou superior no `go.mod`.

## `defer`

`defer` agenda uma chamada de função para rodar quando a função atual terminar, seja por `return` normal, seja por panic. É a forma idiomática de garantir limpeza de recursos:

```go
func lerArquivo(caminho string) error {
	f, err := os.Open(caminho)
	if err != nil {
		return err
	}
	defer f.Close() // roda na saída, aconteça o que acontecer

	// ... usa f
	return nil
}
```

Três detalhes:

- Vários `defer` rodam em ordem inversa (o último agendado roda primeiro)
- Os argumentos são avaliados **na hora do `defer`**, não na hora da execução
- Coloque o `defer` logo depois de conferir o erro da abertura, para a limpeza ficar perto da aquisição

Você vai ver `defer` de novo com `wg.Done()`, `mu.Unlock()` e `cancel()` na seção de [concorrência](/labs/go/concorrencia/01-goroutines-e-waitgroup/).

## Exemplos no repositório

- [03-if-switch](https://github.com/caramelotech/go-labs/tree/main/examples/03-if-switch) - `if/else`, `if` com declaração curta, `switch` com e sem expressão
- [04-for](https://github.com/caramelotech/go-labs/tree/main/examples/04-for) - as formas do `for`, `range` em slice, `break` e `continue`

Para fixar, resolva o exercício FizzBuzz em [Exercícios](/labs/go/pratica/01-exercicios/).

## Referências

- [A Tour of Go - Flow control statements](https://go.dev/tour/flowcontrol/1) - Go Team, inglês
- [Effective Go - Control structures](https://go.dev/doc/effective_go#control-structures) - Go Team, inglês
- [Fixing For Loops in Go 1.22](https://go.dev/blog/loopvar-preview) - David Chase e Russ Cox, Go Blog, inglês
- [Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover) - Andrew Gerrand, Go Blog, inglês
