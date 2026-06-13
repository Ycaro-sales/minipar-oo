// Package semantic implements the semantic-analysis pass of the Minipar
// compiler. It consumes the AST produced by the parser and reports semantic
// errors (undeclared names, type mismatches, bad control flow, broken
// interface contracts) as "linha N: ..." strings, matching the parser.
package semantic

import (
	"fmt"

	"minipar/ast"
	"minipar/internal/symtab"
)

// classInfo holds a class's resolved members for lookups during checking.
type classInfo struct {
	name       string
	implements []string
	fields     map[string]*Type
	methods    map[string]*Type // method name -> KindFunc type
}

// ifaceInfo holds an interface's required method signatures.
type ifaceInfo struct {
	name    string
	methods map[string]*Type
}

// Analyzer walks the AST and accumulates semantic errors.
type Analyzer struct {
	global *scope
	scope  *scope
	errors []string

	classes    map[string]*classInfo
	interfaces map[string]*ifaceInfo

	// context tracked while descending into bodies
	currentReturn *Type      // expected return type of the enclosing function
	currentClass  *classInfo // enclosing class, for `Self` typing
	loopDepth     int        // >0 while inside a for/while/do
}

// New returns a ready-to-use Analyzer.
func New() *Analyzer {
	g := symtab.NewGlobal[*Type]()
	return &Analyzer{
		global:     g,
		scope:      g,
		classes:    map[string]*classInfo{},
		interfaces: map[string]*ifaceInfo{},
	}
}

// GlobalScope returns the root of the retained scope tree built during
// analysis, so later passes (e.g. code generation) can resolve names again.
func (a *Analyzer) GlobalScope() *symtab.Scope[*Type] { return a.global }

// Analyze checks an entire program and returns the collected error messages.
// The slice is empty when the program is semantically valid.
func (a *Analyzer) Analyze(prog *ast.Program) []string {
	if prog == nil {
		return a.errors
	}
	a.collectDeclarations(prog) // pass 1: signatures and type names
	a.checkDeclarations(prog)   // pass 2: bodies
	return a.errors
}

func (a *Analyzer) errorf(line int, format string, args ...any) {
	a.errors = append(a.errors, fmt.Sprintf("linha %d: ", line)+fmt.Sprintf(format, args...))
}

// ==========================================
// Pass 1 — collect top-level declarations
// ==========================================

func (a *Analyzer) collectDeclarations(prog *ast.Program) {
	// First register every class/interface *name* so member signatures can
	// reference types declared later in the file.
	for _, d := range prog.Declarations {
		switch n := d.(type) {
		case *ast.ClassDecl:
			a.classes[n.Name] = &classInfo{
				name: n.Name, implements: n.Implements,
				fields: map[string]*Type{}, methods: map[string]*Type{},
			}
			a.defineGlobal(&symbol{Name: n.Name, Kind: symtab.Class, Line: n.Line,
				Type: &Type{Name: n.Name, Kind: KindClass}})
		case *ast.InterfaceDecl:
			a.interfaces[n.Name] = &ifaceInfo{name: n.Name, methods: map[string]*Type{}}
			a.defineGlobal(&symbol{Name: n.Name, Kind: symtab.Interface, Line: n.Line,
				Type: &Type{Name: n.Name, Kind: KindInterface}})
		}
	}

	// Now resolve member and top-level signatures (types are all known).
	for _, d := range prog.Declarations {
		switch n := d.(type) {
		case *ast.FuncDecl:
			a.defineGlobal(&symbol{Name: n.Name, Kind: symtab.Func, Line: n.Line,
				Type: a.funcType(n.Params, n.ReturnType)})
		case *ast.VarDecl:
			a.defineGlobal(&symbol{Name: n.Name, Kind: symtab.Var, Line: n.Line,
				Type: a.resolveType(n.Type, n.Line)})
		case *ast.ClassDecl:
			a.collectClassMembers(n)
		case *ast.InterfaceDecl:
			info := a.interfaces[n.Name]
			for _, m := range n.Methods {
				info.methods[m.Name] = a.funcType(m.Params, m.ReturnType)
			}
		}
	}
}

