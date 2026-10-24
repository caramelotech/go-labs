# 21 - Ponteiros

Ponteiro guarda o **endereço** de uma variável. Serve para ler/alterar o original em vez de uma cópia.

## Ideia central

```go
x := 10
p := &x  // endereço de x
*p = 20  // altera x via ponteiro

dobraPonteiro(&n) // função recebe *int e muda n
```

## Como executar

```bash
go run .
```

Saída esperada (o endereço muda a cada execução):

```
x: 10
endereço (&x): 0x...
valor via ponteiro (*p): 10
depois de *p = 20, x: 20
após dobraCopia: 5
após dobraPonteiro: 10
new(int): 0
após *pi = 42: 42
```

## Conceitos

| Sintaxe | Significado |
| ------- | ----------- |
| `&x` | endereço de `x` (`*T`) |
| `*p` | valor apontado por `p` |
| `*int` | tipo “ponteiro para int” |
| `new(T)` | aloca `T` zerado e devolve `*T` |

Argumentos de função são **sempre** passados por valor (cópia). Para mutar o original, passe um ponteiro.

Em Go quase não se faz aritmética de ponteiros (diferente de C).

## Experimente

1. Imprima `p == &x` (deve ser `true`).
2. Passe uma struct por valor e por ponteiro para uma função que altera um campo.
3. Declare `var p *int` e veja o panic ao fazer `*p` (nil).

**[⬅️ Voltar aos Exemplos](../README.md)**
