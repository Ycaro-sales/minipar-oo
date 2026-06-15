package semantic

// Testes de semântica para arrays estáticos e tuplas.
// Cobrem: declaração, acesso por índice, alteração (arr[i]=x) e built-in len.
//
// Estratégia: onde o parser real ainda aceita a sintaxe, usa-se analyze().
// Onde o parser ainda não suporta (anotação [i32], tuplas, arr[i]=x),
// constrói-se o AST diretamente para testar o semântico de forma isolada.

import (
	"testing"

	"minipar/ast"
)

// ==========================================
// DECLARAÇÃO DE ARRAYS
// ==========================================

// TestArrayLiteralHomogeneous verifica que um literal homogêneo resolve para
// o tipo correto do elemento.
func TestArrayLiteralHomogeneous(t *testing.T) {
	cases := map[string]string{
		"inteiros":  `func f() { let xs = [1, 2, 3]; }`,
		"strings":   `func f() { let ss = ["a", "b"]; }`,
		"booleanos": `func f() { let bs = [true, false]; }`,
	}
	for name, src := range cases {
		t.Run(name, func(t *testing.T) {
			if errs := analyze(t, src); len(errs) != 0 {
				t.Errorf("esperava 0 erros, recebeu: %v", errs)
			}
		})
	}
}

// TestArrayLiteralHeterogeneousError verifica que lista com tipos misturados
// gera erro semântico.
func TestArrayLiteralHeterogeneousError(t *testing.T) {
	errs := analyze(t, `func f() { let xs = [1, "a"]; }`)
	if !containsAny(errs, "heterogênea") {
		t.Errorf("esperava erro de lista heterogênea, recebeu: %v", errs)
	}
}

// TestArrayLiteralEmpty verifica que uma lista vazia é aceita (elemento: any).
func TestArrayLiteralEmpty(t *testing.T) {
	if errs := analyze(t, `func f() { let xs = []; }`); len(errs) != 0 {
		t.Errorf("lista vazia: esperava 0 erros, recebeu: %v", errs)
	}
}

// TestArrayAnnotationViaAST verifica que o semântico, ao receber um VarDecl
// com tipo "[i32]", resolve o elemento corretamente (sem passar pelo parser).
// Isso testa resolveType("[i32]") que ainda não existe — fase vermelha.
func TestArrayAnnotationViaAST(t *testing.T) {
	// let arr: [i32] = [1, 2, 3]
	prog := &ast.Program{
		Declarations: []ast.Declaration{
			&ast.VarDecl{
				Line: 1,
				Name: "arr",
				Type: "[i32]",
				Value: &ast.ListLiteral{
					Line: 1,
					Elements: []ast.Expression{
						&ast.IntLiteral{Line: 1, Value: 1},
						&ast.IntLiteral{Line: 1, Value: 2},
						&ast.IntLiteral{Line: 1, Value: 3},
					},
				},
			},
		},
	}
	errs := New().Analyze(prog)
	if len(errs) != 0 {
		t.Errorf("let arr: [i32] = [1,2,3]: esperava 0 erros, recebeu: %v", errs)
	}
}

// TestArrayAnnotationTypeMismatch verifica que atribuir strings a [i32] gera erro.
func TestArrayAnnotationTypeMismatch(t *testing.T) {
	prog := &ast.Program{
		Declarations: []ast.Declaration{
			&ast.VarDecl{
				Line: 1,
				Name: "arr",
				Type: "[i32]",
				Value: &ast.ListLiteral{
					Line:     1,
					Elements: []ast.Expression{&ast.StringLiteral{Line: 1, Value: "oops"}},
				},
			},
		},
	}
	errs := New().Analyze(prog)
	if len(errs) == 0 {
		t.Error("esperava erro de tipo, recebeu 0 erros")
	}
}

// ==========================================
// ACESSO POR ÍNDICE
// ==========================================

// TestArrayIndexAccess verifica que arr[0] tem o tipo do elemento.
func TestArrayIndexAccess(t *testing.T) {
	if errs := analyze(t, `func f() { let xs = [1, 2, 3]; let x = xs[0]; }`); len(errs) != 0 {
		t.Errorf("xs[0]: esperava 0 erros, recebeu: %v", errs)
	}
}

// TestArrayIndexNotInt verifica que índice não-inteiro gera erro.
func TestArrayIndexNotInt(t *testing.T) {
	errs := analyze(t, `func f() { let xs = [1, 2]; let i = xs["a"]; }`)
	if !containsAny(errs, "índice") {
		t.Errorf("xs[\"a\"]: esperava erro de índice, recebeu: %v", errs)
	}
}

