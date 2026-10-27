# HTTP com net/http

A biblioteca padrão do Go traz um servidor HTTP pronto para produção no pacote `net/http`. Desde o Go 1.22, o roteador padrão entende método HTTP e parâmetros no caminho, então dá para construir uma API inteira sem framework. Frameworks como chi, Gin e Echo existem e são populares, mas todos se apoiam nas interfaces que esta nota apresenta.

## A interface `http.Handler`

Todo o servidor gira em torno de uma interface de um método:

```go
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
```

- `r *http.Request` traz tudo sobre a requisição: método, URL, cabeçalhos, corpo e o `r.Context()`
- `w http.ResponseWriter` é por onde a resposta sai: `w.Header()`, `w.WriteHeader(status)` e `w.Write(corpo)`

Na prática você raramente cria uma struct para isso. O tipo `http.HandlerFunc` transforma uma função comum num `Handler`:

```go
func handlePedido(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pedido ok")
}
```

## Roteamento com `ServeMux`

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /livros", listarLivros)
mux.HandleFunc("GET /livros/{id}", buscarLivro)
mux.HandleFunc("POST /livros", criarLivro)
mux.HandleFunc("DELETE /livros/{id}", removerLivro)

log.Fatal(http.ListenAndServe(":8080", mux))
```

Os padrões de rota (Go 1.22+) aceitam:

- **Método no início**: `GET /livros` só casa com GET (e HEAD). Um método não registrado recebe `405 Method Not Allowed` automaticamente
- **Curingas**: `{id}` casa com um segmento do caminho; `{caminho...}` casa com todo o resto
- **`{$}`**: casa só com o caminho exato terminado em barra, como `GET /{$}` para a raiz

O valor do curinga é lido com `r.PathValue`:

```go
func buscarLivro(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	// ...
}
```

Quando dois padrões casam com a mesma requisição, vence o **mais específico**: `GET /livros/novo` ganha de `GET /livros/{id}`.

## JSON na entrada e na saída

```go
type Livro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

func criarLivro(w http.ResponseWriter, r *http.Request) {
	var l Livro
	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	// ... salvar

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(l)
}
```

A ordem importa: cabeçalhos primeiro, depois `WriteHeader`, depois o corpo. Escrever no corpo antes de `WriteHeader` envia `200 OK` automaticamente, e mudanças posteriores no status são ignoradas. As tags `json:"..."` são explicadas em [Structs](/labs/go/fundamentos/08-structs/).

Cada requisição roda na sua própria goroutine. Se os handlers compartilham um armazenamento em memória, ele precisa de proteção, como mostra [Mutex e Estado Compartilhado](/labs/go/concorrencia/05-mutex-e-estado-compartilhado/).

## Middlewares

Um middleware é uma função que recebe um `Handler` e devolve outro, que faz algo antes e depois de chamar o original:

```go
func MiddlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("request", "path", r.URL.Path, "duracao", time.Since(inicio))
	})
}

handler := MiddlewareLog(mux)
http.ListenAndServe(":8080", handler)
```

Middlewares se encadeiam como camadas de cebola: `Recuperar(Autenticar(MiddlewareLog(mux)))`. Autenticação, log, CORS, recuperação de panic e limites de taxa costumam ser implementados assim.

### Capturando o status da resposta

O `ResponseWriter` não informa qual status foi enviado. Para logar o status, o middleware envolve o writer original com [embedding](/labs/go/linguagem/01-metodos-e-receivers/) e sobrescreve só o `WriteHeader`:

```go
type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}
```

Todos os outros métodos (`Header`, `Write`) são promovidos do writer embutido, sem nenhum código extra.

### Passando dados para o handler

O middleware passa dados adiante colocando-os no context da requisição e chamando `next.ServeHTTP(w, r.WithContext(ctx))`. É assim que o exemplo 23 entrega um logger já configurado para os handlers; veja [Logging com slog](/labs/go/backend/02-logging-com-slog/) e [Context](/labs/go/concorrencia/04-context/).

## Testando com `httptest`

O pacote `net/http/httptest` testa handlers sem abrir porta de verdade:

```go
func TestHandlePedido(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/pedido?id=7", nil)
	rec := httptest.NewRecorder()

	handlePedido(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
}
```

Para testar a pilha completa, com middlewares e roteamento, `httptest.NewServer(handler)` sobe um servidor local numa porta aleatória e devolve a URL em `srv.URL`.

## Antes de ir para produção

`http.ListenAndServe` é ótimo para estudar, mas não define nenhum timeout: um cliente lento consegue segurar uma conexão para sempre. Em produção, configure um `http.Server`:

```go
srv := &http.Server{
	Addr:              ":8080",
	Handler:           handler,
	ReadHeaderTimeout: 5 * time.Second,
	ReadTimeout:       10 * time.Second,
	WriteTimeout:      10 * time.Second,
	IdleTimeout:       60 * time.Second,
}
```

E implemente **graceful shutdown**: ao receber o sinal de parada (`SIGTERM`, `Ctrl+C`), chame `srv.Shutdown(ctx)`, que para de aceitar conexões novas e espera as requisições em andamento terminarem.

No lado do cliente, nunca use `http.Get` sem prazo para chamar outro serviço. Use um `http.Client` com `Timeout`, ou requisições com context, como em [Context](/labs/go/concorrencia/04-context/).

## Exemplos no repositório

- [23-slog-contexto](https://github.com/caramelotech/go-labs/tree/main/examples/23-slog-contexto) - `ServeMux` com `GET /pedido`, um middleware de log com `statusWriter`, handler e service em camadas, e testes de ponta a ponta com `httptest.NewServer`. Rode `go run . serve` para subir em `:8080`
- [20-context](https://github.com/caramelotech/go-labs/tree/main/examples/20-context) - `caso_http_cancel.go` mostra `r.Context()` sendo cancelado quando o cliente desiste da requisição

Aplique tudo no Projeto 03, a API REST de livros, em [Projetos](/labs/go/pratica/02-projetos/).

## Referências

- [Routing Enhancements for Go 1.22](https://go.dev/blog/routing-enhancements) - Jonathan Amsterdam, Go Blog, inglês
- [Pacote net/http](https://pkg.go.dev/net/http) - documentação oficial, inglês
- [Writing Web Applications](https://go.dev/doc/articles/wiki/) - Go Team, inglês
- [Pacote net/http/httptest](https://pkg.go.dev/net/http/httptest) - documentação oficial, inglês
