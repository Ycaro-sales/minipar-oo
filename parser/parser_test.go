package parser

import (
	"testing"

	"minipar/ast"
)

func parseOne(t *testing.T, src string) ast.Declaration {
	t.Helper()
	prog, errs := ParseProgram(src)
	if len(errs) > 0 {
		t.Fatalf("parse errors: %v", errs)
	}
	if len(prog.Declarations) == 0 {
		t.Fatal("no declarations parsed")
	}
	return prog.Declarations[0]
}

func parseStmt(t *testing.T, src string) ast.Statement {
	t.Helper()
	// wrap in a func body so the parser sees a statement context
	wrapped := "func _test() { " + src + " }"
	decl := parseOne(t, wrapped)
	fn, ok := decl.(*ast.FuncDecl)
	if !ok {
		t.Fatalf("expected FuncDecl, got %T", decl)
	}
	if len(fn.Body.Statements) == 0 {
		t.Fatal("no statements in body")
	}
	return fn.Body.Statements[0]
}

// ==========================================
// VarDecl
// ==========================================

func TestParseVarDeclWithType(t *testing.T) {
	decl := parseOne(t, "let x: i32 = 5;")
	v, ok := decl.(*ast.VarDecl)
	if !ok {
		t.Fatalf("want *ast.VarDecl, got %T", decl)
	}
	if v.Name != "x" {
		t.Errorf("name: want x, got %s", v.Name)
	}
	if v.Type != "i32" {
		t.Errorf("type: want i32, got %s", v.Type)
	}
	lit, ok := v.Value.(*ast.IntLiteral)
	if !ok {
		t.Fatalf("value: want *ast.IntLiteral, got %T", v.Value)
	}
	if lit.Value != 5 {
		t.Errorf("value: want 5, got %d", lit.Value)
	}
}

func TestParseVarDeclNoType(t *testing.T) {
	decl := parseOne(t, "let y = true;")
	v := decl.(*ast.VarDecl)
	if v.Type != "" {
		t.Errorf("type: want empty, got %s", v.Type)
	}
	b := v.Value.(*ast.BooleanLiteral)
	if !b.Value {
		t.Error("value: want true")
	}
}

// ==========================================
// FuncDecl
// ==========================================

func TestParseFuncDeclNoParams(t *testing.T) {
	decl := parseOne(t, "func main() { }")
	f := decl.(*ast.FuncDecl)
	if f.Name != "main" {
		t.Errorf("name: want main, got %s", f.Name)
	}
	if len(f.Params) != 0 {
		t.Errorf("params: want 0, got %d", len(f.Params))
	}
}

func TestParseFuncDeclWithParams(t *testing.T) {
	decl := parseOne(t, "func add(a: i32, b: i32) -> i32 { }")
	f := decl.(*ast.FuncDecl)
	if len(f.Params) != 2 {
		t.Fatalf("params: want 2, got %d", len(f.Params))
	}
	if f.Params[0].Name != "a" || f.Params[0].Type != "i32" {
		t.Errorf("param[0]: want a:i32, got %s:%s", f.Params[0].Name, f.Params[0].Type)
	}
	if f.ReturnType != "i32" {
		t.Errorf("return: want i32, got %s", f.ReturnType)
	}
}

// ==========================================
// ClassDecl
// ==========================================

func TestParseClassDeclBasic(t *testing.T) {
	src := `class Dog { name: string }`
	decl := parseOne(t, src)
	c := decl.(*ast.ClassDecl)
	if c.Name != "Dog" {
		t.Errorf("name: want Dog, got %s", c.Name)
	}
	if len(c.Members) != 1 {
		t.Fatalf("members: want 1, got %d", len(c.Members))
	}
}

func TestParseClassDeclImplements(t *testing.T) {
	src := `class Duck implements (Speaker) { }`
	decl := parseOne(t, src)
	c := decl.(*ast.ClassDecl)
	if len(c.Implements) != 1 || c.Implements[0] != "Speaker" {
		t.Errorf("implements: want [Speaker], got %v", c.Implements)
	}
}

func TestParseClassDeclMultipleInterfaces(t *testing.T) {
	src := `class Foo implements (A, B) { }`
	c := parseOne(t, src).(*ast.ClassDecl)
	if len(c.Implements) != 2 {
		t.Errorf("implements: want 2, got %d", len(c.Implements))
	}
}

