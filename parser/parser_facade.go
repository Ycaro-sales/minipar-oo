package parser

import (
	"fmt"
	"minipar/ast"

	"github.com/antlr4-go/antlr/v4"
)

type MiniparFacade struct {
	errors []string
}

// ==========================================
// PONTO DE ENTRADA DO SEMÂNTICO
// ==========================================

func ParseProgram(code string) (*ast.Program, []string) {
	facade := &MiniparFacade{}

	input := antlr.NewInputStream(code)
	lexer := NewMiniparLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewMiniparParser(stream)

	p.RemoveErrorListeners()
	errorListener := &CustomErrorListener{facade: facade}
	p.AddErrorListener(errorListener)

	tree := p.Program()

	if len(facade.errors) > 0 {
		return nil, facade.errors
	}

	astRoot := facade.buildAST(tree)
	return astRoot, nil
}

// ==========================================
// TRATAMENTO DE ERROS
// ==========================================

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	facade *MiniparFacade
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	errStr := fmt.Sprintf("Erro Sintático na linha %d, coluna %d: %s", line, column, msg)
	c.facade.errors = append(c.facade.errors, errStr)
}

// ==========================================
// CONVERSOR DA AST (ANTLR -> GO PURO)
// ==========================================

func (f *MiniparFacade) buildAST(ctx IProgramContext) *ast.Program {
	program := &ast.Program{
		Declarations: []ast.Node{},
	}

	for _, declCtx := range ctx.AllDeclaration() {

		// 1. Canais
		if channelStmt := declCtx.Channel_stmt(); channelStmt != nil {
			if cChan := channelStmt.C_channel_stmt(); cChan != nil {
				program.Declarations = append(program.Declarations, &ast.ChannelStmt{
					Line:     cChan.GetStart().GetLine(),
					ChanType: "c_channel",
					Name:     cChan.Id(0).GetText(),
					Comp1:    cChan.Id(1).GetText(),
					Comp2:    cChan.Id(2).GetText(),
				})
			}
		}

		// 2. Classes
		if classCtx := declCtx.Class_decl(); classCtx != nil {
			classNode := &ast.ClassDecl{
				Line:    classCtx.GetStart().GetLine(),
				Name:    classCtx.Id(0).GetText(),
				Members: []ast.Node{},
			}
			if len(classCtx.AllId()) > 1 {
				classNode.Extends = classCtx.Id(1).GetText()
			}

			// Varrendo variáveis dentro da classe
			for _, memberCtx := range classCtx.AllClass_member() {
				if varDeclCtx := memberCtx.Var_decl(); varDeclCtx != nil {
					varNode := &ast.VarDecl{
						Line: varDeclCtx.GetStart().GetLine(),
						Type: varDeclCtx.Type_().GetText(),
						Name: varDeclCtx.Id().GetText(),
					}
					if varDeclCtx.Expr() != nil {
						varNode.Value = f.buildExpression(varDeclCtx.Expr())
					}
					classNode.Members = append(classNode.Members, varNode)
				}
			}
			program.Declarations = append(program.Declarations, classNode)
		}

		// 3. Funções
		if funcCtx := declCtx.Func_decl(); funcCtx != nil {
			funcNode := &ast.FuncDecl{
				Line:       funcCtx.GetStart().GetLine(),
				Name:       funcCtx.Id().GetText(),
				ReturnType: funcCtx.Type_().GetText(),
				Body:       f.buildBlock(funcCtx.Block()),
			}
			program.Declarations = append(program.Declarations, funcNode)
		}

		// 4. Variáveis Globais
		if varCtx := declCtx.Var_decl(); varCtx != nil {
			varNode := &ast.VarDecl{
				Line: varCtx.GetStart().GetLine(),
				Type: varCtx.Type_().GetText(),
				Name: varCtx.Id().GetText(),
			}
			if varCtx.Expr() != nil {
				varNode.Value = f.buildExpression(varCtx.Expr())
			}
			program.Declarations = append(program.Declarations, varNode)
		}
	}

	return program
}

