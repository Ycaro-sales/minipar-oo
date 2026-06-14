package parser

import (
	"os"
	"path/filepath"
	"testing"

	"minipar/ast"
)

// Estes testes especificam a forma da AST produzida para um subconjunto
// representativo dos exemplos em tests/. Eles conferem tipos de nó, nomes,
// operadores, precedência/associatividade e aridade — não apenas a ausência de
// erros sintáticos (já coberta por TestExamplesParse).

// ==========================================
// Helpers
// ==========================================

func parseExample(t *testing.T, name string) *ast.Program {
	t.Helper()
	src, err := os.ReadFile(filepath.Join("..", "tests", name))
	if err != nil {
		t.Fatalf("não foi possível ler %s: %v", name, err)
	}
	prog, errs := ParseProgram(string(src))
	if len(errs) != 0 {
		t.Fatalf("erros sintáticos inesperados em %s: %v", name, errs)
	}
	return prog
}

// mustType faz type-assert para T, falhando o teste em caso de incompatibilidade.
func mustType[T any](t *testing.T, v any) T {
	t.Helper()
	got, ok := v.(T)
	if !ok {
		t.Fatalf("esperava %T, recebeu %T", *new(T), v)
	}
	return got
}

func binExpr(t *testing.T, e ast.Expression, op ast.Op) *ast.BinaryExpr {
	t.Helper()
	b := mustType[*ast.BinaryExpr](t, e)
	if b.Operator != op {
		t.Fatalf("operador: esperava %v, recebeu %v", op, b.Operator)
	}
	return b
}

func intLit(t *testing.T, e ast.Expression, v int64) {
	t.Helper()
	lit := mustType[*ast.IntLiteral](t, e)
	if lit.Value != v {
		t.Errorf("IntLiteral: esperava %d, recebeu %d", v, lit.Value)
	}
}

func ident(t *testing.T, e ast.Expression, name string) {
	t.Helper()
	id := mustType[*ast.Identifier](t, e)
	if id.Value != name {
		t.Errorf("Identifier: esperava %q, recebeu %q", name, id.Value)
	}
}

func funcCall(t *testing.T, e any, name string, argc int) *ast.FuncCall {
	t.Helper()
	c := mustType[*ast.FuncCall](t, e)
	if c.Name != name {
		t.Errorf("FuncCall: esperava nome %q, recebeu %q", name, c.Name)
	}
	if len(c.Args) != argc {
		t.Errorf("FuncCall %q: esperava %d args, recebeu %d", name, argc, len(c.Args))
	}
	return c
}

func funcDecl(t *testing.T, d ast.Declaration, name string, params int, ret string) *ast.FuncDecl {
	t.Helper()
	fn := mustType[*ast.FuncDecl](t, d)
	if fn.Name != name {
		t.Errorf("FuncDecl: esperava nome %q, recebeu %q", name, fn.Name)
	}
	if len(fn.Params) != params {
		t.Errorf("FuncDecl %q: esperava %d params, recebeu %d", name, params, len(fn.Params))
	}
	if fn.ReturnType != ret {
		t.Errorf("FuncDecl %q: esperava retorno %q, recebeu %q", name, ret, fn.ReturnType)
	}
	return fn
}

// ==========================================
// ex5: func count + while + main
// ==========================================

func TestNodesEx5(t *testing.T) {
	prog := parseExample(t, "ex5.minipar")
	if len(prog.Declarations) != 2 {
		t.Fatalf("declarações: esperava 2, recebeu %d", len(prog.Declarations))
	}

	count := funcDecl(t, prog.Declarations[0], "count", 1, "void")
	if count.Params[0].Name != "n" || count.Params[0].Type != "i32" {
		t.Errorf("param: esperava {n i32}, recebeu %+v", count.Params[0])
	}
	if len(count.Body.Statements) != 1 {
		t.Fatalf("corpo de count: esperava 1 stmt, recebeu %d", len(count.Body.Statements))
	}
	while := mustType[*ast.WhileStmt](t, count.Body.Statements[0])
	cond := binExpr(t, while.Condition, ast.OpGeq)
	ident(t, cond.Left, "n")
	intLit(t, cond.Right, 0)
	if len(while.Block.Statements) != 2 {
		t.Fatalf("bloco do while: esperava 2 stmts, recebeu %d", len(while.Block.Statements))
	}
	print := mustType[*ast.PrintStmt](t, while.Block.Statements[0])
	if len(print.Args) != 1 {
		t.Errorf("print: esperava 1 arg, recebeu %d", len(print.Args))
	}
	asg := mustType[*ast.Assignment](t, while.Block.Statements[1])
	if asg.Name != "n" {
		t.Errorf("assignment: esperava alvo n, recebeu %s", asg.Name)
	}
	sub := binExpr(t, asg.Value, ast.OpSub)
	ident(t, sub.Left, "n")
	intLit(t, sub.Right, 1)

	main := funcDecl(t, prog.Declarations[1], "main", 0, "")
	if len(main.Body.Statements) != 2 {
		t.Fatalf("corpo de main: esperava 2 stmts, recebeu %d", len(main.Body.Statements))
	}
	v := mustType[*ast.VarDecl](t, main.Body.Statements[0])
	if v.Name != "num" || v.Type != "i32" {
		t.Errorf("var: esperava num:i32, recebeu %s:%s", v.Name, v.Type)
	}
	intLit(t, v.Value, 10)
	call := funcCall(t, main.Body.Statements[1], "count", 1)
	ident(t, call.Args[0], "num")
}

// ==========================================
// ex8: while com if/continue
// ==========================================

