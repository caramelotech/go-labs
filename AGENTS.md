# AGENTS.md

Orientações especializadas para agentes de IA trabalhando com este repositório de conteúdo.

## Contexto do Projeto

Este é o repositório de **conteúdo** do Go Labs (Caramelo Tech):

- **Conteúdo:** 100% em `notes/`, Markdown puro em português brasileiro, sem frontmatter
- **Publicação:** o repositório hub [labs](https://github.com/caramelotech/labs) busca as notas daqui e publica em `https://caramelotech.com.br/labs/go/`
- **Sem build de site:** não há dependências Node - apenas Markdown e exemplos Go em `examples/`
- **Seções:** ver "Seções de conteúdo" abaixo

Veja o [README.md](README.md) para a visão geral do repositório.

## Seções de conteúdo

Cada pasta em `notes/` é uma seção do site, com um escopo próprio. Use o escopo para decidir onde uma nota nova entra. A ideia é não manter uma lista de tópicos por pasta aqui nem no `index.md` - ao adicionar conteúdo dentro de um escopo já descrito, esses arquivos não precisam mudar.

- **`fundamentos/`** - primeiros passos com Go: instalar e configurar o ambiente, entender o que é a linguagem e escrever os primeiros programas (sintaxe, tipos, funções, controle de fluxo, coleções, structs, ponteiros, pacotes e módulos). Para quem nunca programou em Go.
- **`linguagem/`** - a linguagem a fundo, depois da sintaxe básica: métodos e receivers, interfaces e composição, tratamento de erros, testes e, quando houver material, generics e biblioteca padrão.
- **`concorrencia/`** - goroutines, channels, `select`, padrões como worker pool e pipeline, `context` e sincronização com o pacote `sync`.
- **`backend/`** - Go aplicado a serviços: HTTP com `net/http`, middlewares, logging e observabilidade, e no futuro banco de dados e deploy.
- **`pratica/`** - exercícios (`01-exercicios.md`) e projetos (`02-projetos.md`) para aplicar o conteúdo do lab.
- **`roadmap.md`** (nota solta) - trilha de estudos de Go, do primeiro programa à produção. É um guia de percurso, mantido à parte e independente da cobertura atual das notas.
- **`recursos.md`** (nota solta) - materiais externos para aprofundar (livros, cursos, documentação, comunidades).

As notas citam os exemplos de `examples/` que as acompanham com link absoluto para o GitHub (`https://github.com/caramelotech/go-labs/tree/main/examples/<pasta>`), já que os exemplos não são publicados no site.

## Tarefas Comuns

### Adicionar uma nova nota

1. **Escolha a pasta pelo escopo:** consulte "Seções de conteúdo" e pegue a pasta cujo escopo cobre o tema. Se nenhuma cobre e há material para 2-3 notas relacionadas, avalie uma pasta nova (ver "Criar nova seção de tema").
2. **Nomeie com prefixo numérico** para controlar a ordem na barra lateral: se a pasta já tem `02-*.md`, crie `03-nome-do-topico.md`
3. **Primeira linha = título:** comece o arquivo com `# Título da Nota` - o site usa esse H1 como título da página
4. **Sem frontmatter:** escreva direto o Markdown

Exemplo de nota nova (`notes/fundamentos/03-sintaxe.md`):

```markdown
# Sintaxe Básica

## Introdução

Conteúdo aqui...
```

### Criar nova seção de tema

1. Crie a subpasta em `notes/nova-secao/` (nome em minúsculas) com ao menos uma nota
2. Adicione a seção em `sidebar.json`:
   ```json
   { "label": "Título da Seção", "directory": "nova-secao" }
   ```
3. Descreva o escopo da nova pasta em "Seções de conteúdo" e no `notes/index.md`

### Adicionar um exemplo de código

1. Crie uma pasta em `examples/` com prefixo numérico: `examples/02-variaveis/`
2. Coloque o código em `main.go` (pacote `main`) e um `README.md` explicando objetivo, como executar e saída esperada
3. Garanta que o código roda com `go run .` e está formatado com `gofmt`

## Regras de Conteúdo

- Use **hífens (-)**, não travessões (—)
- NÃO use `---` para separar seções (apenas para notas/atribuições no final do arquivo)
- Apenas um `# H1` por arquivo, na primeira linha
- Imagens ficam junto das notas e são referenciadas com caminho relativo em sintaxe Markdown: `![descrição](./assets/img.png)` - nunca `<img>` HTML nem caminho absoluto
- Links para outras notas do site: caminho completo `/labs/go/<secao>/<nota>/` (slug em minúsculas, igual ao nome da pasta)
- Blocos de código Go usam a linguagem `go` e devem estar formatados como o `gofmt` deixaria (tabs para indentação)

## Publicação

- Push em `main` alterando `notes/` ou `sidebar.json` dispara o workflow `notify-hub.yml`, que aciona o rebuild do site no hub
- O workflow requer o secret `HUB_DISPATCH_TOKEN` configurado no repositório
- Para validar como a nota fica no site, rode no clone do hub: `npm run fetch:local && npm run build`

## Git Conventions

- **NUNCA** fazer `git commit` ou `git push` automaticamente
- Apenas executar comandos git quando explicitamente solicitado pelo usuário
- Comunicar claramente o que será commitado antes de executar

Para informações adicionais, ver o [README.md](README.md).