func (a *Analyzer) collectClassMembers(n *ast.ClassDecl) {
	info := a.classes[n.Name]
	for _, member := range n.Members {
		switch m := member.(type) {
		case *ast.FieldDecl:
			info.fields[m.Name] = a.resolveType(m.Type, m.Line)
		case *ast.MethodDecl:
			info.methods[m.Name] = a.funcType(m.Params, m.ReturnType)
		}
	}
}

func (a *Analyzer) defineGlobal(sym *symbol) {
	if prev, ok := a.global.Define(sym); !ok {
		a.errorf(sym.Line, "'%s' já declarado na linha %d", sym.Name, prev.Line)
	}
}

// funcType builds a KindFunc Type from a parameter list and a return spec.
func (a *Analyzer) funcType(params []ast.Param, ret string) *Type {
	ps := make([]*Type, len(params))
	for i, p := range params {
		if p.Type == "" {
			ps[i] = tAny // untyped param (e.g. `self`) defaults to any
		} else {
			ps[i] = a.resolveType(p.Type, 0)
		}
	}
	rt := tVoid
	if ret != "" {
		rt = a.resolveType(ret, 0)
	}
	return &Type{Name: "func", Kind: KindFunc, Params: ps, Return: rt}
}

// resolveType maps a raw AST type string to a resolved *Type, reporting unknown
// names. An empty string means "infer later" and yields any.
func (a *Analyzer) resolveType(name string, line int) *Type {
	if name == "" {
		return tAny
	}
	if t, ok := primitives[name]; ok {
		return t
	}
	if _, ok := a.classes[name]; ok {
		return &Type{Name: name, Kind: KindClass}
	}
	if _, ok := a.interfaces[name]; ok {
		return &Type{Name: name, Kind: KindInterface}
	}
	a.errorf(line, "tipo desconhecido: '%s'", name)
	return tInvalid
}

// ==========================================
// Pass 2 — check bodies
// ==========================================

func (a *Analyzer) checkDeclarations(prog *ast.Program) {
	for _, d := range prog.Declarations {
		switch n := d.(type) {
		case *ast.VarDecl:
			a.checkGlobalVar(n)
		case *ast.FuncDecl:
			a.checkFunc(n.Params, n.ReturnType, n.Body)
		case *ast.ClassDecl:
			a.checkClass(n)
		case *ast.InterfaceDecl:
			// signatures only; nothing to check in a body
		}
	}
}

func (a *Analyzer) checkClass(n *ast.ClassDecl) {
	info := a.classes[n.Name]

	// Every declared interface must exist and be fully implemented.
	for _, iname := range n.Implements {
		iface, ok := a.interfaces[iname]
		if !ok {
			a.errorf(n.Line, "classe '%s' implementa interface desconhecida '%s'", n.Name, iname)
			continue
		}
		for mname, sig := range iface.methods {
			impl, ok := info.methods[mname]
			if !ok {
				a.errorf(n.Line, "classe '%s' não implementa o método '%s' exigido por '%s'",
					n.Name, mname, iname)
				continue
			}
			if !sameSignature(sig, impl) {
				a.errorf(n.Line, "método '%s' de '%s' não corresponde à assinatura de '%s'",
					mname, n.Name, iname)
			}
		}
	}

	prevScope, prevClass := a.scope, a.currentClass
	a.scope = a.global.Push(symtab.ClassScope)
	a.currentClass = info

	// Bind every field first so a method body can reference a field declared
	// later in the class. Method bodies open Function scopes parented off this
	// Class scope, so bare field names resolve by walking outward.
	for _, member := range n.Members {
		if f, ok := member.(*ast.FieldDecl); ok {
			a.define(&symbol{Name: f.Name, Kind: symtab.Field, Type: info.fields[f.Name], Line: f.Line})
		}
	}

	for _, member := range n.Members {
		switch m := member.(type) {
		case *ast.FieldDecl:
			if m.Value != nil {
				vt := a.checkExpr(m.Value)
				ft := info.fields[m.Name]
				if !assignable(ft, vt) {
					a.errorf(m.Line, "não é possível atribuir '%s' ao campo '%s' do tipo '%s'",
						vt, m.Name, ft)
				}
			}
		case *ast.MethodDecl:
			a.checkFunc(m.Params, m.ReturnType, m.Body)
		case *ast.ConstructorDecl:
			a.checkFunc(nil, "", m.Body)
		}
	}
	a.scope, a.currentClass = prevScope, prevClass
}

