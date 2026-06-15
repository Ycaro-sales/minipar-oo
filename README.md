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
- **Arrays estáticos `[T]`:** declaração, literal `[v0, v1, ...]`, acesso `arr[i]`, mutação `arr[i] = x`, passagem para funções (struct com ponteiro compartilhado), `for (x in arr)`
- **Tuplas `(T0, T1, ...)`:** declaração, literal `(v0, v1)`, acesso por índice constante; imutáveis por design
- **Funções built-in:** `print`, `input`, `len`, `to_string`, `to_number`, `isalpha`, `isnum` — geração C completa (ver tabela abaixo)

### ✅ Funções built-in (geração C implementada)

| Função | Retorno | Geração C |
|--------|---------|-----------|
| `print(x)` | `void` | `printf` via instrução TAC `PRINT` |
| `input([prompt])` | `string` | helper `mp_input()` emitido no preâmbulo; prompt via `printf` |
| `len(x)` | `int` | traduzido para `strlen(x)` (`<string.h>`); comprimento de array via `ARRAY_LEN` → `.len` |
| `to_string(x)` | `string` | helper `static char* to_string(long long)` com `snprintf` emitido no preâmbulo |
| `to_number(x)` | `int` | traduzido para `atoi(x)` (`<stdlib.h>`) |
| `isalpha(x)` | `bool` | traduzido para `isalpha(x)` de `<ctype.h>` |
| `isnum(x)` | `bool` | traduzido para `isdigit(x)` de `<ctype.h>` |
| `arr.append(elem)` | `void` | não suportado (exige array dinâmico; fora do escopo atual) |

> **Nota:** `<ctype.h>` é incluído automaticamente apenas quando `isalpha` ou `isnum` é usado no programa. Helpers próprios (`to_string`, `mp_input`) são emitidos no preâmbulo somente quando usados, evitando warnings de funções estáticas não utilizadas.

### ⚠️ TAC gerado, mas geração C incompleta

| Construção | Situação |
|---|---|
| `par { }` | Emite comentário `/* BEGIN_PAR — pthreads: TODO */`; corpo executado sequencialmente |
| `seq { }` | Emite comentários; funcionalmente correto (já é sequencial) |

### 🚫 Não suportado (trabalhos futuros)

Estas construções foram avaliadas e **não serão implementadas** na versão atual por exigirem infraestrutura de memória dinâmica que está fora do escopo do projeto.

| Construção | Motivo |
|---|---|
| `dict {"a": 1}` | Exige tabela hash — memória dinâmica |
| `set {1, 2, 3}` | Exige tabela hash — memória dinâmica |
| `arr.append(elem)` | Exige array dinâmico com realloc |
| Função literal `func(x) { }` | Parser aceita, semântico não resolve o tipo |
| Indexação de String com char** | Alteração no C Generator - ARRAY_GET em char* deve emitir `s[1]` |



---

## Gramática

Ver [`BNF.bnf`](BNF.bnf) para a gramática formal e [`DEFINITIONS.md`](DEFINITIONS.md) para a especificação da linguagem.

---

## Relatório

Ver [Relatório Compilador Minpar LSMVS](relatorio/relatorio_lsmvs.pdf) para detalhamento das estruturas do projeto.
