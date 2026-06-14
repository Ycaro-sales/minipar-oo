package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"minipar/compiler"
	"minipar/lexer"
	"minipar/parser"
)

func main() {
	printAST := flag.Bool("ast", false, "print the AST as JSON instead of the generated TAC")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "uso: minipar [-ast] <arquivo.minipar>")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	src, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "erro ao ler arquivo: %v\n", err)
		os.Exit(1)
	}

	p := parser.NewParser(func(s string) lexer.Tokenizer {
		return lexer.New(s)
	})
	c := compiler.New(p)

	program, code, erros := c.Compile(string(src))

	if len(erros) > 0 {
		fmt.Fprintln(os.Stderr, "erros de compilação:")
		for _, e := range erros {
			fmt.Fprintf(os.Stderr, "  %s\n", e)
		}
		os.Exit(1)
	}

	if *printAST {
		out, err := json.MarshalIndent(program, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "erro ao serializar AST: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(out))
		return
	}

	fmt.Print(code)
}