// checkFunc opens a new scope, binds parameters, then checks the body against
// the declared return type.
func (a *Analyzer) checkFunc(params []ast.Param, ret string, body *ast.BlockStmt) {
	prevScope, prevReturn := a.scope, a.currentReturn
	a.scope = prevScope.Push(symtab.FunctionScope)
	a.currentReturn = tVoid
	if ret != "" {
		a.currentReturn = a.resolveType(ret, 0)
	}

	for _, p := range params {
		t := tAny
		if p.Type != "" {
			t = a.resolveType(p.Type, 0)
		}
		a.define(&symbol{Name: p.Name, Kind: symtab.Param, Type: t})
	}
	if body != nil {
		a.checkBlock(body, false)
	}

	a.scope, a.currentReturn = prevScope, prevReturn
}

func (a *Analyzer) define(sym *symbol) {
	if prev, ok := a.scope.Define(sym); !ok {
		a.errorf(sym.Line, "'%s' já declarado neste escopo (linha %d)", sym.Name, prev.Line)
	}
}

// ==========================================
// Statements
// ==========================================

// checkBlock checks statements in a nested Block scope. ownScope=false reuses
// the caller's scope (used by checkFunc, which already opened one for params).
func (a *Analyzer) checkBlock(b *ast.BlockStmt, ownScope bool) {
	if ownScope {
		prev := a.scope
		a.scope = prev.Push(symtab.BlockScope)
		defer func() { a.scope = prev }()
	}
	for _, s := range b.Statements {
		a.checkStmt(s)
	}
}

func (a *Analyzer) checkStmt(s ast.Statement) {
	switch n := s.(type) {
	case *ast.VarDecl:
		a.checkVarDecl(n)
	case *ast.Assignment:
		a.checkAssignment(n)
	case *ast.IfStmt:
		a.expectBool(n.Condition, "condição do if")
		a.checkBlock(n.Then, true)
		if n.Else != nil {
			a.checkBlock(n.Else, true)
		}
	case *ast.WhileStmt:
		a.expectBool(n.Condition, "condição do while")
		a.loopDepth++
		a.checkBlock(n.Block, true)
		a.loopDepth--
	case *ast.DoStmt:
		a.loopDepth++
		a.checkBlock(n.Block, true)
		a.loopDepth--
		a.expectBool(n.Condition, "condição do do-while")
	case *ast.ForStmt:
		a.checkFor(n)
	case *ast.SwitchStmt:
		a.checkSwitch(n)
	case *ast.SeqStmt:
		a.checkBlock(n.Block, true)
	case *ast.ParStmt:
		a.checkBlock(n.Block, true)
	case *ast.ReturnStmt:
		a.checkReturn(n)
	case *ast.BreakStmt:
		if a.loopDepth == 0 {
			a.errorf(n.Line, "'break' fora de um laço")
		}
	case *ast.ContinueStmt:
		if a.loopDepth == 0 {
			a.errorf(n.Line, "'continue' fora de um laço")
		}
	case *ast.PrintStmt:
		for _, arg := range n.Args {
			a.checkExpr(arg)
		}
	case *ast.InputStmt:
		if n.Prompt != nil {
			a.checkExpr(n.Prompt)
		}
	case *ast.ExprStmt:
		a.checkExpr(n.Expression)
	case *ast.FuncCall:
		a.checkExpr(n)
	case *ast.MethodCall:
		a.checkExpr(n)
	case *ast.ChannelStmt:
		a.checkChannel(n)
	case *ast.PassStmt, *ast.GotoStmt:
		// nothing to type-check
	}
}

// checkGlobalVar checks a top-level var. The symbol was already inserted in
// pass 1, so this only validates the initializer and, when the type was left
// to inference, back-fills the symbol's type from the value.
func (a *Analyzer) checkGlobalVar(n *ast.VarDecl) {
	if n.Value == nil {
		return
	}
	valueType := a.checkExpr(n.Value)
	if n.Type != "" {
		declared := a.resolveType(n.Type, n.Line)
		if !assignable(declared, valueType) {
			a.errorf(n.Line, "não é possível atribuir valor '%s' à variável '%s' do tipo '%s'",
				valueType, n.Name, declared)
		}
		return
	}
	if sym := a.global.Resolve(n.Name); sym != nil {
		sym.Type = valueType // inferred
	}
}

