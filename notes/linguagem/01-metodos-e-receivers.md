# Métodos e Receivers

Um método é uma função ligada a um tipo. Em vez de uma classe com dados e comportamento juntos, Go separa as duas coisas: a struct guarda os dados e os métodos são declarados à parte, apontando para o tipo através de um **receiver**.

## Declarando um método

O receiver fica entre `func` e o nome do método:

```go
type Retangulo struct {
	Largura, Altura float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

r := Retangulo{Largura: 4, Altura: 3}
fmt.Println(r.Area()) // 12
```

Dentro do método, `r` é uma variável comum. Por convenção o nome do receiver é curto (uma ou duas letras, derivadas do tipo) e é o mesmo em todos os métodos do tipo. Nada de `this` ou `self`.

Métodos podem ser declarados em qualquer tipo definido no mesmo pacote, não só em structs:

```go
type Celsius float64

func (c Celsius) Fahrenheit() float64 {
	return float64(c)*9/5 + 32
}
```

## Value receiver ou pointer receiver

O receiver pode ser um valor `(r Retangulo)` ou um ponteiro `(r *Retangulo)`:

```go
// Value receiver: o método recebe uma CÓPIA
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

// Pointer receiver: o método recebe o ORIGINAL
func (r *Retangulo) Escalar(fator float64) {
	r.Largura *= fator
	r.Altura *= fator
}
```

A armadilha clássica é tentar alterar o estado num value receiver:

```go
func (c Carteira) Depositar(valor int) {
	c.saldo += valor // altera a cópia; o original não muda
}
```

O código compila, roda, e o saldo nunca sobe. Se o método precisa persistir uma mudança, o receiver é ponteiro.

### Chamada automática

Você não precisa escrever `(&r).Escalar(2)`. Se `r` é uma variável, Go pega o endereço sozinho quando o método pede ponteiro, e desreferencia sozinho quando o método pede valor. `r.Escalar(2)` funciona nos dois casos.

### Qual escolher

| Use value receiver `(t T)` quando                        | Use pointer receiver `(t *T)` quando                               |
| -------------------------------------------------------- | ------------------------------------------------------------------ |
| O método só lê o receiver                                | O método altera o receiver                                         |
| O tipo é pequeno (poucos campos)                         | O tipo contém `sync.Mutex` ou outro campo que não pode ser copiado |
| O tipo representa um valor imutável, como um `time.Time` | A struct é grande e copiá-la a cada chamada pesa                   |

Os exemplos deste lab seguem a regra de **preferir value receiver** e usar ponteiro só quando um dos motivos da direita aparece.

### Consistência dentro do tipo

Se **um** método do tipo precisa de ponteiro, use ponteiro em **todos** os métodos desse tipo, inclusive nos que só leem. Assim o conjunto de métodos não muda conforme a variável seja `T` ou `*T`, o que importa quando o tipo é usado através de uma [interface](/labs/go/linguagem/02-interfaces/). Misturar os dois no mesmo tipo é o que o [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments#receiver-type) pede para evitar.

No exemplo [14-receivers](https://github.com/caramelotech/go-labs/tree/main/examples/14-receivers), `Perfil` usa só value receiver e `Carteira` usa só pointer receiver, inclusive em `Saldo()`, que só lê.

## Method sets e interfaces

O **method set** de um tipo é o conjunto de métodos que ele tem para fins de interface:

- O tipo `T` tem só os métodos com receiver `T`
- O tipo `*T` tem os métodos com receiver `T` **e** os com receiver `*T`

Na prática: se um método que a interface exige usa pointer receiver, só um `*T` satisfaz a interface, não um `T`. Esse é um erro de compilação comum, com uma mensagem do tipo `Carteira does not implement Depositor (method Depositar has pointer receiver)`.

## `String()` e o `fmt.Stringer`

Um método `String() string` faz o tipo satisfazer a interface `fmt.Stringer`, e as funções de `fmt` passam a usá-lo para imprimir o valor:

```go
func (r Retangulo) String() string {
	return fmt.Sprintf("%.0fx%.0f", r.Largura, r.Altura)
}

fmt.Println(r) // 4x3
```

## Método é função com receiver

Um método pode ser referenciado como função comum, recebendo o receiver como primeiro argumento. É o que o exemplo 11 mostra com `Retangulo.Area(r)`:

```go
area := Retangulo.Area // tipo: func(Retangulo) float64
fmt.Println(area(r))
```

Também dá para "prender" o método a um valor específico, gerando uma função sem argumento extra: `f := r.Area` e depois `f()`.

## Composição com embedding

Go não tem herança. Para reaproveitar comportamento, uma struct pode **embutir** outro tipo, declarando um campo sem nome. Os campos e métodos do tipo embutido são **promovidos** e ficam acessíveis direto:

```go
type Animal struct {
	Nome string
}

func (a Animal) Apresentar() string {
	return "Eu sou " + a.Nome
}

type Cachorro struct {
	Animal // embutido
	Raca   string
}

c := Cachorro{Animal: Animal{Nome: "Caramelo"}, Raca: "SRD"}
fmt.Println(c.Apresentar()) // Eu sou Caramelo
fmt.Println(c.Nome)         // campo promovido
```

Isso é composição, não herança: `Cachorro` **tem um** `Animal`, e um `Cachorro` não pode ser passado onde se espera um `Animal`. O exemplo `statusWriter` em [HTTP com net/http](/labs/go/backend/01-http-com-net-http/) usa embedding para envolver um `http.ResponseWriter` e sobrescrever um único método.

## Exemplos no repositório

- [11-metodos](https://github.com/caramelotech/go-labs/tree/main/examples/11-metodos) - primeiro contato: `Area` e `String` com value receiver, `Escalar` com pointer receiver e `Retangulo.Area(r)` como função
- [14-receivers](https://github.com/caramelotech/go-labs/tree/main/examples/14-receivers) - o pacote `perfil` (só valor) e o pacote `carteira` (só ponteiro). Mude `Depositar` para `(c Carteira)` e veja o saldo parar de subir

## Referências

- [Go Code Review Comments - Receiver Type](https://go.dev/wiki/CodeReviewComments#receiver-type) - Go Wiki, inglês
- [Go FAQ - Should I define methods on values or pointers?](https://go.dev/doc/faq#methods_on_values_or_pointers) - Go Team, inglês
- [A Tour of Go - Methods](https://go.dev/tour/methods/1) - Go Team, inglês
- [Effective Go - Embedding](https://go.dev/doc/effective_go#embedding) - Go Team, inglês
