package tac

import (
	"fmt"
	"strings"

	"minipar/ast"
)

// ==========================================
// TAC DEFINITION
// ==========================================

// Struct for storing TAC's temporary variable for instruction result, operator
// and arguments
type TAC struct {
	op     string
	arg1   string
	arg2   string
	result string
}

// String formats TAC's elements in string for easier debugging.
func (t TAC) String() string {
	// if there is no 2nd argument, the operation is either a...
	if t.arg2 == "" {
		// ...label or goto...
		if t.op == "LABEL" || t.op == "GOTO" {
			return fmt.Sprintf("%s %s", t.op, t.arg1)
		}
		// ...or an unary expression
		return fmt.Sprintf("%s %s -> %s", t.op, t.arg1, t.result)
	}

	// otherwise, is a binary expression
	return fmt.Sprintf("%s %s %s -> %s", t.op, t.arg1, t.arg2, t.result)
}

// ==========================================
//  CODE GENERATOR DEFINITION
// ==========================================

type LoopLabels struct {
	startLabel string
	endLabel   string
}

type TACGenerator struct {
	instructions []TAC
	tempCount    int
	labelCount   int
	tempTypes    map[string]string // temp name -> resolved type, from the decorated AST
	loopStack    []LoopLabels
}

// New returns a ready-to-use TACGenerator.
func New() *TACGenerator {
	return &TACGenerator{tempTypes: map[string]string{}}
}

// TempType returns the resolved type recorded for a temporary, or "" if unknown.
func (gen *TACGenerator) TempType(name string) string { return gen.tempTypes[name] }

// TempTypes returns a copy of the temp-name → resolved-type map for use by
// downstream passes (e.g. the C code generator).
func (gen *TACGenerator) TempTypes() map[string]string {
	out := make(map[string]string, len(gen.tempTypes))
	for k, v := range gen.tempTypes {
		out[k] = v
	}
	return out
}

// Instruction is the exported view of a single TAC instruction, for use by
// downstream passes (e.g. the C code generator).
type Instruction struct {
	Op     string
	Arg1   string
	Arg2   string
	Result string
}

// RawInstructions returns the generated instructions as a slice of Instruction,
// for use by downstream passes that need structured access rather than the
// formatted string from Instructions().
func (gen *TACGenerator) RawInstructions() []Instruction {
	out := make([]Instruction, len(gen.instructions))
	for i, t := range gen.instructions {
		out[i] = Instruction{Op: t.op, Arg1: t.arg1, Arg2: t.arg2, Result: t.result}
	}
	return out
}

// Instructions returns the generated TAC as one instruction per line.
func (gen *TACGenerator) Instructions() string {
	var b strings.Builder
	for _, ins := range gen.instructions {
		b.WriteString(ins.String())
		b.WriteByte('\n')
	}
	return b.String()
}

// newTemp creates a new temporary variable and returns it
func (gen *TACGenerator) newTemp() string {
	temp := fmt.Sprintf("t%d", gen.tempCount)
	gen.tempCount++
	return temp
}

// newTypedTemp creates a temporary and records its resolved type (when known),
// so later passes can do type-aware work without re-deriving types.
func (gen *TACGenerator) newTypedTemp(typ string) string {
	temp := gen.newTemp()
	if typ != "" {
		gen.tempTypes[temp] = typ
	}
	return temp
}

// newLabel creates a new label and returns it.
// Used in loops and functions
func (gen *TACGenerator) newLabel() string {
	label := fmt.Sprintf("L%d", gen.labelCount)
	gen.labelCount++
	return label
}

// opToStr converts an AST operator constant to a TAC mnemonic
func (gen *TACGenerator) opToString(op ast.Op) string {
	switch op {
	// Operadores Lógicos
	case ast.OpOr:
		return "OR"
	case ast.OpAnd:
		return "AND"
	case ast.OpNot:
		return "NOT"

	// Operadores Relacionais (Comparações)
	case ast.OpEq:
		return "EQ"
	case ast.OpNeq:
		return "NEQ"
	case ast.OpLt:
		return "LT" // Less Than
	case ast.OpGt:
		return "GT" // Greater Than
	case ast.OpLeq:
		return "LEQ" // Less or Equal
	case ast.OpGeq:
		return "GEQ" // Greater or Equal

	// Operadores Aritméticos
	case ast.OpAdd:
		return "ADD"
	case ast.OpSub:
		return "SUB"
	case ast.OpMul:
		return "MUL"
	case ast.OpDiv:
		return "DIV"
	case ast.OpMod:
		return "MOD"
	case ast.OpNeg:
		return "NEG" // Menos unário (negativo)

	default:
		return "UNKNOWN"
	}
}

