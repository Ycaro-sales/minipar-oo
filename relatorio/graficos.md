### Diagrama de Casos de Uso

```mermaid
flowchart LR
    Dev([Desenvolvedor])

    Dev --> UC1[Compilar para binário nativo]
    Dev --> UC2[Emitir lista de tokens]
    Dev --> UC3[Emitir AST como JSON]
    Dev --> UC4[Emitir TAC]
    Dev --> UC5[Emitir código C gerado]
    Dev --> UC6[Usar Minipar Studio]

    UC1 --> GCC[[GCC]]
    UC6 --> UC1
```

### Arquitetura de Componentes

```mermaid
flowchart LR
    SRC["Fonte .minipar"]

    subgraph Pipeline
        LEX["lexer/\nTokenizador"]
        PAR["parser/\nParser Desc. Rec."]
        AST["ast/\nNós da AST"]
        SEM["semantic/\nAnalisador Semântico"]
        SYM["internal/symtab/\nTabela de Símbolos"]
        TAC["tac/\nGerador de TAC"]
        CGEN["cgen/\nGerador de Código C"]
    end

    COMP["compiler/\nOrquestrador"]
    MCC["mcc\nCLI"]
    GCC["gcc"]
    BIN["Binário nativo"]

    SRC --> LEX --> PAR --> AST --> SEM --> TAC --> CGEN --> GCC --> BIN
    SEM <--> SYM
    CGEN --> SYM
    MCC --> COMP
    COMP --> LEX
    COMP --> PAR
    COMP --> SEM
    COMP --> TAC
    COMP --> CGEN
```

### Diagramas de Classe por Componente

#### Lexer e Parser

```mermaid
classDiagram
    class Lexer {
        -input string
        -position int
        -nextPosition int
        -character byte
        -line int
        +New(input string) Lexer
        +NextToken() Token
        +Line() int
        -readCharacter()
        -peekCharacter() byte
        -skipWhiteSpace()
        -skipSingleLineComment()
        -skipMultiLineComment()
        -readIdentifier() string
        -readNumber() string
        -readString() string
    }

    class Token {
        +Type TokenType
        +Literal string
        +Line int
    }

    class Tokenizer {
        <<interface>>
        +NextToken() Token
    }

    class Parser {
        <<interface>>
        +ParseProgram(src string) Program
    }

    class miniparParser {
        -factory LexerFactory
        -currentToken Token
        -peekToken Token
        -errors []string
        +ParseProgram(src string) Program
        -parseDeclaration() Declaration
        -parseStatement() Statement
        -parseExpression() Expression
        -parseOrExpr() Expression
        -parseAndExpr() Expression
        -parseComparisonExpr() Expression
        -parseAdditiveExpr() Expression
        -parseMultiplicativeExpr() Expression
        -parseUnaryExpr() Expression
        -parsePostfixExpr() Expression
        -parsePrimaryExpr() Expression
    }

    Lexer ..|> Tokenizer
    miniparParser ..|> Parser
    miniparParser --> Tokenizer
    Lexer --> Token
```

#### Analisador Semântico e Tabela de Símbolos

```mermaid
classDiagram
    class TACGenerator {
        -instructions []TAC
        -tempCount int
        -labelCount int
        -tempTypes map~string~
        -loopStack []LoopLabels
        +New() TACGenerator
        +Generate(node Node) string
        +Instructions() string
        +RawInstructions() []Instruction
        +TempTypes() map~string~
        -newTemp() string
        -newTypedTemp(typ string) string
        -newLabel() string
        -emit(op arg1 arg2 result string)
    }

    class Instruction {
        +Op string
        +Arg1 string
        +Arg2 string
        +Result string
    }

    class CGenerator {
        -instrs []Instruction
        -tempTypes map~string~
        -varTypes map~string~
        -globalVars map~bool~
        -localVars map~bool~
        -paramBuf []string
        -currentClass string
        -insideFunc bool
        +New(instrs tempTypes globalScope) CGenerator
        +Generate() string
        -declPrefix(name string) string
        -resolveType(name string) string
        -emitBinary(result arg1 op arg2 string)
        -emitUnary(result op arg string)
    }

    TACGenerator --> Instruction
    CGenerator --> Instruction
```

#### Gerador TAC e Gerador C

```mermaid
classDiagram
    class TACGenerator {
        -instructions []TAC
        -tempCount int
        -labelCount int
        -tempTypes map~string~
        -loopStack []LoopLabels
        +New() TACGenerator
        +Generate(node Node) string
        +Instructions() string
        +RawInstructions() []Instruction
        +TempTypes() map~string~
        -newTemp() string
        -newTypedTemp(typ string) string
        -newLabel() string
        -emit(op arg1 arg2 result string)
    }

    class Instruction {
        +Op string
        +Arg1 string
        +Arg2 string
        +Result string
    }

    class CGenerator {
        -instrs []Instruction
        -tempTypes map~string~
        -varTypes map~string~
        -globalVars map~bool~
        -localVars map~bool~
        -paramBuf []string
        -currentClass string
        -insideFunc bool
        +New(instrs tempTypes globalScope) CGenerator
        +Generate() string
        -declPrefix(name string) string
        -resolveType(name string) string
        -emitBinary(result arg1 op arg2 string)
        -emitUnary(result op arg string)
    }

    TACGenerator --> Instruction
    CGenerator --> Instruction
```
