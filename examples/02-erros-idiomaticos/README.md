# 02 - Erros Idiomáticos

Exemplo bem simples, separado em dois tipos de erro: **sentinela** e **estruturado**.

## Arquivos

| Arquivo | O que tem |
| ------- | --------- |
| `erro_sentinela.go` | `ErrNaoEncontrado` + ideia do `errors.Is` |
| `erro_estruturado.go` | `ErrValidacao` + ideia do `errors.As` |
| `usuario.go` | Busca que devolve sentinela |
| `idade.go` | Validação que devolve estruturado |
| `caso_sucesso.go` | Caminho sem erro |
| `caso_sentinela.go` | Casos de `Is` e wrapping |
| `caso_estruturado.go` | Caso de `As` |
| `main.go` | Só chama os casos |

## Como executar

```bash
go run .
go test -v
```

## Ideia central

```text
SENTINELA     → valor fixo (errors.New)
                inspecionar com errors.Is

ESTRUTURADO   → tipo com campos (struct)
                inspecionar com errors.As

%w            → envolve e mantém a cadeia
%v / %s       → só texto; perde Is/As
```

## Relacionado

- [Exercício 04 - Divisão Segura](../../notes/pratica/01-exercicios.md)
- Roadmap: seção Tratamento de Erros
