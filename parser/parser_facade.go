package parser

import (
	"fmt"
	"strconv"
	"strings"

	"minipar/ast"
	"minipar/lexer"
)

type miniparParser struct {
	factory      LexerFactory
	l            lexer.Tokenizer
	currentToken lexer.Token
	peekToken    lexer.Token
	errors       []string
}

func (p *miniparParser) ParseProgram(src string) (*ast.Program, []string) {
	p.l = p.factory(src)
	p.errors = []string{}
	p.nextToken()
	p.nextToken()

	program := &ast.Program{Declarations: []ast.Declaration{}}

	for !p.currentTokenIs(lexer.EOF) {
		decl := p.parseDeclaration()
		if decl != nil {
			program.Declarations = append(program.Declarations, decl)
		}
		p.nextToken()
	}

	return program, p.errors
}

// ==========================================
// Token helpers
// ==========================================

func (p *miniparParser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *miniparParser) currentTokenIs(t lexer.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *miniparParser) peekTokenIs(t lexer.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *miniparParser) expectPeek(t lexer.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.errors = append(p.errors, fmt.Sprintf(
		"linha %d: esperava '%s', recebi '%s' ('%s')",
		p.peekToken.Line, t, p.peekToken.Type, p.peekToken.Literal,
	))
	return false
}

// synchronize skips tokens until a safe recovery point.
func (p *miniparParser) synchronize() {
	for !p.currentTokenIs(lexer.EOF) {
		if p.currentTokenIs(lexer.SEMICOLON) {
			p.nextToken()
			return
		}
		switch p.peekToken.Type {
		case lexer.CLASS, lexer.INTERFACE, lexer.FUNC, lexer.LET,
			lexer.IF, lexer.WHILE, lexer.FOR, lexer.DO,
			lexer.RETURN, lexer.SEQ, lexer.PAR:
			return
		}
		p.nextToken()
	}
}

// ==========================================
// Declarations
// ==========================================

func (p *miniparParser) parseDeclaration() ast.Declaration {
	switch p.currentToken.Type {
	case lexer.CLASS:
		return p.parseClassDecl()
	case lexer.INTERFACE:
		return p.parseInterfaceDecl()
	case lexer.FUNC:
		return p.parseFuncDecl()
	case lexer.LET:
		return p.parseVarDecl()
	case lexer.C_CHANNEL, lexer.S_CHANNEL:
		return p.parseChannelAsDecl()
	default:
		p.errors = append(p.errors, fmt.Sprintf(
			"linha %d: declaração inválida: '%s'", p.currentToken.Line, p.currentToken.Literal,
		))
		p.synchronize()
		return nil
	}
}

// parseVarDecl: "let" <id> (":" <type>)? "=" <expr> ";"
func (p *miniparParser) parseVarDecl() *ast.VarDecl {
	line := p.currentToken.Line

	if !p.expectPeek(lexer.IDENT) {
		p.synchronize()
		return nil
	}
	name := p.currentToken.Literal

	var typeName string
	if p.peekTokenIs(lexer.COLON) {
		p.nextToken() // ":"
		p.nextToken() // type token
		typeName = p.currentToken.Literal
	}

	if !p.expectPeek(lexer.ASSIGN) {
		p.synchronize()
		return nil
	}
	p.nextToken()

	value := p.parseExpression()
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}

	return &ast.VarDecl{Line: line, Name: name, Type: typeName, Value: value}
}

// parseFieldDecl: <id> ":" <type> ("=" <expr>)? ";"?
func (p *miniparParser) parseFieldDecl() *ast.FieldDecl {
	line := p.currentToken.Line
	name := p.currentToken.Literal

	if !p.expectPeek(lexer.COLON) {
		return nil
	}
	p.nextToken()
	typeName := p.currentToken.Literal

	var value ast.Expression
	if p.peekTokenIs(lexer.ASSIGN) {
		p.nextToken()
		p.nextToken()
		value = p.parseExpression()
	}
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return &ast.FieldDecl{Line: line, Name: name, Type: typeName, Value: value}
}