func (gen *TACGenerator) getZeroForType(nodeType string) string {
	switch nodeType {
	case "i8", "i16", "i32", "i64", "u8", "u16", "u32", "u64":
		return "0"
	case "f16", "f32", "f64":
		return "0.0"
	case "string":
		return "\"\""
	case "char":
		return "''"
	case "bool":
		return "false"
	default:
		return ""
	}
}

// emit generates a TAC struct and appends it to the TACGenerator's instructions slice
func (gen *TACGenerator) emit(op, arg1, arg2, result string) {
	gen.instructions = append(gen.instructions, TAC{
		op:     op,
		arg1:   arg1,
		arg2:   arg2,
		result: result,
	})
}

// ==========================================
// GENERATE — main dispatcher
// ==========================================

// Generate receives an AST Node and generates corresponding TAC instructions.
// When the AST Node is an Expression, Generate also returns the temp variable
// which contains the expression's result.
func (gen *TACGenerator) Generate(node ast.Node) string {
	if node == nil {
		return ""
	}

	switch n := node.(type) {

	/* PROGRAM */
	case *ast.Program:
		return gen.genProgram(n)

	/* LITERALS - apenas retornam seus valores para a chamada superior */
	case *ast.IntLiteral:
		return fmt.Sprintf("%d", n.Value)

	case *ast.FloatLiteral:
		return fmt.Sprintf("%f", n.Value)

	case *ast.CharLiteral:
		return fmt.Sprintf("%c", n.Value)

	case *ast.StringLiteral:
		return fmt.Sprintf("\"%s\"", n.Value)

	case *ast.BooleanLiteral:
		return fmt.Sprintf("%t", n.Value)

	case *ast.ListLiteral:
		return gen.genListLiteral(n)

	case *ast.TupleLiteral:
		return gen.genTupleLiteral(n)

	// TODO implement dict literal

	// TODO implement set literal

	// TODO implement function literal

	case *ast.SelfExpr:
		return "self"

	case *ast.Identifier:
		return n.Value

	/* EXPRESSIONS */
	case *ast.UnaryExpr:
		return gen.genUnaryExpr(n)

	case *ast.BinaryExpr:
		return gen.genBinaryExpr(n)

	case *ast.FuncCall:
		return gen.genFuncCall(n)

	case *ast.MethodCall:
		return gen.genMethodCall(n)

	case *ast.IndexExpr:
		return gen.genIndexExpr(n)

	case *ast.ObjCreation:
		return gen.genObjCreation(n)

	case *ast.ExprStmt:
		return gen.genExprStmt(n)

	/* STATEMENTS */
	case *ast.Assignment:
		gen.genAssignment(n)
		return ""

	case *ast.IndexAssignment:
		gen.genIndexAssignment(n)
		return ""

	case *ast.VarDecl:
		gen.genVarDecl(n)
		return ""

	case *ast.BlockStmt:
		return gen.genBlock(n)

	case *ast.IfStmt:
		return gen.genIfStmt(n)

	case *ast.WhileStmt:
		return gen.genWhileStmt(n)

	case *ast.DoStmt:
		return gen.genDoStmt(n)

	case *ast.ForStmt:
		return gen.genForStmt(n)

	case *ast.SwitchStmt:
		return gen.genSwitchStmt(n)

	case *ast.BreakStmt:
		return gen.genBreakStmt(n)

	case *ast.ContinueStmt:
		return gen.genContinueStmt(n)

	case *ast.PassStmt:
		return gen.genPassStmt(n)

	case *ast.GotoStmt:
		return gen.genGotoStmt(n)

	case *ast.ReturnStmt:
		return gen.genReturnStmt(n)

	case *ast.PrintStmt:
		return gen.genPrintStmt(n)

	case *ast.InputStmt:
		return gen.genInputStmt(n)

	case *ast.SeqStmt:
		return gen.genSeqStmt(n)

	case *ast.ParStmt:
		return gen.genParStmt(n)

	case *ast.ChannelStmt:
		return gen.genChannelStmt(n)

	/* DECLARATIONS */
	case *ast.FuncDecl:
		return gen.genFuncDecl(n)

	case *ast.ClassDecl:
		return gen.genClassDecl(n)

	case *ast.InterfaceDecl:
		return gen.genInterfaceDecl(n)

	default:
		return "ERRO"
	}
}

// ==========================================
// EXPRESSIONS
// ==========================================

