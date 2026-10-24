# 05 - Funções

Funções com um retorno, múltiplos retornos, retornos nomeados e parâmetros variádicos.

## Ideia central

```go
func soma(a, b int) int
func dividir(a, b int) (int, int)     // múltiplos retornos
func coordenadas() (x, y int)         // retornos nomeados
func media(nums ...int) float64       // variádica
```

## Como executar

```bash
go run .
```

Saída esperada:

```
soma: 5
dividir 17/5 → quociente: 3 resto: 2
coordenadas: 10 20
média: 5
média vazia: 0
```

## Conceitos

- **Mesmo tipo nos parâmetros:** `a, b int` em vez de `a int, b int`.
- **Múltiplos retornos:** idiomático para “resultado + erro” (erro vem no próximo exemplo da trilha avançada).
- **Retornos nomeados:** úteis em funções curtas; em funções longas, prefira `return x, y` explícito.
- **`...T`:** a função recebe um slice; o caller passa `media(1, 2, 3)` ou `media(slice...)`.

## Experimente

1. Escreva `max(a, b int) int`.
2. Faça `minmax(nums ...int) (min, max int)`.
3. Ignore um retorno com `_`: `q, _ := dividir(10, 3)`.

**[⬅️ Voltar aos Exemplos](../README.md)**
