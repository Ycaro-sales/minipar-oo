package main

import (
	"encoding/json"
	"fmt"
	"os"

	"minipar/compiler"
	"minipar/lexer"
	"minipar/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "uso: minipar <arquivo.minipar>")
		os.Exit(1)
	}

	src, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "erro ao ler arquivo: %v\n", err)
		os.Exit(1)
	}

	p := parser.NewParser(func(s string) lexer.Tokenizer {
		return lexer.New(s)
	})
	c := compiler.New(p)

	program, erros := c.Compile(string(src))

	if len(erros) > 0 {
		fmt.Fprintln(os.Stderr, "erros sintáticos:")
		for _, e := range erros {
			fmt.Fprintf(os.Stderr, "  %s\n", e)
		}
		os.Exit(1)
	}

	out, err := json.MarshalIndent(program, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "erro ao serializar AST: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}