// ==========================================
// ALTERAÇÃO POR ÍNDICE (arr[i] = x)
// ==========================================

// TestArrayIndexAssignmentViaAST verifica que arr[0] = 5 é aceito quando arr é [int].
// Usa AST direto pois o parser ainda não emite IndexAssignment.
func TestArrayIndexAssignmentViaAST(t *testing.T) {
	// func f() { let arr = [1,2,3]; arr[0] = 5; }
	arrDecl := &ast.VarDecl{
		Line: 1, Name: "arr",
		Value: &ast.ListLiteral{
			Line:     1,
			Elements: []ast.Expression{&ast.IntLiteral{Value: 1}, &ast.IntLiteral{Value: 2}, &ast.IntLiteral{Value: 3}},
		},
	}
	indexAssign := &ast.IndexAssignment{
		Line:   2,
		Object: &ast.Identifier{Line: 2, Value: "arr"},
		Index:  &ast.IntLiteral{Line: 2, Value: 0},
		Value:  &ast.IntLiteral{Line: 2, Value: 5},
	}
	prog := &ast.Program{
		Declarations: []ast.Declaration{
			&ast.FuncDecl{
				Line: 1, Name: "f", ReturnType: "",
				Body: &ast.BlockStmt{Statements: []ast.Statement{arrDecl, indexAssign}},
			},
		},
	}
	errs := New().Analyze(prog)
	if len(errs) != 0 {
		t.Errorf("arr[0]=5: esperava 0 erros, recebeu: %v", errs)
	}
}

// TestArrayIndexAssignmentTypeMismatchViaAST verifica que arr[0]="x" gera erro
// quando arr é [int].
func TestArrayIndexAssignmentTypeMismatchViaAST(t *testing.T) {
	arrDecl := &ast.VarDecl{
		Line: 1, Name: "arr",
		Value: &ast.ListLiteral{
			Line:     1,
			Elements: []ast.Expression{&ast.IntLiteral{Value: 1}},
		},
	}
	indexAssign := &ast.IndexAssignment{
		Line:   2,
		Object: &ast.Identifier{Line: 2, Value: "arr"},
		Index:  &ast.IntLiteral{Line: 2, Value: 0},
		Value:  &ast.StringLiteral{Line: 2, Value: "oops"},
	}
	prog := &ast.Program{
		Declarations: []ast.Declaration{
			&ast.FuncDecl{
				Line: 1, Name: "f", ReturnType: "",
				Body: &ast.BlockStmt{Statements: []ast.Statement{arrDecl, indexAssign}},
			},
		},
	}
	errs := New().Analyze(prog)
	if len(errs) == 0 {
		t.Error("arr[0]=\"x\" em [int]: esperava erro de tipo, recebeu 0 erros")
	}
}

// ==========================================
// TUPLAS
// ==========================================

// TestTupleLiteralViaAST verifica que uma tupla heterogênea é aceita.
func TestTupleLiteralViaAST(t *testing.T) {
	// let t = (1, "a")
	prog := &ast.Program{
		Declarations: []ast.Declaration{
			&ast.VarDecl{
				Line: 1, Name: "tup",
				Value: &ast.TupleLiteral{
					Line: 1,
					Elements: []ast.Expression{
						&ast.IntLiteral{Line: 1, Value: 1},
						&ast.StringLiteral{Line: 1, Value: "a"},
					},
				},
			},
		},
	}
	errs := New().Analyze(prog)
	if len(errs) != 0 {
		t.Errorf("tupla (1, \"a\"): esperava 0 erros, recebeu: %v", errs)
	}
}

