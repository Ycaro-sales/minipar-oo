package parser

import (
	"fmt"
	"minipar/ast"
	"minipar/lexer" // Importa o Lexer do seu amigo
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

// expectPeek é a função mais importante ("consume" do PDF).
// Ela olha o próximo token. Se for o que a gente quer, ela avança.
// Se não for, ela gera um erro na tela e não avança.
func (f *MiniparFacade) expectPeek(t lexer.TokenType) bool {
	if f.peekTokenIs(t) {
		f.nextToken()
		return true
	}
	f.peekError(t)
	return false
}

// peekError formata a mensagem de erro bonitinha.
func (f *MiniparFacade) peekError(t lexer.TokenType) {
	msg := fmt.Sprintf("Erro sintático: esperava o token '%s', mas recebi '%s' (valor: '%s')",
		t, f.peekToken.Type, f.peekToken.Literal)
	f.errors = append(f.errors, msg)
}

// ==========================================
// 2. O PONTO DE ENTRADA (ParseProgram)
// ==========================================

// ParseProgram é a função que o main.go chama.
func ParseProgram(code string) (*ast.Program, []string) {
	// Inicia o lexer do seu amigo
	l := lexer.New(code)

	// Cria a nossa Facade
	f := &MiniparFacade{
		l:      l,
		errors: []string{},
	}

	// Chama nextToken duas vezes para preencher o currentToken e o peekToken
	f.nextToken()
	f.nextToken()

	// Cria o molde da árvore raiz
	program := &ast.Program{
		Declarations: []ast.Node{},
	}

	// O Loop Principal (Equivalente à regra: <program> ::= <declaration>+)
	for !f.currentTokenIs(lexer.EOF) {
		decl := f.parseDeclaration() // Vai tentar ler uma declaração global

		if decl != nil {
			program.Declarations = append(program.Declarations, decl)
		}

		f.nextToken()
	}

	// Se achou erros, devolve a lista de erros e árvore nula
	if len(f.errors) > 0 {
		return nil, f.errors
	}

	return program, nil
}

// ==========================================
// 3. REGRAS GLOBAIS (Declarations)
// ==========================================

// Regra: <declaration> ::= <class_decl> | <func_decl> | <var_decl> | <channel_decl>
func (f *MiniparFacade) parseDeclaration() ast.Node {
	switch f.currentToken.Type {
	case lexer.CLASS:
		return f.parseClassDecl()
	case lexer.FUNC:
		return f.parseFuncDecl()
	case lexer.C_CHANNEL, lexer.S_CHANNEL:
		return f.parseChannelDecl()
	// Como variáveis globais começam com um tipo (number, string, etc):
	case lexer.TYPE_NUMBER, lexer.TYPE_STRING, lexer.TYPE_BOOL, lexer.TYPE_VOID, lexer.TYPE_LIST, lexer.TYPE_DICT:
		return f.parseVarDecl()
	default:
		// Se não é nenhuma palavra reservada global, checamos se é uma classe (tipo customizado)
		if f.currentToken.Type == lexer.IDENT {
			// Pode ser "Neuronio peso = 0"
			return f.parseVarDecl()
		}
		return nil
	}
}

// ==========================================
// 4. LENDO VARIÁVEIS E CLASSES
// ==========================================

// Regra: <var_decl> ::= <type> <id> ("=" <expr>)?
func (f *MiniparFacade) parseVarDecl() *ast.VarDecl {
	decl := &ast.VarDecl{
		// O currentToken agora é o TIPO (ex: "number", "string" ou um ID de classe)
		Type: f.currentToken.Literal,
	}

	// O próximo token OBRIGATORIAMENTE tem que ser o NOME da variável (<id>)
	if !f.expectPeek(lexer.IDENT) {
		return nil
	}
	decl.Name = f.currentToken.Literal

	// Verifica se tem a parte opcional da atribuição ("=" <expr>)?
	if f.peekTokenIs(lexer.ASSIGN) {
		f.nextToken() // Pula para o "="
		f.nextToken() // Pula para o início da expressão/valor

		// Chama a função que lê expressões matemáticas ou valores
		decl.Value = f.parseExpression()
	}

	// NOTA: Como a sua BNF NÃO tem ";" no final de variáveis,
	// a função acaba perfeitamente aqui!
	return decl
}

// Regra: <class_decl> ::= "class" <id> ("extends" <id>)? "{" <class_member>+ "}"
func (f *MiniparFacade) parseClassDecl() *ast.ClassDecl {
	classNode := &ast.ClassDecl{
		Members: []ast.Node{},
	}

	// Espera o nome da classe
	if !f.expectPeek(lexer.IDENT) {
		return nil
	}
	classNode.Name = f.currentToken.Literal

	// Checa a parte opcional: ("extends" <id>)?
	if f.peekTokenIs(lexer.EXTENDS) {
		f.nextToken()                   // Pula para o "extends"
		if !f.expectPeek(lexer.IDENT) { // Espera o nome da superclasse
			return nil
		}
		classNode.Extends = f.currentToken.Literal
	}

	// Espera abrir as chaves "{"
	if !f.expectPeek(lexer.LBRACE) {
		return nil
	}

	f.nextToken() // Entra no primeiro token do miolo da classe

	// Lê os membros da classe até achar a chave de fechar "}"
	for !f.currentTokenIs(lexer.RBRACE) && !f.currentTokenIs(lexer.EOF) {

		// <class_member> ::= <var_decl> | <method_decl>
		// ATENÇÃO: Por enquanto, vamos ler apenas variáveis.
		// Quando formos fazer a Issue de Funções, adicionamos um 'if' aqui
		// para checar se o membro tem um "(" (o que indica que é um método).
		member := f.parseVarDecl()

		if member != nil {
			classNode.Members = append(classNode.Members, member)
		}

		f.nextToken() // Avança para a próxima declaração da classe
	}

	return classNode
}

// ==========================================
// 5. PLACEHOLDER DE EXPRESSÕES
// ==========================================

// parseExpression será o coração da Issue #23.
// Por enquanto, deixamos esse "boneco" apenas capturando o valor bruto
// para o código compilar e não travar a árvore.
// Regra: <stmt> ::= <compound_stmt> | <simple_stmt>
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

	// Declarações de variáveis locais
	case lexer.TYPE_NUMBER, lexer.TYPE_STRING, lexer.TYPE_BOOL, lexer.TYPE_VOID:
		return f.parseVarDecl()

	// --- NOVOS COMANDOS SIMPLES ---
	case lexer.PRINT:
		return f.parsePrintStmt()
	case lexer.RETURN:
		return f.parseReturnStmt()

	default:
		// Se o token for um identificador (ex: "peso", "rede")
		// Pode ser uma atribuição (peso = 10) ou chamada de método/função
		if f.currentToken.Type == lexer.IDENT {
			return f.parseIdentifierStmt()
		}
		return nil
	}
}

