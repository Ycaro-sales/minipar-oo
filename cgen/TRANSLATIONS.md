# TAC → C: Tabela de Traduções

Este documento descreve como cada instrução TAC (Three Address Code) do compilador
minipar é traduzida para código C pelo gerador `cgen`.

---

## Preâmbulo

Todo arquivo C gerado começa com os seguintes includes:

```c
#include <stdint.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
```

---

## Mapeamento de Tipos

| Tipo Minipar        | Tipo C       |
|---------------------|--------------|
| `i8`                | `int8_t`     |
| `i16`               | `int16_t`    |
| `i32`               | `int32_t`    |
| `i64`               | `int64_t`    |
| `u8`                | `uint8_t`    |
| `u16`               | `uint16_t`   |
| `u32`               | `uint32_t`   |
| `u64`               | `uint64_t`   |
| `f16` / `f32` / `float` | `float` |
| `f64`               | `double`     |
| `bool`              | `bool`       |
| `char`              | `char`       |
| `string`            | `char*`      |
| `void`              | `void`       |
| `int` (genérico)    | `int`        |
| `any`               | `void*`      |

---

## Declaração de Variáveis

O TAC não emite uma instrução dedicada para declaração de variáveis nomeadas.
O gerador C resolve os tipos consultando a tabela de símbolos do analisador semântico
(`Analyzer.GlobalScope()`), que é uma árvore retida acessível após a análise.

**Regra:** na primeira ocorrência de `ASSIGN value -> x`, se `x` é uma variável
nomeada ainda não declarada no C gerado, emite `<tipo_c> x = value;`.
Nas ocorrências seguintes (reatribuição), emite apenas `x = value;`.

Temporários (`t0`, `t1`, ...) seguem a mesma lógica, mas o tipo vem do mapa
`TempTypes()` em vez da tabela de símbolos.

```
TAC:    ASSIGN 42 -> x          (1ª ocorrência de x)
C:      int32_t x = 42;

TAC:    ASSIGN t0 -> x          (reatribuição)
C:      x = t0;

TAC:    ASSIGN 1 -> t0          (temporário: tipo vem de TempTypes)
C:      int t0 = 1;
```

---

## Aritmética e Lógica

| TAC                  | C                        |
|----------------------|--------------------------|
| `ADD a b -> r`       | `<decl?> r = a + b;`     |
| `SUB a b -> r`       | `<decl?> r = a - b;`     |
| `MUL a b -> r`       | `<decl?> r = a * b;`     |
| `DIV a b -> r`       | `<decl?> r = a / b;`     |
| `MOD a b -> r`       | `<decl?> r = a % b;`     |
| `NEG a -> r`         | `<decl?> r = -a;`        |
| `AND a b -> r`       | `<decl?> r = a && b;`    |
| `OR  a b -> r`       | `<decl?> r = a \|\| b;`  |
| `NOT a -> r`         | `<decl?> r = !a;`        |

`<decl?>` indica que a declaração do tipo é emitida na primeira ocorrência do
resultado (ex: `int32_t r = a + b;`).

---

## Comparação

| TAC                  | C                          |
|----------------------|----------------------------|
| `EQ  a b -> r`       | `<decl?> r = (a == b);`    |
| `NEQ a b -> r`       | `<decl?> r = (a != b);`    |
| `LT  a b -> r`       | `<decl?> r = (a < b);`     |
| `GT  a b -> r`       | `<decl?> r = (a > b);`     |
| `LEQ a b -> r`       | `<decl?> r = (a <= b);`    |
| `GEQ a b -> r`       | `<decl?> r = (a >= b);`    |

---

## Controle de Fluxo

| TAC                    | C                          |
|------------------------|----------------------------|
| `LABEL name`           | `name:;`  (coluna 0)       |
| `GOTO  name`           | `goto name;`               |
| `IF_FALSE cond -> L`   | `if (!cond) goto L;`       |
| `RETURN value`         | `return value;`            |
| `RETURN` (sem valor)   | `return;`  (ou `return 0;` dentro de `main`) |

**Estrutura gerada para `if/else`:**

```
TAC:
  <cond> -> t0
  IF_FALSE t0 -> L0
  <then-body>
  GOTO L1
  LABEL L0
  <else-body>
  LABEL L1

C:
  bool t0 = <cond>;
  if (!t0) goto L0;
  <then-body>
  goto L1;
L0:;
  <else-body>
L1:;
```

**Estrutura gerada para `while`:**

```
TAC:
  LABEL L0
  <cond> -> t0
  IF_FALSE t0 -> L1
  <body>
  GOTO L0
  LABEL L1

C:
  L0:;
  bool t0 = <cond>;
  if (!t0) goto L1;
  <body>
  goto L0;
L1:;
```

**`break` e `continue`** são traduzidos diretamente:

| TAC                  | C           |
|----------------------|-------------|
| `GOTO <end_label>`   | `break;`    |
| `GOTO <start_label>` | `continue;` |

> Nota: o TAC usa `GOTO` para break/continue. O gerador C emite `goto` e o
> compilador C elimina o goto desnecessário via otimização, ou podemos emitir
> `break`/`continue` diretamente se rastrearmos os labels de loop.

---

## Funções

