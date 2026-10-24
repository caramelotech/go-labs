# 10 - Visibilidade (exportado vs não exportado)

Em Go não existem as palavras `public` / `private`. A regra é só a **primeira letra**:

- **Maiúscula** → exportado (visível em outros pacotes)
- **minúscula** → não exportado (só no mesmo pacote)

## Ideia central

```go
type Conta struct {
	Titular string // ok fora do pacote
	saldo   int    // só dentro de package conta
}

func New(...) Conta      // exportada
func extrato(...) string // privada ao pacote
```

## Arquivos

| Arquivo | Papel |
| ------- | ----- |
| `conta/conta.go` | tipo e funções; mistura exportado / privado |
| `main.go` | pacote `main` consumindo só a API pública |

## Como executar

```bash
go run .
```

Saída esperada:

```
titular: Caramelo
saldo: 100
após depósito: 150
resumo: Caramelo: R$ 150
```

## Conceitos

- A regra vale para **tipos, campos, funções, métodos e constantes**.
- Pacotes diferentes = fronteira de visibilidade. Arquivos do **mesmo** pacote veem tudo (maiúsculo e minúsculo).
- Encapsular `saldo` evita que outro pacote coloque saldo negativo direto no campo.

## Experimente

1. Descomente `fmt.Println(c.saldo)` e leia o erro do compilador.
2. Renomeie `Resumo` para `resumo` e veja o `main` quebrar.
3. Crie `Sacar` exportado que só debita se houver saldo.

**[⬅️ Voltar aos Exemplos](../README.md)**