// TestTupleIndexAccessViaAST verifica que t[0] retorna o tipo do elemento na
// posição 0 e t[1] o da posição 1.
func TestTupleIndexAccessViaAST(t *testing.T) {
	// let tup = (1, "a"); let x = tup[0]; let s = tup[1];
	tupDecl := &ast.VarDecl{
		Line: 1, Name: "tup",
		Value: &ast.TupleLiteral{
			Line: 1,
			Elements: []ast.Expression{
				&ast.IntLiteral{Line: 1, Value: 1},
				&ast.StringLiteral{Line: 1, Value: "a"},
			},
		},
	}
	accessX := &ast.VarDecl{
		Line: 2, Name: "x",
		Value: &ast.IndexExpr{
			Line:   2,
			Object: &ast.Identifier{Line: 2, Value: "tup"},
			Index:  &ast.IntLiteral{Line: 2, Value: 0},
		},
	}
	accessS := &ast.VarDecl{
		Line: 3, Name: "s",
		Value: &ast.IndexExpr{
			Line:   3,
			Object: &ast.Identifier{Line: 3, Value: "tup"},
			Index:  &ast.IntLiteral{Line: 3, Value: 1},
		},
	}
	prog := &ast.Program{
		Declarations: []ast.Declaration{
			&ast.FuncDecl{
				Line: 1, Name: "f", ReturnType: "",
				Body: &ast.BlockStmt{Statements: []ast.Statement{tupDecl, accessX, accessS}},
			},
		},
	}
	errs := New().Analyze(prog)
	if len(errs) != 0 {
		t.Errorf("tup[0]/tup[1]: esperava 0 erros, recebeu: %v", errs)
	}
}

// TestTupleIndexOutOfBoundsViaAST verifica que acesso além do tamanho gera erro.
func TestTupleIndexOutOfBoundsViaAST(t *testing.T) {
	tupDecl := &ast.VarDecl{
		Line: 1, Name: "tup",
		Value: &ast.TupleLiteral{
			Line:     1,
			Elements: []ast.Expression{&ast.IntLiteral{Value: 1}},
		},
	}
	access := &ast.VarDecl{
		Line: 2, Name: "x",
		Value: &ast.IndexExpr{
			Line:   2,
			Object: &ast.Identifier{Line: 2, Value: "tup"},
			Index:  &ast.IntLiteral{Line: 2, Value: 5}, // fora dos limites
		},
	}
	prog := &ast.Program{
		Declarations: []ast.Declaration{
			&ast.FuncDecl{
				Line: 1, Name: "f", ReturnType: "",
				Body: &ast.BlockStmt{Statements: []ast.Statement{tupDecl, access}},
			},
		},
	}
	errs := New().Analyze(prog)
	if len(errs) == 0 {
		t.Error("tup[5] em tupla de 1 elemento: esperava erro de índice, recebeu 0 erros")
	}
}

// TestTupleImmutableViaAST verifica que tentar atribuir a tup[0] gera erro.
func TestTupleImmutableViaAST(t *testing.T) {
	tupDecl := &ast.VarDecl{
		Line: 1, Name: "tup",
		Value: &ast.TupleLiteral{
			Line:     1,
			Elements: []ast.Expression{&ast.IntLiteral{Value: 1}},
		},
	}
	assign := &ast.IndexAssignment{
		Line:   2,
		Object: &ast.Identifier{Line: 2, Value: "tup"},
		Index:  &ast.IntLiteral{Line: 2, Value: 0},
		Value:  &ast.IntLiteral{Line: 2, Value: 99},
	}
	prog := &ast.Program{
		Declarations: []ast.Declaration{
			&ast.FuncDecl{
				Line: 1, Name: "f", ReturnType: "",
				Body: &ast.BlockStmt{Statements: []ast.Statement{tupDecl, assign}},
			},
		},
	}
	errs := New().Analyze(prog)
	if !containsAny(errs, "imutável") {
		t.Errorf("tup[0]=x: esperava erro de imutabilidade, recebeu: %v", errs)
	}
}

// ==========================================
// LEN
// ==========================================

// TestLenBuiltinOnArray verifica que len(arr) analisa sem erros.
func TestLenBuiltinOnArray(t *testing.T) {
	if errs := analyze(t, `func f() { let xs = [1, 2, 3]; let n = len(xs); }`); len(errs) != 0 {
		t.Errorf("len(arr): esperava 0 erros, recebeu: %v", errs)
	}
}

// ==========================================
// INTEGRAÇÃO PARSER → SEMÂNTICO
// Testes que usam o pipeline real (analyze) com a nova sintaxe habilitada.
// ==========================================

// TestIntegrationArrayAnnotation verifica declaração com anotação [i32]
// passando pelo parser real e analisador semântico.
func TestIntegrationArrayAnnotation(t *testing.T) {
	cases := map[string]string{
		"anotação [i32]":     `func f() { let arr: [i32] = [1, 2, 3]; }`,
		"anotação [string]":  `func f() { let ss: [string] = ["a", "b"]; }`,
		"anotação [bool]":    `func f() { let bs: [bool] = [true, false]; }`,
		"param array":        `func f(arr: [i32]) {}`,
		"retorno array":      `func f() -> [i32] { return [1, 2] }`,
	}
	for name, src := range cases {
		t.Run(name, func(t *testing.T) {
			if errs := analyze(t, src); len(errs) != 0 {
				t.Errorf("esperava 0 erros, recebeu: %v", errs)
			}
		})
	}
}

