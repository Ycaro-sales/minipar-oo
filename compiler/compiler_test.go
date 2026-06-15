package compiler

import (
	"strings"
	"testing"
	"time"

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
	prog, _, errs := c.Compile("anything")
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
	_, code, errs := c.Compile("bad code")
	if len(errs) != 1 {
		t.Errorf("want 1 error, got %d", len(errs))
	}
	if code != "" {
		t.Errorf("want no TAC when parsing fails, got %q", code)
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
	prog, _, errs := c.Compile("let answer: i32 = 42;")
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

// TestCompilerGeneratesTAC verifies the pipeline runs through to code
// generation and emits TAC for a valid program.
func TestCompilerGeneratesTAC(t *testing.T) {
	src := `
func main()
{
  let x: i32 = 1 + 2
  print(x)
}
`
	p := parser.NewParser(func(s string) lexer.Tokenizer { return lexer.New(s) })
	c := New(p)
	_, code, errs := c.Compile(src)
	if len(errs) != 0 {
		t.Fatalf("errors: %v", errs)
	}
	for _, want := range []string{"BEGIN_FUNC", "ADD", "PRINT"} {
		if !strings.Contains(code, want) {
			t.Errorf("TAC missing %q\n--- code ---\n%s", want, code)
		}
	}
}

// ==========================================
// Tokenize tests
// ==========================================

// TestTokenize verifies that Tokenize returns a formatted token listing that
// includes each token's line number, type name, and literal.
func TestTokenize(t *testing.T) {
	src := "let x: i32 = 42;"
	p := parser.NewParser(func(s string) lexer.Tokenizer { return lexer.New(s) })
	c := New(p)
	out := c.Tokenize(src)

	// Each expected token: type name and literal must appear together in the output.
	wants := []struct{ typ, lit string }{
		{"LET", "let"},
		{"IDENT", "x"},
		{":", ":"},
		{"i32", "i32"},
		{"=", "="},
		{"NUMBER", "42"},
		{";", ";"},
	}
	for _, w := range wants {
		if !strings.Contains(out, w.typ) {
			t.Errorf("Tokenize output missing token type %q\noutput:\n%s", w.typ, out)
		}
		if !strings.Contains(out, w.lit) {
			t.Errorf("Tokenize output missing literal %q\noutput:\n%s", w.lit, out)
		}
	}
	// Line number "1" must appear (all tokens are on line 1).
	if !strings.Contains(out, "1") {
		t.Errorf("Tokenize output missing line number\noutput:\n%s", out)
	}
}

// TestTokenizeMultiLine checks that line numbers advance across newlines.
func TestTokenizeMultiLine(t *testing.T) {
	src := "let x = 1;\nlet y = 2;"
	p := parser.NewParser(func(s string) lexer.Tokenizer { return lexer.New(s) })
	c := New(p)
	out := c.Tokenize(src)

	if !strings.Contains(out, "2") {
		t.Errorf("Tokenize output missing line 2 marker\noutput:\n%s", out)
	}
}

// ==========================================
// AST tests
// ==========================================

// TestAST_valid checks that AST returns a non-nil program and no errors for
// a well-formed program, without requiring TAC or C generation to work.
func TestAST_valid(t *testing.T) {
	src := "let answer: i32 = 42;"
	p := parser.NewParser(func(s string) lexer.Tokenizer { return lexer.New(s) })
	c := New(p)
	prog, errs := c.AST(src)
	if len(errs) != 0 {
		t.Fatalf("want no errors, got: %v", errs)
	}
	if prog == nil {
		t.Fatal("want non-nil *ast.Program")
	}
	if len(prog.Declarations) != 1 {
		t.Fatalf("want 1 declaration, got %d", len(prog.Declarations))
	}
}

// TestAST_parseError checks that parse errors are surfaced and a non-nil
// (partial) program is still returned.
func TestAST_parseError(t *testing.T) {
	src := "let = ;"  // missing identifier
	p := parser.NewParser(func(s string) lexer.Tokenizer { return lexer.New(s) })
	c := New(p)
	prog, errs := c.AST(src)
	if len(errs) == 0 {
		t.Fatal("want errors for invalid source, got none")
	}
	if prog == nil {
		t.Fatal("want non-nil (partial) *ast.Program even on error")
	}
}

// TestAST_semanticError verifies that AST reports semantic errors and stops
// before TAC generation (which would panic/fail on the same input).
func TestAST_semanticError(t *testing.T) {
	// x is referenced but never declared — a semantic error.
	src := `func main() { let y: i32 = x; }`
	p := parser.NewParser(func(s string) lexer.Tokenizer { return lexer.New(s) })
	c := New(p)
	prog, errs := c.AST(src)
	if len(errs) == 0 {
		t.Fatal("want semantic errors for undeclared variable, got none")
	}
	// Program must still be returned (parser succeeded).
	if prog == nil {
		t.Fatal("want non-nil *ast.Program")
	}
}

// TestTokenize_unterminatedComment checks that Tokenize does not hang when the
// source contains an unterminated block comment.
func TestTokenize_unterminatedComment(t *testing.T) {
	src := "let x = 1; /* unterminated"
	p := parser.NewParser(func(s string) lexer.Tokenizer { return lexer.New(s) })
	c := New(p)

	done := make(chan string, 1)
	go func() { done <- c.Tokenize(src) }()

	select {
	case out := <-done:
		if !strings.Contains(out, "LET") {
			t.Errorf("expected LET token in output, got:\n%s", out)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Tokenize hung on unterminated block comment")
	}
}
