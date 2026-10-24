# 16 - Testes em Tabela

O padrão idiomático de testes em Go: uma fatia de casos e um loop com `t.Run`.

## Sobre

Em vez de copiar o mesmo `if` para cada entrada, você declara os casos em uma tabela e reutiliza a asserção. Isso:

- Deixa claro o que está sendo coberto
- Facilita adicionar regressões (só mais uma linha na tabela)
- Combina bem com subtestes (`t.Run`) no relatório do `go test`

## Como executar

```bash
go run .
go test -v
```

Saída do programa:

```
texto: "Go go GO labs"
contagem: map[go:3 labs:1]
```

## Conceitos

### Tabela + subteste

```go
testes := []struct {
	nome  string
	texto string
	quer  map[string]int
}{
	{nome: "vazio", texto: "", quer: map[string]int{}},
	// ...
}

for _, tt := range testes {
	t.Run(tt.nome, func(t *testing.T) {
		// asserção única
	})
}
```

### `maps.Equal` (Go 1.21+)

Compara dois maps sem escrever loop manual. Para slices, use `slices.Equal`.

## Experimente

1. Adicione um caso com pontuação colada (`"Go,"`) e decida o comportamento esperado
2. Rode `go test -cover` e veja a cobertura
3. Quebre de propósito a função e observe qual subteste falha

## Relacionado

- [Exercício 05 - Contador de Palavras](../../notes/pratica/01-exercicios.md)
- Roadmap: seção Testes e Qualidade
