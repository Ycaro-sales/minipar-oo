package semantic

import (
	"strings"
	"testing"

	"minipar/ast"
	"minipar/lexer"
	"minipar/parser"
)

// analyze parses src with the real frontend and runs the analyzer, failing the
// test on parse errors so semantic assertions aren't masked by syntax issues.
func analyze(t *testing.T, src string) []string {
	t.Helper()
	prog, perrs := parser.ParseProgram(src)
	if len(perrs) > 0 {
		t.Fatalf("erros de parsing inesperados: %v", perrs)
	}
	return New().Analyze(prog)
}

// TestAnalyzerDecoratesTypes verifies the analyzer records each expression's
// resolved type on the AST node (consumed later by code generation).
func TestAnalyzerDecoratesTypes(t *testing.T) {
	prog, perrs := parser.ParseProgram(`let x = 1 + 2; let b = 1 < 2;`)
	if len(perrs) > 0 {
		t.Fatalf("erros de parsing inesperados: %v", perrs)
	}
	if errs := New().Analyze(prog); len(errs) > 0 {
		t.Fatalf("erros semânticos inesperados: %v", errs)
	}

	sum := prog.Declarations[0].(*ast.VarDecl).Value.(*ast.BinaryExpr)
	if got := sum.ResolvedType(); got != "int" {
		t.Errorf("tipo de '1 + 2': esperava \"int\", recebeu %q", got)
	}
	cmp := prog.Declarations[1].(*ast.VarDecl).Value.(*ast.BinaryExpr)
	if got := cmp.ResolvedType(); got != "bool" {
		t.Errorf("tipo de '1 < 2': esperava \"bool\", recebeu %q", got)
	}
}

func TestValidPrograms(t *testing.T) {
	cases := map[string]string{
		"var tipada":        `let x: i32 = 42;`,
		"var inferida":      `let x = 10; let y = x + 1;`,
		"concat string":     `let s: string = "a" + "b";`,
		"func e chamada":    `func soma(a: i32, b: i32) -> i32 { return a + b; } func main() { let r: i32 = soma(1, 2); }`,
		"if booleano":       `func f() { if (1 < 2) { let x = 1; } }`,
		"while e break":     `func f() { while (true) { break; } }`,
		"for sobre lista":   `func f() { for (i in [1, 2, 3]) { print(i); } }`,
		"interface ok":      `interface Falante { func fala(self); } class Cao implements (Falante) { func fala(self) { print("woof"); } }`,
		"campo no metodo":   `class Ponto { x: i32 func soma(self) -> i32 { return x + 1; } }`,
	}
	for name, src := range cases {
		t.Run(name, func(t *testing.T) {
			if errs := analyze(t, src); len(errs) != 0 {
				t.Errorf("esperava 0 erros, recebeu: %v", errs)
			}
		})
	}
}

func TestSemanticErrors(t *testing.T) {
	cases := []struct {
		name string
		src  string
		want string // substring expected in the first error
	}{
		{"var nao declarada", `func f() { x = 1; }`, "não declarada"},
		{"identificador livre", `let y = z;`, "não declarado"},
		{"tipo incompativel", `let x: string = 1;`, "não é possível atribuir"},
		{"redeclaracao global", `func f() {} func f() {}`, "já declarado"},
		{"tipo desconhecido", `let x: Foo = 1;`, "tipo desconhecido"},
		{"if nao booleano", `func f() { if (1 + 1) { let a = 1; } }`, "deve ser 'bool'"},
		{"break fora de laco", `func f() { break; }`, "fora de um laço"},
		{"retorno incompativel", `func f() -> i32 { return "x"; }`, "incompatível"},
		{"aritmetica em string", `func f() { let a = "x" - 1; }`, "requer números"},
		{"chamada de nao funcao", `func f() { let v = 3; v(1); }`, "não é uma função"},
		{"aridade errada", `func g(a: i32) {} func f() { g(1, 2); }`, "argumento(s)"},
		{"interface nao implementada", `interface I { func m(self) -> i32; } class C implements (I) { }`, "não implementa"},
		{"self fora de metodo", `func f() { return Self; }`, "fora de um método"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			errs := analyze(t, tc.src)
			if len(errs) == 0 {
				t.Fatalf("esperava um erro contendo %q, não houve erros", tc.want)
			}
			if !containsAny(errs, tc.want) {
				t.Errorf("esperava erro contendo %q, recebeu: %v", tc.want, errs)
			}
		})
	}
}

func TestNilProgram(t *testing.T) {
	if errs := New().Analyze(nil); len(errs) != 0 {
		t.Errorf("programa nil deveria gerar 0 erros, recebeu: %v", errs)
	}
}

func containsAny(errs []string, sub string) bool {
	for _, e := range errs {
		if strings.Contains(e, sub) {
			return true
		}
	}
	return false
}

// ensure the real lexer factory is referenced (keeps imports honest if the
// convenience wrapper changes).
var _ = func() lexer.Tokenizer { return lexer.New("") }