// Regra: <block> ::= <stmt>*
// Especialista em ler tudo que está entre chaves { ... }
func (f *MiniparFacade) parseBlock() *ast.BlockStmt {
	block := &ast.BlockStmt{
		Statements: []ast.Statement{},
	}

	f.nextToken() // Avança para o primeiro token dentro das chaves

	// Loop contínuo: lê comandos até achar a chave de fechamento '}'
	for !f.currentTokenIs(lexer.RBRACE) && !f.currentTokenIs(lexer.EOF) {
		stmt := f.parseStatement() // Manda pro Roteador

		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		f.nextToken()
	}

	return block
}

// Regra: <seq_stmt> ::= "seq" "{" <block> "}"
func (f *MiniparFacade) parseSeqStmt() *ast.SeqStmt {
	stmt := &ast.SeqStmt{}

	// Espera abrir a chave '{' do bloco sequencial
	if !f.expectPeek(lexer.LBRACE) {
		return nil
	}

	// Terceiriza a leitura do miolo para o parseBlock!
	stmt.Block = f.parseBlock()

	return stmt
}

// Regra: <par_stmt> ::= "par" "{" <block> "}"
func (f *MiniparFacade) parseParStmt() *ast.ParStmt {
	stmt := &ast.ParStmt{}

	// Espera abrir a chave '{' do bloco paralelo
	if !f.expectPeek(lexer.LBRACE) {
		return nil
	}

	// Terceiriza a leitura do miolo
	stmt.Block = f.parseBlock()

	return stmt
}

