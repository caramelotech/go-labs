# Structs

Go não tem classes. Para agrupar dados relacionados, você define uma **struct**: um tipo com campos nomeados. O comportamento entra depois, na forma de métodos, assunto de [Métodos e Receivers](/labs/go/linguagem/01-metodos-e-receivers/).

## Definindo uma struct

```go
type Endereco struct {
	Cidade string
	UF     string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco // struct dentro de struct
}
```

`type Nome struct { ... }` cria um tipo novo. A primeira letra decide a visibilidade: `Pessoa` e `Nome` são exportados, `pessoa` e `nome` ficariam visíveis só dentro do pacote.

## Criando valores

A forma recomendada é o **literal com nomes de campo**:

```go
p := Pessoa{
	Nome:  "Caramelo",
	Idade: 3,
	Endereco: Endereco{
		Cidade: "São Paulo",
		UF:     "SP",
	},
}
```

Campos omitidos recebem o zero value:

```go
outra := Pessoa{Nome: "Ana"} // Idade = 0, Endereco.Cidade = ""
var vazia Pessoa             // todos os campos no zero value
```

Existe também o literal posicional, `Pessoa{"Ana", 30, Endereco{}}`, mas ele quebra quando alguém adiciona ou reordena campos. O `go vet` avisa quando ele é usado com structs de outros pacotes.

## Acessando e alterando campos

```go
fmt.Println(p.Nome)
fmt.Println(p.Endereco.Cidade)
p.Idade++
```

Atribuir uma struct a outra variável **copia todos os campos**. Alterar a cópia não mexe no original, o que vale também para structs passadas como argumento. Para uma função alterar a struct de quem chamou, ela precisa receber um ponteiro, como mostra a nota de [ponteiros](/labs/go/fundamentos/09-ponteiros/).

## Funções construtoras

Go não tem construtor embutido. A convenção é uma função `New` ou `NewTipo` que valida a entrada e devolve o valor pronto:

```go
func NovaPessoa(nome string, idade int) (Pessoa, error) {
	if idade < 0 {
		return Pessoa{}, errors.New("idade não pode ser negativa")
	}
	return Pessoa{Nome: nome, Idade: idade}, nil
}
```

Se o zero value da struct já é utilizável, nem precisa de construtor.

## Structs anônimas

Para dados locais e pontuais, dá para declarar e usar uma struct sem dar nome a ela:

```go
ponto := struct {
	X, Y int
}{X: 2, Y: 5}
```

Esse formato é muito usado nas tabelas de casos de teste, como em [Testes](/labs/go/linguagem/04-testes/).

## Comparação

Duas structs podem ser comparadas com `==` quando todos os campos são comparáveis. A comparação é campo a campo:

```go
a := Endereco{Cidade: "Recife", UF: "PE"}
b := Endereco{Cidade: "Recife", UF: "PE"}
fmt.Println(a == b) // true
```

Se a struct tiver um slice, map ou função, `==` não compila.

## Tags de campo

Uma string depois do tipo do campo é uma **tag**: metadado que bibliotecas leem em tempo de execução. O caso mais comum é o `encoding/json`:

```go
type Livro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor,omitempty"`
}
```

Você vai usar tags assim no projeto da API de livros, em [Projetos](/labs/go/pratica/02-projetos/).

## Exemplo no repositório

O exemplo [08-structs](https://github.com/caramelotech/go-labs/tree/main/examples/08-structs) define `Pessoa` com um `Endereco` aninhado, mostra os campos omitidos no zero value, altera um campo e cria uma struct anônima. Como exercício, adicione um campo `Email`, crie um `[]Pessoa` e percorra com `range`.

## Referências

- [A Tour of Go - Structs](https://go.dev/tour/moretypes/2) - Go Team, inglês
- [Effective Go - Composite literals](https://go.dev/doc/effective_go#composite_literals) - Go Team, inglês
- [The Go Programming Language Specification - Struct types](https://go.dev/ref/spec#Struct_types) - Go Team, inglês
- [Go by Example - Structs](https://gobyexample.com/structs) - Mark McGranaghan e Eli Bendersky, inglês
