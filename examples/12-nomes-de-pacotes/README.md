# 12 - Nomes de pacotes

Como nomear pacotes e o que exportar deles - o hábito que mais aparece no código Go do dia a dia.

Baseado em [Effective Go - Package names](https://go.dev/doc/effective_go#package-names).

## Sobre

Quando alguém importa um pacote, o **nome do pacote** vira o prefixo de tudo:

```go
import "bytes"

buf := bytes.Buffer{} // pacote + tipo
```

Por isso o nome precisa ser bom: curto, claro e fácil de digitar. Este exemplo mostra o jeito idiomático com um pacote `saudacao` e um `main` que o usa.

## Arquivos

| Arquivo | O que tem |
| ------- | --------- |
| `saudacao/saudacao.go` | Pacote de biblioteca com nomes sem "gaguejar" |
| `main.go` | Importa e usa `saudacao.Ola`, `saudacao.New`, `Texto` |
| `go.mod` | Módulo do exemplo |

## Como executar

```bash
go run .
```

Saída esperada:

```
Olá, Caramelo!
texto: bem-vindo ao lab
```

## Regras (Effective Go)

### 1. Nome curto, minúsculo, uma palavra

```text
bom:  saudacao, bufio, http, json
ruim: saudacao_utils, SaudacaoLib, packageSaudacao
```

Sem underscore e sem `mixedCaps` no **nome do pacote**. Prefira brevidade: quem importa digita esse nome o tempo todo.

### 2. Nome do pacote = nome da pasta

O diretório `saudacao/` declara `package saudacao`. O caminho de import pode ser longo (`.../saudacao`), mas o identificador local continua só `saudacao` - não `nomes_de_pacotes_saudacao`.

### 3. Não gagueje com o nome do pacote

Quem usa já escreve `saudacao.` na frente. Evite repetir:

| Evite | Prefira | Quem importa vê |
| ----- | ------- | --------------- |
| `SaudacaoMensagem` | `Mensagem` | `saudacao.Mensagem` |
| `NewMensagem` | `New` | `saudacao.New` |
| `SaudacaoOla` | `Ola` | `saudacao.Ola` |

O mesmo padrão da biblioteca padrão: `bufio.Reader` (não `BufReader`), `ring.New` (não `NewRing`).

### 4. Getter sem `Get`

Campo não exportado `texto` → método `Texto()`, não `GetTexto()`. Setter, se existir, fica `SetTexto`.

### 5. Colisão de nomes é rara

O nome do pacote não precisa ser único no mundo. O caminho do import decide *qual* pacote é. Se dois imports colidirem, o import ganha um alias local:

```go
import (
	crand "crypto/rand"
	"math/rand"
)
```

Evite `import .` - esconde de onde veio cada nome.

## Experimente

1. Renomeie o tipo para `SaudacaoMensagem` e use no `main` - note como fica redundante: `saudacao.SaudacaoMensagem`
2. Troque `Texto` por `GetTexto` - funciona, mas foge do idioma da linguagem
3. Mude a pasta para `saudacao_util` e o `package` para o mesmo nome - o compilador aceita, a comunidade não recomenda

## Relacionado

- Anterior na trilha de fundamentos: [01 - Hello World](../01-hello-world/)
- Próximo: [13 - Interface no consumidor](../13-interface-consumidor/)
- Documento oficial: [Effective Go - Package names](https://go.dev/doc/effective_go#package-names)
- Nota: [Introdução ao Go](../../notes/fundamentos/02-introducao.md)