func (gen *TACGenerator) genUnaryExpr(node *ast.UnaryExpr) string {
	arg1 := gen.Generate(node.Right)
	op := gen.opToString(node.Operator)
	result := gen.newTypedTemp(node.ResolvedType())

	gen.emit(op, arg1, "", result)
	return result
}

func (gen *TACGenerator) genBinaryExpr(node *ast.BinaryExpr) string {
	arg1 := gen.Generate(node.Left)
	arg2 := gen.Generate(node.Right)
	op := gen.opToString(node.Operator)
	result := gen.newTypedTemp(node.ResolvedType())

	gen.emit(op, arg1, arg2, result)
	return result
}

func (gen *TACGenerator) genFuncCall(node *ast.FuncCall) string {
	for _, arg := range node.Args {
		gen.emit("PARAM", gen.Generate(arg), "", "")
	}
	result := gen.newTypedTemp(node.ResolvedType())
	gen.emit("CALL", node.Name, fmt.Sprint(len(node.Args)), result)
	return result
}

func (gen *TACGenerator) genMethodCall(node *ast.MethodCall) string {
	obj := gen.Generate(node.Object)
	for _, arg := range node.Args {
		gen.emit("PARAM", gen.Generate(arg), "", "")
	}
	result := gen.newTypedTemp(node.ResolvedType())
	gen.emit("METHOD_CALL", obj+"."+node.Method, fmt.Sprint(len(node.Args)), result)
	return result
}

func (gen *TACGenerator) genIndexExpr(node *ast.IndexExpr) string {
	obj := gen.Generate(node.Object)
	idx := gen.Generate(node.Index)
	result := gen.newTypedTemp(node.ResolvedType())
	gen.emit("ARRAY_GET", obj, idx, result)
	return result
}

func (gen *TACGenerator) genListLiteral(node *ast.ListLiteral) string {
	count := fmt.Sprintf("%d", len(node.Elements))
	elemType := ""
	if rt := node.ResolvedType(); len(rt) > 2 {
		elemType = rt[1 : len(rt)-1] // strip "[" and "]"
	}
	result := gen.newTypedTemp(node.ResolvedType())
	gen.emit("ARRAY_NEW", count, elemType, result)
	for i, elem := range node.Elements {
		val := gen.Generate(elem)
		gen.emit("ARRAY_SET", result, fmt.Sprintf("%d", i), val)
	}
	return result
}

func (gen *TACGenerator) genTupleLiteral(node *ast.TupleLiteral) string {
	count := fmt.Sprintf("%d", len(node.Elements))
	result := gen.newTypedTemp(node.ResolvedType())
	gen.emit("TUPLE_NEW", count, "", result)
	for i, elem := range node.Elements {
		val := gen.Generate(elem)
		gen.emit("TUPLE_SET", result, fmt.Sprintf("%d", i), val)
	}
	return result
}

func (gen *TACGenerator) genIndexAssignment(node *ast.IndexAssignment) {
	obj := gen.Generate(node.Object)
	idx := gen.Generate(node.Index)
	val := gen.Generate(node.Value)
	gen.emit("ARRAY_SET", obj, idx, val)
}

func (gen *TACGenerator) genObjCreation(node *ast.ObjCreation) string {
	for _, arg := range node.Args {
		gen.emit("PARAM", gen.Generate(arg), "", "")
	}
	result := gen.newTypedTemp(node.ResolvedType())
	gen.emit("NEW_OBJ", node.Class, fmt.Sprint(len(node.Args)), result)
	return result
}

func (gen *TACGenerator) genExprStmt(node *ast.ExprStmt) string {
	return gen.Generate(node.Expression)
}

// ==========================================
// STATEMENTS
// ==========================================

// genAssignment emits a TAC for assignment operation
func (gen *TACGenerator) genAssignment(node *ast.Assignment) {
	arg1 := gen.Generate(node.Value)
	gen.emit("ASSIGN", arg1, "", node.Name)
}

func (gen *TACGenerator) genVarDecl(node *ast.VarDecl) {
	if node.Value != nil {
		initializer := gen.Generate(node.Value)
		// For composite types ([T] and (T0,...)), propagate the declared type
		// to the initializer temp. This ensures cgen uses "[i32]" rather than
		// the literal-inferred "[int]" when emitting struct typedefs.
		if len(node.Type) >= 2 && (node.Type[0] == '[' || node.Type[0] == '(') {
			if _, isTemp := gen.tempTypes[initializer]; isTemp {
				gen.tempTypes[initializer] = node.Type
			}
		}
		gen.emit("ASSIGN", initializer, "", node.Name)
	} else {
		zeroInitiation := gen.getZeroForType(node.Type)
		gen.emit("ASSIGN", zeroInitiation, "", node.Name)
	}
}

