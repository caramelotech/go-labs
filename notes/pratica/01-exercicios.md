# Exercícios

Pratique o que aprendeu nos conteúdos e exemplos de código.

## Como usar

1. Leia o enunciado do exercício
2. Tente resolver por conta própria antes de buscar ajuda
3. Revise as notas de [Fundamentos](/labs/go/fundamentos/01-ambiente/) se precisar relembrar algum conceito
4. Rode `go fmt` e `go vet` na sua solução antes de considerar o exercício concluído
5. Compare sua solução com a de outras pessoas da comunidade

## Exercício 01 - Olá, Go

**Nível:** iniciante

**Objetivo:** Criar um módulo e executar o primeiro programa Go.

**Instrução:**

1. Crie uma pasta `ola-go` e inicialize um módulo com `go mod init`
2. Escreva um `main.go` que imprima `Olá, Caramelo Tech!`
3. Execute com `go run .` e depois gere o binário com `go build`

**Critérios de sucesso:**

- [ ] A pasta contém um `go.mod` válido
- [ ] `go run .` imprime a mensagem corretamente
- [ ] `go build` gera um executável que imprime a mesma mensagem

## Exercício 02 - Variáveis e Tipos

**Nível:** iniciante

**Objetivo:** Declarar variáveis de diferentes tipos e formatar a saída.

**Instrução:**

1. Declare uma variável `string` com seu nome usando `:=`
2. Declare uma variável `int` com sua idade usando `var`
3. Declare uma constante `float64` com sua altura em metros
4. Imprima uma frase combinando as três com `fmt.Printf`

Saída esperada:

```
Meu nome é Ana, tenho 25 anos e meço 1.68m.
```

**Critérios de sucesso:**

- [ ] Usa `:=`, `var` e `const` pelo menos uma vez cada
- [ ] Usa `fmt.Printf` com os verbos `%s`, `%d` e `%.2f`
- [ ] O programa compila sem variáveis não usadas

## Exercício 03 - FizzBuzz

**Nível:** iniciante

**Objetivo:** Praticar laços e condicionais.

**Instrução:**
Imprima os números de 1 a 100. Para múltiplos de 3, imprima `Fizz` no lugar do número; para múltiplos de 5, `Buzz`; para múltiplos de ambos, `FizzBuzz`.

**Critérios de sucesso:**

- [ ] Usa um único laço `for`
- [ ] Usa `switch` sem expressão (`switch { case ...: }`) em vez de vários `if`
- [ ] A saída está correta para 3, 5, 15 e 100

## Exercício 04 - Divisão Segura

**Nível:** iniciante

**Objetivo:** Escrever funções com múltiplos retornos e tratar erros como valores.

**Instrução:**
Crie a função `dividir(a, b float64) (float64, error)` que retorna um erro quando `b` for zero. No `main`, chame a função com valores válidos e inválidos e trate o erro.

**Critérios de sucesso:**

- [ ] A função retorna `errors.New` (ou `fmt.Errorf`) quando `b == 0`
- [ ] O `main` verifica `if err != nil` antes de usar o resultado
- [ ] A mensagem de erro é exibida sem derrubar o programa

## Exercício 05 - Contador de Palavras

**Nível:** intermediário

**Objetivo:** Praticar slices, maps e o pacote `strings`.

**Instrução:**
Dado um texto, conte quantas vezes cada palavra aparece (ignorando maiúsculas e minúsculas) e imprima as palavras em ordem alfabética com suas contagens.

**Critérios de sucesso:**

- [ ] Usa `strings.Fields` e `strings.ToLower`
- [ ] Usa um `map[string]int` para as contagens
- [ ] Ordena as chaves com o pacote `slices` ou `sort` antes de imprimir
- [ ] Tem ao menos um teste com `go test` validando a contagem
