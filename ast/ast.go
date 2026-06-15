package ast

// Op represents a typed operator, avoiding string comparisons in later passes.
type Op int

const (
	OpOr Op = iota
	OpAnd
	OpEq
	OpNeq
	OpLt
	OpGt
	OpLeq
	OpGeq
	OpAdd
	OpSub
	OpMul
	OpDiv
	OpMod
	OpNeg
	OpNot
)

// Node is any node in the AST.
type Node interface {
	GetLine() int
}

// Statement executes an action but does not produce a value.
type Statement interface {
	Node
	stmtNode()
}

// Expression produces a value.
type Expression interface {
	Node
	exprNode()
	// ResolvedType reports the type the semantic analyzer inferred for this
	// expression, as a canonical type-name string (empty until analysis runs).
	ResolvedType() string
	SetResolvedType(string)
}

// ExprMeta carries analysis results attached to an expression by later passes.
// Expression nodes embed it so the semantic analyzer can annotate them with a
// resolved type that code generation reads back. The type is stored as a string
// (the analyzer's Type.String()) to keep ast free of a semantic import.
type ExprMeta struct{ resolvedType string }

func (m *ExprMeta) ResolvedType() string     { return m.resolvedType }
func (m *ExprMeta) SetResolvedType(t string) { m.resolvedType = t }

// Declaration is a top-level declaration (class, interface, func, var).
type Declaration interface {
	Node
	declNode()
}

// ==========================================
// Program root
// ==========================================

type Program struct {
	Declarations []Declaration
}

func (p *Program) GetLine() int {
	if len(p.Declarations) > 0 {
		return p.Declarations[0].GetLine()
	}
	return 0
}

// ==========================================
// Shared types
// ==========================================

type Param struct {
	Name string
	Type string
}

// ==========================================
// Declarations
// ==========================================

type ClassDecl struct {
	Line       int
	Name       string
	Implements []string
	Members    []Node
}

func (c *ClassDecl) GetLine() int { return c.Line }
func (c *ClassDecl) declNode()    {}

type InterfaceDecl struct {
	Line    int
	Name    string
	Methods []*InterfaceMethod
}

func (i *InterfaceDecl) GetLine() int { return i.Line }
func (i *InterfaceDecl) declNode()    {}

type InterfaceMethod struct {
	Line       int
	Name       string
	Params     []Param
	ReturnType string
}

func (i *InterfaceMethod) GetLine() int { return i.Line }

type FuncDecl struct {
	Line       int
	Name       string
	Params     []Param
	ReturnType string
	Body       *BlockStmt
}

func (f *FuncDecl) GetLine() int { return f.Line }
func (f *FuncDecl) declNode()    {}

type VarDecl struct {
	Line  int
	Type  string
	Name  string
	Value Expression
}

func (v *VarDecl) GetLine() int { return v.Line }
func (v *VarDecl) declNode()    {}
func (v *VarDecl) stmtNode()    {}

// ==========================================
// Class members
// ==========================================

type FieldDecl struct {
	Line  int
	Name  string
	Type  string
	Value Expression
}

func (f *FieldDecl) GetLine() int { return f.Line }

type ConstructorDecl struct {
	Line int
	Name string
	Body *BlockStmt
}

func (c *ConstructorDecl) GetLine() int { return c.Line }

type MethodDecl struct {
	Line       int
	Name       string
	Params     []Param
	ReturnType string
	Body       *BlockStmt
}

func (m *MethodDecl) GetLine() int { return m.Line }

// ==========================================
// Statements
// ==========================================

type BlockStmt struct {
	Line       int
	Statements []Statement
}

func (b *BlockStmt) GetLine() int { return b.Line }
func (b *BlockStmt) stmtNode()    {}

type IfStmt struct {
	Line      int
	Condition Expression
	Then      *BlockStmt
	Else      *BlockStmt
}

func (i *IfStmt) GetLine() int { return i.Line }
func (i *IfStmt) stmtNode()    {}

type WhileStmt struct {
	Line      int
	Condition Expression
	Block     *BlockStmt
}

func (w *WhileStmt) GetLine() int { return w.Line }
func (w *WhileStmt) stmtNode()    {}

type DoStmt struct {
	Line      int
	Block     *BlockStmt
	Condition Expression
}

func (d *DoStmt) GetLine() int { return d.Line }
func (d *DoStmt) stmtNode()    {}

type ForStmt struct {
	Line  int
	Var   string
	Iter  Expression
	Block *BlockStmt
}

func (f *ForStmt) GetLine() int { return f.Line }
func (f *ForStmt) stmtNode()    {}

type SwitchStmt struct {
	Line  int
	Expr  Expression
	Cases []*CaseClause
}

func (s *SwitchStmt) GetLine() int { return s.Line }
func (s *SwitchStmt) stmtNode()    {}

type CaseClause struct {
	Line  int
	Value Expression
	Block *BlockStmt
}

func (c *CaseClause) GetLine() int { return c.Line }

type SeqStmt struct {
	Line  int
	Block *BlockStmt
}

func (s *SeqStmt) GetLine() int { return s.Line }
func (s *SeqStmt) stmtNode()    {}

type ParStmt struct {
	Line  int
	Block *BlockStmt
}

func (p *ParStmt) GetLine() int { return p.Line }
func (p *ParStmt) stmtNode()    {}

type ChannelStmt struct {
	Line     int
	ChanType string
	Name     string
	Args     []Expression
}

func (c *ChannelStmt) GetLine() int { return c.Line }
func (c *ChannelStmt) stmtNode()    {}
func (c *ChannelStmt) declNode()    {} // also satisfies Declaration (top-level channel)

type Assignment struct {
	Line  int
	Name  string
	Value Expression
}

