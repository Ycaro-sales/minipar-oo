# Nós da AST

Resumo de todos os nós definidos em [`ast/ast.go`](./ast/ast.go).

## Interfaces base

Toda a árvore é construída sobre quatro interfaces. `Node` é a raiz; as outras
três marcam o papel sintático de cada nó (um mesmo tipo pode satisfazer mais de
uma — ex.: `VarDecl` é declaração **e** comando).

| Interface     | Método marcador | Significado                                            |
|---------------|-----------------|--------------------------------------------------------|
| `Node`        | `GetLine() int` | Qualquer nó da AST; expõe a linha de origem.           |
| `Statement`   | `stmtNode()`    | Executa uma ação, não produz valor.                    |
| `Expression`  | `exprNode()`    | Produz um valor.                                       |
| `Declaration` | `declNode()`    | Declaração de topo (classe, interface, função, var).   |

## Operadores

`Op` é um enum (`int`) que tipa os operadores, evitando comparação de strings em
passes posteriores.

| Categoria   | Valores                                              |
|-------------|------------------------------------------------------|
| Lógicos     | `OpOr`, `OpAnd`, `OpNot`                              |
| Comparação  | `OpEq`, `OpNeq`, `OpLt`, `OpGt`, `OpLeq`, `OpGeq`     |
| Aritméticos | `OpAdd`, `OpSub`, `OpMul`, `OpDiv`, `OpMod`, `OpNeg`  |

## Raiz

| Nó        | Campos                       | Descrição                              |
|-----------|------------------------------|----------------------------------------|
| `Program` | `Declarations []Declaration` | Raiz da árvore; lista de declarações.  |

## Tipos compartilhados

| Tipo    | Campos                  | Descrição                                       |
|---------|-------------------------|-------------------------------------------------|
| `Param` | `Name string`, `Type string` | Parâmetro de função/método (não é um `Node`). |

## Declarações (`Declaration`)

| Nó                | Campos principais                                          | Descrição                                              |
|-------------------|------------------------------------------------------------|-------------------------------------------------------|
| `ClassDecl`       | `Name`, `Implements []string`, `Members []Node`            | Declaração de classe; membros são campos/métodos/construtor. |
| `InterfaceDecl`   | `Name`, `Methods []*InterfaceMethod`                       | Declaração de interface.                               |
| `FuncDecl`        | `Name`, `Params []Param`, `ReturnType string`, `Body *BlockStmt` | Declaração de função de topo.                   |
| `VarDecl`         | `Name`, `Type string`, `Value Expression`                  | Declaração de variável (`let`). Também é `Statement`. |

## Membros de classe / interface

Estes nós são `Node` (têm `GetLine`), mas não são `Declaration` nem `Statement`;
aparecem apenas dentro de `ClassDecl.Members` ou `InterfaceDecl.Methods`.

| Nó                | Campos principais                                          | Descrição                                  |
|-------------------|------------------------------------------------------------|--------------------------------------------|
| `InterfaceMethod` | `Name`, `Params []Param`, `ReturnType string`              | Assinatura de método em uma interface.     |
| `FieldDecl`       | `Name`, `Type string`, `Value Expression`                  | Atributo de classe (com valor opcional).   |
| `ConstructorDecl` | `Name`, `Body *BlockStmt`                                  | Construtor de classe.                      |
| `MethodDecl`      | `Name`, `Params []Param`, `ReturnType string`, `Body *BlockStmt` | Método de classe.                    |

## Comandos (`Statement`)

