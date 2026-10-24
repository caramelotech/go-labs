# 23 - Métodos

Método é uma função com **receiver**: o valor (ou ponteiro) à esquerda da chamada.

## Ideia central

```go
func (r Retangulo) Area() float64      // value - leitura / cópia
func (r *Retangulo) Escalar(fator float64) // pointer - mutação

r.Area()
r.Escalar(2)
```

Este exemplo é o **primeiro contato**. O aprofundamento (quando preferir cada um, consistência no tipo) está em [10-receivers](../10-receivers/).

## Como executar

```bash
go run .
```

Saída esperada:

```
retângulo: 4x3
área: 12
após escalar x2: 8x6
nova área: 48
Area como função: 48
```

## Conceitos

- Receiver **por valor** `(r Retangulo)`: cópia; bom para ler campos.
- Receiver **por ponteiro** `(r *Retangulo)`: original; necessário para mutar.
- Go permite `r.Escalar(2)` mesmo com `r` sendo valor - o compilador usa `&r` quando o método pede ponteiro (se `r` for addressable).
- `String()` no estilo de `fmt.Stringer` - `Println` pode usar isso em tipos que implementam a interface (aqui chamamos explicitamente).

## Experimente

1. Troque `Escalar` para value receiver e veja que o retângulo não muda.
2. Adicione `Perimetro() float64`.
3. Siga para [10-receivers](../10-receivers/) para a regra do lab com mais detalhes.

**[⬅️ Voltar aos Exemplos](../README.md)**
