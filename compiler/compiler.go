package compiler

import (
	"minipar/ast"
	"minipar/parser"
	"minipar/semantic"
	"minipar/tac"
)

// Compiler orchestrates the front-end pipeline (parse → semantic analysis →
// TAC generation). Inject any Parser implementation.
type Compiler struct {
	parser parser.Parser
}

func New(p parser.Parser) *Compiler {
	return &Compiler{parser: p}
}

// Compile runs the full pipeline: it parses src, runs semantic analysis when
// parsing succeeds, and generates three-address code when analysis is clean.
// It returns the AST, the generated TAC (empty if any stage reported errors),
// and the collected error messages. A non-nil AST is always returned (partial
// tree on error).
func (c *Compiler) Compile(src string) (*ast.Program, string, []string) {
	program, errs := c.parser.ParseProgram(src)
	// Each later stage assumes the previous one produced a well-formed result,
	// so stop at the first stage that reports errors to avoid cascading noise.
	if len(errs) > 0 {
		return program, "", errs
	}
	if errs = semantic.New().Analyze(program); len(errs) > 0 {
		return program, "", errs
	}
	gen := tac.New()
	gen.Generate(program)
	return program, gen.Instructions(), errs
}