| Nó             | Campos principais                                  | Descrição                                            |
|----------------|----------------------------------------------------|------------------------------------------------------|
| `BlockStmt`    | `Statements []Statement`                           | Bloco `{ ... }`.                                     |
| `IfStmt`       | `Condition`, `Then *BlockStmt`, `Else *BlockStmt`  | Condicional (`Else` opcional).                       |
| `WhileStmt`    | `Condition`, `Block *BlockStmt`                     | Laço `while`.                                        |
| `DoStmt`       | `Block *BlockStmt`, `Condition`                     | Laço `do { } while`.                                 |
| `ForStmt`      | `Var string`, `Iter Expression`, `Block *BlockStmt`| Laço `for (x in iter)`.                              |
| `SwitchStmt`   | `Expr Expression`, `Cases []*CaseClause`           | Seleção `switch`.                                    |
| `CaseClause`   | `Value Expression`, `Block *BlockStmt`             | Um `case` do switch (não é `Statement` por si só).   |
| `SeqStmt`      | `Block *BlockStmt`                                  | Bloco sequencial `seq { }`.                          |
| `ParStmt`      | `Block *BlockStmt`                                  | Bloco paralelo `par { }`.                            |
| `ChannelStmt`  | `ChanType string`, `Name string`, `Args []Expression` | Declaração de canal (`s_channel`/`c_channel`).    |
| `Assignment`   | `Name string`, `Value Expression`                  | Atribuição `x = expr`.                               |
| `PrintStmt`    | `Args []Expression`                                | Comando `print(...)`.                                |
| `InputStmt`    | `Prompt Expression`                                | Comando `input(...)` (prompt opcional).              |
| `ReturnStmt`   | `Value Expression`                                 | `return` (valor opcional).                           |
| `BreakStmt`    | `Line int`                                         | `break`.                                             |
| `ContinueStmt` | `Line int`                                         | `continue`.                                          |
| `PassStmt`     | `Line int`                                         | `pass` (no-op).                                      |
| `GotoStmt`     | `Label string`                                     | `goto label`.                                        |
| `ExprStmt`     | `Expression Expression`                            | Expressão usada como comando.                        |

## Expressões (`Expression`)

| Nó            | Campos principais                                      | Descrição                                          |
|---------------|--------------------------------------------------------|----------------------------------------------------|
| `BinaryExpr`  | `Left`, `Operator Op`, `Right`                         | Operação binária.                                  |
| `UnaryExpr`   | `Operator Op`, `Right`                                 | Operação unária (`-x`, `!x`).                      |
| `Identifier`  | `Value string`                                         | Referência a um nome.                              |
| `FuncCall`    | `Name string`, `Args []Expression`                     | Chamada `f(args)`. Também é `Statement`.           |
| `MethodCall`  | `Object Expression`, `Method string`, `Args []Expression` | Chamada `obj.m(args)`. Também é `Statement`.     |
| `IndexExpr`   | `Object Expression`, `Index Expression`                | Indexação `obj[i]`.                                |
| `ObjCreation` | `Class string`, `Args []Expression`                    | Instanciação de objeto.                            |
| `FuncLiteral` | `Params []Param`, `ReturnType string`, `Body *BlockStmt` | Função anônima (closure).                        |

## Literais (`Expression`)

| Nó               | Campos principais                          | Descrição                       |
|------------------|--------------------------------------------|---------------------------------|
| `IntLiteral`     | `Value int64`                              | Literal inteiro.                |
| `FloatLiteral`   | `Value float64`                            | Literal de ponto flutuante.     |
| `StringLiteral`  | `Value string`                             | Literal de string.              |
| `CharLiteral`    | `Value rune`                               | Literal de caractere.           |
| `BooleanLiteral` | `Value bool`                               | `true` / `false`.               |
| `ListLiteral`    | `Elements []Expression`                    | Literal de lista `[...]`.       |
| `DictLiteral`    | `Pairs map[Expression]Expression`          | Literal de dicionário `{k: v}`. |
| `SelfExpr`       | `Line int`                                 | Referência `Self`.              |

## Observações

- Todo nó carrega `Line int` (direta ou indiretamente) e implementa `GetLine()`
  para diagnósticos com número de linha.
- Os tipos são armazenados como `string` crua (ex.: `Type`, `ReturnType`),
  resolvidos para tipos concretos apenas na análise semântica.
- Alguns nós cumprem papel duplo: `VarDecl` (`Declaration` + `Statement`),
  `FuncCall` e `MethodCall` (`Expression` + `Statement`).
