# 06 - Slices

Slice é a forma usual de trabalhar com sequências. Array existe, mas tem tamanho fixo no tipo.

## Ideia central

```go
nums := []int{10, 20, 30}
nums = append(nums, 40)
meio := nums[1:4] // do índice 1 até 3
```

`len` = quantos elementos; `cap` = quanto cabe sem realocar o array interno.

## Como executar

```bash
go run .
```

Saída esperada (a `cap` após `append` pode variar um pouco conforme a implementação, mas o restante é estável):

```
literal: [10 20 30] len: 3 cap: 3
make: [0 0 0]
após append: [10 20 30 40 50] len: 5 cap: 6
fatiamento [1:4]: [20 30 40]
array: [a b c]
slice do array: [a b c]
original: [10 20 30 40 50]
cópia alterada: [999 20 30 40 50]
```

## Conceitos

| Operação | Efeito |
| -------- | ------ |
| `make([]T, n)` | slice com `n` zero values |
| `append` | adiciona no fim (pode realocar) |
| `s[i:j]` | nova view; compartilha o array de fundo |
| `copy(dst, src)` | copia elementos (não compartilha) |

Fatias compartilham memória: alterar `meio[0]` pode alterar `nums`. Use `copy` quando quiser independência.

## Experimente

1. Imprima `len` e `cap` a cada `append` num loop.
2. Faça `nums[1:3]` e mude um elemento - veja o original.
3. Declare `var s []int` (nil) e faça `append(s, 1)`.

**[⬅️ Voltar aos Exemplos](../README.md)**
