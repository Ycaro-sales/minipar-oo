# Minipar

Compilador da linguagem **Minipar** — uma linguagem com suporte a orientação a objetos, programação funcional e paralelismo. Implementado em Go com pipeline recursivo-descendente.

## Pipeline

```
Fonte (.minipar) → Lexer → Parser → AST → Analisador Semântico → TAC → Gerador C
```

## Uso

```bash
# Compilar o binário
go build -o minipar .

# Imprimir o TAC gerado
./minipar tests/ex1.minipar

# Imprimir a AST como JSON
./minipar -ast tests/ex1.minipar

# Gerar código C
./minipar -c tests/ex1.minipar

# Compilar o C gerado e executar
./minipar -c tests/ex1.minipar | gcc -x c - -o out && ./out

# Rodar todos os testes
go test ./...
```

## Pacotes

| Pacote | Responsabilidade |
|--------|-----------------|
| `lexer/` | Tokenizador hand-written |
| `parser/` | Parser recursivo-descendente |
| `ast/` | Definições dos nós da AST e operadores |
| `semantic/` | Analisador semântico em dois passes (coleta de assinaturas + verificação de corpos) |
| `internal/symtab/` | Tabela de símbolos aninhada e retida (genérica sobre o payload de tipo) |
| `tac/` | Gerador de Three-Address Code |
| `cgen/` | Gerador de código C a partir do TAC |
| `compiler/` | Orquestrador do pipeline com injeção de dependência |

---

## Estado de Implementação

### ✅ Funcionando (TAC + geração C)

- **Tipos primitivos:** `i8`–`i64`, `u8`–`u64`, `f16`–`f64`, `bool`, `char`, `string`, `void`, `any`
- **Variáveis:** declaração com tipo explícito, inicialização, reatribuição, variáveis globais
- **Aritmética:** `+`, `-`, `*`, `/`, `%`, negação unária `-`
- **Lógica:** `and`, `or`, `!`
- **Comparação:** `==`, `!=`, `<`, `>`, `<=`, `>=`
- **Controle de fluxo:** `if/else`, `while`, `do-while`, `break`, `continue`
- **`switch`:** gerado via gotos em C (semânticamente correto)
- **Funções:** declaração, parâmetros tipados, retorno, chamadas, recursão
- **I/O:** `print` (com resolução de formato printf por tipo)
- **OOP:** `class` com campos, construtor, métodos; name mangling `Classe_metodo`
- **Variáveis globais:** declaradas fora de funções

### ⚠️ Funções built-in (semântico ok, C não implementado)

O analisador semântico registra as seguintes funções nativas. No C gerado elas precisam de implementação via macro ou função auxiliar emitida no preâmbulo — hoje o cgen não gera esse código.

| Função | Retorno | Equivalente C sugerido |
|--------|---------|------------------------|
| `print(x)` | `void` | `printf` (já funciona via instrução TAC `PRINT`) |
| `input()` | `string` | `scanf` / `fgets` — TAC emite `INPUT`, cgen descarta |
| `len(x)` | `int` | `strlen(x)` para strings; `sizeof(arr)/sizeof(arr[0])` para arrays |
| `to_string(x)` | `string` | `sprintf` em buffer estático |
| `to_number(x)` | `int` | `atoi(x)` |
| `isalpha(x)` | `bool` | `isalpha(x)` de `<ctype.h>` |
| `isnum(x)` | `bool` | `isdigit(x)` de `<ctype.h>` |
| `arr.append(elem)` | `void` | não suportado (exige array dinâmico; fora do escopo atual) |

### ⚠️ TAC gerado, mas geração C incompleta

| Construção | Situação |
|---|---|
| `input()` | TAC emite `INPUT`, cgen descarta silenciosamente |
| `ARRAY_GET arr idx` | TAC emite, cgen descarta silenciosamente |
| `ARRAY_LEN arr` | TAC emite, cgen descarta silenciosamente |
| `METHOD_CALL obj.m n` | TAC emite, cgen descarta (exceto em contexto de classe) |
| `NEW_OBJ Classe n` | TAC emite, cgen descarta silenciosamente |
| `par { }` | Emite comentário `/* BEGIN_PAR — pthreads: TODO */`; corpo executado sequencialmente |
| `seq { }` | Emite comentários; funcionalmente correto (já é sequencial) |
| Canais (`s_channel`, `c_channel`) | TAC emite `CHAN_DECL`; cgen emite apenas comentário |

### ❌ Em progresso (parse ou semântico incompletos)

| Construção | Problema |
|---|---|
| Literal de array `[1, 2, 3]` | Parser aceita, mas semântico não resolve o tipo → TAC emite `ASSIGN ERRO` |
| Tipo array `[i32]` em parâmetros | Resolve para `any`; TAC não sabe o tipo do elemento |
| Index-assignment `arr[i] = x` | Gramática só permite `<id> = <expr>`; arrays são somente-leitura hoje |
| Arrays por referência | Não modelado; necessário para `swap` e algoritmos in-place |
| Função literal `func(x) { ... }` | Parser aceita, mas semântico não resolve o tipo → TAC emite `ASSIGN ERRO` |
| `for (x in arr) { }` | TAC gerado, mas `arr` tem tipo `ERRO` quando é literal de array |

### 🚫 Não suportado (decisão de escopo)

Estas construções foram avaliadas e **não serão implementadas** na versão atual por exigirem infraestrutura de memória dinâmica (tabela hash) que está fora do escopo do projeto.

| Construção | Motivo |
|---|---|
| `dict {"a": 1}` | Exige tabela hash — memória dinâmica |
| `set {1, 2, 3}` | Exige tabela hash — memória dinâmica |
| `tuple (1, "a")` | Parser confunde com expressão parentesizada; baixa prioridade |
| `arr.append(elem)` | Exige array dinâmico com realloc |

---

## Próximos passos

1. **Semântico para arrays** — resolver tipo de literais `[1, 2, 3]` e parâmetros `[i32]`
2. **TAC para arrays** — emitir `ARRAY_NEW`, `ARRAY_GET`, `ARRAY_LEN` corretamente
3. **Geração C para arrays estáticos** — `int arr[] = {1, 2, 3}`, acesso por índice, `len` via `sizeof`
4. **Index-assignment e arrays por referência** — habilitar `arr[i] = x` e passagem por ponteiro (pré-requisitos do quicksort)
5. **`for-in`** — depende de arrays funcionando
6. **`input()`** — implementar `scanf`/`fgets` no cgen
7. **`len`, `to_string`, `to_number`, `isalpha`, `isnum`** — emitir funções auxiliares C no preâmbulo do arquivo gerado
8. **Paralelismo real** — substituir os stubs `BEGIN_PAR`/`END_PAR` por `pthread_create`

---

## Gramática

Ver [`BNF.bnf`](BNF.bnf) para a gramática formal e [`DEFINITIONS.md`](DEFINITIONS.md) para a especificação da linguagem.
