# 07 - Maps

Map é o dicionário de Go: chave tipada → valor tipado.

## Ideia central

```go
idades := map[string]int{"Ana": 28}
idades["Diana"] = 41

v, ok := idades["Ana"] // ok == false se a chave não existe
delete(idades, "Ana")
```

## Como executar

```bash
go run .
```

A ordem do `range` **não é estável** - as linhas do range podem aparecer em qualquer ordem.

Saída (exemplo; o bloco `range` pode variar):

```
mapa: map[Ana:28 Bruno:34 Caramelo:3]
após insert/update: map[Ana:29 Bruno:34 Caramelo:3 Diana:41]
Ana: 29 existe? true
Zé: 0 existe? false
após delete Bruno: map[Ana:29 Caramelo:3 Diana:41]
--- range ---
...
estoque: map[café:10] len: 1
```

## Conceitos

- **Chave ausente:** leitura devolve o zero value do tipo do valor (`0` para `int`). Por isso use `v, ok := m[k]`.
- **`delete`:** seguro mesmo se a chave não existir.
- **`len(m)`:** quantidade de pares.
- Mapa **não é seguro** para escrita concorrente - isso fica para o exemplo `21-sync-mutex-map`.

## Experimente

1. Conte quantas vezes cada letra aparece numa palavra (`map[rune]int`).
2. Itere só as chaves: `for k := range m`.
3. Tente `var m map[string]int; m["x"] = 1` - entenda o panic e use `make`.

**[⬅️ Voltar aos Exemplos](../README.md)**
