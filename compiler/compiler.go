package compiler

import (
	"fmt"
	"strings"

	"minipar/ast"
	"minipar/cgen"
	"minipar/lexer"
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

// Tokenize lexes src and returns a formatted listing of all tokens (excluding
// the final EOF), one per line in the format:
//
//	<line>  <type>  <literal>
func (c *Compiler) Tokenize(src string) string {
	l := lexer.New(src)
	var b strings.Builder
	for {
		tok := l.NextToken()
		if tok.Type == lexer.EOF {
			break
		}
		fmt.Fprintf(&b, "%d\t%s\t%s\n", tok.Line, tok.Type, tok.Literal)
		// The lexer never transitions out of the unterminated-comment error
		// state, so ILLEGAL would repeat forever without this guard.
		if tok.Type == lexer.ILLEGAL {
			break
		}
	}
	return b.String()
}

// AST parses src and runs semantic analysis, stopping before TAC generation.
// It returns the (possibly partial) AST and any accumulated error messages.
// Use this when only parse/semantic information is needed and TAC generation
// might not yet be supported for the given constructs.
func (c *Compiler) AST(src string) (*ast.Program, []string) {
	program, errs := c.parser.ParseProgram(src)
	if len(errs) > 0 {
		return program, errs
	}
	errs = semantic.New().Analyze(program)
	return program, errs
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

// CompileToC runs the full pipeline and returns generated C code instead of TAC.
func (c *Compiler) CompileToC(src string) (*ast.Program, string, []string) {
	program, errs := c.parser.ParseProgram(src)
	if len(errs) > 0 {
		return program, "", errs
	}
	an := semantic.New()
	if errs = an.Analyze(program); len(errs) > 0 {
		return program, "", errs
	}
	gen := tac.New()
	gen.Generate(program)
	cg := cgen.New(gen.RawInstructions(), gen.TempTypes(), an.GlobalScope())
	return program, cg.Generate(), nil
}