func (p *miniparParser) parseClassDecl() *ast.ClassDecl {
	line := p.currentToken.Line
	node := &ast.ClassDecl{Members: []ast.Node{}}

	if !p.expectPeek(lexer.IDENT) {
		return nil
	}
	node.Line = line
	node.Name = p.currentToken.Literal

	if p.peekTokenIs(lexer.IMPLEMENTS) {
		p.nextToken()
		if !p.expectPeek(lexer.LPAREN) {
			return nil
		}
		for !p.peekTokenIs(lexer.RPAREN) && !p.peekTokenIs(lexer.EOF) {
			p.nextToken()
			if p.currentTokenIs(lexer.IDENT) {
				node.Implements = append(node.Implements, p.currentToken.Literal)
			}
			if p.peekTokenIs(lexer.COMMA) {
				p.nextToken()
			}
		}
		if !p.expectPeek(lexer.RPAREN) {
			return nil
		}
	}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	p.nextToken()

	for !p.currentTokenIs(lexer.RBRACE) && !p.currentTokenIs(lexer.EOF) {
		var member ast.Node
		if p.currentTokenIs(lexer.FUNC) {
			member = p.parseMethodDecl()
		} else if p.currentTokenIs(lexer.IDENT) && p.peekTokenIs(lexer.LBRACE) {
			member = p.parseConstructorDecl()
		} else if p.currentTokenIs(lexer.IDENT) {
			member = p.parseFieldDecl()
		}
		if member != nil {
			node.Members = append(node.Members, member)
		}
		p.nextToken()
	}

	return node
}

func (p *miniparParser) parseConstructorDecl() *ast.ConstructorDecl {
	node := &ast.ConstructorDecl{Line: p.currentToken.Line, Name: p.currentToken.Literal}
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	node.Body = p.parseBlock()
	return node
}

func (p *miniparParser) parseMethodDecl() *ast.MethodDecl {
	line := p.currentToken.Line

	if !p.expectPeek(lexer.IDENT) {
		return nil
	}
	name := p.currentToken.Literal

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	params := p.parseParams()
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	var returnType string
	if p.peekTokenIs(lexer.ARROW) {
		p.nextToken()
		p.nextToken()
		returnType = p.currentToken.Literal
	}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	body := p.parseBlock()
	return &ast.MethodDecl{Line: line, Name: name, Params: params, ReturnType: returnType, Body: body}
}

func (p *miniparParser) parseInterfaceDecl() *ast.InterfaceDecl {
	line := p.currentToken.Line

	if !p.expectPeek(lexer.IDENT) {
		return nil
	}
	node := &ast.InterfaceDecl{Line: line, Name: p.currentToken.Literal}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	p.nextToken()

	for !p.currentTokenIs(lexer.RBRACE) && !p.currentTokenIs(lexer.EOF) {
		if p.currentTokenIs(lexer.FUNC) {
			m := p.parseInterfaceMethod()
			if m != nil {
				node.Methods = append(node.Methods, m)
			}
		}
		p.nextToken()
	}
	return node
}

func (p *miniparParser) parseInterfaceMethod() *ast.InterfaceMethod {
	line := p.currentToken.Line

	if !p.expectPeek(lexer.IDENT) {
		return nil
	}
	name := p.currentToken.Literal

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	params := p.parseParams()
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	var returnType string
	if p.peekTokenIs(lexer.ARROW) {
		p.nextToken()
		p.nextToken()
		returnType = p.currentToken.Literal
	}

	return &ast.InterfaceMethod{Line: line, Name: name, Params: params, ReturnType: returnType}
}