func (a *Analyzer) checkVarDecl(n *ast.VarDecl) {
	var valueType *Type = tAny
	if n.Value != nil {
		valueType = a.checkExpr(n.Value)
	}

	declared := valueType // inferred when no annotation is present
	if n.Type != "" {
		declared = a.resolveType(n.Type, n.Line)
		if n.Value != nil && !assignable(declared, valueType) {
			a.errorf(n.Line, "não é possível atribuir valor '%s' à variável '%s' do tipo '%s'",
				valueType, n.Name, declared)
		}
	}
	a.define(&symbol{Name: n.Name, Kind: symtab.Var, Type: declared, Line: n.Line})
}

func (a *Analyzer) checkAssignment(n *ast.Assignment) {
	valueType := a.checkExpr(n.Value)

	// Resolution walks the scope chain; inside a method this reaches the
	// enclosing Class scope, so bare field writes resolve here too.
	sym := a.scope.Resolve(n.Name)
	if sym == nil {
		a.errorf(n.Line, "atribuição a variável não declarada: '%s'", n.Name)
		return
	}
	if !assignable(sym.Type, valueType) {
		a.errorf(n.Line, "não é possível atribuir '%s' a '%s' do tipo '%s'",
			valueType, n.Name, sym.Type)
	}
}

func (a *Analyzer) checkFor(n *ast.ForStmt) {
	iter := a.checkExpr(n.Iter)
	elem := tAny
	if iter.Kind == KindArray {
		elem = iter.Elem
	} else if iter.Kind != KindAny && iter.Kind != KindString && iter.Kind != KindInvalid {
		a.errorf(n.Line, "tipo '%s' não é iterável", iter)
	}

	prev := a.scope
	a.scope = prev.Push(symtab.BlockScope)
	a.define(&symbol{Name: n.Var, Kind: symtab.Var, Type: elem, Line: n.Line})
	a.loopDepth++
	for _, s := range n.Block.Statements {
		a.checkStmt(s)
	}
	a.loopDepth--
	a.scope = prev
}

func (a *Analyzer) checkSwitch(n *ast.SwitchStmt) {
	subject := a.checkExpr(n.Expr)
	for _, c := range n.Cases {
		ct := a.checkExpr(c.Value)
		if !assignable(subject, ct) && !assignable(ct, subject) {
			a.errorf(c.Line, "case do tipo '%s' incompatível com o switch do tipo '%s'", ct, subject)
		}
		a.checkBlock(c.Block, true)
	}
}

func (a *Analyzer) checkReturn(n *ast.ReturnStmt) {
	if n.Value == nil {
		if a.currentReturn.Kind != KindVoid {
			a.errorf(n.Line, "função deve retornar '%s', mas o return está vazio", a.currentReturn)
		}
		return
	}
	got := a.checkExpr(n.Value)
	if a.currentReturn.Kind == KindVoid {
		a.errorf(n.Line, "return com valor em função sem tipo de retorno")
		return
	}
	if !assignable(a.currentReturn, got) {
		a.errorf(n.Line, "retorno do tipo '%s' incompatível com o tipo declarado '%s'",
			got, a.currentReturn)
	}
}

func (a *Analyzer) checkChannel(n *ast.ChannelStmt) {
	for _, arg := range n.Args {
		a.checkExpr(arg)
	}
	a.define(&symbol{Name: n.Name, Kind: symtab.Var, Type: tAny, Line: n.Line})
}

// ==========================================
// Expressions
// ==========================================

