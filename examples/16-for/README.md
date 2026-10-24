# 16 - for

Em Go existe **só** o `for`. Ele cobre loop clássico, “while” e iteração com `range`.

## Ideia central

```go
for i := 0; i < 3; i++ { ... }   // clássico
for n > 0 { ... }                 // só condição
for i, v := range slice { ... }   // índice e valor
for _, v := range slice { ... }   // só valor
```

## Como executar

```bash
go run .
```

Saída esperada:

```
--- clássico ---
i = 0
i = 1
i = 2
--- condição ---
n = 3
n = 2
n = 1
--- range slice ---
0: maçã
1: banana
2: uva
--- só valores ---
maçã
banana
uva
--- break / continue ---
número: 1
número: 3
```

## Conceitos

| Forma | Equivalente mental |
| ----- | ------------------ |
| `for init; cond; post` | `for` do C |
| `for cond` | `while` |
| `for {}` | loop infinito (até `break`) |
| `for range` | percorre slice, map, string, channel |

`continue` pula para a próxima iteração; `break` encerra o loop.

## Experimente

1. Some os números de 1 a 10 com um `for` clássico.
2. Use `range` numa `string` e imprima cada `rune`.
3. Faça um `for {}` com `break` quando um contador chegar a 3.

**[⬅️ Voltar aos Exemplos](../README.md)**
