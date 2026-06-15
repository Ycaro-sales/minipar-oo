package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"minipar/compiler"
	"minipar/lexer"
	"minipar/parser"
)

// emitFlag holds the configuration for a single intermediate-artifact flag.
// enabled is true when the flag was supplied; path is empty when output should
// go to stdout, or the file path when the user wrote -flag=path.
type emitFlag struct {
	enabled bool
	path    string
}

// config captures the parsed command-line state.
type config struct {
	tokens  emitFlag
	ast     emitFlag
	tac     emitFlag
	c       emitFlag
	src     string // required: path to the .minipar source file
	binName string // optional: output binary name (defaults to src basename)
	help    bool
}

// usage prints the command-line help to w.
func usage(w *os.File) {
	fmt.Fprintln(w, `mcc - Minipar Compiler Collection

Usage:
  mcc [options] <path/to/code.minipar> [binary_name]

Options (without =path output goes to stdout; with =path output is saved to that file):
  -tokens[=<path>]   Dump token list
  -ast[=<path>]      Dump AST as JSON
  -tac[=<path>]      Dump Three-Address Code
  -c[=<path>]        Dump generated C source
  -h, --help         Show this help

Emit flags are additive. The native binary is always produced via gcc.
binary_name defaults to the source basename without .minipar.`)
}

// parseFlag splits a single '-name' or '-name=val' / '--name=val' argument
// into its name and optional value. Only exactly one or two leading dashes are
// accepted; anything else (---x, bare -, empty after strip) returns the
// original arg as the name so the caller surfaces a clean "unknown flag" error.
func parseFlag(arg string) (name, val string) {
	var s string
	switch {
	case strings.HasPrefix(arg, "--"):
		s = arg[2:]
	case strings.HasPrefix(arg, "-"):
		s = arg[1:]
	default:
		return arg, ""
	}
	// After stripping one prefix, a leading dash means malformed (e.g. ---x).
	if strings.HasPrefix(s, "-") || s == "" {
		return arg, ""
	}
	if idx := strings.IndexByte(s, '='); idx >= 0 {
		return s[:idx], s[idx+1:]
	}
	return s, ""
}

// parseArgs parses a slice of CLI arguments (typically os.Args[1:]) into a
// config. Flags (starting with '-') are parsed first; remaining positionals
// are the source file (required) and an optional binary name.
// Returns an error if the source file is missing (unless -h/--help was given).
func parseArgs(args []string) (config, error) {
	var cfg config
	var positionals []string

	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			positionals = append(positionals, arg)
			continue
		}

		name, val := parseFlag(arg)
		switch name {
		case "tokens":
			cfg.tokens = emitFlag{enabled: true, path: val}
		case "ast":
			cfg.ast = emitFlag{enabled: true, path: val}
		case "tac":
			cfg.tac = emitFlag{enabled: true, path: val}
		case "c":
			cfg.c = emitFlag{enabled: true, path: val}
		case "h", "help":
			cfg.help = true
		default:
			return config{}, fmt.Errorf("flag desconhecida: %s", arg)
		}
	}

	if cfg.help {
		return cfg, nil
	}

	if len(positionals) == 0 {
		return config{}, errors.New("arquivo fonte obrigatório: mcc [opções] <arquivo.minipar> [binary_name]")
	}
	if len(positionals) > 2 {
		return config{}, fmt.Errorf("argumentos inesperados: %v", positionals[2:])
	}
	cfg.src = positionals[0]
	if len(positionals) == 2 {
		cfg.binName = positionals[1]
	}
	return cfg, nil
}

// writeArtifact writes content to path. When path is empty it writes to stdout.
func writeArtifact(path, content string) error {
	if path == "" {
		_, err := fmt.Print(content)
		return err
	}
	return os.WriteFile(path, []byte(content), 0644)
}

