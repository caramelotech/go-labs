# Maps

Um map associa **chaves** a **valores**, como um dicionário. Os dois lados são tipados: `map[string]int` mapeia texto para inteiro e não aceita outra coisa.

## Criando um map

```go
idades := map[string]int{
	"Ana":      28,
	"Caramelo": 3,
}

estoque := make(map[string]int) // map vazio, pronto para receber dados
```

A chave pode ser de qualquer tipo **comparável** com `==`: números, strings, booleanos, ponteiros, arrays e structs formadas só por esses tipos. Slices, maps e funções não servem como chave.

## Operações básicas

```go
idades["Diana"] = 41 // insere
idades["Ana"] = 29   // atualiza
fmt.Println(idades["Ana"])
delete(idades, "Ana")    // remove (não faz nada se a chave não existir)
fmt.Println(len(idades)) // quantidade de pares
```

## "Comma ok": a chave existe?

Ler uma chave que não existe **não dá erro**: devolve o zero value do tipo do valor. Para `map[string]int`, isso é `0`, e aí fica impossível saber se a pessoa tem idade 0 ou se não está no map. A solução é a forma de dois valores:

```go
idade, ok := idades["Zé"]
if !ok {
	fmt.Println("Zé não está cadastrado")
}
```

O padrão aparece também como declaração curta no `if`:

```go
if idade, ok := idades["Ana"]; ok {
	fmt.Println("Ana tem", idade)
}
```

O zero value às vezes é exatamente o que você quer. Contar ocorrências fica simples porque a primeira leitura de uma chave nova devolve `0`:

```go
contagem := make(map[string]int)
for _, palavra := range strings.Fields(texto) {
	contagem[palavra]++
}
```

## Percorrendo com `range`

```go
for nome, idade := range idades {
	fmt.Printf("%s tem %d anos\n", nome, idade)
}
```

A ordem de iteração **não é garantida** e muda de uma execução para outra de propósito, para ninguém depender dela. Se precisar de ordem, ordene as chaves antes:

```go
for _, nome := range slices.Sorted(maps.Keys(idades)) {
	fmt.Println(nome, idades[nome])
}
```

`maps.Keys` (Go 1.23+) devolve um iterador com as chaves, e `slices.Sorted` coleta esse iterador em um slice ordenado.

## Armadilhas comuns

### Escrever em map `nil` causa panic

```go
var m map[string]int
fmt.Println(m["x"]) // ok: leitura devolve 0
m["x"] = 1          // panic: assignment to entry in nil map
```

Um map declarado com `var` sem valor é `nil`. Leitura funciona, escrita não. Inicialize com `make` ou com um literal.

### Valores struct não são endereçáveis

```go
type Ponto struct{ X, Y int }

pontos := map[string]Ponto{"a": {1, 2}}
pontos["a"].X = 10 // não compila
```

O valor guardado no map não tem endereço fixo, porque o map pode reorganizar a memória ao crescer. Leia, altere a cópia e grave de volta, ou use `map[string]*Ponto`.

### Maps não são seguros para concorrência

Duas goroutines escrevendo no mesmo map ao mesmo tempo derrubam o programa com `fatal error: concurrent map writes`. A proteção com `sync.Mutex` ou `sync.Map` é assunto de [Mutex e Estado Compartilhado](/labs/go/concorrencia/05-mutex-e-estado-compartilhado/).

## Map como conjunto

Go não tem um tipo "set". O idioma é usar um map com valor `struct{}`, que não ocupa memória:

```go
vistos := make(map[string]struct{})
vistos["go"] = struct{}{}

if _, ok := vistos["go"]; ok {
	fmt.Println("já vi")
}
```

`map[string]bool` também funciona e é mais fácil de ler; escolha um e mantenha no projeto.

## O pacote `maps`

O pacote `maps` (Go 1.21+) tem utilitários como `maps.Equal` (compara dois maps), `maps.Clone` (cópia rasa) e os iteradores `maps.Keys` e `maps.Values`. Para esvaziar um map existente, use a função embutida `clear(m)`.

## Exemplo no repositório

O exemplo [07-maps](https://github.com/caramelotech/go-labs/tree/main/examples/07-maps) cria um map com literal, insere, atualiza, usa o "comma ok", remove com `delete`, percorre com `range` e cria outro com `make`. Rode algumas vezes e observe a ordem do `range` mudar.

O exemplo [16-testes-em-tabela](https://github.com/caramelotech/go-labs/tree/main/examples/16-testes-em-tabela) usa um map para contar palavras e `maps.Equal` nos testes. Ele é o ponto de partida do exercício Contador de Palavras, em [Exercícios](/labs/go/pratica/01-exercicios/).

## Referências

- [Go maps in action](https://go.dev/blog/maps) - Andrew Gerrand, Go Blog, inglês
- [A Tour of Go - Maps](https://go.dev/tour/moretypes/19) - Go Team, inglês
- [Effective Go - Maps](https://go.dev/doc/effective_go#maps) - Go Team, inglês
- [Pacote maps](https://pkg.go.dev/maps) - documentação oficial, inglês
