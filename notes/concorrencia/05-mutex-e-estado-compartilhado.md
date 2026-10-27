# Mutex e Estado Compartilhado

Channels são a ferramenta preferida para passar dados entre goroutines, mas nem todo problema é um fluxo de dados. Um contador de acessos, um cache em memória ou um catálogo consultado por vários handlers HTTP são **estado compartilhado**: uma variável que várias goroutines leem e escrevem. Para esses casos, o pacote `sync` oferece travas.

## O problema: corrida de dados

```go
contador := 0
var wg sync.WaitGroup

for range 1000 {
	wg.Add(1)
	go func() {
		defer wg.Done()
		contador++ // ler, somar e gravar: três passos, não um
	}()
}
wg.Wait()
fmt.Println(contador) // quase nunca dá 1000
```

`contador++` parece uma operação só, mas são três: ler o valor, somar 1 e gravar. Duas goroutines podem ler o mesmo valor ao mesmo tempo e gravar o mesmo resultado, perdendo um incremento. O `go run -race .` aponta esse problema na hora.

## `sync.Mutex`

Um mutex (de _mutual exclusion_) garante que só uma goroutine por vez execute o trecho protegido:

```go
type Contador struct {
	mu sync.Mutex
	n  int
}

func (c *Contador) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *Contador) Valor() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}
```

Boas práticas:

- Coloque o mutex **junto dos dados que ele protege**, na mesma struct, logo acima dos campos
- Use `defer mu.Unlock()` logo depois do `Lock()`, para destravar mesmo se a função retornar cedo
- **Proteja também as leituras**. Ler sem trava enquanto outra goroutine escreve é corrida de dados
- Mantenha o trecho travado curto. Nada de chamada de rede ou de disco com o mutex travado
- O zero value do `sync.Mutex` já está pronto para uso, não precisa de construtor

### Mutex não pode ser copiado

Copiar uma struct que contém um mutex copia o estado da trava, e as duas cópias passam a proteger coisas diferentes. Por isso structs com `sync.Mutex` usam **pointer receiver em todos os métodos** (veja [Métodos e Receivers](/labs/go/linguagem/01-metodos-e-receivers/)) e são passadas por ponteiro. O `go vet` detecta cópias acidentais com a checagem `copylocks`.

## `sync.RWMutex`

Quando o dado é **muito lido e pouco escrito**, como um catálogo de configurações, o `RWMutex` permite vários leitores ao mesmo tempo:

```go
type Catalogo struct {
	mu    sync.RWMutex
	itens map[string]string
}

func (c *Catalogo) Get(chave string) (string, bool) {
	c.mu.RLock() // vários leitores juntos
	defer c.mu.RUnlock()
	v, ok := c.itens[chave]
	return v, ok
}

func (c *Catalogo) Set(chave, valor string) {
	c.mu.Lock() // escritor exclusivo
	defer c.mu.Unlock()
	c.itens[chave] = valor
}
```

Leitores não bloqueiam uns aos outros. Um `Lock` de escrita espera todos os leitores atuais saírem e impede novos leitores até terminar. O `RWMutex` tem um custo maior por operação, então com escritas frequentes ele pode ficar mais lento que um `Mutex` simples. Comece com `Mutex` e troque só se a leitura dominar.

## `sync.Map`

`sync.Map` é um map seguro para concorrência, mas **não** é um substituto geral para `map` com mutex. A documentação do pacote diz que ele é otimizado para dois casos:

1. Chaves escritas **uma vez** e lidas muitas vezes, como um cache que só cresce
2. Goroutines que trabalham com **conjuntos de chaves separados**

```go
var cache sync.Map

cache.Store("versão", "1.22")
v, ok := cache.Load("versão")
atual, carregou := cache.LoadOrStore("cfg", "prod") // não sobrescreve se já existe
cache.Range(func(chave, valor any) bool {
	fmt.Println(chave, valor)
	return true // continuar
})
```

A API usa `any`, então cada leitura exige type assertion e o compilador não protege os tipos. Fora dos dois casos acima, `map` + `Mutex` é mais claro e costuma ser mais rápido.

## Outras ferramentas do pacote `sync`

- **`sync.Once`**: executa uma inicialização uma única vez, mesmo com várias goroutines chamando. `sync.OnceValue` (Go 1.21+) faz o mesmo devolvendo um valor
- **`sync/atomic`**: tipos como `atomic.Int64` e `atomic.Bool` para contadores e flags simples, sem mutex
- **`sync.Pool`**: reaproveita objetos temporários para reduzir alocações em código de alto desempenho

## Channel ou mutex?

A [Go Wiki](https://go.dev/wiki/MutexOrChannel) dá uma regra prática: use o que for mais simples e expressivo para o problema.

| Situação                                       | Prefira |
| ---------------------------------------------- | ------- |
| Passar a posse de um dado para outra goroutine | channel |
| Distribuir trabalho entre workers              | channel |
| Comunicar resultados assíncronos               | channel |
| Cache, contador, estado interno de uma struct  | mutex   |
| Proteger um map acessado por vários handlers   | mutex   |

## Exemplo no repositório

O exemplo [21-sync-mutex-map](https://github.com/caramelotech/go-labs/tree/main/examples/21-sync-mutex-map) tem o `Contador` com `Mutex` incrementado por 8 goroutines, o `Catalogo` com `RWMutex` (um escritor e seis leitores) e o `sync.Map` nos dois cenários da documentação.

```bash
cd examples/21-sync-mutex-map
go run .
go test -race ./...
```

Comente o `Lock`/`Unlock` de `Contador.Inc` e rode `go test -race` para ver o detector acusar a corrida.

O Projeto 03, a API REST de livros em [Projetos](/labs/go/pratica/02-projetos/), usa exatamente esse padrão: um armazenamento em memória protegido por `sync.Mutex`.

## Referências

- [Pacote sync](https://pkg.go.dev/sync) - documentação oficial, inglês
- [Go Wiki - Use a sync.Mutex or a channel?](https://go.dev/wiki/MutexOrChannel) - Go Wiki, inglês
- [A Tour of Go - sync.Mutex](https://go.dev/tour/concurrency/9) - Go Team, inglês
- [The Go Memory Model](https://go.dev/ref/mem) - Go Team, inglês
