# 14 - Value vs pointer receiver

Quando o método usa `(s Servico)` e quando usa `(s *Servico)`. Regra deste lab: **prefira value receiver**; só use ponteiro quando precisar.

## Ideia central

```text
(s Servico)   → método trabalha numa CÓPIA  (leitura, tipo pequeno)
(s *Servico)  → método trabalha no ORIGINAL (mutação, sync.Mutex, tipo grande)
```

```go
func (p Perfil) Nome() string      // value - não muda o struct
func (c *Carteira) Depositar(int)  // pointer - altera c.saldo
```

## Arquivos

| Arquivo | Papel |
| ------- | ----- |
| `perfil/perfil.go` | Só value receiver - lê campos |
| `carteira/carteira.go` | Pointer - `Depositar` / `Sacar` mudam saldo |
| `main.go` | Mostra os dois lados |

## Como executar

```bash
go run .
```

Saída esperada:

```
=== value receiver (perfil) ===
Caramelo (3 anos)
nome: Caramelo idade: 3

=== pointer receiver (carteira) ===
saldo inicial: 100
após depósito: 150
após saque: 120
```

## Quando usar value `(s Servico)`

Use por padrão quando:

1. O método **não altera** o receiver
2. O struct é **pequeno** (poucos campos, sem mutex)
3. Você quer deixar claro que a operação é só leitura

`perfil` inteiro segue isso: `New` devolve valor, métodos com `(p Perfil)`.

## Quando usar pointer `(s *Servico)`

Use quando:

1. O método **precisa mutar** o receiver (`Depositar`, `Sacar`)
2. O tipo contém algo que **não pode ser copiado** (`sync.Mutex`, etc.)
3. O struct é **grande** e você quer evitar copiar sempre (caso menos comum no início)

## Consistência no mesmo tipo

Se **um** método do tipo precisa de ponteiro, os outros do mesmo tipo também usam ponteiro. Assim o method set não muda conforme a variável seja `T` ou `*T`.

Por isso `Carteira.Saldo` também é `(c *Carteira)`, mesmo só lendo.

Tipos diferentes podem (e devem) escolher diferente: `Perfil` 100% valor, `Carteira` 100% ponteiro.

## Armadilha clássica

```go
func (c Carteira) Depositar(valor int) {
	c.saldo += valor // altera a CÓPIA - o original não muda
}
```

Com value receiver, a mutação some ao sair do método. Se precisa persistir a mudança → ponteiro.

## Experimente

1. Em `carteira`, mude `Depositar` para `(c Carteira)` e rode - o saldo "não sobe"
2. Em `perfil`, mude `Nome` para `(p *Perfil)` - funciona, mas o ponteiro é desnecessário
3. Compare com [13 - Interface no consumidor](../13-interface-consumidor/): `Servico` e `email.Cliente` usam value receiver

## Relacionado

- Anterior: [13 - Interface no consumidor](../13-interface-consumidor/)
- [Code Review Comments - Receiver Type](https://go.dev/wiki/CodeReviewComments#receiver-type)