// checkExpr returns the type of an expression, reporting errors as it goes.
// It never returns nil: on error it yields tInvalid to halt cascading.
func (a *Analyzer) checkExpr(e ast.Expression) *Type {
	switch n := e.(type) {
	case *ast.IntLiteral:
		return primitives["int"]
	case *ast.FloatLiteral:
		return primitives["float"]
	case *ast.StringLiteral:
		return tString
	case *ast.CharLiteral:
		return tChar
	case *ast.BooleanLiteral:
		return tBool
	case *ast.Identifier:
		return a.checkIdentifier(n)
	case *ast.SelfExpr:
		if a.currentClass == nil {
			a.errorf(n.Line, "'self' usado fora de um método")
			return tInvalid
		}
		return &Type{Name: a.currentClass.name, Kind: KindClass}
	case *ast.BinaryExpr:
		return a.checkBinary(n)
	case *ast.UnaryExpr:
		return a.checkUnary(n)
	case *ast.FuncCall:
		return a.checkCall(n)
	case *ast.MethodCall:
		return a.checkMethodCall(n)
	case *ast.IndexExpr:
		return a.checkIndex(n)
	case *ast.ObjCreation:
		return a.checkObjCreation(n)
	case *ast.ListLiteral:
		return a.checkList(n)
	case *ast.FuncLiteral:
		return a.funcType(n.Params, n.ReturnType)
	case *ast.DictLiteral:
		return tAny
	case nil:
		return tInvalid
	}
	return tAny
}

func (a *Analyzer) checkIdentifier(n *ast.Identifier) *Type {
	// Resolution walks outward to enclosing scopes, including the Class scope
	// for bare field references inside a method body.
	if sym := a.scope.Resolve(n.Value); sym != nil {
		return sym.Type
	}
	a.errorf(n.Line, "identificador não declarado: '%s'", n.Value)
	return tInvalid
}

func (a *Analyzer) checkBinary(n *ast.BinaryExpr) *Type {
	lt := a.checkExpr(n.Left)
	rt := a.checkExpr(n.Right)
	if lt.Kind == KindInvalid || rt.Kind == KindInvalid {
		return tInvalid
	}

	switch n.Operator {
	case ast.OpAnd, ast.OpOr:
		if lt.Kind != KindBool || rt.Kind != KindBool {
			a.errorf(n.Line, "operador lógico requer operandos 'bool', recebeu '%s' e '%s'", lt, rt)
		}
		return tBool
	case ast.OpEq, ast.OpNeq:
		return tBool
	case ast.OpLt, ast.OpGt, ast.OpLeq, ast.OpGeq:
		if !(lt.isNumeric() && rt.isNumeric()) && lt.Name != rt.Name {
			a.errorf(n.Line, "comparação inválida entre '%s' e '%s'", lt, rt)
		}
		return tBool
	default: // arithmetic
		if lt.Kind == KindString && rt.Kind == KindString && n.Operator == ast.OpAdd {
			return tString // string concatenation
		}
		if !lt.isNumeric() || !rt.isNumeric() {
			a.errorf(n.Line, "operação aritmética requer números, recebeu '%s' e '%s'", lt, rt)
			return tInvalid
		}
		if lt.Kind == KindFloat || rt.Kind == KindFloat {
			return primitives["float"]
		}
		return lt
	}
}

func (a *Analyzer) checkUnary(n *ast.UnaryExpr) *Type {
	rt := a.checkExpr(n.Right)
	switch n.Operator {
	case ast.OpNot:
		if rt.Kind != KindBool && rt.Kind != KindInvalid {
			a.errorf(n.Line, "operador '!' requer 'bool', recebeu '%s'", rt)
		}
		return tBool
	case ast.OpNeg:
		if !rt.isNumeric() && rt.Kind != KindInvalid {
			a.errorf(n.Line, "operador '-' requer número, recebeu '%s'", rt)
			return tInvalid
		}
		return rt
	}
	return tInvalid
}

func (a *Analyzer) checkCall(n *ast.FuncCall) *Type {
	// A call whose name is a known class is object construction (`Foo(...)`).
	if _, ok := a.classes[n.Name]; ok {
		return a.callClass(n.Name, n.Args, n.Line)
	}
	sym := a.scope.Resolve(n.Name)
	if sym == nil {
		a.errorf(n.Line, "chamada a função não declarada: '%s'", n.Name)
		a.checkArgs(n.Args)
		return tInvalid
	}
	if sym.Type.Kind != KindFunc {
		a.errorf(n.Line, "'%s' não é uma função", n.Name)
		a.checkArgs(n.Args)
		return tInvalid
	}
	a.checkCallArgs(n.Name, sym.Type, n.Args, n.Line)
	return sym.Type.Return
}

