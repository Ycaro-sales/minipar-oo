package tac

import (
	"fmt"
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

// toString formats TAC's elements in string for easier debbuging
func (t TAC) toString() string {
	// if there is no 2nd argument, the operation is either a...
	if t.arg2 == "" {
		// ...label or goto...
		if t.op == "LABEL" || t.op == "GOTO" {
			return fmt.Sprint("%s %s", t.op, t.arg1)
		}
		// ...or an unary expression
		return fmt.Sprint("%s = %s %s", t.result, t.op, t.arg1)
	}

	// otherwise, is a binary expression
	return fmt.Sprint("%s = %s %s %s", t.result, t.arg1, t.op, t.arg2)
}

// ==========================================
//  CODE GENERATOR DEFINITION
// ==========================================

type Symbol struct {
	symbolName string
	symbolType string
}

type LoopLabels struct {
	startLabel string
	endLabel   string
}

type TACGenerator struct {
	instructions []TAC
	tempCount    int
	labelCount   int
	SymbolTable  map[string]Symbol
	loopStack    []LoopLabels
}

func (gen *TACGenerator) newTemp() string {
	temp := fmt.Sprintf("f%d", gen.tempCount)
	gen.tempCount++
	return temp
}

func (gen *TACGenerator) newLabel() string {
	label := fmt.Sprintf("f%d", gen.tempCount)
	gen.tempCount++
	return label
}

// opToStr converte a constante de operador da AST para um mnemônico puro no TAC
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

// emit generates a TAC struct and appends to the TACGenerator's instructions slice
func (gen *TACGenerator) emit(op, arg1, arg2, result string) {
	gen.instructions = append(gen.instructions, TAC{
		op:     op,
		arg1:   arg1,
		arg2:   arg2,
		result: result,
	})
}

func (gen *TACGenerator) Generate(node ast.Node) string {
	if node == nil {
		return ""
	}

	switch n := node.(type) {

	/* LITERALS - apenas retornam seus valores para a chamada supeior */
	case *ast.IntLiteral:
		return fmt.Sprint("%d", n.Value)

	case *ast.FloatLiteral:
		return fmt.Sprint("%f", n.Value)

	case *ast.CharLiteral:
		return fmt.Sprint("%c", n.Value)

	case *ast.StringLiteral:
		return fmt.Sprint("\"%s\"", n.Value)

	case *ast.BooleanLiteral:
		return fmt.Sprint("%b", n.Value)

	// TODO implement array literal

	// TODO implement dict literal

	// TODO implement set literal

	// TODO implement tuple literal

	case *ast.SelfExpr:
		return "self"

	case *ast.Identifier:
		return n.Value

	/* EXPRESSIONS */
	case *ast.UnaryExpr:
		return gen.genUnaryExpr(n)

	case *ast.BinaryExpr:
		return gen.genBinaryExpr(n)

	/* STATEMENTS */

	/* FUNCTIONS */

	/* DATA STRUCTURES */

	/* OBJECT ORIENTATION */

	default:
		return "ERRO"
	}
}

func (gen *TACGenerator) genUnaryExpr(node *ast.UnaryExpr) string {
	arg1 := gen.Generate(node.Right)
	op := gen.opToString(node.Operator)
	result := gen.newTemp()

	gen.emit(op, arg1, "", result)
	return result
}

func (gen *TACGenerator) genBinaryExpr(node *ast.BinaryExpr) string {
	arg1 := gen.Generate(node.Left)
	arg2 := gen.Generate(node.Right)
	op := gen.opToString(node.Operator)
	result := gen.newTemp()

	gen.emit(op, arg1, arg2, result)
	return result
}

// genAssignment emits a TAC for assignment operation
func (gen *TACGenerator) genAssignment(node *ast.Assignment) {
	arg1 := gen.Generate(node.Value)
	op := "ASSIGN"

	gen.emit(op, arg1, "", node.Name)
}

func (gen *TACGenerator) genVarDecl(node *ast.VarDecl) {

}