// ==========================================
// InterfaceDecl
// ==========================================

func TestParseInterfaceDecl(t *testing.T) {
	src := `interface Speaker { func speak(self) -> void }`
	decl := parseOne(t, src)
	iface := decl.(*ast.InterfaceDecl)
	if iface.Name != "Speaker" {
		t.Errorf("name: want Speaker, got %s", iface.Name)
	}
	if len(iface.Methods) != 1 {
		t.Fatalf("methods: want 1, got %d", len(iface.Methods))
	}
	if iface.Methods[0].Name != "speak" {
		t.Errorf("method name: want speak, got %s", iface.Methods[0].Name)
	}
}

// ==========================================
// Statements
// ==========================================

func TestParseIfStmtNoElse(t *testing.T) {
	stmt := parseStmt(t, "if (x) { }")
	s := stmt.(*ast.IfStmt)
	if s.Else != nil {
		t.Error("else: want nil")
	}
}

func TestParseIfStmtWithElse(t *testing.T) {
	stmt := parseStmt(t, "if (x) { } else { }")
	s := stmt.(*ast.IfStmt)
	if s.Else == nil {
		t.Error("else: want non-nil")
	}
}

func TestParseWhileStmt(t *testing.T) {
	stmt := parseStmt(t, "while (i < 10) { }")
	s := stmt.(*ast.WhileStmt)
	if s.Condition == nil {
		t.Error("condition: want non-nil")
	}
}

func TestParseDoStmt(t *testing.T) {
	stmt := parseStmt(t, "do { } while (x);")
	s := stmt.(*ast.DoStmt)
	if s.Block == nil || s.Condition == nil {
		t.Error("do stmt: block and condition must be non-nil")
	}
}

func TestParseForStmt(t *testing.T) {
	stmt := parseStmt(t, "for (i in arr) { }")
	s := stmt.(*ast.ForStmt)
	if s.Var != "i" {
		t.Errorf("var: want i, got %s", s.Var)
	}
}

func TestParseSwitchStmt(t *testing.T) {
	stmt := parseStmt(t, "switch (x) { 1 => { } 2 => { } }")
	s := stmt.(*ast.SwitchStmt)
	if len(s.Cases) != 2 {
		t.Errorf("cases: want 2, got %d", len(s.Cases))
	}
}

func TestParseSeqStmt(t *testing.T) {
	stmt := parseStmt(t, "seq { }")
	if _, ok := stmt.(*ast.SeqStmt); !ok {
		t.Fatalf("want *ast.SeqStmt, got %T", stmt)
	}
}

func TestParseParStmt(t *testing.T) {
	stmt := parseStmt(t, "par { }")
	if _, ok := stmt.(*ast.ParStmt); !ok {
		t.Fatalf("want *ast.ParStmt, got %T", stmt)
	}
}

func TestParseReturnStmtWithValue(t *testing.T) {
	stmt := parseStmt(t, "return 42;")
	s := stmt.(*ast.ReturnStmt)
	if s.Value == nil {
		t.Error("value: want non-nil")
	}
}

func TestParseReturnStmtVoid(t *testing.T) {
	stmt := parseStmt(t, "return;")
	s := stmt.(*ast.ReturnStmt)
	if s.Value != nil {
		t.Error("value: want nil for void return")
	}
}

func TestParseBreakStmt(t *testing.T) {
	stmt := parseStmt(t, "break;")
	if _, ok := stmt.(*ast.BreakStmt); !ok {
		t.Fatalf("want *ast.BreakStmt, got %T", stmt)
	}
}

func TestParseContinueStmt(t *testing.T) {
	stmt := parseStmt(t, "continue;")
	if _, ok := stmt.(*ast.ContinueStmt); !ok {
		t.Fatalf("want *ast.ContinueStmt, got %T", stmt)
	}
}

func TestParsePassStmt(t *testing.T) {
	stmt := parseStmt(t, "pass;")
	if _, ok := stmt.(*ast.PassStmt); !ok {
		t.Fatalf("want *ast.PassStmt, got %T", stmt)
	}
}

func TestParseGotoStmt(t *testing.T) {
	stmt := parseStmt(t, "goto myLabel;")
	s := stmt.(*ast.GotoStmt)
	if s.Label != "myLabel" {
		t.Errorf("label: want myLabel, got %s", s.Label)
	}
}