O TAC armazena o tipo de retorno em `Arg2` de `BEGIN_FUNC` (adicionado na
Fase 1 das modificações ao TAC).

```
TAC:
  BEGIN_FUNC soma i32
  PARAM_DECL a i32
  PARAM_DECL b i32
  <corpo>
  END_FUNC soma

C:
  int32_t soma(int32_t a, int32_t b) {
      <corpo>
  }
```

**Chamada de função:**

```
TAC:
  PARAM 2
  PARAM 3
  CALL soma 2 -> t0

C:
  int32_t t0 = soma(2, 3);
```

Os `PARAM` são acumulados num buffer e descarregados na instrução `CALL`.

**`main` especial:** o tipo de retorno em C é sempre `int` e um `return 0;`
é emitido automaticamente antes do `}` final se não houver `RETURN` explícito.

---

## I/O

### `PRINT`

O formato `printf` depende do tipo do argumento. Resolução do tipo:

1. Se `arg1` é temporário `tN` → consulta `TempTypes()`
2. Se `arg1` é variável nomeada → consulta tabela de símbolos (mapa `varTypes`)
3. Se `arg1` é literal string (começa e termina com `"`) → `%s`
4. Se `arg1` é literal numérico com `.` → `%f`
5. Default → `%d`

| Tipo Minipar          | Especificador |
|-----------------------|---------------|
| `i8` / `i16` / `i32` / `int` | `%d`  |
| `i64`                 | `%ld`         |
| `u8` / `u16` / `u32` | `%u`          |
| `u64`                 | `%lu`         |
| `f32` / `float`       | `%f`          |
| `f64`                 | `%lf`         |
| `bool`                | `%d`          |
| `char`                | `%c`          |
| `string`              | `%s`          |

```
TAC:    PRINT x          (x: i32)
C:      printf("%d\n", x);

TAC:    PRINT "hello"
C:      printf("%s\n", "hello");
```

### `INPUT`

```
TAC:    INPUT prompt -> r
C:      printf("%s", prompt);
        scanf("%s", r);   /* tipo depende de r */
```

---

## Orientação a Objetos

O gerador C usa **flat functions com name mangling** para representar classes e métodos.

### Definição de classe

```
TAC:
  BEGIN_CLASS Ponto
  FIELD x i32
  FIELD y i32
  END_CLASS Ponto

C:
  typedef struct {
      int32_t x;
      int32_t y;
  } Ponto;
```

### Construtor

```
TAC:
  BEGIN_CTOR Ponto
  <corpo>
  END_CTOR Ponto

C:
  void Ponto_init(Ponto* self) {
      <corpo>
  }
```

### Método

```
TAC:
  BEGIN_METHOD move i32     (Arg2 = tipo de retorno)
  PARAM_DECL dx i32
  <corpo>
  END_METHOD move

C:
  int32_t Ponto_move(Ponto* self, int32_t dx) {
      <corpo>
  }
```

O gerador rastreia a classe atual em `currentClass` enquanto está entre
`BEGIN_CLASS` e `END_CLASS`.

### Instanciação e chamada de método

```
TAC:
  NEW_OBJ Ponto 0 -> p
C:
  Ponto p; Ponto_init(&p);

TAC:
  PARAM 5
  METHOD_CALL p.move 1 -> r
C:
  int32_t r = Ponto_move(&p, 5);
```

### Interface

```
TAC:    INTERFACE IMovivel
C:      /* interface IMovivel (sem equivalente direto em C) */
```

---

## Concorrência (stubs)

O gerador emite stubs comentados. A implementação completa via pthreads fica como TODO.

| TAC                  | C                                           |
|----------------------|---------------------------------------------|
| `BEGIN_SEQ`          | `/* BEGIN_SEQ */`                           |
| `END_SEQ`            | `/* END_SEQ */`                             |
| `BEGIN_PAR`          | `/* BEGIN_PAR — pthreads: TODO */`          |
| `END_PAR`            | `/* END_PAR */`                             |
| `CHAN_DECL t c n`    | `/* channel c (type: t) — não implementado */` |

O código dentro de blocos `par {}` é emitido sequencialmente, claramente marcado
com comentários.

---

## Arrays

| TAC                      | C                                       |
|--------------------------|-----------------------------------------|
| `ARRAY_GET arr idx -> r` | `<decl?> r = arr[idx];`                 |
| `ARRAY_LEN arr -> r`     | `<decl?> r = sizeof(arr)/sizeof(arr[0]);` |

---

## Exemplo Completo

**Fonte Minipar:**
```
func soma(a: i32, b: i32) -> i32 {
    return a + b
}

func main() {
    let x: i32 = soma(2, 3)
    print(x)
}
```

**TAC gerado:**
```
BEGIN_FUNC soma i32
PARAM_DECL a i32
PARAM_DECL b i32
ADD a b -> t0
RETURN t0
END_FUNC soma
BEGIN_FUNC main void
PARAM 2
PARAM 3
CALL soma 2 -> t1
ASSIGN t1 -> x
PRINT x
END_FUNC main
```

**C gerado:**
```c
#include <stdint.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int32_t soma(int32_t a, int32_t b) {
    int32_t t0 = a + b;
    return t0;
}

int main() {
    int32_t t1 = soma(2, 3);
    int32_t x = t1;
    printf("%d\n", x);
    return 0;
}
```
