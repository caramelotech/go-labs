# Funções

Funções são declaradas com `func`, recebem parâmetros tipados e podem devolver mais de um valor. Elas também são valores comuns: dá para guardar uma função em variável, passar como argumento e devolver de outra função.

## Declaração

```go
func soma(a, b int) int {
	return a + b
}
```

- Os tipos vêm **depois** dos nomes: `a int`, não `int a`
- Parâmetros seguidos do mesmo tipo podem compartilhar a declaração: `a, b int`
- O tipo de retorno fica depois dos parênteses

Funções com nome começando em maiúscula (`Soma`) são exportadas para outros pacotes. Isso é detalhado em [Pacotes e Visibilidade](/labs/go/fundamentos/10-pacotes-e-visibilidade/).

## Múltiplos retornos

Uma função pode devolver vários valores, listados entre parênteses:

```go
func dividir(a, b int) (int, int) {
	return a / b, a % b
}

q, r := dividir(17, 5) // q = 3, r = 2
q, _ = dividir(10, 3)  // ignora o resto
```

O uso mais importante disso é o par **resultado + erro**, que substitui exceções em Go:

```go
func dividirSeguro(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}
	return a / b, nil
}
```

Por convenção, o `error` é sempre o **último** valor de retorno.

## Retornos nomeados

Os valores de retorno podem ter nome. Eles viram variáveis locais, começando no zero value, e um `return` sem argumentos (o _naked return_) devolve o valor atual delas:

```go
func coordenadas() (x, y int) {
	x = 10
	y = 20
	return
}
```

Nomes de retorno ajudam a documentar o que é cada valor, como em `func minmax(nums []int) (min, max int)`. O _naked return_, por outro lado, prejudica a leitura em funções longas, porque obriga quem lê a voltar até a assinatura. Use só em funções curtas.

## Funções variádicas

Um parâmetro `...T` aceita zero ou mais argumentos do tipo `T`. Dentro da função, ele é um slice `[]T`:

```go
func media(nums ...int) float64 {
	if len(nums) == 0 {
		return 0
	}
	total := 0
	for _, n := range nums {
		total += n
	}
	return float64(total) / float64(len(nums))
}

media(2, 4, 6, 8) // 5
media()           // 0

notas := []int{7, 8, 9}
media(notas...) // espalha o slice nos argumentos
```

Só o último parâmetro pode ser variádico. `fmt.Println` e `append` são exemplos conhecidos.

## Passagem por valor

Em Go, **todo argumento é copiado** ao entrar na função. Alterar um parâmetro `int` ou uma `struct` dentro da função não muda a variável de quem chamou. Para alterar o original, a função recebe um ponteiro, assunto da nota de [ponteiros](/labs/go/fundamentos/09-ponteiros/).

Slices e maps parecem fugir dessa regra porque a função consegue alterar os elementos, mas o que é copiado ali é um pequeno "cabeçalho" que aponta para os mesmos dados. As notas de [slices](/labs/go/fundamentos/06-arrays-e-slices/) e [maps](/labs/go/fundamentos/07-maps/) explicam o porquê.

## Funções como valores e closures

Funções têm tipo (por exemplo, `func(int) int`) e podem ser usadas como qualquer outro valor:

```go
func aplicar(nums []int, f func(int) int) []int {
	resultado := make([]int, 0, len(nums))
	for _, n := range nums {
		resultado = append(resultado, f(n))
	}
	return resultado
}

dobro := func(n int) int { return n * 2 }
aplicar([]int{1, 2, 3}, dobro) // [2 4 6]
```

Uma função anônima que usa variáveis de fora dela é uma **closure**: ela "captura" essas variáveis e continua enxergando-as depois que a função externa terminou.

```go
func contador() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

proximo := contador()
proximo() // 1
proximo() // 2
```

Closures aparecem muito em Go: em `sort.Slice`, em handlers HTTP e no lançamento de goroutines com `go func() { ... }()`.

## Exemplo no repositório

O exemplo [05-funcoes](https://github.com/caramelotech/go-labs/tree/main/examples/05-funcoes) tem `soma`, `dividir` com dois retornos, `coordenadas` com retornos nomeados e `media` variádica. Como exercício, escreva `minmax(nums ...int) (min, max int)` no mesmo arquivo.

O par resultado + erro é praticado no exercício Divisão Segura, em [Exercícios](/labs/go/pratica/01-exercicios/).

## Referências

- [A Tour of Go - Functions](https://go.dev/tour/basics/4) - Go Team, inglês
- [Effective Go - Functions](https://go.dev/doc/effective_go#functions) - Go Team, inglês
- [Go Code Review Comments - Named Result Parameters](https://go.dev/wiki/CodeReviewComments#named-result-parameters) - Go Wiki, inglês
- [Go by Example - Closures](https://gobyexample.com/closures) - Mark McGranaghan e Eli Bendersky, inglês
