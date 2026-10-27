# Variáveis, Constantes e Tipos

Go tem tipagem estática: toda variável tem um tipo definido em tempo de compilação, e esse tipo não muda. Ao mesmo tempo, o compilador infere o tipo na maioria das declarações, então o código fica curto sem perder a segurança.

## Declarando variáveis

Existem três formas de declarar uma variável:

```go
var nome string = "Caramelo" // tipo explícito
var ativo = true             // tipo inferido a partir do valor
peso := 12.5                 // declaração curta: só dentro de funções
```

| Forma         | Quando usar                                                               |
| ------------- | ------------------------------------------------------------------------- |
| `var x T = v` | Quando o tipo precisa ficar explícito (ex: `float32` em vez de `float64`) |
| `var x = v`   | Em escopo de pacote, onde `:=` não é permitido                            |
| `var x T`     | Quando a variável começa vazia e recebe valor depois                      |
| `x := v`      | Dentro de funções - é a forma mais usada no dia a dia                     |

Dá para declarar várias de uma vez, o que é comum no topo de arquivos:

```go
var (
	host  = "localhost"
	porta = 8080
)
```

Duas regras que pegam quem vem de outras linguagens:

- **Variável local não usada é erro de compilação**, não aviso. Se precisar ignorar um valor, use o identificador em branco `_`.
- `:=` declara pelo menos uma variável nova. `x := 1` seguido de `x := 2` no mesmo escopo não compila; para reatribuir, use `x = 2`.

## Zero values

Em Go não existe variável "não inicializada". Toda variável declarada sem valor recebe o **zero value** do seu tipo:

| Tipo                                                   | Zero value                      |
| ------------------------------------------------------ | ------------------------------- |
| `int`, `float64` e demais números                      | `0`                             |
| `string`                                               | `""`                            |
| `bool`                                                 | `false`                         |
| ponteiros, slices, maps, channels, funções, interfaces | `nil`                           |
| structs                                                | cada campo com o seu zero value |

Isso tem uma consequência de design: na biblioteca padrão, muitos tipos funcionam sem construtor, direto do zero value. Um `var wg sync.WaitGroup` ou um `var buf bytes.Buffer` já estão prontos para uso. O [Effective Go](https://go.dev/doc/effective_go#allocation_new) recomenda seguir a mesma ideia nos seus tipos sempre que possível.

## Tipos básicos

| Categoria       | Tipos                                                          | Observação                                                 |
| --------------- | -------------------------------------------------------------- | ---------------------------------------------------------- |
| Inteiros        | `int`, `int8`, `int16`, `int32`, `int64`                       | `int` tem 32 ou 64 bits conforme a plataforma              |
| Sem sinal       | `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`       | Use quando o valor nunca for negativo e o tamanho importar |
| Ponto flutuante | `float32`, `float64`                                           | `float64` é o padrão da inferência                         |
| Complexos       | `complex64`, `complex128`                                      | Raros fora de cálculo científico                           |
| Texto           | `string`, `byte` (alias de `uint8`), `rune` (alias de `int32`) | Veja a seção sobre strings abaixo                          |
| Lógico          | `bool`                                                         | Só aceita `true` e `false`, sem "truthy"                   |

Na dúvida, use `int` para inteiros e `float64` para números com casas decimais.

## Conversão de tipos

Go **nunca** converte tipos automaticamente, nem entre `int` e `int64`. A conversão é sempre explícita, com a sintaxe `T(valor)`:

```go
total := 10
quantidade := 4
media := float64(total) / float64(quantidade) // 2.5
```

Para converter entre texto e número, use o pacote `strconv` - `string(65)` não gera `"65"`, gera `"A"` (o caractere de código 65):

```go
n, err := strconv.Atoi("42") // string -> int, pode falhar
s := strconv.Itoa(42)        // int -> string
f, err := strconv.ParseFloat("3.14", 64)
```

Repare que `Atoi` devolve um erro junto com o valor. Esse padrão aparece em toda a linguagem e é assunto da nota de [tratamento de erros](/labs/go/linguagem/03-tratamento-de-erros/).

## Constantes e `iota`

Constantes são valores fixos, calculados em tempo de compilação. Só podem ser números, strings, runes ou booleanos:

```go
const lab = "Go Labs"
const pi = 3.14159
```

Constantes numéricas sem tipo declarado são **untyped**: elas se adaptam ao contexto em que são usadas, por isso `const limite = 10` funciona tanto com `int` quanto com `float64`.

Para enumerações, Go usa `iota`, um contador que começa em 0 e incrementa a cada linha do bloco `const`:

```go
type Status int

const (
	Pendente  Status = iota // 0
	Aprovado                // 1
	Cancelado               // 2
)
```

## Strings, bytes e runes

Uma `string` em Go é uma sequência **imutável de bytes**, normalmente em UTF-8. Isso explica um comportamento que surpreende no começo:

```go
s := "café"
fmt.Println(len(s)) // 5, não 4: o "é" ocupa 2 bytes em UTF-8
```

- `len(s)` conta **bytes**
- `s[i]` devolve um **byte**
- `for _, r := range s` percorre **runes** (caracteres Unicode), decodificando o UTF-8 para você

```go
for i, r := range "café" {
	fmt.Printf("%d: %c\n", i, r) // 0: c, 1: a, 2: f, 3: é
}
```

Uma `rune` é escrita entre aspas simples (`'C'`) e guarda o código Unicode do caractere. Para contar caracteres, use `utf8.RuneCountInString(s)`. Para montar strings em laços, prefira `strings.Builder` a concatenar com `+`, que cria uma string nova a cada operação.

## Exemplo no repositório

O exemplo [02-variaveis-tipos](https://github.com/caramelotech/go-labs/tree/main/examples/02-variaveis-tipos) declara variáveis nas três formas, usa uma `rune` e imprime os zero values de `int`, `string`, `bool` e `float64`:

```bash
cd examples/02-variaveis-tipos
go run .
```

Depois de rodar, tente reatribuir a constante `lab` e leia a mensagem do compilador.

## Referências

- [A Tour of Go - Basics](https://go.dev/tour/basics/1) - Go Team, inglês
- [Effective Go - Constants](https://go.dev/doc/effective_go#constants) - Go Team, inglês
- [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings) - Rob Pike, Go Blog, inglês
- [The Go Programming Language Specification - Constants](https://go.dev/ref/spec#Constants) - Go Team, inglês
