package cgen_test

import (
	"strings"
	"testing"

	"minipar/cgen"
	"minipar/lexer"
	"minipar/parser"
	"minipar/semantic"
	"minipar/tac"
)

// compileToC runs the full pipeline (parse → semantic → TAC → C) and returns
// the generated C source as a string.
func compileToC(t *testing.T, src string) string {
	t.Helper()
	p := parser.NewParser(func(s string) lexer.Tokenizer { return lexer.New(s) })
	prog, parseErrs := p.ParseProgram(src)
	if len(parseErrs) != 0 {
		t.Fatalf("parse errors: %v", parseErrs)
	}
	an := semantic.New()
	if errs := an.Analyze(prog); len(errs) != 0 {
		t.Fatalf("semantic errors: %v", errs)
	}
	gen := tac.New()
	gen.Generate(prog)
	cg := cgen.New(gen.RawInstructions(), gen.TempTypes(), an.GlobalScope())
	return cg.Generate()
}

// assertContains fails the test if any of the wanted substrings are absent from out.
func assertContains(t *testing.T, out string, wants ...string) {
	t.Helper()
	for _, want := range wants {
		if !strings.Contains(out, want) {
			t.Errorf("expected output to contain %q\n--- output ---\n%s", want, out)
		}
	}
}

// assertNotContains fails the test if any of the unwanted substrings appear in out.
func assertNotContains(t *testing.T, out string, unwanted ...string) {
	t.Helper()
	for _, s := range unwanted {
		if strings.Contains(out, s) {
			t.Errorf("expected output NOT to contain %q\n--- output ---\n%s", s, out)
		}
	}
}

// ==========================================
// PREÂMBULO E ESTRUTURA BÁSICA
// ==========================================

func TestIncludes(t *testing.T) {
	out := compileToC(t, `func main() {}`)
	assertContains(t, out,
		"#include <stdint.h>",
		"#include <stdbool.h>",
		"#include <stdio.h>",
		"#include <stdlib.h>",
		"#include <string.h>",
	)
}

func TestEmptyMain(t *testing.T) {
	out := compileToC(t, `func main() {}`)
	assertContains(t, out, "int main()", "return 0;")
}

// ==========================================
// DECLARAÇÃO DE VARIÁVEIS
// ==========================================

func TestVarDeclInt(t *testing.T) {
	out := compileToC(t, `func main() { let x: i32 = 42 }`)
	assertContains(t, out, "int32_t x = 42;")
}

func TestVarDeclI64(t *testing.T) {
	out := compileToC(t, `func main() { let n: i64 = 0 }`)
	assertContains(t, out, "int64_t n = 0;")
}

func TestVarDeclBool(t *testing.T) {
	out := compileToC(t, `func main() { let b: bool = true }`)
	assertContains(t, out, "bool b = true;")
}

func TestVarDeclString(t *testing.T) {
	out := compileToC(t, `func main() { let s: string = "hi" }`)
	assertContains(t, out, `char* s = "hi";`)
}

func TestVarDeclFloat(t *testing.T) {
	out := compileToC(t, `func main() { let f: f32 = 3.14 }`)
	// O TAC formata floats com 6 casas decimais (formato padrão do Go)
	assertContains(t, out, "float f = 3.140000;")
}

// Reatribuição não deve redeclarar a variável.
func TestAssignmentNoDuplicateDecl(t *testing.T) {
	out := compileToC(t, `
func main() {
    let x: i32 = 1
    x = 2
}`)
	// "int32_t x" deve aparecer apenas uma vez
	count := strings.Count(out, "int32_t x")
	if count != 1 {
		t.Errorf("want exactly 1 declaration of 'int32_t x', got %d\n--- output ---\n%s", count, out)
	}
	assertContains(t, out, "x = 2;")
}

