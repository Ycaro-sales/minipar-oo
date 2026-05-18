package ast

// ==========================================
// INTERFACES BASE
// ==========================================

type Node interface {
	GetLine() int
}

type Statement interface {
	Node
	stmtNode()
}

type Expression interface {
	Node
	exprNode()
}

// ==========================================
// A RAIZ DO PROGRAMA
// ==========================================

type Program struct {
	Declarations []Node
}

func (p *Program) GetLine() int {
	if len(p.Declarations) > 0 {
		return p.Declarations[0].GetLine()
	}
	return 0
}

// ==========================================
// CLASSES E FUNÇÕES
// ==========================================

type ClassDecl struct {
	Line    int
	Name    string
	Extends string
	Members []Node // Pode ser VarDecl ou MethodDecl
}

func (c *ClassDecl) GetLine() int { return c.Line }

type FuncDecl struct {
	Line       int
	Name       string
	ReturnType string
	Body       *BlockStmt
}

func (f *FuncDecl) GetLine() int { return f.Line }

// ==========================================
// REAL PARALELISMO E CANAIS
// ==========================================

type BlockStmt struct {
	Line       int
	Statements []Statement
}

func (b *BlockStmt) GetLine() int { return b.Line }
func (b *BlockStmt) stmtNode()    {}

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
	Comp1    string
	Comp2    string
}

func (c *ChannelStmt) GetLine() int { return c.Line }
func (c *ChannelStmt) stmtNode()    {}

// ==========================================
// FLUXO DE CONTROLE (IF / WHILE)
// ==========================================

type IfStmt struct {
	Line      int
	Condition Expression
	Block     *BlockStmt
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

// ==========================================
// VARIÁVEIS E ATRIBUIÇÕES
// ==========================================

type VarDecl struct {
	Line  int
	Type  string
	Name  string
	Value Expression // Pode ser nil
}

func (v *VarDecl) GetLine() int { return v.Line }
func (v *VarDecl) stmtNode()    {} // Serve como Statement também

type Assignment struct {
	Line  int
	Name  string
	Value Expression
}

func (a *Assignment) GetLine() int { return a.Line }
func (a *Assignment) stmtNode()    {}

// ==========================================
// EXPRESSÕES E LITERAIS
// ==========================================

type BinaryExpr struct {
	Line     int
	Left     Expression
	Operator string
	Right    Expression
}

func (b *BinaryExpr) GetLine() int { return b.Line }
func (b *BinaryExpr) exprNode()    {}

type Identifier struct {
	Line  int
	Value string
}

func (i *Identifier) GetLine() int { return i.Line }
func (i *Identifier) exprNode()    {}

type IntegerLiteral struct {
	Line  int
	Value string // Mantido como string para simplificar a conversão inicial
}

func (i *IntegerLiteral) GetLine() int { return i.Line }
func (i *IntegerLiteral) exprNode()    {}

type StringLiteral struct {
	Line  int
	Value string
}

func (s *StringLiteral) GetLine() int { return s.Line }
func (s *StringLiteral) exprNode()    {}

type BooleanLiteral struct {
	Line  int
	Value bool
}

func (b *BooleanLiteral) GetLine() int { return b.Line }
func (b *BooleanLiteral) exprNode()    {}