func (p *miniparParser) parseFuncDecl() *ast.FuncDecl {
	line := p.currentToken.Line

	if !p.expectPeek(lexer.IDENT) {
		return nil
	}
	name := p.currentToken.Literal

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	params := p.parseParams()
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	var returnType string
	if p.peekTokenIs(lexer.ARROW) {
		p.nextToken()
		p.nextToken()
		returnType = p.currentToken.Literal
	}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	body := p.parseBlock()
	return &ast.FuncDecl{Line: line, Name: name, Params: params, ReturnType: returnType, Body: body}
}

func (p *miniparParser) parseChannelAsDecl() *ast.VarDecl {
	stmt := p.parseChannelStmt()
	if stmt == nil {
		return nil
	}
	return &ast.VarDecl{Line: stmt.Line, Name: stmt.Name, Type: stmt.ChanType}
}

// ==========================================
// Params
// ==========================================

func (p *miniparParser) parseParams() []ast.Param {
	var params []ast.Param
	if p.peekTokenIs(lexer.RPAREN) {
		return params
	}
	p.nextToken()
	for {
		if !p.currentTokenIs(lexer.IDENT) && p.currentToken.Type != lexer.SELF {
			break
		}
		param := ast.Param{Name: p.currentToken.Literal}
		if p.peekTokenIs(lexer.COLON) {
			p.nextToken()
			p.nextToken()
			param.Type = p.currentToken.Literal
		}
		params = append(params, param)
		if !p.peekTokenIs(lexer.COMMA) {
			break
		}
		p.nextToken()
		p.nextToken()
	}
	return params
}

func (p *miniparParser) parseArgs() []ast.Expression {
	var args []ast.Expression
	if p.peekTokenIs(lexer.RPAREN) {
		return args
	}
	p.nextToken()
	args = append(args, p.parseExpression())
	for p.peekTokenIs(lexer.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression())
	}
	return args
}

// ==========================================
// Block & Statements
// ==========================================

