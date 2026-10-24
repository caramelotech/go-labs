# 03 - if e switch

Condicionais em Go: `if` (com declaração curta) e `switch` (com e sem expressão).

## Ideia central

```go
if n := len(s); n > 5 { ... } // n só existe dentro do if/else

switch nota {
case 8, 9:
	// vários valores no mesmo case
default:
}
```

Não existe `while` nem ternário (`? :`). Para condições encadeadas, use `switch` sem expressão.

## Como executar

```bash
go run .
```

Saída esperada:

```
menor de idade
palavra longa, letras: 8
nota: muito bom
faixa: adolescente
```

## Conceitos

- **Declaração curta no `if`:** inicializa e testa na mesma linha; o escopo fica limitado ao bloco.
- **`switch` com expressão:** compara o valor com cada `case` (sem `break` obrigatório - Go não “cai” no próximo case).
- **`switch` sem expressão:** cada `case` é uma condição booleana.

## Experimente

1. Mude `idade` para `20` e observe as duas saídas relacionadas.
2. Adicione `case 0:` no switch de nota.
3. Use `fallthrough` num case (cuidado: quase nunca é o que você quer).

**[⬅️ Voltar aos Exemplos](../README.md)**