func (a *Assignment) GetLine() int { return a.Line }
func (a *Assignment) stmtNode()    {}

// IndexAssignment represents arr[idx] = value, where the left-hand side is an
// index expression. Only valid on mutable arrays; tuples are immutable.
type IndexAssignment struct {
	Line   int
	Object Expression
	Index  Expression
	Value  Expression
}

func (ia *IndexAssignment) GetLine() int { return ia.Line }
func (ia *IndexAssignment) stmtNode()    {}

type PrintStmt struct {
	Line int
	Args []Expression
}

func (p *PrintStmt) GetLine() int { return p.Line }
func (p *PrintStmt) stmtNode()    {}

type InputStmt struct {
	Line   int
	Prompt Expression
}

func (i *InputStmt) GetLine() int { return i.Line }
func (i *InputStmt) stmtNode()    {}

type ReturnStmt struct {
	Line  int
	Value Expression
}

func (r *ReturnStmt) GetLine() int { return r.Line }
func (r *ReturnStmt) stmtNode()    {}

type BreakStmt struct{ Line int }

func (b *BreakStmt) GetLine() int { return b.Line }
func (b *BreakStmt) stmtNode()    {}

type ContinueStmt struct{ Line int }

func (c *ContinueStmt) GetLine() int { return c.Line }
func (c *ContinueStmt) stmtNode()    {}

type PassStmt struct{ Line int }

func (p *PassStmt) GetLine() int { return p.Line }
func (p *PassStmt) stmtNode()    {}

type GotoStmt struct {
	Line  int
	Label string
}

func (g *GotoStmt) GetLine() int { return g.Line }
func (g *GotoStmt) stmtNode()    {}

type ExprStmt struct {
	Line       int
	Expression Expression
}

func (e *ExprStmt) GetLine() int { return e.Line }
func (e *ExprStmt) stmtNode()    {}

// ==========================================
// Expressions
// ==========================================

type BinaryExpr struct {
	ExprMeta
	Line     int
	Left     Expression
	Operator Op
	Right    Expression
}

func (b *BinaryExpr) GetLine() int { return b.Line }
func (b *BinaryExpr) exprNode()    {}

type UnaryExpr struct {
	ExprMeta
	Line     int
	Operator Op
	Right    Expression
}

func (u *UnaryExpr) GetLine() int { return u.Line }
func (u *UnaryExpr) exprNode()    {}

type Identifier struct {
	ExprMeta
	Line  int
	Value string
}

func (i *Identifier) GetLine() int { return i.Line }
func (i *Identifier) exprNode()    {}

type FuncCall struct {
	ExprMeta
	Line int
	Name string
	Args []Expression
}

func (f *FuncCall) GetLine() int { return f.Line }
func (f *FuncCall) exprNode()    {}
func (f *FuncCall) stmtNode()    {}

type MethodCall struct {
	ExprMeta
	Line   int
	Object Expression
	Method string
	Args   []Expression
}

func (m *MethodCall) GetLine() int { return m.Line }
func (m *MethodCall) exprNode()    {}
func (m *MethodCall) stmtNode()    {}

type IndexExpr struct {
	ExprMeta
	Line   int
	Object Expression
	Index  Expression
}

func (i *IndexExpr) GetLine() int { return i.Line }
func (i *IndexExpr) exprNode()    {}

type ObjCreation struct {
	ExprMeta
	Line  int
	Class string
	Args  []Expression
}

func (o *ObjCreation) GetLine() int { return o.Line }
func (o *ObjCreation) exprNode()    {}

type FuncLiteral struct {
	ExprMeta
	Line       int
	Params     []Param
	ReturnType string
	Body       *BlockStmt
}

func (f *FuncLiteral) GetLine() int { return f.Line }
func (f *FuncLiteral) exprNode()    {}

// ==========================================
// Literals
// ==========================================

type IntLiteral struct {
	ExprMeta
	Line  int
	Value int64
}

func (i *IntLiteral) GetLine() int { return i.Line }
func (i *IntLiteral) exprNode()    {}

type FloatLiteral struct {
	ExprMeta
	Line  int
	Value float64
}

func (f *FloatLiteral) GetLine() int { return f.Line }
func (f *FloatLiteral) exprNode()    {}

type StringLiteral struct {
	ExprMeta
	Line  int
	Value string
}

func (s *StringLiteral) GetLine() int { return s.Line }
func (s *StringLiteral) exprNode()    {}

type CharLiteral struct {
	ExprMeta
	Line  int
	Value rune
}

func (c *CharLiteral) GetLine() int { return c.Line }
func (c *CharLiteral) exprNode()    {}

type BooleanLiteral struct {
	ExprMeta
	Line  int
	Value bool
}

func (b *BooleanLiteral) GetLine() int { return b.Line }
func (b *BooleanLiteral) exprNode()    {}

type ListLiteral struct {
	ExprMeta
	Line     int
	Elements []Expression
}

func (l *ListLiteral) GetLine() int { return l.Line }
func (l *ListLiteral) exprNode()    {}

// TupleLiteral is an immutable fixed-size collection with heterogeneous element
// types, written (e1, e2, ...). Access is by constant integer index.
type TupleLiteral struct {
	ExprMeta
	Line     int
	Elements []Expression
}

func (t *TupleLiteral) GetLine() int { return t.Line }
func (t *TupleLiteral) exprNode()    {}

type DictLiteral struct {
	ExprMeta
	Line  int
	Pairs map[Expression]Expression
}

func (d *DictLiteral) GetLine() int { return d.Line }
func (d *DictLiteral) exprNode()    {}

type SelfExpr struct {
	ExprMeta
	Line int
}

func (s *SelfExpr) GetLine() int { return s.Line }
func (s *SelfExpr) exprNode()    {}
