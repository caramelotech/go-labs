# 13 - Interface no consumidor

Como um pacote fala com outro sem se acoplar: a **interface fica no consumidor**, o provedor só exporta um tipo concreto, e o **`main` injeta** a dependência.

## Ideia central

```text
pedido/     → define Notificador e usa via New(n Notificador)
email/      → Cliente concreto com método Notificar (não importa pedido)
console/    → outra implementação, mesma assinatura
main.go     → cria o concreto e passa para pedido.New (DI)
```

Em Go a satisfação da interface é **implícita**: se o método tem o nome e a assinatura certos, o tipo serve. O provedor não declara `implements`.

## Arquivos

| Arquivo | Papel |
| ------- | ----- |
| `pedido/pedido.go` | Consumidor: interface + regra de negócio |
| `pedido/pedido_test.go` | Testa com fake - sem e-mail de verdade |
| `email/email.go` | Provedor concreto |
| `console/console.go` | Outro provedor (troca no `main`) |
| `main.go` | Composição / injeção de dependência |

## Como executar

```bash
go run .
go test ./...
```

Saída típica:

```
=== injeção com email.Cliente ===
[email noreply@caramelo.dev -> aluno@exemplo.com] pedido 42 confirmado

=== mesma regra, outro provedor (console.Saida) ===
[console -> aluno@exemplo.com] pedido 43 confirmado
```

## Por que a interface fica no consumidor

`pedido` precisa "avisar alguém". Ele define **só o que precisa**:

```go
type Notificador interface {
	Notificar(destino, mensagem string) error
}
```

`email` não precisa saber que existe `pedido`. Se amanhã o aviso for SMS, webhook ou console, `pedido` não muda - só o `main` troca o que injeta.

Evite o inverso (interface no provedor + pacote `notificador` genérico cedo demais): o consumidor acaba dependendo de um contrato que não é dele.

## Injeção no `main`

```go
mail := email.New("noreply@caramelo.dev")
svc, err := pedido.New(mail)
```

- `email.New` devolve `email.Cliente` (concreto, valor)
- `pedido.New` recebe `pedido.Notificador` (interface)
- o compilador aceita porque `Cliente` tem `Notificar` com value receiver

Dependência obrigatória valida em `New`, não no meio do `Confirmar`.
`Servico.Confirmar` e `Cliente.Notificar` usam `(s Servico)` / `(c Cliente)` - não mutam o struct. Ver [14 - Receivers](../14-receivers/).

## Teste sem o provedor real

No teste do consumidor, um fake local basta:

```go
type fakeNotificador struct{ /* ... */ }

func (f *fakeNotificador) Notificar(destino, mensagem string) error { /* ... */ }
```

`pedido` é testado sem importar `email` nem `console`.

## Fluxo de dependências

```text
main ──► email / console   (concreto)
main ──► pedido            (injeta o concreto)
pedido ─✘─► email          (não importa)
email  ─✘─► pedido         (não importa)
```

Setas de import só apontam para dentro a partir do `main` (e dos testes do consumidor). Sem ciclo.

## Experimente

1. Comente a injeção com `email` e rode só com `console.Saida`
2. Faça `Notificar` do e-mail retornar erro e veja o wrap em `Confirmar`
3. Em `pedido.New(nil)` - o `main`/`New` deve falhar cedo
4. Rode `go test ./pedido` - o fake cobre o caminho feliz e o erro

## Relacionado

- Anterior: [12 - Nomes de pacotes](../12-nomes-de-pacotes/)
- Próximo: [14 - Value vs pointer receiver](../14-receivers/)
- Depois: [15 - Erros idiomáticos](../15-erros-idiomaticos/)
