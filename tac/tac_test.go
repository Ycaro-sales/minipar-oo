package tac

import (
	"strings"
	"testing"

	"minipar/ast"
	"minipar/lexer"
	"minipar/parser"
	"minipar/semantic"
)

// parseProgram builds an AST from source using the real lexer + parser,
// mirroring the wiring in main.go and the compiler package.
func parseProgram(t *testing.T, src string) *ast.Program {
	t.Helper()
	p := parser.NewParser(func(s string) lexer.Tokenizer { return lexer.New(s) })
	prog, errs := p.ParseProgram(src)
	if len(errs) != 0 {
		t.Fatalf("parse errors: %v", errs)
	}
	return prog
}

// analyze runs semantic analysis, which decorates expression nodes with their
// resolved types. The TAC generator reads those back, so tests that assert on
// type information must analyze before generating.
func analyze(t *testing.T, prog *ast.Program) {
	t.Helper()
	if errs := semantic.New().Analyze(prog); len(errs) != 0 {
		t.Fatalf("semantic errors: %v", errs)
	}
}

// TestGenerateRoundTrip exercises the full source → AST → TAC path. It guards
// the four bugs fixed alongside it: the missing constructor (nil-map panic), the
// i64 zero-value typo, the unreachable output, and the unguarded loop stack.
func TestGenerateRoundTrip(t *testing.T) {
	src := `
func main()
{
  let total: i64 = 0
  let i: i32 = 0
  while (i < 10) {
      total = total + i
      i = i + 1
      if (i == 5) {
          break
      }
  }
  print("done", total)
}
`
	prog := parseProgram(t, src)

	gen := New() // bug #1: must not panic on the nil SymbolTable
	gen.Generate(prog)

	out := gen.Instructions() // bug #3: output must be reachable
	if strings.TrimSpace(out) == "" {
		t.Fatal("Instructions() returned no output")
	}

	for _, want := range []string{"BEGIN_FUNC", "LABEL", "IF_FALSE", "GOTO", "PRINT"} {
		if !strings.Contains(out, want) {
			t.Errorf("expected TAC to contain %q\n--- output ---\n%s", want, out)
		}
	}
}

// TestTempTypes verifies the generator reads types off the decorated AST: each
// temporary holding an expression result is tagged with that expression's
// resolved type. This is the consumer that replaced TAC's own symbol table.
func TestTempTypes(t *testing.T) {
	src := `
func main()
{
  let x: i32 = 1 + 2
  let b: bool = x < 5
}
`
	prog := parseProgram(t, src)
	analyze(t, prog) // populate ResolvedType on expression nodes

	gen := New()
	gen.Generate(prog)

	// t0 holds (1 + 2) → int; t1 holds (x < 5) → bool.
	if got := gen.TempType("t0"); got != "int" {
		t.Errorf("TempType(t0): want %q, got %q", "int", got)
	}
	if got := gen.TempType("t1"); got != "bool" {
		t.Errorf("TempType(t1): want %q, got %q", "bool", got)
	}
}

// TestI64ZeroValue verifies an uninitialized i64 var gets a numeric zero, not an
// empty string (regression for the "i65" typo in getZeroForType).
func TestI64ZeroValue(t *testing.T) {
	if z := New().getZeroForType("i64"); z != "0" {
		t.Errorf("getZeroForType(i64): want %q, got %q", "0", z)
	}
}

// TestBreakOutsideLoopNoPanic ensures a break with an empty loop stack is a
// no-op rather than an index-out-of-range panic.
func TestBreakOutsideLoopNoPanic(t *testing.T) {
	gen := New()
	gen.genBreakStmt(nil)
	gen.genContinueStmt(nil)
	if out := gen.Instructions(); out != "" {
		t.Errorf("want no instructions for break/continue outside a loop, got %q", out)
	}
}
