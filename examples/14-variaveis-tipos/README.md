# 14 - Variáveis e tipos

Como declarar variáveis e constantes, e o que é o *zero value* de cada tipo.

## Ideia central

```go
var nome string = "Caramelo" // explícito
var ativo = true             // inferência
peso := 12.5                 // forma curta (só em funções)
const lab = "Go Labs"        // constante
```

Tipos básicos neste exemplo: `string`, `int`, `bool`, `float64`, `rune`.

## Como executar

```bash
go run .
```

Saída esperada:

```
nome: Caramelo
idade: 3
ativo: true
peso: 12.5
inicial: C código: 67
lab: Go Labs
--- zero values ---
int=0 string="" bool=false float64=0
```

## Conceitos

| Forma | Quando usar |
| ----- | ----------- |
| `var x T = v` | tipo explícito, útil em escopo de pacote |
| `var x = v` | deixa o compilador inferir |
| `x := v` | dentro de funções - a mais comum |
| `const` | valor fixo em tempo de compilação |

**Zero value:** variável declarada sem valor inicial recebe o “vazio” do tipo (`0`, `""`, `false`).

## Experimente

1. Troque `peso` para `float32` e veja o que muda na declaração.
2. Tente reatribuir `lab` - o compilador deve recusar.
3. Declare `var r rune = '🐶'` e imprima com `fmt.Printf("%c\n", r)`.

**[⬅️ Voltar aos Exemplos](../README.md)**
