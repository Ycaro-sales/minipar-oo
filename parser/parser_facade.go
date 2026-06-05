package parser

import (
	"fmt"
	"minipar/ast"
	"minipar/lexer"
)

// ==========================================
// 1. A ESTRUTURA DA FACADE (O MOTOR)
// ==========================================

type MiniparFacade struct {
	l            *lexer.Lexer
	currentToken lexer.Token // O token que estamos lendo agora
	peekToken    lexer.Token // O próximo token (a "espiadinha")
	errors       []string    // Lista de erros sintáticos
}

// nextToken avança os ponteiros do parser, consumindo um token do Lexer.
func (f *MiniparFacade) nextToken() {
	f.currentToken = f.peekToken
	f.peekToken = f.l.NextToken()
}

// currentTokenIs checa se o token atual é do tipo esperado.
func (f *MiniparFacade) currentTokenIs(t lexer.TokenType) bool {
	return f.currentToken.Type == t
}

// peekTokenIs checa se o PRÓXIMO token é do tipo esperado.
func (f *MiniparFacade) peekTokenIs(t lexer.TokenType) bool {
	return f.peekToken.Type == t
}

// expectPeek é a função mais importante.
// Ela olha o próximo token. Se for o que a gente quer, ela avança.
func (f *MiniparFacade) expectPeek(t lexer.TokenType) bool {
	if f.peekTokenIs(t) {
		f.nextToken()
		return true
	}
	f.peekError(t)
	return false
}

// peekError formata a mensagem de erro.
func (f *MiniparFacade) peekError(t lexer.TokenType) {
	msg := fmt.Sprintf("Erro sintático: esperava o token '%s', mas recebi '%s' (valor: '%s')",
		t, f.peekToken.Type, f.peekToken.Literal)
	f.errors = append(f.errors, msg)
}

// ==========================================
// 2. O PONTO DE ENTRADA (ParseProgram)
// ==========================================

func ParseProgram(code string) (*ast.Program, []string) {
	l := lexer.New(code)
	f := &MiniparFacade{
		l:      l,
		errors: []string{},
	}

	f.nextToken()
	f.nextToken()

	program := &ast.Program{
		Declarations: []ast.Node{},
	}

	for !f.currentTokenIs(lexer.EOF) {
		decl := f.parseDeclaration()

		if decl != nil {
			program.Declarations = append(program.Declarations, decl)
		}

		f.nextToken()
	}

	if len(f.errors) > 0 {
		return nil, f.errors
	}

	return program, nil
}

// ==========================================
// 3. REGRAS GLOBAIS (Declarations)
// ==========================================

// ==========================================
// 3. REGRAS GLOBAIS (Declarations)
// ==========================================

func (f *MiniparFacade) parseDeclaration() ast.Node {
	switch f.currentToken.Type {
	case lexer.CLASS:
		return f.parseClassDecl()
	case lexer.FUNC:
		return f.parseFuncDecl()
	case lexer.C_CHANNEL, lexer.S_CHANNEL:
		return f.parseChannelDecl()
	case lexer.TYPE_NUMBER, lexer.TYPE_STRING, lexer.TYPE_BOOL, lexer.TYPE_VOID, lexer.TYPE_LIST, lexer.TYPE_DICT:
		return f.parseVarDecl()

	// CORREÇÃO: O seu código de teste possui blocos (seq/par) no escopo global.
	// O porteiro precisa deixá-los passar redirecionando para os Comandos (Statements)!
	case lexer.SEQ, lexer.PAR, lexer.IF, lexer.WHILE, lexer.PRINT, lexer.RETURN:
		return f.parseStatement()

	default:
		// A MÁGICA: Como saber se a palavra solta é uma Classe (Neuronio peso = 1)
		// ou uma simples Atribuição (peso = 1)? Nós espiamos o próximo token!
		if f.currentToken.Type == lexer.IDENT {

			if f.peekTokenIs(lexer.IDENT) {
				return f.parseVarDecl() // Ex: "Neuronio peso" (Variável)
			}
			return f.parseStatement() // Ex: "peso = 10" (Comando/Atribuição)

		}
		return nil
	}
}

// ==========================================
// 4. LENDO VARIÁVEIS E CLASSES
// ==========================================

func (f *MiniparFacade) parseVarDecl() *ast.VarDecl {
	decl := &ast.VarDecl{
		Type: f.currentToken.Literal,
	}

	if !f.expectPeek(lexer.IDENT) {
		return nil
	}
	decl.Name = f.currentToken.Literal

	if f.peekTokenIs(lexer.ASSIGN) {
		f.nextToken()
		f.nextToken()
		decl.Value = f.parseExpression()
	}
	return decl
}