func (gen *TACGenerator) genProgram(node *ast.Program) string {
	for _, decl := range node.Declarations {
		gen.Generate(decl)
	}
	return ""
}

func (gen *TACGenerator) genBlock(node *ast.BlockStmt) string {
	for _, stmt := range node.Statements {
		gen.Generate(stmt)
	}
	return ""
}

func (gen *TACGenerator) genIfStmt(node *ast.IfStmt) string {
	elseLabel := gen.newLabel()
	endLabel := gen.newLabel()

	cond := gen.Generate(node.Condition)
	gen.emit("IF_FALSE", cond, "", elseLabel)
	gen.genBlock(node.Then)
	gen.emit("GOTO", endLabel, "", "")
	gen.emit("LABEL", elseLabel, "", "")
	if node.Else != nil {
		gen.genBlock(node.Else)
	}
	gen.emit("LABEL", endLabel, "", "")
	return ""
}

func (gen *TACGenerator) genWhileStmt(node *ast.WhileStmt) string {
	startLabel := gen.newLabel()
	endLabel := gen.newLabel()

	gen.emit("LABEL", startLabel, "", "")
	cond := gen.Generate(node.Condition)
	gen.emit("IF_FALSE", cond, "", endLabel)

	gen.loopStack = append(gen.loopStack, LoopLabels{startLabel: startLabel, endLabel: endLabel})
	gen.genBlock(node.Block)
	gen.loopStack = gen.loopStack[:len(gen.loopStack)-1]

	gen.emit("GOTO", startLabel, "", "")
	gen.emit("LABEL", endLabel, "", "")
	return ""
}

func (gen *TACGenerator) genDoStmt(node *ast.DoStmt) string {
	startLabel := gen.newLabel()
	endLabel := gen.newLabel()

	gen.emit("LABEL", startLabel, "", "")

	gen.loopStack = append(gen.loopStack, LoopLabels{startLabel: startLabel, endLabel: endLabel})
	gen.genBlock(node.Block)
	gen.loopStack = gen.loopStack[:len(gen.loopStack)-1]

	cond := gen.Generate(node.Condition)
	gen.emit("IF_FALSE", cond, "", endLabel)
	gen.emit("GOTO", startLabel, "", "")
	gen.emit("LABEL", endLabel, "", "")
	return ""
}

func (gen *TACGenerator) genForStmt(node *ast.ForStmt) string {
	startLabel := gen.newLabel()
	endLabel := gen.newLabel()

	iter := gen.Generate(node.Iter)
	idx := gen.newTypedTemp("int")
	gen.emit("ASSIGN", "0", "", idx)
	length := gen.newTypedTemp("int")
	gen.emit("ARRAY_LEN", iter, "", length)

	gen.emit("LABEL", startLabel, "", "")
	cond := gen.newTypedTemp("bool")
	gen.emit("LT", idx, length, cond)
	gen.emit("IF_FALSE", cond, "", endLabel)

	elem := gen.newTemp()
	gen.emit("ARRAY_GET", iter, idx, elem)
	gen.emit("ASSIGN", elem, "", node.Var)

	gen.loopStack = append(gen.loopStack, LoopLabels{startLabel: startLabel, endLabel: endLabel})
	gen.genBlock(node.Block)
	gen.loopStack = gen.loopStack[:len(gen.loopStack)-1]

	gen.emit("ADD", idx, "1", idx)
	gen.emit("GOTO", startLabel, "", "")
	gen.emit("LABEL", endLabel, "", "")
	return ""
}

func (gen *TACGenerator) genSwitchStmt(node *ast.SwitchStmt) string {
	switchVal := gen.Generate(node.Expr)
	endLabel := gen.newLabel()

	for _, c := range node.Cases {
		caseVal := gen.Generate(c.Value)
		cond := gen.newTypedTemp("bool")
		gen.emit("EQ", switchVal, caseVal, cond)
		nextLabel := gen.newLabel()
		gen.emit("IF_FALSE", cond, "", nextLabel)
		gen.genBlock(c.Block)
		gen.emit("GOTO", endLabel, "", "")
		gen.emit("LABEL", nextLabel, "", "")
	}

	gen.emit("LABEL", endLabel, "", "")
	return ""
}

func (gen *TACGenerator) genBreakStmt(node *ast.BreakStmt) string {
	if len(gen.loopStack) == 0 {
		return ""
	}
	top := gen.loopStack[len(gen.loopStack)-1]
	gen.emit("GOTO", top.endLabel, "", "")
	return ""
}

