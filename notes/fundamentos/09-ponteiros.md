# Ponteiros

Um ponteiro guarda o **endereço de memória** de uma variável. Com ele, uma função consegue ler e alterar o valor original em vez de trabalhar numa cópia. Ponteiros em Go são bem mais simples que em C: não há aritmética de ponteiros e a memória é liberada pelo garbage collector.

## `&` e `*`

| Sintaxe | Significado                                      |
| ------- | ------------------------------------------------ |
| `&x`    | endereço de `x` (o resultado tem tipo `*T`)      |
| `*T`    | tipo "ponteiro para T"                           |
| `*p`    | o valor guardado no endereço `p` (desreferência) |

```go
x := 10
p := &x // p é *int e aponta para x

fmt.Println(*p) // 10
*p = 20         // altera x através do ponteiro
fmt.Println(x)  // 20
```

## Por que ponteiros existem: passagem por valor

Em Go, argumentos são sempre copiados. Uma função que recebe `int` altera só a cópia dela:

```go
func dobraCopia(n int) {
	n = n * 2 // some quando a função termina
}

func dobraPonteiro(n *int) {
	*n = *n * 2 // altera o valor no endereço recebido
}

n := 5
dobraCopia(n)     // n continua 5
dobraPonteiro(&n) // n vira 10
```

## Ponteiros para structs

Com structs, Go simplifica a sintaxe: `p.Nome` funciona mesmo quando `p` é um ponteiro, sem precisar escrever `(*p).Nome`.

```go
func aniversario(p *Pessoa) {
	p.Idade++
}

ana := Pessoa{Nome: "Ana", Idade: 30}
aniversario(&ana) // ana.Idade = 31
```

É comum criar a struct já como ponteiro com `&Pessoa{...}`, principalmente em construtores.

## `nil` e `new`

O zero value de um ponteiro é `nil`, que não aponta para nada. Desreferenciar `nil` derruba o programa:

```go
var p *int
fmt.Println(p == nil) // true
fmt.Println(*p)       // panic: invalid memory address or nil pointer dereference
```

Antes de usar um ponteiro que pode ser `nil`, confira. Melhor ainda é evitar que ele seja `nil`, validando na construção, como faz o `pedido.New` do exemplo [13-interface-consumidor](https://github.com/caramelotech/go-labs/tree/main/examples/13-interface-consumidor).

`new(T)` aloca um `T` com zero value e devolve o `*T`. Na prática, `&T{}` é mais usado para structs.

```go
pi := new(int) // *int apontando para 0
*pi = 42
```

## Devolver ponteiro para variável local é seguro

Em C, devolver o endereço de uma variável local é um bug clássico. Em Go não:

```go
func NovoContador() *Contador {
	c := Contador{}
	return &c // seguro
}
```

O compilador faz _escape analysis_: percebe que `c` precisa viver depois da função e a aloca no heap. O garbage collector libera quando ninguém mais apontar para ela.

## Quando usar ponteiro

Use ponteiro quando:

- A função ou método **precisa alterar** o valor original
- O tipo **não pode ser copiado**, como structs que contêm `sync.Mutex`
- A struct é grande o bastante para a cópia pesar (caso menos comum do que parece)
- Você precisa representar "ausência de valor" com `nil`

Para tipos pequenos que só são lidos, prefira valores. Slices, maps, channels e funções já carregam uma referência interna, então raramente faz sentido usar `*[]T` ou `*map[K]V`.

A escolha entre valor e ponteiro em métodos tem regras próprias, explicadas em [Métodos e Receivers](/labs/go/linguagem/01-metodos-e-receivers/).

## Exemplo no repositório

O exemplo [09-ponteiros](https://github.com/caramelotech/go-labs/tree/main/examples/09-ponteiros) mostra `&` e `*`, compara `dobraCopia` com `dobraPonteiro` e usa `new(int)`. Como exercício, declare `var p *int` e tente `*p` para ver o panic de perto.

## Referências

- [A Tour of Go - Pointers](https://go.dev/tour/moretypes/1) - Go Team, inglês
- [Go FAQ - When are function parameters passed by value?](https://go.dev/doc/faq#pass_by_value) - Go Team, inglês
- [Go FAQ - How do I know whether a variable is allocated on the heap or the stack?](https://go.dev/doc/faq#stack_or_heap) - Go Team, inglês
- [Effective Go - Allocation with new](https://go.dev/doc/effective_go#allocation_new) - Go Team, inglês