func (f *MiniparFacade) parseClassDecl() *ast.ClassDecl {
	classNode := &ast.ClassDecl{
		Members: []ast.Node{},
	}

	if !f.expectPeek(lexer.IDENT) {
		return nil
	}
	classNode.Name = f.currentToken.Literal

	if f.peekTokenIs(lexer.EXTENDS) {
		f.nextToken()
		if !f.expectPeek(lexer.IDENT) {
			return nil
		}
		classNode.Extends = f.currentToken.Literal
	}

	if !f.expectPeek(lexer.LBRACE) {
		return nil
	}

	f.nextToken() // Entra no primeiro token da classe

	for !f.currentTokenIs(lexer.RBRACE) && !f.currentTokenIs(lexer.EOF) {

		var member ast.Node

		// A CORREÇÃO 1 AQUI: Roteador da Classe diferencia métodos e variáveis!
		if f.currentTokenIs(lexer.FUNC) {
			member = f.parseFuncDecl()
		} else {
			member = f.parseVarDecl()
		}

		if member != nil {
			classNode.Members = append(classNode.Members, member)
		}

		f.nextToken()
	}

	return classNode
}

// ==========================================
// 5. BLOCOS DE COMANDOS E PARALELISMO
// ==========================================

func (f *MiniparFacade) parseStatement() ast.Statement {
	switch f.currentToken.Type {
	case lexer.IF:
		return f.parseIfStmt()
	case lexer.WHILE:
		return f.parseWhileStmt()
	case lexer.SEQ:
		return f.parseSeqStmt()
	case lexer.PAR:
		return f.parseParStmt()
	case lexer.TYPE_NUMBER, lexer.TYPE_STRING, lexer.TYPE_BOOL, lexer.TYPE_VOID:
		return f.parseVarDecl()
	case lexer.PRINT:
		return f.parsePrintStmt()
	case lexer.RETURN:
		return f.parseReturnStmt()
	default:
		if f.currentToken.Type == lexer.IDENT {
			return f.parseIdentifierStmt()
		}
		return nil
	}
}

func (f *MiniparFacade) parseBlock() *ast.BlockStmt {
	block := &ast.BlockStmt{
		Statements: []ast.Statement{},
	}

	f.nextToken()

	for !f.currentTokenIs(lexer.RBRACE) && !f.currentTokenIs(lexer.EOF) {
		stmt := f.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		f.nextToken()
	}

	return block
}

func (f *MiniparFacade) parseSeqStmt() *ast.SeqStmt {
	stmt := &ast.SeqStmt{}
	if !f.expectPeek(lexer.LBRACE) {
		return nil
	}
	stmt.Block = f.parseBlock()
	return stmt
}

func (f *MiniparFacade) parseParStmt() *ast.ParStmt {
	stmt := &ast.ParStmt{}
	if !f.expectPeek(lexer.LBRACE) {
		return nil
	}
	stmt.Block = f.parseBlock()
	return stmt
}

// ==========================================
// 6. CONTROLES DE FLUXO (IF e WHILE)
// ==========================================

func (f *MiniparFacade) parseIfStmt() *ast.IfStmt {
	stmt := &ast.IfStmt{}
	if !f.expectPeek(lexer.LPAREN) {
		return nil
	}
	f.nextToken()
	stmt.Condition = f.parseExpression()
	if !f.expectPeek(lexer.RPAREN) {
		return nil
	}
	if !f.expectPeek(lexer.LBRACE) {
		return nil
	}
	stmt.Block = f.parseBlock()
	return stmt
}

func (f *MiniparFacade) parseWhileStmt() *ast.WhileStmt {
	stmt := &ast.WhileStmt{}
	if !f.expectPeek(lexer.LPAREN) {
		return nil
	}
	f.nextToken()
	stmt.Condition = f.parseExpression()
	if !f.expectPeek(lexer.RPAREN) {
		return nil
	}
	if !f.expectPeek(lexer.LBRACE) {
		return nil
	}
	stmt.Block = f.parseBlock()
	return stmt
}

// ==========================================
// 7. FUNÇÕES E CANAIS
// ==========================================

func (f *MiniparFacade) parseFuncDecl() *ast.FuncDecl {
	decl := &ast.FuncDecl{}
	if !f.expectPeek(lexer.IDENT) {
		return nil
	}
	decl.Name = f.currentToken.Literal
	if !f.expectPeek(lexer.LPAREN) {
		return nil
	}

	for !f.peekTokenIs(lexer.RPAREN) && !f.peekTokenIs(lexer.EOF) {
		f.nextToken()
	}

	if !f.expectPeek(lexer.RPAREN) {
		return nil
	}

	f.nextToken()
	decl.ReturnType = f.currentToken.Literal

	if !f.expectPeek(lexer.LBRACE) {
		return nil
	}
	decl.Body = f.parseBlock()
	return decl
}

func (f *MiniparFacade) parseChannelDecl() *ast.ChannelStmt {
	decl := &ast.ChannelStmt{
		ChanType: f.currentToken.Literal,
	}
	if !f.expectPeek(lexer.IDENT) {
		return nil
	}
	decl.Name = f.currentToken.Literal
	if !f.expectPeek(lexer.LPAREN) {
		return nil
	}

	for !f.peekTokenIs(lexer.RPAREN) && !f.peekTokenIs(lexer.EOF) {
		f.nextToken()
	}

	if !f.expectPeek(lexer.RPAREN) {
		return nil
	}
	if !f.expectPeek(lexer.SEMICOLON) {
		return nil
	}
	return decl
}

