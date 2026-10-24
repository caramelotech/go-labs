# 20 - Structs

Struct agrupa campos. Em Go não há classes: dados ficam em structs; comportamento em funções/métodos.

## Ideia central

```go
type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

p := Pessoa{Nome: "Caramelo", Idade: 3}
```

## Como executar

```bash
go run .
```

Saída esperada:

```
Caramelo tem 3 anos
mora em São Paulo/SP
outra: nome=Ana idade=0 cidade=""
idade depois do aniversário: 4
ponto: {2 5}
```

## Conceitos

- Prefira **literais com nome de campo** (`Nome: "..."`) - a ordem dos campos no tipo pode mudar.
- Campos omitidos recebem **zero value**.
- Structs aninhadas acessam com `p.Endereco.Cidade`.
- Maiúscula no nome do campo (`Nome`) = exportado entre pacotes (veja o exemplo 22).

## Experimente

1. Adicione o campo `Email string` em `Pessoa`.
2. Crie um slice `[]Pessoa` e percorra com `range`.
3. Compare duas structs com `==` (funciona se todos os campos forem comparáveis).

**[⬅️ Voltar aos Exemplos](../README.md)**
