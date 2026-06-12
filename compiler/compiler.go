package compiler

import (
	"minipar/ast"
	"minipar/parser"
)

// Compiler orchestrates parsing. Inject any Parser implementation.
type Compiler struct {
	parser parser.Parser
}

func New(p parser.Parser) *Compiler {
	return &Compiler{parser: p}
}

// Compile parses src and returns the AST alongside any error messages.
// A non-nil AST is always returned (partial tree on error).
func (c *Compiler) Compile(src string) (*ast.Program, []string) {
	return c.parser.ParseProgram(src)
}
