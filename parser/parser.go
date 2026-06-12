package parser

import (
	"minipar/ast"
	"minipar/lexer"
)

// Parser is the interface the compiler depends on.
type Parser interface {
	ParseProgram(src string) (*ast.Program, []string)
}

// LexerFactory builds a Tokenizer from source — inject a mock in tests.
type LexerFactory func(src string) lexer.Tokenizer

// NewParser returns a Parser using the given factory to create lexers.
func NewParser(factory LexerFactory) Parser {
	return &miniparParser{factory: factory}
}

// ParseProgram is a convenience wrapper using the real lexer.
func ParseProgram(src string) (*ast.Program, []string) {
	return NewParser(func(s string) lexer.Tokenizer {
		return lexer.New(s)
	}).ParseProgram(src)
}
