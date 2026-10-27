# Testes

Go traz testes na própria ferramenta: não há framework para instalar nem configuração para escrever. Um arquivo terminado em `_test.go`, uma função começando com `Test` e o comando `go test` bastam.

## O primeiro teste

Ao lado de `contador.go`, crie `contador_test.go`:

```go
package main

import "testing"

func TestContarPalavras(t *testing.T) {
	obtido := ContarPalavras("Go go")
	if obtido["go"] != 2 {
		t.Errorf("ContarPalavras(%q)[go] = %d; quer 2", "Go go", obtido["go"])
	}
}
```

- O arquivo termina em `_test.go` e só é compilado pelo `go test`
- A função começa com `Test`, seguida de uma letra maiúscula, e recebe `t *testing.T`
- Não há `assert`: a verificação é um `if` comum, e a mensagem de falha mostra o que veio e o que se esperava

## Reportando falhas

| Método          | Efeito                                                                |
| --------------- | --------------------------------------------------------------------- |
| `t.Errorf(...)` | marca o teste como falho e **continua** executando                    |
| `t.Fatalf(...)` | marca como falho e **para** o teste na hora                           |
| `t.Helper()`    | numa função auxiliar, faz a falha apontar para a linha de quem chamou |
| `t.Cleanup(f)`  | registra `f` para rodar no fim do teste (como um `defer`)             |
| `t.Skip(...)`   | pula o teste                                                          |

Use `Fatal` quando continuar não faz sentido (o valor é `nil` e a próxima linha vai dar panic) e `Error` para registrar várias falhas de uma vez.

## Testes em tabela

O padrão idiomático para testar vários casos é declarar uma **tabela** (um slice de structs anônimas) e percorrê-la com subtestes:

```go
func TestContarPalavras(t *testing.T) {
	testes := []struct {
		nome  string
		texto string
		quer  map[string]int
	}{
		{nome: "vazio", texto: "", quer: map[string]int{}},
		{nome: "ignora maiúsculas", texto: "Go go GO", quer: map[string]int{"go": 3}},
		{nome: "varias palavras", texto: "a b a", quer: map[string]int{"a": 2, "b": 1}},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			obtido := ContarPalavras(tt.texto)
			if !maps.Equal(obtido, tt.quer) {
				t.Fatalf("ContarPalavras(%q) = %v; quer %v", tt.texto, obtido, tt.quer)
			}
		})
	}
}
```

Os ganhos:

- Adicionar um caso novo, como uma regressão encontrada em produção, é só mais uma linha
- `t.Run` cria um subteste com nome, e um `Fatal` encerra só aquele caso, não a tabela toda
- O relatório mostra exatamente qual caso falhou: `--- FAIL: TestContarPalavras/vazio`

Para comparar maps e slices, use `maps.Equal` e `slices.Equal`. Para structs aninhadas e diferenças legíveis, muita gente usa o pacote [go-cmp](https://pkg.go.dev/github.com/google/go-cmp/cmp).

## Rodando os testes

```bash
go test ./...                    # todos os pacotes do módulo
go test -v                       # mostra cada teste, inclusive os que passaram
go test -run TestContar          # só os testes cujo nome casa com o padrão
go test -run 'TestContar/vazio'  # um subteste específico
go test -cover                   # porcentagem de cobertura
go test -race ./...              # ativa o detector de corrida
```

O `-race` é obrigatório em código concorrente. Ele detecta acessos simultâneos à mesma memória durante a execução, e o `Makefile` em `examples/` roda todos os exemplos com ele (`make test`).

## Testando pela API pública com fakes

Um arquivo de teste pode declarar `package pedido_test` em vez de `package pedido`. Aí ele só enxerga o que o pacote exporta, testando como um usuário de verdade usaria.

Quando o código depende de uma [interface](/labs/go/linguagem/02-interfaces/), o teste pode passar uma implementação falsa, declarada no próprio arquivo de teste:

```go
type fakeNotificador struct {
	destino  string
	mensagem string
	err      error
}

func (f *fakeNotificador) Notificar(destino, mensagem string) error {
	f.destino = destino
	f.mensagem = mensagem
	return f.err
}
```

O fake grava o que recebeu, para o teste conferir depois, e devolve o erro que o teste mandar, para simular falhas. Nenhum e-mail é enviado e nenhuma biblioteca de mocks é necessária.

## Testando HTTP com `httptest`

O pacote `net/http/httptest` sobe um servidor de teste numa porta local (`httptest.NewServer`) ou grava a resposta de um handler em memória (`httptest.NewRecorder`). O exemplo 23 usa `NewServer` para testar o middleware de log de ponta a ponta. Mais detalhes em [HTTP com net/http](/labs/go/backend/01-http-com-net-http/).

## Benchmarks e fuzzing

O `go test` também mede desempenho e gera entradas aleatórias:

```go
func BenchmarkContarPalavras(b *testing.B) {
	for b.Loop() { // Go 1.24+
		ContarPalavras("Go go GO labs")
	}
}

func FuzzContarPalavras(f *testing.F) {
	f.Add("Go go")
	f.Fuzz(func(t *testing.T, texto string) {
		ContarPalavras(texto) // não pode dar panic com nenhuma entrada
	})
}
```

Rode com `go test -bench .` e `go test -fuzz FuzzContarPalavras`. O fuzzing é bom para achar entradas que derrubam parsers e validadores.

## Exemplos no repositório

- [16-testes-em-tabela](https://github.com/caramelotech/go-labs/tree/main/examples/16-testes-em-tabela) - tabela de casos com `t.Run` e `maps.Equal` para `ContarPalavras`
- [13-interface-consumidor](https://github.com/caramelotech/go-labs/tree/main/examples/13-interface-consumidor) - `package pedido_test` com fake de `Notificador`, cobrindo sucesso, erro propagado e dependência `nil`
- [15-erros-idiomaticos](https://github.com/caramelotech/go-labs/tree/main/examples/15-erros-idiomaticos), [20-context](https://github.com/caramelotech/go-labs/tree/main/examples/20-context), [21-sync-mutex-map](https://github.com/caramelotech/go-labs/tree/main/examples/21-sync-mutex-map) e [23-slog-contexto](https://github.com/caramelotech/go-labs/tree/main/examples/23-slog-contexto) - testes de erros, cancelamento, concorrência com `-race` e HTTP com `httptest`

## Referências

- [Add a test - Tutorial oficial](https://go.dev/doc/tutorial/add-a-test) - Go Team, inglês
- [Go Wiki - TableDrivenTests](https://go.dev/wiki/TableDrivenTests) - Go Wiki, inglês
- [Using Subtests and Sub-benchmarks](https://go.dev/blog/subtests) - Marcel van Lohuizen, Go Blog, inglês
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) - Chris James, inglês
- [Pacote testing](https://pkg.go.dev/testing) - documentação oficial, inglês