func (p *miniparParser) parseBlock() *ast.BlockStmt {
	block := &ast.BlockStmt{Line: p.currentToken.Line}
	p.nextToken()

	for !p.currentTokenIs(lexer.RBRACE) && !p.currentTokenIs(lexer.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

func (p *miniparParser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case lexer.LET:
		return p.parseVarDecl()
	case lexer.IF:
		return p.parseIfStmt()
	case lexer.WHILE:
		return p.parseWhileStmt()
	case lexer.DO:
		return p.parseDoStmt()
	case lexer.FOR:
		return p.parseForStmt()
	case lexer.SWITCH:
		return p.parseSwitchStmt()
	case lexer.SEQ:
		return p.parseSeqStmt()
	case lexer.PAR:
		return p.parseParStmt()
	case lexer.S_CHANNEL, lexer.C_CHANNEL:
		return p.parseChannelStmt()
	case lexer.PRINT:
		return p.parsePrintStmt()
	case lexer.INPUT:
		return p.parseInputStmt()
	case lexer.RETURN:
		return p.parseReturnStmt()
	case lexer.BREAK:
		s := &ast.BreakStmt{Line: p.currentToken.Line}
		if p.peekTokenIs(lexer.SEMICOLON) {
			p.nextToken()
		}
		return s
	case lexer.CONTINUE:
		s := &ast.ContinueStmt{Line: p.currentToken.Line}
		if p.peekTokenIs(lexer.SEMICOLON) {
			p.nextToken()
		}
		return s
	case lexer.PASS:
		s := &ast.PassStmt{Line: p.currentToken.Line}
		if p.peekTokenIs(lexer.SEMICOLON) {
			p.nextToken()
		}
		return s
	case lexer.GOTO:
		return p.parseGotoStmt()
	case lexer.FUNC:
		decl := p.parseFuncDecl()
		if decl == nil {
			return nil
		}
		return &ast.ExprStmt{Line: decl.Line, Expression: &ast.FuncLiteral{
			Line: decl.Line, Params: decl.Params,
			ReturnType: decl.ReturnType, Body: decl.Body,
		}}
	case lexer.IDENT:
		return p.parseIdentifierStmt()
	default:
		p.errors = append(p.errors, fmt.Sprintf(
			"linha %d: comando inválido: '%s'", p.currentToken.Line, p.currentToken.Literal,
		))
		p.synchronize()
		return nil
	}
}

func (p *miniparParser) parseIfStmt() *ast.IfStmt {
	line := p.currentToken.Line
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	p.nextToken()
	condition := p.parseExpression()
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	then := p.parseBlock()

	var elseBranch *ast.BlockStmt
	if p.peekTokenIs(lexer.ELSE) {
		p.nextToken()
		if !p.expectPeek(lexer.LBRACE) {
			return nil
		}
		elseBranch = p.parseBlock()
	}
	return &ast.IfStmt{Line: line, Condition: condition, Then: then, Else: elseBranch}
}

func (p *miniparParser) parseWhileStmt() *ast.WhileStmt {
	line := p.currentToken.Line
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	p.nextToken()
	condition := p.parseExpression()
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	return &ast.WhileStmt{Line: line, Condition: condition, Block: p.parseBlock()}
}

func (p *miniparParser) parseDoStmt() *ast.DoStmt {
	line := p.currentToken.Line
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	block := p.parseBlock()
	if !p.expectPeek(lexer.WHILE) {
		return nil
	}
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	p.nextToken()
	condition := p.parseExpression()
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return &ast.DoStmt{Line: line, Block: block, Condition: condition}
}

func (p *miniparParser) parseForStmt() *ast.ForStmt {
	line := p.currentToken.Line
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	if !p.expectPeek(lexer.IDENT) {
		return nil
	}
	varName := p.currentToken.Literal
	if !p.expectPeek(lexer.IN) {
		return nil
	}
	p.nextToken()
	iter := p.parseExpression()
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	return &ast.ForStmt{Line: line, Var: varName, Iter: iter, Block: p.parseBlock()}
}

func (p *miniparParser) parseSwitchStmt() *ast.SwitchStmt {
	line := p.currentToken.Line
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	p.nextToken()
	expr := p.parseExpression()
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	p.nextToken()

	var cases []*ast.CaseClause
	for !p.currentTokenIs(lexer.RBRACE) && !p.currentTokenIs(lexer.EOF) {
		caseLine := p.currentToken.Line
		value := p.parseExpression()
		if !p.expectPeek(lexer.FAT_ARROW) {
			break
		}
		if !p.expectPeek(lexer.LBRACE) {
			break
		}
		block := p.parseBlock()
		cases = append(cases, &ast.CaseClause{Line: caseLine, Value: value, Block: block})
		p.nextToken()
	}
	return &ast.SwitchStmt{Line: line, Expr: expr, Cases: cases}
}

func (p *miniparParser) parseSeqStmt() *ast.SeqStmt {
	line := p.currentToken.Line
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	return &ast.SeqStmt{Line: line, Block: p.parseBlock()}
}

func (p *miniparParser) parseParStmt() *ast.ParStmt {
	line := p.currentToken.Line
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	return &ast.ParStmt{Line: line, Block: p.parseBlock()}
}

func (p *miniparParser) parseChannelStmt() *ast.ChannelStmt {
	line := p.currentToken.Line
	chanType := p.currentToken.Literal
	if !p.expectPeek(lexer.IDENT) {
		return nil
	}
	name := p.currentToken.Literal
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	args := p.parseArgs()
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return &ast.ChannelStmt{Line: line, ChanType: chanType, Name: name, Args: args}
}

func (p *miniparParser) parsePrintStmt() *ast.PrintStmt {
	line := p.currentToken.Line
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	args := p.parseArgs()
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return &ast.PrintStmt{Line: line, Args: args}
}

func (p *miniparParser) parseInputStmt() *ast.InputStmt {
	line := p.currentToken.Line
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	var prompt ast.Expression
	if !p.peekTokenIs(lexer.RPAREN) {
		p.nextToken()
		prompt = p.parseExpression()
	}
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return &ast.InputStmt{Line: line, Prompt: prompt}
}

func (p *miniparParser) parseReturnStmt() *ast.ReturnStmt {
	line := p.currentToken.Line
	p.nextToken()
	var value ast.Expression
	if !p.currentTokenIs(lexer.SEMICOLON) {
		value = p.parseExpression()
		if p.peekTokenIs(lexer.SEMICOLON) {
			p.nextToken()
		}
	}
	return &ast.ReturnStmt{Line: line, Value: value}
}

func (p *miniparParser) parseGotoStmt() *ast.GotoStmt {
	line := p.currentToken.Line
	if !p.expectPeek(lexer.IDENT) {
		return nil
	}
	label := p.currentToken.Literal
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return &ast.GotoStmt{Line: line, Label: label}
}

func (p *miniparParser) parseIdentifierStmt() ast.Statement {
	name := p.currentToken.Literal
	line := p.currentToken.Line

	// assignment: id = expr
	if p.peekTokenIs(lexer.ASSIGN) {
		p.nextToken()
		p.nextToken()
		value := p.parseExpression()
		if p.peekTokenIs(lexer.SEMICOLON) {
			p.nextToken()
		}
		return &ast.Assignment{Line: line, Name: name, Value: value}
	}

	// compound assignment: id += expr etc.
	if p.peekTokenIs(lexer.PLUS_ASSIGN) || p.peekTokenIs(lexer.MINUS_ASSIGN) ||
		p.peekTokenIs(lexer.STAR_ASSIGN) || p.peekTokenIs(lexer.SLASH_ASSIGN) {
		op := p.peekToken.Literal
		p.nextToken()
		p.nextToken()
		right := p.parseExpression()
		// desugar: id op= rhs  →  id = id op rhs
		bop := compoundOp(op)
		value := &ast.BinaryExpr{
			Line:     line,
			Left:     &ast.Identifier{Line: line, Value: name},
			Operator: bop,
			Right:    right,
		}
		if p.peekTokenIs(lexer.SEMICOLON) {
			p.nextToken()
		}
		return &ast.Assignment{Line: line, Name: name, Value: value}
	}

	// func call: id(args)
	if p.peekTokenIs(lexer.LPAREN) {
		p.nextToken()
		args := p.parseArgs()
		if !p.expectPeek(lexer.RPAREN) {
			return nil
		}
		if p.peekTokenIs(lexer.SEMICOLON) {
			p.nextToken()
		}
		return &ast.FuncCall{Line: line, Name: name, Args: args}
	}

	// method call: id.method(args) — start with parsing as expression
	if p.peekTokenIs(lexer.DOT) {
		expr := p.parsePostfixExpr(&ast.Identifier{Line: line, Value: name})
		if p.peekTokenIs(lexer.SEMICOLON) {
			p.nextToken()
		}
		if stmt, ok := expr.(ast.Statement); ok {
			return stmt
		}
		return &ast.ExprStmt{Line: line, Expression: expr}
	}

	return nil
}

// ==========================================
// Expressions
// ==========================================

func (p *miniparParser) parseExpression() ast.Expression {
	return p.parseOrExpr()
}

func (p *miniparParser) parseOrExpr() ast.Expression {
	left := p.parseAndExpr()
	for p.peekTokenIs(lexer.OR) {
		p.nextToken()
		p.nextToken()
		right := p.parseAndExpr()
		left = &ast.BinaryExpr{Line: left.GetLine(), Left: left, Operator: ast.OpOr, Right: right}
	}
	return left
}

func (p *miniparParser) parseAndExpr() ast.Expression {
	left := p.parseComparisonExpr()
	for p.peekTokenIs(lexer.AND) {
		p.nextToken()
		p.nextToken()
		right := p.parseComparisonExpr()
		left = &ast.BinaryExpr{Line: left.GetLine(), Left: left, Operator: ast.OpAnd, Right: right}
	}
	return left
}

func (p *miniparParser) parseComparisonExpr() ast.Expression {
	left := p.parseAdditiveExpr()
	for isComparisonOp(p.peekToken.Type) {
		p.nextToken()
		op := tokenToBinaryOp(p.currentToken.Type)
		p.nextToken()
		right := p.parseAdditiveExpr()
		left = &ast.BinaryExpr{Line: left.GetLine(), Left: left, Operator: op, Right: right}
	}
	return left
}

func (p *miniparParser) parseAdditiveExpr() ast.Expression {
	left := p.parseMultiplicativeExpr()
	for p.peekTokenIs(lexer.PLUS) || p.peekTokenIs(lexer.MINUS) {
		p.nextToken()
		op := tokenToBinaryOp(p.currentToken.Type)
		p.nextToken()
		right := p.parseMultiplicativeExpr()
		left = &ast.BinaryExpr{Line: left.GetLine(), Left: left, Operator: op, Right: right}
	}
	return left
}

func (p *miniparParser) parseMultiplicativeExpr() ast.Expression {
	left := p.parseUnaryExpr()
	for p.peekTokenIs(lexer.ASTERISK) || p.peekTokenIs(lexer.SLASH) || p.peekTokenIs(lexer.MOD) {
		p.nextToken()
		op := tokenToBinaryOp(p.currentToken.Type)
		p.nextToken()
		right := p.parseUnaryExpr()
		left = &ast.BinaryExpr{Line: left.GetLine(), Left: left, Operator: op, Right: right}
	}
	return left
}

func (p *miniparParser) parseUnaryExpr() ast.Expression {
	if p.currentTokenIs(lexer.MINUS) {
		line := p.currentToken.Line
		p.nextToken()
		return &ast.UnaryExpr{Line: line, Operator: ast.OpNeg, Right: p.parseUnaryExpr()}
	}
	if p.currentTokenIs(lexer.BANG) {
		line := p.currentToken.Line
		p.nextToken()
		return &ast.UnaryExpr{Line: line, Operator: ast.OpNot, Right: p.parseUnaryExpr()}
	}
	return p.parsePostfixExpr(p.parsePrimaryExpr())
}

func (p *miniparParser) parsePostfixExpr(left ast.Expression) ast.Expression {
	for {
		if p.peekTokenIs(lexer.DOT) {
			p.nextToken() // "."
			if !p.expectPeek(lexer.IDENT) {
				break
			}
			method := p.currentToken.Literal
			line := p.currentToken.Line
			if p.peekTokenIs(lexer.LPAREN) {
				p.nextToken()
				args := p.parseArgs()
				if !p.expectPeek(lexer.RPAREN) {
					break
				}
				left = &ast.MethodCall{Line: line, Object: left, Method: method, Args: args}
			} else {
				left = &ast.MethodCall{Line: line, Object: left, Method: method}
			}
		} else if p.peekTokenIs(lexer.LBRACKET) {
			p.nextToken()
			p.nextToken()
			index := p.parseExpression()
			if !p.expectPeek(lexer.RBRACKET) {
				break
			}
			left = &ast.IndexExpr{Line: left.GetLine(), Object: left, Index: index}
		} else {
			break
		}
	}
	return left
}

func (p *miniparParser) parsePrimaryExpr() ast.Expression {
	line := p.currentToken.Line

	switch p.currentToken.Type {
	case lexer.NUMBER:
		lit := p.currentToken.Literal
		if strings.Contains(lit, ".") {
			v, _ := strconv.ParseFloat(lit, 64)
			return &ast.FloatLiteral{Line: line, Value: v}
		}
		v, _ := strconv.ParseInt(lit, 10, 64)
		return &ast.IntLiteral{Line: line, Value: v}

	case lexer.STRING:
		return &ast.StringLiteral{Line: line, Value: p.currentToken.Literal}

	case lexer.CHAR:
		r := rune(p.currentToken.Literal[0])
		return &ast.CharLiteral{Line: line, Value: r}

	case lexer.TRUE:
		return &ast.BooleanLiteral{Line: line, Value: true}

	case lexer.FALSE:
		return &ast.BooleanLiteral{Line: line, Value: false}

	case lexer.SELF:
		return &ast.SelfExpr{Line: line}

	case lexer.FUNC:
		return p.parseFuncLiteral()

	case lexer.IDENT:
		name := p.currentToken.Literal
		if p.peekTokenIs(lexer.LPAREN) {
			p.nextToken()
			args := p.parseArgs()
			if !p.expectPeek(lexer.RPAREN) {
				return nil
			}
			return &ast.FuncCall{Line: line, Name: name, Args: args}
		}
		return &ast.Identifier{Line: line, Value: name}

	case lexer.LPAREN:
		p.nextToken()
		expr := p.parseExpression()
		if !p.expectPeek(lexer.RPAREN) {
			return nil
		}
		return expr

	case lexer.LBRACKET:
		return p.parseListLiteral()
	}

	p.errors = append(p.errors, fmt.Sprintf(
		"linha %d: expressão primária inválida: '%s'", line, p.currentToken.Literal,
	))
	return nil
}

func (p *miniparParser) parseFuncLiteral() *ast.FuncLiteral {
	line := p.currentToken.Line
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	params := p.parseParams()
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}
	var returnType string
	if p.peekTokenIs(lexer.ARROW) {
		p.nextToken()
		p.nextToken()
		returnType = p.currentToken.Literal
	}
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	return &ast.FuncLiteral{Line: line, Params: params, ReturnType: returnType, Body: p.parseBlock()}
}

