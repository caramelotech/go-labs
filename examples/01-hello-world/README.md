# 01 - Hello World

Este é o programa mais básico em Go - o famoso "Olá, Mundo!"

## 📝 Sobre

Todo programador começa sua jornada com um programa Hello World. É uma tradição na computação e serve para:

- Verificar se o ambiente de desenvolvimento está configurado
- Entender a estrutura básica de um programa
- Fazer seu primeiro código funcionar!

## 📂 Arquivos

- `go.mod` - declara o módulo e a versão mínima do Go
- `main.go` - código fonte do programa

## 🚀 Como Executar

Dentro desta pasta:

### 1. Execute direto

```bash
go run .
```

### 2. Ou compile e execute o binário

```bash
go build
./01-hello-world        # no Windows: .\01-hello-world.exe
```

### 3. Saída Esperada

```
Olá, Mundo!
Bem-vindo ao Go!
Este é meu primeiro programa.
```

## 🔍 Entendendo o Código

```go
package main
```

- Todo arquivo Go pertence a um pacote
- O pacote `main` é o que gera um programa executável

```go
import "fmt"
```

- Importa o pacote `fmt` da biblioteca padrão, usado para formatar e imprimir texto

```go
func main() {
```

- `func` - palavra-chave para declarar uma função
- `main` - ponto de entrada do programa; a execução começa aqui
- A chave `{` precisa ficar na mesma linha da declaração

```go
fmt.Println("Olá, Mundo!")
```

- `fmt.Println` - imprime o texto e pula linha
- `Println` começa com letra maiúscula porque é **exportada** pelo pacote `fmt`

## 💡 Conceitos Importantes

### 1. Sem ponto e vírgula

O compilador insere os `;` automaticamente no fim das linhas.

### 2. Imports não usados são erro

Se você importar um pacote e não usar, o programa não compila. O mesmo vale para variáveis locais.

### 3. Formatação padrão

Rode `go fmt` e o código fica no formato que toda a comunidade usa - indentação com tabs incluída.

## 🎯 Experimente

1. **Mude a mensagem:**

   ```go
   fmt.Println("Seu nome aqui!")
   ```

2. **Use `Printf` com formatação:**

   ```go
   nome := "Caramelo"
   fmt.Printf("Olá, %s!\n", nome)
   ```

3. **Use `Print` em vez de `Println`:**

   ```go
   fmt.Print("Não pula linha ")
   fmt.Print("fica na mesma linha\n")
   ```

## 🔗 Conceitos Relacionados

- [📖 Configurando o Ambiente](../../notes/fundamentos/01-ambiente.md)
- [📖 Introdução ao Go](../../notes/fundamentos/02-introducao.md)

## ❓ Problemas Comuns

### Erro: "go: command not found" / "'go' não é reconhecido"

**Solução:** o Go não está instalado ou não está no `PATH`.
Veja: [Configurando o Ambiente](../../notes/fundamentos/01-ambiente.md)

### Erro: "imported and not used"

**Solução:** remova o import que não está sendo usado ou use o pacote no código.

### Erro: "syntax error: unexpected semicolon or newline before {"

**Solução:** a chave `{` deve ficar na mesma linha de `func main()`.

**[⬅️ Voltar aos Exemplos](../README.md)**