// ==========================================
// 7. CONTROLES DE FLUXO (IF e WHILE)
// ==========================================

// Regra: <if_stmt> ::= "if" "(" <expr> ")" "{" <block> "}"
func (f *MiniparFacade) parseIfStmt() *ast.IfStmt {
	stmt := &ast.IfStmt{}

	// Espera abrir parênteses "("
	if !f.expectPeek(lexer.LPAREN) {
		return nil
	}

	f.nextToken() // Pula para o início da expressão/condição

	// Chama a função matemática para resolver a condição (ex: ativo == true)
	stmt.Condition = f.parseExpression()

	// Espera fechar parênteses ")"
	if !f.expectPeek(lexer.RPAREN) {
		return nil
	}

	// Espera abrir as chaves "{" do bloco
	if !f.expectPeek(lexer.LBRACE) {
		return nil
	}

	// Terceiriza a leitura do miolo para o parseBlock
	stmt.Block = f.parseBlock()

	return stmt
}

// Regra: <while_stmt> ::= "while" "(" <expr> ")" "{" <block> "}"
func (f *MiniparFacade) parseWhileStmt() *ast.WhileStmt {
	stmt := &ast.WhileStmt{}

	// Espera abrir parênteses "("
	if !f.expectPeek(lexer.LPAREN) {
		return nil
	}

	f.nextToken() // Pula para a expressão

	// Resolve a condição de parada do while
	stmt.Condition = f.parseExpression()

	// Espera fechar parênteses ")"
	if !f.expectPeek(lexer.RPAREN) {
		return nil
	}

	// Espera abrir as chaves "{"
	if !f.expectPeek(lexer.LBRACE) {
		return nil
	}

	// Terceiriza a leitura do miolo
	stmt.Block = f.parseBlock()

	return stmt
}

// ==========================================
// 8. FUNÇÕES E CANAIS (Declarações Globais)
// ==========================================

// Regra: <func_decl> ::= "func" <id> "(" <params>? ")" <type> "{" <block> "}"
func (f *MiniparFacade) parseFuncDecl() *ast.FuncDecl {
	decl := &ast.FuncDecl{}

	// O token atual é o "func". Esperamos que o próximo seja o nome da função (<id>)
	if !f.expectPeek(lexer.IDENT) {
		return nil
	}
	decl.Name = f.currentToken.Literal

	// Espera abrir parênteses "("
	if !f.expectPeek(lexer.LPAREN) {
		return nil
	}

	// Loop para ler os parâmetros: <params> ::= <param> ("," <param>)*
	// Como a sua struct FuncDecl original não tinha um campo 'Params',
	// nós vamos apenas varrer e consumir esses tokens por enquanto para a sintaxe não quebrar.
	for !f.peekTokenIs(lexer.RPAREN) && !f.peekTokenIs(lexer.EOF) {
		f.nextToken()
	}

	// Espera fechar parênteses ")"
	if !f.expectPeek(lexer.RPAREN) {
		return nil
	}

	// AQUI TEM UM DETALHE INCRÍVEL DA SUA BNF:
	// O tipo de retorno vem DEPOIS dos parênteses! (ex: func soma(a, b) number { ... })
	f.nextToken()
	decl.ReturnType = f.currentToken.Literal

	// Espera abrir as chaves "{" do corpo da função
	if !f.expectPeek(lexer.LBRACE) {
		return nil
	}

	// Terceiriza a leitura de tudo que tem dentro da função para o nosso operário parseBlock!
	decl.Body = f.parseBlock()

	return decl
}