// TestIntegrationArrayAnnotationMismatch verifica que [i32] = ["x"] gera erro.
func TestIntegrationArrayAnnotationMismatch(t *testing.T) {
	errs := analyze(t, `func f() { let arr: [i32] = ["oops"]; }`)
	if len(errs) == 0 {
		t.Error("esperava erro de tipo para [i32] = [string], recebeu 0 erros")
	}
}

// TestIntegrationIndexAccess verifica acesso arr[i] via parser real.
func TestIntegrationIndexAccess(t *testing.T) {
	cases := map[string]string{
		"acesso literal":  `func f() { let arr: [i32] = [1,2,3]; let x: i32 = arr[0]; }`,
		"acesso variável": `func f() { let arr: [i32] = [1,2,3]; let i: i32 = 0; let x = arr[i]; }`,
	}
	for name, src := range cases {
		t.Run(name, func(t *testing.T) {
			if errs := analyze(t, src); len(errs) != 0 {
				t.Errorf("%s: esperava 0 erros, recebeu: %v", name, errs)
			}
		})
	}
}

// TestIntegrationIndexAssignment verifica arr[i] = x via parser real.
func TestIntegrationIndexAssignment(t *testing.T) {
	cases := map[string]string{
		"assign literal":  `func f() { let arr: [i32] = [1,2,3]; arr[0] = 10; }`,
		"assign variável": `func f() { let arr: [i32] = [1,2,3]; let i: i32 = 1; arr[i] = 99; }`,
	}
	for name, src := range cases {
		t.Run(name, func(t *testing.T) {
			if errs := analyze(t, src); len(errs) != 0 {
				t.Errorf("%s: esperava 0 erros, recebeu: %v", name, errs)
			}
		})
	}
}

// TestIntegrationIndexAssignmentTypeMismatch verifica erro de tipo em arr[i]=x.
func TestIntegrationIndexAssignmentTypeMismatch(t *testing.T) {
	errs := analyze(t, `func f() { let arr: [i32] = [1,2]; arr[0] = "oops"; }`)
	if len(errs) == 0 {
		t.Error("arr[0]=\"oops\" em [i32]: esperava erro de tipo, recebeu 0 erros")
	}
}

// TestIntegrationTupleLiteral verifica literal de tupla via parser real.
func TestIntegrationTupleLiteral(t *testing.T) {
	if errs := analyze(t, `func f() { let t = (1, "a", true); }`); len(errs) != 0 {
		t.Errorf("tupla (1, \"a\", true): esperava 0 erros, recebeu: %v", errs)
	}
}

// TestIntegrationTupleAnnotation verifica declaração com anotação (i32, string).
func TestIntegrationTupleAnnotation(t *testing.T) {
	if errs := analyze(t, `func f() { let t: (i32, string) = (1, "a"); }`); len(errs) != 0 {
		t.Errorf("let t: (i32, string): esperava 0 erros, recebeu: %v", errs)
	}
}

// TestIntegrationTupleIndexAccess verifica t[0] e t[1] via parser real.
func TestIntegrationTupleIndexAccess(t *testing.T) {
	src := `func f() {
		let t = (42, "hello")
		let n = t[0]
		let s = t[1]
	}`
	if errs := analyze(t, src); len(errs) != 0 {
		t.Errorf("t[0]/t[1]: esperava 0 erros, recebeu: %v", errs)
	}
}

// TestIntegrationTupleImmutable verifica que t[0] = x via parser real gera erro.
func TestIntegrationTupleImmutable(t *testing.T) {
	errs := analyze(t, `func f() { let t = (1, 2); t[0] = 99; }`)
	if !containsAny(errs, "imutável") {
		t.Errorf("t[0]=x: esperava erro de imutabilidade, recebeu: %v", errs)
	}
}

// TestIntegrationArrayParamInFunc verifica que função com parâmetro [i32]
// pode acessar e modificar elementos — base para o quicksort.
func TestIntegrationArrayParamInFunc(t *testing.T) {
	src := `
func swap(arr: [i32], i: i32, j: i32) {
	let tmp: i32 = arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
func main() {
	let arr: [i32] = [3, 1, 2]
	swap(arr, 0, 2)
}`
	if errs := analyze(t, src); len(errs) != 0 {
		t.Errorf("swap: esperava 0 erros, recebeu: %v", errs)
	}
}