func TestNodesEx8(t *testing.T) {
	prog := parseExample(t, "ex8.minipar")
	main := funcDecl(t, prog.Declarations[0], "main", 0, "")
	if len(main.Body.Statements) != 2 {
		t.Fatalf("corpo de main: esperava 2 stmts, recebeu %d", len(main.Body.Statements))
	}
	mustType[*ast.VarDecl](t, main.Body.Statements[0])
	while := mustType[*ast.WhileStmt](t, main.Body.Statements[1])
	wc := binExpr(t, while.Condition, ast.OpLt)
	ident(t, wc.Left, "id")
	intLit(t, wc.Right, 20)

	if len(while.Block.Statements) != 3 {
		t.Fatalf("bloco do while: esperava 3 stmts, recebeu %d", len(while.Block.Statements))
	}
	iff := mustType[*ast.IfStmt](t, while.Block.Statements[0])
	// cond: (id % 2) == 0
	eq := binExpr(t, iff.Condition, ast.OpEq)
	mod := binExpr(t, eq.Left, ast.OpMod)
	ident(t, mod.Left, "id")
	intLit(t, mod.Right, 2)
	intLit(t, eq.Right, 0)
	if iff.Else != nil {
		t.Error("if não deveria ter else")
	}
	if len(iff.Then.Statements) != 2 {
		t.Fatalf("then: esperava 2 stmts, recebeu %d", len(iff.Then.Statements))
	}
	mustType[*ast.Assignment](t, iff.Then.Statements[0])
	mustType[*ast.ContinueStmt](t, iff.Then.Statements[1])
	mustType[*ast.PrintStmt](t, while.Block.Statements[1])
	mustType[*ast.Assignment](t, while.Block.Statements[2])
}

// ==========================================
// ex9: fibonacci recursivo
// ==========================================

func TestNodesEx9(t *testing.T) {
	prog := parseExample(t, "ex9.minipar")
	if len(prog.Declarations) != 2 {
		t.Fatalf("declarações: esperava 2, recebeu %d", len(prog.Declarations))
	}
	fib := funcDecl(t, prog.Declarations[0], "fibonacci", 1, "i32")
	if fib.Params[0].Name != "n" || fib.Params[0].Type != "i32" {
		t.Errorf("param: esperava {n i32}, recebeu %+v", fib.Params[0])
	}
	if len(fib.Body.Statements) != 3 {
		t.Fatalf("corpo de fibonacci: esperava 3 stmts, recebeu %d", len(fib.Body.Statements))
	}
	mustType[*ast.IfStmt](t, fib.Body.Statements[0])
	mustType[*ast.IfStmt](t, fib.Body.Statements[1])
	ret := mustType[*ast.ReturnStmt](t, fib.Body.Statements[2])
	// fibonacci(n - 1) + fibonacci(n - 2)
	add := binExpr(t, ret.Value, ast.OpAdd)
	left := funcCall(t, add.Left, "fibonacci", 1)
	binExpr(t, left.Args[0], ast.OpSub)
	right := funcCall(t, add.Right, "fibonacci", 1)
	binExpr(t, right.Args[0], ast.OpSub)

	main := funcDecl(t, prog.Declarations[1], "main", 0, "")
	print := mustType[*ast.PrintStmt](t, main.Body.Statements[len(main.Body.Statements)-1])
	if len(print.Args) != 4 {
		t.Fatalf("print: esperava 4 args, recebeu %d", len(print.Args))
	}
	funcCall(t, print.Args[3], "fibonacci", 1)
}

// ==========================================
// fatorial_rec: if/else com retorno recursivo
// ==========================================

func TestNodesFatorialRec(t *testing.T) {
	prog := parseExample(t, "fatorial_rec.minipar")
	fat := funcDecl(t, prog.Declarations[0], "fatorial", 1, "i32")
	if len(fat.Body.Statements) != 1 {
		t.Fatalf("corpo de fatorial: esperava 1 stmt, recebeu %d", len(fat.Body.Statements))
	}
	iff := mustType[*ast.IfStmt](t, fat.Body.Statements[0])
	if iff.Else == nil {
		t.Fatal("if deveria ter else")
	}
	// cond: (n == 0) or (n == 1)
	or := binExpr(t, iff.Condition, ast.OpOr)
	binExpr(t, or.Left, ast.OpEq)
	binExpr(t, or.Right, ast.OpEq)
	// then: return 1
	thenRet := mustType[*ast.ReturnStmt](t, iff.Then.Statements[0])
	intLit(t, thenRet.Value, 1)
	// else: return n * fatorial(n - 1)
	elseRet := mustType[*ast.ReturnStmt](t, iff.Else.Statements[0])
	mul := binExpr(t, elseRet.Value, ast.OpMul)
	ident(t, mul.Left, "n")
	call := funcCall(t, mul.Right, "fatorial", 1)
	binExpr(t, call.Args[0], ast.OpSub)
}

// ==========================================
// quicksort in-place: código-alvo, ainda não parseável completamente.
// Pendências no parser que impedem a execução:
//   (1) index-assignment como lvalue: arr[i] = x  (gramática só aceita <id> = <expr>)
//   (2) tipo array em anotações: let arr: [i32]    (parser lê apenas o literal do token)
// O teste é pulado até essas features serem implementadas.
// ==========================================

func TestNodesQuicksort(t *testing.T) {
	t.Skip("quicksort.minipar é código-alvo: requer index-assignment (arr[i]=x) e " +
		"suporte a [tipo] em anotações de variáveis/parâmetros no parser")
}