func TestParsePrintStmt(t *testing.T) {
	stmt := parseStmt(t, `print("hello");`)
	s := stmt.(*ast.PrintStmt)
	if len(s.Args) != 1 {
		t.Errorf("args: want 1, got %d", len(s.Args))
	}
}

func TestParseAssignment(t *testing.T) {
	stmt := parseStmt(t, "x = 10;")
	s := stmt.(*ast.Assignment)
	if s.Name != "x" {
		t.Errorf("name: want x, got %s", s.Name)
	}
}

// ==========================================
// Expressions
// ==========================================

func parseExpr(t *testing.T, src string) ast.Expression {
	t.Helper()
	stmt := parseStmt(t, "return "+src+";")
	return stmt.(*ast.ReturnStmt).Value
}

func TestParseBinaryExprPrecedence(t *testing.T) {
	// 1 + 2 * 3  →  Add(1, Mul(2, 3))
	expr := parseExpr(t, "1 + 2 * 3")
	add, ok := expr.(*ast.BinaryExpr)
	if !ok || add.Operator != ast.OpAdd {
		t.Fatalf("want Add at root, got %T op=%v", expr, add)
	}
	if _, ok := add.Left.(*ast.IntLiteral); !ok {
		t.Error("left of Add should be IntLiteral(1)")
	}
	mul, ok := add.Right.(*ast.BinaryExpr)
	if !ok || mul.Operator != ast.OpMul {
		t.Error("right of Add should be Mul")
	}
}

func TestParseUnaryNeg(t *testing.T) {
	expr := parseExpr(t, "-x")
	u := expr.(*ast.UnaryExpr)
	if u.Operator != ast.OpNeg {
		t.Errorf("operator: want OpNeg, got %v", u.Operator)
	}
}

func TestParseUnaryNot(t *testing.T) {
	expr := parseExpr(t, "!flag")
	u := expr.(*ast.UnaryExpr)
	if u.Operator != ast.OpNot {
		t.Errorf("operator: want OpNot, got %v", u.Operator)
	}
}

func TestParseIntLiteral(t *testing.T) {
	expr := parseExpr(t, "42")
	lit := expr.(*ast.IntLiteral)
	if lit.Value != 42 {
		t.Errorf("value: want 42, got %d", lit.Value)
	}
}

func TestParseFloatLiteral(t *testing.T) {
	expr := parseExpr(t, "3.14")
	lit := expr.(*ast.FloatLiteral)
	if lit.Value != 3.14 {
		t.Errorf("value: want 3.14, got %f", lit.Value)
	}
}

func TestParseStringLiteral(t *testing.T) {
	expr := parseExpr(t, `"hello"`)
	lit := expr.(*ast.StringLiteral)
	if lit.Value != "hello" {
		t.Errorf("value: want hello, got %s", lit.Value)
	}
}

func TestParseBoolLiteral(t *testing.T) {
	expr := parseExpr(t, "true")
	lit := expr.(*ast.BooleanLiteral)
	if !lit.Value {
		t.Error("value: want true")
	}
}

func TestParseCharLiteral(t *testing.T) {
	expr := parseExpr(t, "'a'")
	lit := expr.(*ast.CharLiteral)
	if lit.Value != 'a' {
		t.Errorf("value: want 'a', got %c", lit.Value)
	}
}

func TestParseFuncLiteral(t *testing.T) {
	src := "func(x: i32) -> i32 { return x; }"
	expr := parseExpr(t, src)
	fl := expr.(*ast.FuncLiteral)
	if len(fl.Params) != 1 || fl.Params[0].Name != "x" {
		t.Errorf("params: want [{x i32}], got %v", fl.Params)
	}
	if fl.ReturnType != "i32" {
		t.Errorf("return type: want i32, got %s", fl.ReturnType)
	}
}

// ==========================================
// Error recovery
// ==========================================

func TestParseErrorRecovery(t *testing.T) {
	// Bad declaration followed by a valid one — partial tree + error list
	src := "@@@ invalid\nlet x = 1;"
	prog, errs := ParseProgram(src)
	if len(errs) == 0 {
		t.Error("expected parse errors, got none")
	}
	// The valid declaration should still be in the tree
	found := false
	for _, d := range prog.Declarations {
		if v, ok := d.(*ast.VarDecl); ok && v.Name == "x" {
			found = true
		}
	}
	if !found {
		t.Error("valid VarDecl after bad input should be recovered")
	}
}
