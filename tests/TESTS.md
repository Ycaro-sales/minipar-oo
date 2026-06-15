# Arquivos de teste — `tests/`

## Exemplos gerais (`ex*.minipar`)

| Arquivo | O que testa |
|---------|-------------|
| `ex1.minipar` | Variáveis globais (`a`, `b`), função com `while` + `break` + `if`, reatribuição de global dentro de função, `print` |
| `ex2.minipar` | `s_channel`, função retornando `string` via `to_string`, `while` com `break`, literais `bool`, inteiro negativo (`-1`) |
| `ex3.minipar` | `while` com `break` em `if`, chamada de função com expressão complexa como argumento `(3+4)*6`, variável `string` |
| `ex4.minipar` | Função literal (lambda) atribuída a variável (`let fat = func...`), fatorial iterativo, fibonacci iterativo, bloco `par{}` rodando as duas em paralelo |
| `ex5.minipar` | Contagem regressiva simples: `while` com decremento, função `void` |
| `ex7.minipar` | Indexação de string (`message[index]`), builtins `isalpha`, `isnum`, `to_number`, `to_string`, `switch`, `continue` — mini-calculadora de expressões |
| `ex8.minipar` | `continue` com operador `%` (módulo): imprime somente ímpares de 1 a 19 |
| `ex9.minipar` | Fibonacci recursivo com dois casos base, `if` sem `else`, recursão dupla |

> Não existe `ex6.minipar`.

---

## Algoritmos (`fatorial_rec.minipar`, `quicksort.minipar`)

| Arquivo | O que testa |
|---------|-------------|
| `fatorial_rec.minipar` | Fatorial recursivo com `if/else`, recursão simples, `return` em branches |
| `quicksort.minipar` | Array `[i32]`, quicksort in-place recursivo, funções `swap`/`partition`/`quicksort`, leitura e escrita de elementos de array (`arr[i]`) |

---

## Controle de fluxo (`test_break_continue.minipar`, `test_seq_par.minipar`)

| Arquivo | O que testa |
|---------|-------------|
| `test_break_continue.minipar` | `break` (sai do loop em `i == 5`) e `continue` (pula `j == 3` e `j == 7`) em `while`, com saída impressa para cada caso |
| `test_seq_par.minipar` | Blocos `seq{}` e `par{}`, variável global `counter` incrementada sequencialmente |

---

## Canais / comunicação (`calc_*`, `test_echo_*`)

Estes arquivos vêm em pares servidor/cliente e testam a feature de canais (`s_channel` / `c_channel`).

| Arquivo | Papel | O que testa |
|---------|-------|-------------|
| `calc_server.minipar` | Servidor | `s_channel`, `switch` com as quatro operações aritméticas, parâmetro `f64`, `return` dentro de `switch` |
| `calc_client.minipar` | Cliente | `c_channel`, `send` com múltiplos argumentos, `close` |
| `test_echo_server.minipar` | Servidor | `s_channel` com função que devolve o argumento recebido (`string → string`) |
| `test_echo_client.minipar` | Cliente | `c_channel`, `send` com `string`, `close` |

Para rodar os pares: inicie o servidor em um terminal primeiro, depois o cliente em outro.
