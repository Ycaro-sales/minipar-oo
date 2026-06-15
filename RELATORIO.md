<!--
Build:  pandoc RELATORIO.md -o RELATORIO.pdf \
          --pdf-engine=xelatex \
          -V geometry:margin=2.5cm \
          -V lang=pt-BR

Deps:   sudo apt install pandoc texlive-xetex texlive-lang-portuguese \
                         texlive-latex-recommended texlive-fonts-recommended
-->

---
title: "Relatório Técnico: Desenvolvimento do Compilador para a Linguagem MiniPar v2025.1"
author:
  - Wyvian Gabrielly Cavalcante Valença
  - Ycaro Bruno Souza Sales
  - Felipe Lira da Silva
  - Marcos Mendonça
  - João Gabriel Seixas
date: "2025"
lang: pt-BR
toc: true
toc-depth: 3
numbersections: true
geometry: margin=2.5cm
mainfont: "DejaVu Serif"
monofont: "DejaVu Sans Mono"
---

\begin{center}

\vspace*{2cm}

\textbf{UNIVERSIDADE FEDERAL DE ALAGOAS}

Instituto de Computação

\textbf{CIÊNCIA DA COMPUTAÇÃO}

\vspace{4cm}

Wyvian Gabrielly Cavalcante Valença\\
Ycaro Bruno Souza Sales\\
Felipe Lira da Silva\\
Marcos Mendonça\\
João Gabriel Seixas

\vspace{3cm}

\textbf{Relatório Técnico: Desenvolvimento do Compilador para a Linguagem MiniPar v2025.1}

\vspace{1cm}

GitHub: \texttt{https://github.com/Ycaro-sales/minipar-oo}

Vídeo de demonstração: Não disponível

\vfill

Maceió — AL\\
2025

\end{center}

\newpage

# Introdução

## Objetivo

<!-- TODO: descrever o objetivo geral do projeto — desenvolver um compilador completo para a linguagem MiniPar 2025.1, com geração de código C nativo. -->

## Escopo

<!-- TODO: descrever as fases implementadas -->

### Análise Léxica

<!-- TODO: adicionar principais funcionalidades e caracteristicas -->

### Análise Sintática com Construção de Árvore Sintática Abstrata (AST)

<!-- TODO: adicionar principais funcionalidades e caracteristicas -->

### Análise Semântica

<!-- TODO : adicionar principais funcionalidades e caracteristicas-->

### Tabela de Símbolos

<!-- TODO : adicionar principais funcionalidades e caracteristicas-->

### Geração de Código Intermediário (TAC — Three-Address Code)

<!-- TODO : adicionar principais funcionalidades e caracteristicas-->

### Geração de Código C

<!-- TODO : adicionar principais funcionalidades e caracteristicas-->

### Interface de Linha de Comando (CLI — `mcc`)

<!-- TODO : adicionar principais funcionalidades e caracteristicas-->

# A Linguagem MiniPar 2025.1

<!-- TODO: breve parágrafo introdutório sobre a linguagem -->

## Gramática da Linguagem MiniPar

### Estrutura Geral

<!-- TODO: colar/referenciar produções BNF de BNF.bnf -->

```bnf
<!-- TODO -->
```

### Precedência de Operadores (do menor para o maior)

<!-- TODO: tabela de precedência -->

| Precedência | Operador | Associatividade |
|-------------|----------|-----------------|
| <!-- TODO --> | | |

## Elementos Léxicos

### Palavras-chave

<!-- TODO : adicionar principais funcionalidades e caracteristicas-->

### Tipos

<!-- TODO : adicionar principais funcionalidades e caracteristicas-->

### Literais

<!-- TODO : adicionar principais funcionalidades e caracteristicas-->

### Operadores

<!-- TODO : adicionar principais funcionalidades e caracteristicas-->

### Delimitadores

<!-- TODO : adicionar principais funcionalidades e caracteristicas-->

# Metodologia e Modelagem de Software

## Gerenciamento de Projeto (Backlogs)

### Product Backlog

<!-- TODO: tabela com funcionalidades, descrição e prioridade -->

| Funcionalidade | Descrição | Prioridade |
|----------------|-----------|------------|
| <!-- TODO --> | | |

### Sprint Backlog

<!-- TODO: estruturar divisao das tarefas em sprints -->

## Modelagem UML

### Diagrama de Casos de Uso

<!-- TODO: inserir imagem ou descrição -->

### Arquitetura de Componentes

<!-- TODO: diagrama mostrando os pacotes lexer, parser, ast, semantic, symtab, tac, cgen, compiler -->

### Diagramas de Classe por Componente

<!-- TODO: um diagrama por pacote principal -->

# Arquitetura e Implementação

<!-- TODO: parágrafo introdutório sobre o pipeline:
Source (.minipar) → Lexer → Token stream → Parser → AST → Semantic → TAC → Cgen → C source -->

## Analisador Léxico (`lexer/`)

<!-- TODO: responsabilidade, entrada, saída, pseudocódigo/trechos relevantes -->

```go
// TODO
```

## Analisador Sintático e AST (`parser/`, `ast/`)

### AST

<!-- TODO -->

### Analisador Sintático

<!-- TODO -->

```go
// TODO
```

## Analisador Semântico e Tratamento de Erros (`semantic/`)

<!-- TODO: duas passagens — coleta de assinaturas e verificação de corpos -->

```go
// TODO
```

## Tabela de Símbolos (`internal/symtab/`)

<!-- TODO: estrutura de árvore de escopos aninhados, parametrizada -->

```go
// TODO
```

## Gerador de Código Intermediário / TAC (`tac/`)

<!-- TODO: instruções de três endereços, temporários tipados -->

```
// TODO: exemplo de instrução TAC gerada
```

## Gerador de Código C (`cgen/`)

<!-- TODO: mapeamento TAC → C, referência a TRANSLATIONS.md -->

```c
// TODO: exemplo de trecho C gerado
```

## Orquestração do Pipeline (`compiler/`, `mcc`)

<!-- TODO: como compiler.go une as fases; flags -tokens, -ast, -tac, -c da CLI -->

```go
// TODO
```

# Tecnologias Utilizadas

<!-- TODO: tabela com linguagem, versão, justificativa -->

| Tecnologia | Versão | Papel no Projeto |
|------------|--------|-----------------|
| Go | <!-- TODO --> | Implementação do compilador |
| GCC | <!-- TODO --> | Compilação do código C gerado |
| <!-- TODO --> | | |

# Testes e Validação

## Interface do Compilador (CLI `mcc`)

<!-- TODO: mostrar saídas de -tokens, -ast, -tac, -c para um programa de exemplo -->

```
$ mcc -tac tests/ex1.minipar
# TODO: colar saída
```

## Testes de Integração

<!-- TODO: go test ./... output -->

```
$ go test ./...
# TODO
```

## Execução dos Programas de Teste

### Programa de Teste 1: <!-- TODO: nome -->

<!-- TODO: código-fonte MiniPar + TAC gerado + C gerado + saída de execução -->

```minipar
// TODO
```

### Programa de Teste 2: Fatorial

<!-- TODO -->

```minipar
// TODO
```

### Programa de Teste 3: Fibonacci

<!-- TODO -->

```minipar
// TODO
```

### Programa de Teste 4: Quicksort

<!-- TODO -->

```minipar
// TODO
```

# Conclusão

<!-- TODO: resultados alcançados, dificuldades, trabalhos futuros -->

# Anexos

## Link para o Repositório (GitHub)

<https://github.com/Ycaro-sales/minipar-oo>

## Link para o Vídeo de Apresentação

Não disponível.
