# Configurando o Ambiente

Antes de escrever qualquer código Go, você precisa de três coisas:

1. **Toolchain do Go** - o compilador, a biblioteca padrão e a ferramenta de linha de comando `go`
2. **Editor com suporte a Go** - para ter autocompletar, navegação e formatação automática
3. **Um módulo** - a forma como o Go organiza um projeto e suas dependências

A boa notícia: o Go é distribuído como um pacote único. Não existe uma etapa separada para instalar gerenciador de dependências, formatador ou executor de testes - tudo vem junto no comando `go`.

## Instalando o Go

Baixe o instalador oficial em [go.dev/dl](https://go.dev/dl/) e escolha a versão estável mais recente para o seu sistema.

### Windows

1. Baixe o arquivo `.msi` e execute o instalador
2. O instalador coloca o Go em `C:\Program Files\Go` e já adiciona `C:\Program Files\Go\bin` ao `PATH`
3. Feche e reabra o terminal para que o `PATH` atualizado seja carregado

Se preferir um gerenciador de pacotes:

```bash
winget install GoLang.Go
```

### macOS

Use o instalador `.pkg` do site oficial ou o Homebrew:

```bash
brew install go
```

### Linux

Remova qualquer instalação anterior em `/usr/local/go`, extraia o arquivo baixado e adicione o binário ao `PATH`:

```bash
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.X.Y.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
source ~/.profile
```

Troque `go1.X.Y.linux-amd64.tar.gz` pelo nome do arquivo que você baixou.

## Verificando a instalação

```bash
go version
```

A saída deve ser algo como:

```
go version go1.X.Y linux/amd64
```

Para ver as variáveis de ambiente que o Go está usando (onde ficam os módulos baixados, o cache de build, o sistema e a arquitetura de destino):

```bash
go env
```

Duas variáveis aparecem com frequência em tutoriais:

- **`GOROOT`** - onde o Go está instalado. Você raramente precisa mexer nela.
- **`GOPATH`** - pasta onde o Go guarda módulos baixados e binários instalados com `go install` (por padrão, `~/go`). Desde a chegada dos módulos, **você não precisa colocar seus projetos dentro do `GOPATH`** - eles podem morar em qualquer pasta.

Para que binários instalados com `go install` fiquem disponíveis no terminal, adicione `$(go env GOPATH)/bin` ao seu `PATH`.

## Escolhendo um editor

| Editor            | Como configurar                                                                                   |
| ----------------- | ------------------------------------------------------------------------------------------------- |
| **VS Code**       | Instale a extensão oficial [Go](https://marketplace.visualstudio.com/items?itemName=golang.go)     |
| **GoLand**        | IDE da JetBrains dedicada ao Go, já vem pronta para uso                                           |
| **Neovim / Vim**  | Configure o `gopls` pelo cliente LSP de sua preferência                                           |

Por baixo dos panos, quase todos os editores usam o **`gopls`**, o language server oficial do Go. No VS Code, a extensão oferece instalar o `gopls` e as demais ferramentas na primeira vez que você abre um arquivo `.go` - aceite.

Uma configuração que vale ativar: **formatar ao salvar**. Em Go não se discute estilo de código - todo mundo usa o `gofmt`, e o editor aplica a formatação automaticamente.

## Criando o primeiro módulo

Um **módulo** é um projeto Go: uma pasta com um arquivo `go.mod` que declara o nome do módulo e a versão do Go usada.

```bash
mkdir ola-go
cd ola-go
go mod init example.com/ola-go
```

O nome do módulo costuma ser o caminho do repositório (por exemplo, `github.com/seu-usuario/ola-go`). Para estudos locais, qualquer nome serve.

O comando gera o `go.mod`:

```
module example.com/ola-go

go 1.X
```

Crie um arquivo `main.go` na mesma pasta:

```go
package main

import "fmt"

func main() {
	fmt.Println("Olá, Caramelo Tech!")
}
```

E execute:

```bash
go run .
```

```
Olá, Caramelo Tech!
```

## Comandos essenciais

| Comando          | O que faz                                                         |
| ---------------- | ----------------------------------------------------------------- |
| `go run .`       | Compila e executa o pacote da pasta atual, sem gerar binário      |
| `go build`       | Compila e gera um executável na pasta atual                       |
| `go fmt ./...`   | Formata todos os arquivos do módulo                               |
| `go vet ./...`   | Procura erros comuns que compilam, mas provavelmente são bugs     |
| `go test ./...`  | Executa os testes do módulo                                       |
| `go mod tidy`    | Adiciona dependências que faltam e remove as que não são usadas   |
| `go get pacote`  | Adiciona ou atualiza uma dependência                              |
| `go doc fmt`     | Mostra a documentação de um pacote no terminal                    |

O `./...` significa "esta pasta e todas as subpastas".

## Sem instalar nada: Go Playground

Para testar um trecho rápido, use o [Go Playground](https://go.dev/play/) no navegador. Ele roda o código em um servidor e permite compartilhar o resultado por link - útil para tirar dúvidas em comunidades.

## Problemas comuns

### `go: command not found` ou `'go' não é reconhecido`

O diretório `bin` do Go não está no `PATH`. Reabra o terminal depois da instalação; se continuar, adicione manualmente (`/usr/local/go/bin` no Linux/macOS ou `C:\Program Files\Go\bin` no Windows).

### `go: go.mod file not found in current directory`

Você está rodando um comando de módulo fora de um módulo. Rode `go mod init <nome>` na pasta do projeto, ou entre na pasta que contém o `go.mod`.

## Referências

- [Download and install - documentação oficial do Go](https://go.dev/doc/install)
- [Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started)
- [Go Modules Reference](https://go.dev/ref/mod)
- [gopls - Go language server](https://go.dev/gopls/)