// ==========================================
// 8. COMANDOS SIMPLES
// ==========================================

func (f *MiniparFacade) parsePrintStmt() *ast.PrintStmt {
	stmt := &ast.PrintStmt{}
	if !f.expectPeek(lexer.LPAREN) {
		return nil
	}
	f.nextToken()

	for !f.currentTokenIs(lexer.RPAREN) && !f.currentTokenIs(lexer.EOF) {
		f.nextToken()
	}

	if !f.expectPeek(lexer.SEMICOLON) {
		return nil
	}
	return stmt
}

func (f *MiniparFacade) parseReturnStmt() *ast.ReturnStmt {
	stmt := &ast.ReturnStmt{}
	f.nextToken()
	if !f.currentTokenIs(lexer.SEMICOLON) {
		stmt.Value = f.parseExpression()
		f.nextToken()
	}
	if !f.currentTokenIs(lexer.SEMICOLON) {
		f.peekError(lexer.SEMICOLON)
		return nil
	}
	return stmt
}

func (f *MiniparFacade) parseIdentifierStmt() ast.Statement {
	if f.peekTokenIs(lexer.ASSIGN) {
		stmt := &ast.Assignment{
			Name: f.currentToken.Literal,
		}
		f.nextToken()
		f.nextToken()
		stmt.Value = f.parseExpression()
		return stmt
	}
	return nil
}

// ==========================================
// 9. EXPRESSÕES E PRECEDÊNCIA MATEMÁTICA
// ==========================================

// A CORREÇÃO 2 AQUI: Uso correto de peekTokenIs e nextToken!
func (f *MiniparFacade) parseExpression() ast.Expression {
	return f.parseOrExpr()
}

func (f *MiniparFacade) parseOrExpr() ast.Expression {
	left := f.parseAndExpr()
	for f.peekTokenIs(lexer.OR) {
		f.nextToken()
		operator := f.currentToken.Literal
		f.nextToken()
		right := f.parseAndExpr()
		left = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return left
}

func (f *MiniparFacade) parseAndExpr() ast.Expression {
	left := f.parseComparisonExpr()
	for f.peekTokenIs(lexer.AND) {
		f.nextToken()
		operator := f.currentToken.Literal
		f.nextToken()
		right := f.parseComparisonExpr()
		left = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return left
}

func (f *MiniparFacade) parseComparisonExpr() ast.Expression {
	left := f.parseAdditiveExpr()
	for f.peekTokenIs(lexer.EQ) || f.peekTokenIs(lexer.NOT_EQ) ||
		f.peekTokenIs(lexer.LT) || f.peekTokenIs(lexer.GT) ||
		f.peekTokenIs(lexer.LTE) || f.peekTokenIs(lexer.GTE) {

		f.nextToken()
		operator := f.currentToken.Literal
		f.nextToken()
		right := f.parseAdditiveExpr()
		left = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return left
}

func (f *MiniparFacade) parseAdditiveExpr() ast.Expression {
	left := f.parseMultiplicativeExpr()
	for f.peekTokenIs(lexer.PLUS) || f.peekTokenIs(lexer.MINUS) {
		f.nextToken()
		operator := f.currentToken.Literal
		f.nextToken()
		right := f.parseMultiplicativeExpr()
		left = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return left
}

func (f *MiniparFacade) parseMultiplicativeExpr() ast.Expression {
	left := f.parsePrimaryExpr()
	for f.peekTokenIs(lexer.ASTERISK) || f.peekTokenIs(lexer.SLASH) || f.peekTokenIs(lexer.MOD) {
		f.nextToken()
		operator := f.currentToken.Literal
		f.nextToken()
		right := f.parsePrimaryExpr()
		left = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return left
}

func (f *MiniparFacade) parsePrimaryExpr() ast.Expression {
	switch f.currentToken.Type {
	case lexer.NUMBER:
		return &ast.IntegerLiteral{Line: 1, Value: f.currentToken.Literal}
	case lexer.STRING:
		return &ast.StringLiteral{Line: 1, Value: f.currentToken.Literal}
	case lexer.TRUE:
		return &ast.BooleanLiteral{Line: 1, Value: true}
	case lexer.FALSE:
		return &ast.BooleanLiteral{Line: 1, Value: false}
	case lexer.IDENT:
		return &ast.Identifier{Line: 1, Value: f.currentToken.Literal}
	case lexer.LPAREN:
		f.nextToken()
		expr := f.parseExpression()
		if !f.expectPeek(lexer.RPAREN) {
			return nil
		}
		return expr
	}

	f.errors = append(f.errors, fmt.Sprintf("Erro sintático: Expressão primária não reconhecida: '%s'", f.currentToken.Literal))
	return nil
}