func TestGlobalAssignInsideFuncNoRedecl(t *testing.T) {
	out := compileToC(t, `let a: i32 = 10
func main() {
    a = a + 1
    print(a)
}`)
	// declaração global deve existir exatamente uma vez
	count := strings.Count(out, "int32_t a")
	if count != 1 {
		t.Errorf("want exactly 1 declaration of 'int32_t a', got %d\n--- output ---\n%s", count, out)
	}
	// dentro de main a atribuição não deve ter prefixo de tipo
	assertNotContains(t, out, "int32_t a = t")
}

// ==========================================
// ARITMÉTICA
// ==========================================

func TestArithmeticAdd(t *testing.T) {
	out := compileToC(t, `func main() { let x: i32 = 1 + 2 }`)
	assertContains(t, out, "1 + 2")
	assertContains(t, out, "int32_t x")
}

func TestArithmeticSub(t *testing.T) {
	out := compileToC(t, `func main() { let x: i32 = 10 - 3 }`)
	assertContains(t, out, "10 - 3")
}

func TestArithmeticMul(t *testing.T) {
	out := compileToC(t, `func main() { let x: i32 = 4 * 5 }`)
	assertContains(t, out, "4 * 5")
}

func TestArithmeticDiv(t *testing.T) {
	out := compileToC(t, `func main() { let x: i32 = 10 / 2 }`)
	assertContains(t, out, "10 / 2")
}

func TestArithmeticMod(t *testing.T) {
	out := compileToC(t, `func main() { let x: i32 = 7 % 3 }`)
	assertContains(t, out, "7 % 3")
}

func TestNegation(t *testing.T) {
	out := compileToC(t, `func main() { let x: i32 = 5
                                        let n: i32 = -x }`)
	assertContains(t, out, "= -x")
}

// ==========================================
// LÓGICA E COMPARAÇÃO
// ==========================================

func TestComparisonLt(t *testing.T) {
	out := compileToC(t, `func main() { let x: i32 = 3
                                        let b: bool = x < 5 }`)
	assertContains(t, out, "< 5")
}

func TestComparisonGt(t *testing.T) {
	out := compileToC(t, `func main() { let x: i32 = 3
                                        let b: bool = x > 1 }`)
	assertContains(t, out, "> 1")
}

func TestComparisonEq(t *testing.T) {
	out := compileToC(t, `func main() { let x: i32 = 3
                                        let b: bool = x == 3 }`)
	assertContains(t, out, "== 3")
}

func TestLogicalAnd(t *testing.T) {
	out := compileToC(t, `func main() {
    let a: bool = true
    let b: bool = false
    let r: bool = a and b
}`)
	assertContains(t, out, "&&")
}

func TestLogicalOr(t *testing.T) {
	out := compileToC(t, `func main() {
    let a: bool = true
    let b: bool = false
    let r: bool = a or b
}`)
	assertContains(t, out, "||")
}

func TestLogicalNot(t *testing.T) {
	out := compileToC(t, `func main() {
    let b: bool = true
    let r: bool = !b
}`)
	assertContains(t, out, "= !b")
}

// ==========================================
// CONTROLE DE FLUXO
// ==========================================

func TestIfStmt(t *testing.T) {
	out := compileToC(t, `
func main() {
    let x: i32 = 5
    if (x > 3) { print("maior") }
}`)
	assertContains(t, out, "if (")
}

func TestIfElseStmt(t *testing.T) {
	out := compileToC(t, `
func main() {
    let x: i32 = 5
    if (x > 3) { print("maior") } else { print("menor") }
}`)
	assertContains(t, out, "if (", `"maior"`, `"menor"`)
}

func TestWhileLoop(t *testing.T) {
	out := compileToC(t, `
func main() {
    let i: i32 = 0
    while (i < 10) { i = i + 1 }
}`)
	// O TAC gera while como LABEL + IF_FALSE + GOTO; esperamos algum desses padrões
	hasWhile := strings.Contains(out, "while (") || strings.Contains(out, "goto ")
	if !hasWhile {
		t.Errorf("expected while loop pattern (while or goto)\n--- output ---\n%s", out)
	}
}