// buildBlock varre as instruções dentro das chaves { ... }
func (f *MiniparFacade) buildBlock(blockCtx IBlockContext) *ast.BlockStmt {
	if blockCtx == nil {
		return nil
	}

	block := &ast.BlockStmt{
		Line:       blockCtx.GetStart().GetLine(),
		Statements: []ast.Statement{},
	}

	for _, stmtCtx := range blockCtx.AllStmt() {

		// COMANDOS COMPOSTOS (If, While, Seq, Par)
		if compStmt := stmtCtx.Compound_stmt(); compStmt != nil {

			if seqCtx := compStmt.Seq_stmt(); seqCtx != nil {
				block.Statements = append(block.Statements, &ast.SeqStmt{
					Line:  seqCtx.GetStart().GetLine(),
					Block: f.buildBlock(seqCtx.Block()),
				})
			}

			if parCtx := compStmt.Par_stmt(); parCtx != nil {
				block.Statements = append(block.Statements, &ast.ParStmt{
					Line:  parCtx.GetStart().GetLine(),
					Block: f.buildBlock(parCtx.Block()),
				})
			}

			if ifCtx := compStmt.If_stmt(); ifCtx != nil {
				block.Statements = append(block.Statements, &ast.IfStmt{
					Line:      ifCtx.GetStart().GetLine(),
					Condition: f.buildExpression(ifCtx.Expr()),
					Block:     f.buildBlock(ifCtx.Block()),
				})
			}

			if whileCtx := compStmt.While_stmt(); whileCtx != nil {
				block.Statements = append(block.Statements, &ast.WhileStmt{
					Line:      whileCtx.GetStart().GetLine(),
					Condition: f.buildExpression(whileCtx.Expr()),
					Block:     f.buildBlock(whileCtx.Block()),
				})
			}
		}

		// COMANDOS SIMPLES (Atribuições, Declarações locais)
		if simpStmt := stmtCtx.Simple_stmt(); simpStmt != nil {

			if varCtx := simpStmt.Var_decl(); varCtx != nil {
				varNode := &ast.VarDecl{
					Line: varCtx.GetStart().GetLine(),
					Type: varCtx.Type_().GetText(),
					Name: varCtx.Id().GetText(),
				}
				if varCtx.Expr() != nil {
					varNode.Value = f.buildExpression(varCtx.Expr())
				}
				block.Statements = append(block.Statements, varNode)
			}

			if assignCtx := simpStmt.Assignment(); assignCtx != nil {
				block.Statements = append(block.Statements, &ast.Assignment{
					Line:  assignCtx.GetStart().GetLine(),
					Name:  assignCtx.Id().GetText(),
					Value: f.buildExpression(assignCtx.Expr()),
				})
			}
		}
	}

	return block
}

// buildExpression resolve os valores e contas matemáticas
func (f *MiniparFacade) buildExpression(exprCtx IExprContext) ast.Expression {
	if exprCtx == nil {
		return nil
	}
	// O ANTLR gera uma árvore profunda para a precedência.
	// Para pegar o texto literal e quebrar em algo simples para o Semântico:

	text := exprCtx.GetText()
	line := exprCtx.GetStart().GetLine()

	// Checagens simplificadas de literais puros
	if text == "true" || text == "false" {
		return &ast.BooleanLiteral{Line: line, Value: text == "true"}
	}

	// Tenta checar se é um ID puro ou Número (como um fallback seguro)
	// Num cenário real profundo, você navegaria até exprCtx.Or_expr().And_expr()... etc
	// Para o escopo deste facade inicial, encapsulamos o literal.
	if len(text) > 0 && text[0] == '"' {
		return &ast.StringLiteral{Line: line, Value: text}
	}

	if text[0] >= '0' && text[0] <= '9' {
		return &ast.IntegerLiteral{Line: line, Value: text}
	}

	// Se não for literal puro, devolvemos como um Identifier genérico ou
	// poderíamos implementar a recursão completa de BinaryExpr aqui.
	return &ast.Identifier{Line: line, Value: text}
}
