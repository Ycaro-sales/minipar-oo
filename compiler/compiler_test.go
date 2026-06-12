package compiler

import (
	"testing"

	"minipar/ast"
	"minipar/lexer"
	"minipar/parser"
)

// ==========================================
// Mock parser for DI testing
// ==========================================

type mockParser struct {
	program *ast.Program
	errs    []string
}

func (m *mockParser) ParseProgram(_ string) (*ast.Program, []string) {
	return m.program, m.errs
}

// ==========================================
// Compiler tests
// ==========================================

func TestCompilerSuccess(t *testing.T) {
	expected := &ast.Program{Declarations: []ast.Declaration{}}
	c := New(&mockParser{program: expected})
	prog, errs := c.Compile("anything")
	if len(errs) != 0 {
		t.Errorf("want no errors, got %v", errs)
	}
	if prog != expected {
		t.Error("returned program should be the mock's program")
	}
}

func TestCompilerPassesErrors(t *testing.T) {
	mock := &mockParser{
		program: &ast.Program{},
		errs:    []string{"linha 1: erro sintático"},
	}
	c := New(mock)
	_, errs := c.Compile("bad code")
	if len(errs) != 1 {
		t.Errorf("want 1 error, got %d", len(errs))
	}
}

func TestCompilerWithMockLexer(t *testing.T) {
	// Verify that NewParser + LexerFactory injection works end-to-end
	var capturedSrc string
	factory := func(s string) lexer.Tokenizer {
		capturedSrc = s
		return lexer.New(s)
	}
	p := parser.NewParser(factory)
	c := New(p)
	c.Compile("let x = 1;")
	if capturedSrc != "let x = 1;" {
		t.Errorf("factory received wrong src: %q", capturedSrc)
	}
}

func TestCompilerRealParser(t *testing.T) {
	p := parser.NewParser(func(s string) lexer.Tokenizer { return lexer.New(s) })
	c := New(p)
	prog, errs := c.Compile("let answer: i32 = 42;")
	if len(errs) != 0 {
		t.Fatalf("errors: %v", errs)
	}
	if len(prog.Declarations) != 1 {
		t.Fatalf("want 1 declaration, got %d", len(prog.Declarations))
	}
	v, ok := prog.Declarations[0].(*ast.VarDecl)
	if !ok {
		t.Fatalf("want *ast.VarDecl, got %T", prog.Declarations[0])
	}
	if v.Name != "answer" {
		t.Errorf("name: want answer, got %s", v.Name)
	}
}