// Regra: <channel_stmt> ::= <s_channel_stmt> | <c_channel_stmt>
//
//	<c_channel_stmt> ::= "c_channel" <id> "(" <args>? ")" ";"
func (f *MiniparFacade) parseChannelDecl() *ast.ChannelStmt {
	decl := &ast.ChannelStmt{
		// O token atual já é "c_channel" ou "s_channel"
		ChanType: f.currentToken.Literal,
	}

	// Espera o nome do canal (<id>)
	if !f.expectPeek(lexer.IDENT) {
		return nil
	}
	decl.Name = f.currentToken.Literal

	// A nova BNF exige abrir parênteses "("
	if !f.expectPeek(lexer.LPAREN) {
		return nil
	}

	// Loop para ler os argumentos opcionais (ex: os computadores da rede)
	for !f.peekTokenIs(lexer.RPAREN) && !f.peekTokenIs(lexer.EOF) {
		f.nextToken()
	}

	// Espera fechar parênteses ")"
	if !f.expectPeek(lexer.RPAREN) {
		return nil
	}

	// OBRIGATÓRIO NA BNF: Canais precisam terminar com ponto e vírgula ";"
	if !f.expectPeek(lexer.SEMICOLON) {
		return nil
	}

	return decl
}

// ==========================================
// 9. COMANDOS SIMPLES (Print, Return, Atribuição)
// ==========================================

// Regra: <print_call> ::= "print" "(" <args>? ")" ";"
func (f *MiniparFacade) parsePrintStmt() *ast.PrintStmt {
	stmt := &ast.PrintStmt{}

	// Espera abrir parênteses "("
	if !f.expectPeek(lexer.LPAREN) {
		return nil
	}
	f.nextToken() // Entra no miolo dos parênteses

	// Loop para ler os argumentos (ex: print("Olá", 10))
	for !f.currentTokenIs(lexer.RPAREN) && !f.currentTokenIs(lexer.EOF) {
		// Por enquanto só consome os tokens para não travar
		f.nextToken()
	}

	// O currentToken agora é ")".
	// A BNF de vocês EXIGE ponto e vírgula no final do print!
	if !f.expectPeek(lexer.SEMICOLON) {
		return nil
	}

	return stmt
}

// Regra: "return" <expr>? ";"
func (f *MiniparFacade) parseReturnStmt() *ast.ReturnStmt {
	stmt := &ast.ReturnStmt{}

	f.nextToken() // Pula a palavra "return"

	// Se o próximo não for ponto e vírgula, significa que tem uma expressão junto!
	if !f.currentTokenIs(lexer.SEMICOLON) {
		stmt.Value = f.parseExpression()
		f.nextToken() // Pula para o ponto e vírgula
	}

	// BNF EXIGE ponto e vírgula no final do return!
	if !f.currentTokenIs(lexer.SEMICOLON) {
		f.peekError(lexer.SEMICOLON)
		return nil
	}

	return stmt
}

// Lida com comandos que começam com um nome (Variáveis ou Chamadas)
func (f *MiniparFacade) parseIdentifierStmt() ast.Statement {
	// Espiamos o próximo token para saber qual regra aplicar

	// Regra: <assignment> ::= <id> "=" <expr>
	if f.peekTokenIs(lexer.ASSIGN) {
		stmt := &ast.Assignment{
			Name: f.currentToken.Literal,
		}

		f.nextToken() // Pula para o "="
		f.nextToken() // Pula para o valor

		stmt.Value = f.parseExpression()
		return stmt
	}

	// Aqui no futuro adicionaremos:
	// se peekToken == "(" -> parseFuncCall()
	// se peekToken == "." -> parseMethodCall()

	return nil
}
// ==========================================
// 10. EXPRESSÕES E PRECEDÊNCIA (Chefão Final)
// ==========================================

