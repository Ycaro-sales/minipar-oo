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

### ❌ Não funciona (parse ou semântico falham)

| Construção | Problema |
|---|---|
| Literal de array `[1, 2, 3]` | Parser aceita, mas semântico não resolve o tipo → TAC emite `ASSIGN ERRO` |
| Literal de dict `{"a": 1}` | Erro de parse: `expressão primária inválida: '{'` |
| Literal de set `{1, 2, 3}` | Erro de parse: `expressão primária inválida: '{'` |
| Literal de tupla `(1, 2)` | Erro de parse: confundido com expressão parentesizada |
| Função literal `func(x) { ... }` | Parser aceita, mas semântico não resolve o tipo → TAC emite `ASSIGN ERRO` |
| `for (x in arr) { }` | TAC gerado, mas `arr` tem tipo `ERRO` quando é literal de array |
| Tipo array `[i32]` em parâmetros | Resolve para `any`; TAC não sabe o tipo do elemento |

---

## Próximos passos

1. **Tipos compostos no semântico** — resolver tipos de literais de array, dict, set e tupla
2. **TAC para tipos compostos** — substituir `ASSIGN ERRO` por instruções adequadas (`ARRAY_NEW`, `DICT_NEW`, etc.)
3. **Geração C para tipos compostos** — representar arrays como ponteiro + tamanho, ou usar uma struct dinâmica
4. **Paralelismo real** — substituir os stubs `BEGIN_PAR`/`END_PAR` por `pthread_create`
5. **`input()`** — implementar `scanf` correspondente no cgen
6. **`for-in`** — depende de arrays funcionando
7. **Dict e Set** — corrigir parse de literais com `{`
8. **Tuplas** — corrigir parse de `(expr, expr)`

---

## Gramática

Ver [`BNF.bnf`](BNF.bnf) para a gramática formal e [`DEFINITIONS.md`](DEFINITIONS.md) para a especificação da linguagem.