func (p *miniparParser) parseListLiteral() *ast.ListLiteral {
	line := p.currentToken.Line
	var elems []ast.Expression
	if p.peekTokenIs(lexer.RBRACKET) {
		p.nextToken()
		return &ast.ListLiteral{Line: line}
	}
	p.nextToken()
	elems = append(elems, p.parseExpression())
	for p.peekTokenIs(lexer.COMMA) {
		p.nextToken()
		p.nextToken()
		elems = append(elems, p.parseExpression())
	}
	if !p.expectPeek(lexer.RBRACKET) {
		return nil
	}
	return &ast.ListLiteral{Line: line, Elements: elems}
}

// ==========================================
// Helpers
// ==========================================

func isComparisonOp(t lexer.TokenType) bool {
	switch t {
	case lexer.EQ, lexer.NOT_EQ, lexer.LT, lexer.GT, lexer.LTE, lexer.GTE:
		return true
	}
	return false
}

func tokenToBinaryOp(t lexer.TokenType) ast.Op {
	switch t {
	case lexer.OR:
		return ast.OpOr
	case lexer.AND:
		return ast.OpAnd
	case lexer.EQ:
		return ast.OpEq
	case lexer.NOT_EQ:
		return ast.OpNeq
	case lexer.LT:
		return ast.OpLt
	case lexer.GT:
		return ast.OpGt
	case lexer.LTE:
		return ast.OpLeq
	case lexer.GTE:
		return ast.OpGeq
	case lexer.PLUS:
		return ast.OpAdd
	case lexer.MINUS:
		return ast.OpSub
	case lexer.ASTERISK:
		return ast.OpMul
	case lexer.SLASH:
		return ast.OpDiv
	case lexer.MOD:
		return ast.OpMod
	}
	return ast.OpAdd
}

func compoundOp(literal string) ast.Op {
	switch literal {
	case "+=":
		return ast.OpAdd
	case "-=":
		return ast.OpSub
	case "*=":
		return ast.OpMul
	case "/=":
		return ast.OpDiv
	}
	return ast.OpAdd
}
