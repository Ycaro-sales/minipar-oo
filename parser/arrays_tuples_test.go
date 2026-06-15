package parser

import (
	"testing"

	"minipar/ast"
)

// ==========================================
// ANOTAÇÕES DE TIPO COMPOSTAS (parseType)
// ==========================================

// TestVarDeclArrayType verifica que "let arr: [i32] = ..." produz VarDecl com
// Type == "[i32]" e value como ListLiteral.
func TestVarDeclArrayType(t *testing.T) {
	decl := parseOne(t, `let arr: [i32] = [1, 2, 3]`)
	v := mustType[*ast.VarDecl](t, decl)
	if v.Type != "[i32]" {
		t.Errorf("type: want [i32], got %q", v.Type)
	}
	list := mustType[*ast.ListLiteral](t, v.Value)
	if len(list.Elements) != 3 {
		t.Errorf("elements: want 3, got %d", len(list.Elements))
	}
}

// TestFuncParamArrayType verifica que parâmetro com tipo array é parseado corretamente.
func TestFuncParamArrayType(t *testing.T) {
	decl := parseOne(t, `func f(arr: [i32], n: i32) {}`)
	fn := mustType[*ast.FuncDecl](t, decl)
	if len(fn.Params) != 2 {
		t.Fatalf("params: want 2, got %d", len(fn.Params))
	}
	if fn.Params[0].Type != "[i32]" {
		t.Errorf("param[0].Type: want [i32], got %q", fn.Params[0].Type)
	}
	if fn.Params[1].Type != "i32" {
		t.Errorf("param[1].Type: want i32, got %q", fn.Params[1].Type)
	}
}

// TestFuncReturnArrayType verifica que tipo de retorno array é parseado.
func TestFuncReturnArrayType(t *testing.T) {
	decl := parseOne(t, `func f() -> [i32] { return [] }`)
	fn := mustType[*ast.FuncDecl](t, decl)
	if fn.ReturnType != "[i32]" {
		t.Errorf("return type: want [i32], got %q", fn.ReturnType)
	}
}

// TestVarDeclTupleType verifica anotação de tupla em var decl.
func TestVarDeclTupleType(t *testing.T) {
	decl := parseOne(t, `let t: (i32, string) = (1, "a")`)
	v := mustType[*ast.VarDecl](t, decl)
	if v.Type != "(i32, string)" {
		t.Errorf("type: want (i32, string), got %q", v.Type)
	}
}

// ==========================================
// LITERAL DE TUPLA
// ==========================================

// TestTupleLiteralParsed verifica que (expr, expr) produz TupleLiteral.
func TestTupleLiteralParsed(t *testing.T) {
	stmt := parseStmt(t, `let t = (1, "a")`)
	v := mustType[*ast.VarDecl](t, stmt)
	tup := mustType[*ast.TupleLiteral](t, v.Value)
	if len(tup.Elements) != 2 {
		t.Errorf("elements: want 2, got %d", len(tup.Elements))
	}
	mustType[*ast.IntLiteral](t, tup.Elements[0])
	mustType[*ast.StringLiteral](t, tup.Elements[1])
}

// TestTupleLiteralThreeElements verifica tupla com 3 elementos.
func TestTupleLiteralThreeElements(t *testing.T) {
	stmt := parseStmt(t, `let t = (1, 2, 3)`)
	v := mustType[*ast.VarDecl](t, stmt)
	tup := mustType[*ast.TupleLiteral](t, v.Value)
	if len(tup.Elements) != 3 {
		t.Errorf("elements: want 3, got %d", len(tup.Elements))
	}
}

// TestParenExprNotTuple verifica que (expr) sem vírgula não é tupla.
func TestParenExprNotTuple(t *testing.T) {
	stmt := parseStmt(t, `let x = (1 + 2)`)
	v := mustType[*ast.VarDecl](t, stmt)
	if _, isTuple := v.Value.(*ast.TupleLiteral); isTuple {
		t.Error("(1 + 2) não deveria ser TupleLiteral")
	}
	// Deve ser uma BinaryExpr
	mustType[*ast.BinaryExpr](t, v.Value)
}

// ==========================================
// ATRIBUIÇÃO POR ÍNDICE (arr[i] = x)
// ==========================================

// TestIndexAssignmentParsed verifica que arr[i] = x produz IndexAssignment.
func TestIndexAssignmentParsed(t *testing.T) {
	stmt := parseStmt(t, `arr[0] = 10`)
	ia := mustType[*ast.IndexAssignment](t, stmt)
	ident(t, ia.Object, "arr")
	intLit(t, ia.Index, 0)
	intLit(t, ia.Value, 10)
}

// TestIndexAssignmentWithVariable verifica arr[i] = x com variáveis.
func TestIndexAssignmentWithVariable(t *testing.T) {
	stmt := parseStmt(t, `arr[i] = x`)
	ia := mustType[*ast.IndexAssignment](t, stmt)
	ident(t, ia.Object, "arr")
	ident(t, ia.Index, "i")
	ident(t, ia.Value, "x")
}

// TestIndexAssignmentExprValue verifica arr[i] = expr complexa.
func TestIndexAssignmentExprValue(t *testing.T) {
	stmt := parseStmt(t, `arr[i] = x + 1`)
	ia := mustType[*ast.IndexAssignment](t, stmt)
	ident(t, ia.Object, "arr")
	mustType[*ast.BinaryExpr](t, ia.Value)
}

// TestArrayReadNotAssignment verifica que arr[i] sem = é ExprStmt (não IndexAssignment).
func TestArrayReadNotAssignment(t *testing.T) {
	// arr[i] como statement autônomo deve ser ExprStmt contendo IndexExpr
	// (ex: resultado descartado, raro mas válido)
	// Testamos via acesso em expressão: let x = arr[i]
	stmt := parseStmt(t, `let x = arr[i]`)
	v := mustType[*ast.VarDecl](t, stmt)
	mustType[*ast.IndexExpr](t, v.Value)
}
