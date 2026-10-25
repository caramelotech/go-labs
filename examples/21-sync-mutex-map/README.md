# 21 - Mutex, RWMutex e sync.Map

Como proteger estado compartilhado entre goroutines: exclusão mútua (`Mutex`), leitores paralelos (`RWMutex`) e o mapa concorrente especializado (`sync.Map`).

## Quando usar cada um

| Tipo | Use quando | Evite quando |
| ---- | ---------- | ------------ |
| `sync.Mutex` | Estado compartilhado genérico; leituras e escritas misturadas | Precisa só de comunicação (prefira channel) |
| `sync.RWMutex` | **Muitas** leituras e **poucas** escritas no mesmo dado | Escritas frequentes - o overhead pode perder para `Mutex` |
| `sync.Map` | (1) chave escrita **uma vez**, lida muitas; ou (2) goroutines com **chaves disjuntas** | Mapa “normal” com tipos fortes - use `map` + `Mutex`/`RWMutex` |

Regra prática: comece com `Mutex` (ou `map` + `Mutex`). Só troque se o perfil de acesso pedir.

## Arquivos

| Arquivo | O que tem |
| ------- | --------- |
| `caso_mutex.go` | `Contador` com `Mutex` |
| `caso_rwmutex.go` | `Catalogo` (`map`) com `RWMutex` |
| `caso_syncmap.go` | `sync.Map` write-once + chaves por worker |
| `main_test.go` | Contagem, catálogo e `LoadOrStore` |

## Como executar

```bash
go run .
go test -race ./...
```

## Conceitos

### Mutex - um de cada vez

```go
c.mu.Lock()
defer c.mu.Unlock()
c.n++
```

Simples e previsível. O `Mutex` **não pode ser copiado** - structs que o embutem usam pointer receiver.

### RWMutex - vários leitores OU um escritor

```go
c.mu.RLock()  // Get
c.mu.Lock()   // Set
```

Leitores não se bloqueiam entre si. Um `Lock` espera todos os `RUnlock` e bloqueia novos `RLock`.

### sync.Map - não é o `map` padrão com trava

API tipada como `any` (`Store`, `Load`, `LoadOrStore`, `Delete`, `Range`). Brilha nos dois casos da documentação do pacote `sync`; fora disso, `map` + mutex costuma ser mais claro e mais rápido.

```go
cache.Store("versão", "1.22")
v, ok := cache.Load("versão")
atual, loaded := cache.LoadOrStore("cfg", "prod") // não sobrescreve se já existe
```

## Experimente

1. Em `Contador.Inc`, comente o `Lock`/`Unlock` e rode `go test -race` - o detector acusa
2. Troque o `Catalogo` para `Mutex` puro - continua correto; `RWMutex` só ajuda se leitura dominar
3. Tente usar `sync.Map` como mapa genérico tipado - note os type assert em todo `Load`

## Relacionado

- Anterior na concorrência: [17 - Goroutines e WaitGroup](../17-goroutines-waitgroup/)
- [18 - Channels](../18-channels/) - comunicação em vez de memória compartilhada
- Docs: [sync](https://pkg.go.dev/sync)