func TestBreak(t *testing.T) {
	out := compileToC(t, `
func main() {
    let i: i32 = 0
    while (i < 10) {
        i = i + 1
        if (i == 5) { break }
    }
}`)
	hasBreak := strings.Contains(out, "break;") || strings.Contains(out, "goto L")
	if !hasBreak {
		t.Errorf("expected break pattern\n--- output ---\n%s", out)
	}
}

// ==========================================
// FUNÇÕES
// ==========================================

func TestFuncDecl(t *testing.T) {
	out := compileToC(t, `
func soma(a: i32, b: i32) -> i32 {
    return a + b
}
func main() {}`)
	assertContains(t, out, "int32_t soma(", "int32_t a", "int32_t b")
}

func TestFuncReturnType(t *testing.T) {
	out := compileToC(t, `
func dobro(x: i32) -> i32 { return x + x }
func main() {}`)
	assertContains(t, out, "int32_t dobro(")
}

func TestFuncReturnVoid(t *testing.T) {
	out := compileToC(t, `
func imprime(x: i32) { print(x) }
func main() {}`)
	assertContains(t, out, "void imprime(")
}

func TestFuncReturnValue(t *testing.T) {
	out := compileToC(t, `
func f() -> i32 { return 42 }
func main() {}`)
	assertContains(t, out, "return 42")
}

func TestFuncCall(t *testing.T) {
	out := compileToC(t, `
func soma(a: i32, b: i32) -> i32 { return a + b }
func main() {
    let r: i32 = soma(2, 3)
}`)
	assertContains(t, out, "soma(2, 3)")
}

func TestMainReturnZero(t *testing.T) {
	out := compileToC(t, `func main() {}`)
	assertContains(t, out, "return 0;")
}

// ==========================================
// I/O — PRINT
// ==========================================

func TestPrintStringLiteral(t *testing.T) {
	out := compileToC(t, `func main() { print("hello") }`)
	assertContains(t, out, `printf("%s\n", "hello")`)
}

func TestPrintInt(t *testing.T) {
	out := compileToC(t, `func main() {
    let x: i32 = 5
    print(x)
}`)
	assertContains(t, out, `printf("%d\n", x)`)
}

func TestPrintI64(t *testing.T) {
	out := compileToC(t, `func main() {
    let n: i64 = 100
    print(n)
}`)
	assertContains(t, out, `printf("%ld\n", n)`)
}

func TestPrintBool(t *testing.T) {
	out := compileToC(t, `func main() {
    let b: bool = true
    print(b)
}`)
	assertContains(t, out, `printf("%d\n", b)`)
}

// ==========================================
// CLASSES (OOP)
// ==========================================

func TestClassStruct(t *testing.T) {
	out := compileToC(t, `
class Ponto {
    x: i32
    y: i32
}
func main() {}`)
	assertContains(t, out, "typedef struct {", "int32_t x;", "int32_t y;", "} Ponto;")
}

func TestClassConstructor(t *testing.T) {
	out := compileToC(t, `
class Ponto {
    x: i32
    Ponto { print("init") }
}
func main() {}`)
	assertContains(t, out, "Ponto_init(", "Ponto* self")
}

func TestClassMethod(t *testing.T) {
	out := compileToC(t, `
class Calc {
    func dobro(x: i32) -> i32 { return x + x }
}
func main() {}`)
	assertContains(t, out, "Calc_dobro(", "Calc* self")
}

// ==========================================
// CONCORRÊNCIA (stubs)
// ==========================================

func TestParBlock(t *testing.T) {
	out := compileToC(t, `
func main() {
    par {
        print("a")
        print("b")
    }
}`)
	assertContains(t, out, "BEGIN_PAR")
}

func TestSeqBlock(t *testing.T) {
	out := compileToC(t, `
func main() {
    seq {
        print("a")
        print("b")
    }
}`)
	assertContains(t, out, "BEGIN_SEQ")
}
