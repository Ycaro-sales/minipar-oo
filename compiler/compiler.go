package compiler

import (
	"minipar/ast"
	"minipar/parser"
	"minipar/semantic"
)

// Compiler orchestrates the front-end pipeline (parse → semantic analysis).
// Inject any Parser implementation.
type Compiler struct {
	parser parser.Parser
}

func New(p parser.Parser) *Compiler {
	return &Compiler{parser: p}
}

// Compile parses src and, when parsing succeeds, runs semantic analysis.
// It returns the AST alongside any error messages (syntactic or semantic).
// A non-nil AST is always returned (partial tree on error).
func (c *Compiler) Compile(src string) (*ast.Program, []string) {
	program, errs := c.parser.ParseProgram(src)
	// Semantic analysis assumes a well-formed tree; skip it when the parser
	// already reported errors to avoid noise from a partial AST.
	if len(errs) == 0 {
		errs = semantic.New().Analyze(program)
	}
	return program, errs
}
