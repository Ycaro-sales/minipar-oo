package ast

// ==========================================
// 1. INTERFACES BASE (Contratos)
// ==========================================

// Node representa qualquer nó da nossa Árvore Sintática Abstrata.
type Node interface {
	GetLine() int
}

// Statement representa um comando que executa uma ação, mas não gera um valor final.
type Statement interface {
	Node
	stmtNode()
}

// Expression representa uma instrução matemática, variável ou literal que devolve um valor.
type Expression interface {
	Node
	exprNode()
}

// ==========================================
// 2. RAIZ DO PROGRAMA
// ==========================================

// Program é o topo da árvore. Guarda tudo o que está no escopo global.
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
// 3. DECLARAÇÕES GLOBAIS
// ==========================================

type ClassDecl struct {
	Line    int
	Name    string
	Extends string
	Members []Node
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
// 4. COMANDOS (STATEMENTS)
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
	ChanType string // "c_channel" ou "s_channel"
	Name     string
}

func (c *ChannelStmt) GetLine() int { return c.Line }
func (c *ChannelStmt) stmtNode()    {}

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

type VarDecl struct {
	Line  int
	Type  string
	Name  string
	Value Expression // Opcional
}

func (v *VarDecl) GetLine() int { return v.Line }
func (v *VarDecl) stmtNode()    {}

type Assignment struct {
	Line  int
	Name  string
	Value Expression
}

func (a *Assignment) GetLine() int { return a.Line }
func (a *Assignment) stmtNode()    {}

type PrintStmt struct {
	Line int
	Args []Expression // Print pode receber múltiplos argumentos
}

func (p *PrintStmt) GetLine() int { return p.Line }
func (p *PrintStmt) stmtNode()    {}

type ReturnStmt struct {
	Line  int
	Value Expression // Opcional (pode ser um return vazio)
}

func (r *ReturnStmt) GetLine() int { return r.Line }
func (r *ReturnStmt) stmtNode()    {}

// ==========================================
// 5. EXPRESSÕES (EXPRESSIONS)
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
	Value string
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