func (a *Analyzer) checkMethodCall(n *ast.MethodCall) *Type {
	obj := a.checkExpr(n.Object)
	if obj.Kind == KindInvalid || obj.Kind == KindAny {
		a.checkArgs(n.Args)
		return tAny
	}

	var sig *Type
	switch obj.Kind {
	case KindClass:
		if info, ok := a.classes[obj.Name]; ok {
			sig = info.methods[n.Method]
		}
	case KindInterface:
		if info, ok := a.interfaces[obj.Name]; ok {
			sig = info.methods[n.Method]
		}
	}
	if sig == nil {
		a.errorf(n.Line, "tipo '%s' não possui método '%s'", obj, n.Method)
		a.checkArgs(n.Args)
		return tInvalid
	}
	a.checkCallArgs(n.Method, sig, n.Args, n.Line)
	return sig.Return
}

func (a *Analyzer) checkObjCreation(n *ast.ObjCreation) *Type {
	return a.callClass(n.Class, n.Args, n.Line)
}

// callClass type-checks construction of a class and returns its instance type.
func (a *Analyzer) callClass(name string, args []ast.Expression, line int) *Type {
	if _, ok := a.classes[name]; !ok {
		a.errorf(line, "classe desconhecida: '%s'", name)
		a.checkArgs(args)
		return tInvalid
	}
	a.checkArgs(args) // constructor arity is not modeled in the AST yet
	return &Type{Name: name, Kind: KindClass}
}

func (a *Analyzer) checkIndex(n *ast.IndexExpr) *Type {
	obj := a.checkExpr(n.Object)
	idx := a.checkExpr(n.Index)
	if idx.Kind != KindInt && idx.Kind != KindInvalid && idx.Kind != KindAny {
		a.errorf(n.Line, "índice deve ser inteiro, recebeu '%s'", idx)
	}
	switch obj.Kind {
	case KindArray:
		return obj.Elem
	case KindString:
		return tChar
	case KindAny, KindInvalid:
		return tAny
	}
	a.errorf(n.Line, "tipo '%s' não é indexável", obj)
	return tInvalid
}

func (a *Analyzer) checkList(n *ast.ListLiteral) *Type {
	if len(n.Elements) == 0 {
		return &Type{Name: "[]", Kind: KindArray, Elem: tAny}
	}
	elem := a.checkExpr(n.Elements[0])
	for _, e := range n.Elements[1:] {
		et := a.checkExpr(e)
		if !assignable(elem, et) && !assignable(et, elem) {
			a.errorf(n.Line, "lista heterogênea: '%s' e '%s'", elem, et)
		}
	}
	return &Type{Name: "[" + elem.String() + "]", Kind: KindArray, Elem: elem}
}

// ==========================================
// Argument helpers
// ==========================================

func (a *Analyzer) checkArgs(args []ast.Expression) {
	for _, arg := range args {
		a.checkExpr(arg)
	}
}

// checkCallArgs verifies arity and per-argument assignability against a signature.
func (a *Analyzer) checkCallArgs(name string, sig *Type, args []ast.Expression, line int) {
	params := sig.Params
	got := make([]*Type, len(args))
	for i, arg := range args {
		got[i] = a.checkExpr(arg)
	}
	if len(args) != len(params) {
		a.errorf(line, "'%s' espera %d argumento(s), recebeu %d", name, len(params), len(args))
		return
	}
	for i := range args {
		if !assignable(params[i], got[i]) {
			a.errorf(line, "argumento %d de '%s': esperado '%s', recebido '%s'",
				i+1, name, params[i], got[i])
		}
	}
}

// expectBool checks that a condition expression is boolean.
func (a *Analyzer) expectBool(e ast.Expression, ctx string) {
	t := a.checkExpr(e)
	if t.Kind != KindBool && t.Kind != KindInvalid && t.Kind != KindAny {
		a.errorf(e.GetLine(), "%s deve ser 'bool', recebeu '%s'", ctx, t)
	}
}

// sameSignature compares two function types for interface conformance.
func sameSignature(want, got *Type) bool {
	if len(want.Params) != len(got.Params) {
		return false
	}
	for i := range want.Params {
		if !assignable(want.Params[i], got.Params[i]) {
			return false
		}
	}
	return assignable(want.Return, got.Return)
}