// outputName returns the binary output name derived from the source path when
// no explicit name was provided. Falls back to "<basename>.out" when the source
// does not end in ".minipar" to avoid overwriting the source file.
func outputName(src, binName string) string {
	if binName != "" {
		return binName
	}
	base := filepath.Base(src)
	out := strings.TrimSuffix(base, ".minipar")
	if out == base {
		out = base + ".out"
	}
	return out
}

// absPath resolves p to an absolute path for collision detection; returns p
// unchanged when os.Getwd fails (best-effort).
func absPath(p string) string {
	if abs, err := filepath.Abs(p); err == nil {
		return abs
	}
	return p
}

// buildBinary compiles the C source into a native binary using gcc.
func buildBinary(cSource, outName string) error {
	cmd := exec.Command("gcc", "-x", "c", "-", "-o", outName)
	cmd.Stdin = strings.NewReader(cSource)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// reportErrors prints compilation errors to stderr in the established format.
func reportErrors(errs []string) {
	fmt.Fprintln(os.Stderr, "erros de compilação:")
	for _, e := range errs {
		fmt.Fprintf(os.Stderr, "  %s\n", e)
	}
}

// run is the core logic. It returns an exit code: 0 for success, 1 for error.
func run(cfg config) int {
	if cfg.help {
		usage(os.Stdout)
		return 0
	}

	src, err := os.ReadFile(cfg.src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "erro ao ler arquivo: %v\n", err)
		return 1
	}
	srcStr := string(src)

	p := parser.NewParser(func(s string) lexer.Tokenizer {
		return lexer.New(s)
	})
	c := compiler.New(p)

	// -tokens: does not require a successful parse.
	if cfg.tokens.enabled {
		if err := writeArtifact(cfg.tokens.path, c.Tokenize(srcStr)); err != nil {
			fmt.Fprintf(os.Stderr, "erro ao escrever tokens: %v\n", err)
			return 1
		}
	}

	// -ast: parse + semantic only (no TAC/C).
	if cfg.ast.enabled {
		prog, errs := c.AST(srcStr)
		if len(errs) > 0 {
			reportErrors(errs)
			return 1
		}
		out, err := json.MarshalIndent(prog, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "erro ao serializar AST: %v\n", err)
			return 1
		}
		if err := writeArtifact(cfg.ast.path, string(out)+"\n"); err != nil {
			fmt.Fprintf(os.Stderr, "erro ao escrever AST: %v\n", err)
			return 1
		}
	}

	// -tac: full pipeline up to TAC.
	if cfg.tac.enabled {
		_, code, errs := c.Compile(srcStr)
		if len(errs) > 0 {
			reportErrors(errs)
			return 1
		}
		if err := writeArtifact(cfg.tac.path, code); err != nil {
			fmt.Fprintf(os.Stderr, "erro ao escrever TAC: %v\n", err)
			return 1
		}
	}

	// Always generate C and build the native binary.
	// When -c is set, also emit the C source to the requested destination.
	_, cCode, errs := c.CompileToC(srcStr)
	if len(errs) > 0 {
		reportErrors(errs)
		return 1
	}

	if cfg.c.enabled {
		if err := writeArtifact(cfg.c.path, cCode); err != nil {
			fmt.Fprintf(os.Stderr, "erro ao escrever código C: %v\n", err)
			return 1
		}
	}

	outName := outputName(cfg.src, cfg.binName)
	if absSrc, absOut := absPath(cfg.src), absPath(outName); absSrc == absOut {
		fmt.Fprintf(os.Stderr, "erro: o nome do binário (%q) colidiria com o arquivo fonte; use um binary_name explícito\n", outName)
		return 1
	}
	if err := buildBinary(cCode, outName); err != nil {
		fmt.Fprintf(os.Stderr, "erro ao compilar com gcc: %v\n", err)
		return 1
	}

	return 0
}

func main() {
	cfg, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		usage(os.Stderr)
		os.Exit(1)
	}
	os.Exit(run(cfg))
}
