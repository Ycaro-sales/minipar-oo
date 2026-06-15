# MCC - Minipar Compiler Collection

Minipar's compiler tools to compile minipar code.

## Usage

```
mcc [options] <path/to/code.minipar> [binary_name]
```

`binary_name` is optional. When omitted, the output binary takes the basename of
the source file without the `.minipar` extension (e.g. `prog.minipar` → `./prog`).

Requires `gcc` in `PATH` for binary generation.

## Options

All emit flags are **additive**: any combination may be used together. The native
binary is **always generated** in addition to any requested intermediate artifacts.

    -tokens[=<path/to/tokens>]
        Dumps the token list for the source file.
        Without a path, prints to stdout.
        With a path, saves to that file instead.

    -ast[=<path/to/json>]
        Dumps the Abstract Syntax Tree as JSON.
        Without a path, prints to stdout.
        With a path, saves to that file instead.

    -tac[=<path/to/file>]
        Dumps the Three-Address Code instructions.
        Without a path, prints to stdout.
        With a path, saves to that file instead.

    -c[=<path/to/c>]
        Dumps the generated C source code.
        Without a path, prints to stdout.
        With a path, saves to that file instead.

    -h, --help
        Prints this usage message.

## Examples

```bash
# Compile to a native binary (default)
mcc prog.minipar

# Compile to a named binary
mcc prog.minipar app

# Dump tokens to stdout, also build ./prog
mcc -tokens prog.minipar

# Save AST to a file and dump TAC to stdout, also build ./prog
mcc -ast=ast.json -tac prog.minipar

# Save generated C code to a file, also build ./prog
mcc -c=out.c prog.minipar

# All intermediate stages combined
mcc -tokens -ast=ast.json -tac=out.tac -c=out.c prog.minipar app
```