func (gen *TACGenerator) genContinueStmt(node *ast.ContinueStmt) string {
	if len(gen.loopStack) == 0 {
		return ""
	}
	top := gen.loopStack[len(gen.loopStack)-1]
	gen.emit("GOTO", top.startLabel, "", "")
	return ""
}

func (gen *TACGenerator) genPassStmt(node *ast.PassStmt) string {
	return ""
}

func (gen *TACGenerator) genGotoStmt(node *ast.GotoStmt) string {
	gen.emit("GOTO", node.Label, "", "")
	return ""
}

func (gen *TACGenerator) genReturnStmt(node *ast.ReturnStmt) string {
	if node.Value != nil {
		gen.emit("RETURN", gen.Generate(node.Value), "", "")
	} else {
		gen.emit("RETURN", "", "", "")
	}
	return ""
}

func (gen *TACGenerator) genPrintStmt(node *ast.PrintStmt) string {
	for _, arg := range node.Args {
		gen.emit("PRINT", gen.Generate(arg), "", "")
	}
	return ""
}

func (gen *TACGenerator) genInputStmt(node *ast.InputStmt) string {
	prompt := ""
	if node.Prompt != nil {
		prompt = gen.Generate(node.Prompt)
	}
	result := gen.newTemp()
	gen.emit("INPUT", prompt, "", result)
	return result
}

func (gen *TACGenerator) genSeqStmt(node *ast.SeqStmt) string {
	gen.emit("BEGIN_SEQ", "", "", "")
	gen.genBlock(node.Block)
	gen.emit("END_SEQ", "", "", "")
	return ""
}

func (gen *TACGenerator) genParStmt(node *ast.ParStmt) string {
	gen.emit("BEGIN_PAR", "", "", "")
	gen.genBlock(node.Block)
	gen.emit("END_PAR", "", "", "")
	return ""
}

func (gen *TACGenerator) genChannelStmt(node *ast.ChannelStmt) string {
	for _, arg := range node.Args {
		gen.emit("PARAM", gen.Generate(arg), "", "")
	}
	gen.emit("CHAN_DECL", node.ChanType, node.Name, fmt.Sprint(len(node.Args)))
	return ""
}

// ==========================================
// FUNCTIONS AND DECLARATIONS
// ==========================================

func (gen *TACGenerator) genFuncDecl(node *ast.FuncDecl) string {
	gen.emit("BEGIN_FUNC", node.Name, node.ReturnType, "")
	for _, param := range node.Params {
		gen.emit("PARAM_DECL", param.Name, param.Type, "")
	}
	gen.genBlock(node.Body)
	gen.emit("END_FUNC", node.Name, "", "")
	return ""
}

// ==========================================
// OBJECT ORIENTATION
// ==========================================

func (gen *TACGenerator) genClassDecl(node *ast.ClassDecl) string {
	gen.emit("BEGIN_CLASS", node.Name, strings.Join(node.Implements, ","), "")
	for _, member := range node.Members {
		switch m := member.(type) {
		case *ast.FieldDecl:
			gen.genFieldDecl(m)
		case *ast.ConstructorDecl:
			gen.genConstructorDecl(m)
		case *ast.MethodDecl:
			gen.genMethodDecl(m)
		}
	}
	gen.emit("END_CLASS", node.Name, "", "")
	return ""
}

func (gen *TACGenerator) genFieldDecl(node *ast.FieldDecl) string {
	val := ""
	if node.Value != nil {
		val = gen.Generate(node.Value)
	}
	gen.emit("FIELD", node.Name, node.Type, val)
	return ""
}

func (gen *TACGenerator) genConstructorDecl(node *ast.ConstructorDecl) string {
	gen.emit("BEGIN_CTOR", node.Name, "", "")
	gen.genBlock(node.Body)
	gen.emit("END_CTOR", node.Name, "", "")
	return ""
}

func (gen *TACGenerator) genMethodDecl(node *ast.MethodDecl) string {
	gen.emit("BEGIN_METHOD", node.Name, node.ReturnType, "")
	for _, param := range node.Params {
		gen.emit("PARAM_DECL", param.Name, param.Type, "")
	}
	gen.genBlock(node.Body)
	gen.emit("END_METHOD", node.Name, "", "")
	return ""
}

func (gen *TACGenerator) genInterfaceDecl(node *ast.InterfaceDecl) string {
	gen.emit("INTERFACE", node.Name, "", "")
	return ""
}