// Regra: <expr> ::= <or_expr>
func (f *MiniparFacade) parseExpression() ast.Expression {
	return f.parseOrExpr()
}

// Regra: <or_expr> ::= <and_expr> ("or" <and_expr>)*
func (f *MiniparFacade) parseOrExpr() ast.Expression {
	left := f.parseAndExpr()

	for f.currentTokenIs(lexer.OR) {
		operator := f.currentToken.Literal
		f.nextToken() // Consome o "or"
		right := f.parseAndExpr()
		left = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return left
}

// Regra: <and_expr> ::= <comparison_expr> ("and" <comparison_expr>)*
func (f *MiniparFacade) parseAndExpr() ast.Expression {
	left := f.parseComparisonExpr()

	for f.currentTokenIs(lexer.AND) {
		operator := f.currentToken.Literal
		f.nextToken() // Consome o "and"
		right := f.parseComparisonExpr()
		left = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return left
}

// Regra: <comparison_expr> ::= <additive_expr> (("==" | "!=" | "<" | ">" | "<=" | ">=") <additive_expr>)*
func (f *MiniparFacade) parseComparisonExpr() ast.Expression {
	left := f.parseAdditiveExpr()

	for f.currentTokenIs(lexer.EQ) || f.currentTokenIs(lexer.NOT_EQ) ||
		f.currentTokenIs(lexer.LT) || f.currentTokenIs(lexer.GT) ||
		f.currentTokenIs(lexer.LTE) || f.currentTokenIs(lexer.GTE) {
		
		operator := f.currentToken.Literal
		f.nextToken() // Consome o operador
		right := f.parseAdditiveExpr()
		left = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return left
}

// Regra: <additive_expr> ::= <multiplicative_expr> (("+" | "-") <multiplicative_expr>)*
func (f *MiniparFacade) parseAdditiveExpr() ast.Expression {
	left := f.parseMultiplicativeExpr()

	for f.currentTokenIs(lexer.PLUS) || f.currentTokenIs(lexer.MINUS) {
		operator := f.currentToken.Literal
		f.nextToken() // Consome o + ou -
		right := f.parseMultiplicativeExpr()
		left = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return left
}

// Regra: <multiplicative_expr> ::= <unary_expr> (("*" | "/" | "%") <unary_expr>)*
func (f *MiniparFacade) parseMultiplicativeExpr() ast.Expression {
	left := f.parsePrimaryExpr() // Simplificamos pulando o Unary por enquanto para testar

	for f.currentTokenIs(lexer.ASTERISK) || f.currentTokenIs(lexer.SLASH) || f.currentTokenIs(lexer.MOD) {
		operator := f.currentToken.Literal
		f.nextToken() // Consome o *, / ou %
		right := f.parsePrimaryExpr()
		left = &ast.BinaryExpr{Left: left, Operator: operator, Right: right}
	}
	return left
}

// Regra: <primary_expr> ::= <numero> | <string> | "true" | "false" | <id> | "(" <expr> ")"
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
		// Para simplificar agora, tratamos IDs puros. 
		// (Futuramente aqui entra a checagem se é chamada de método ou função)
		return &ast.Identifier{Line: 1, Value: f.currentToken.Literal}
	
	case lexer.LPAREN:
		f.nextToken() // Consome o "("
		expr := f.parseExpression() // Volta pro topo da escada!
		if !f.expectPeek(lexer.RPAREN) {
			return nil
		}
		return expr
	}

	f.errors = append(f.errors, fmt.Sprintf("Erro sintático: Expressão primária não reconhecida: '%s'", f.currentToken.Literal))
	return nil
}
