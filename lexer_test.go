package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {

	input := `
		# Esse é um comentário de linha que deve ser ignorado
		class Gato {
			func miar() void {
				print("miau");
			}
		}

		/* Esse é um comentário
		   de bloco gigante que também deve
		   ser ignorado pelo seu Lexer */
		
		patas = 4;
		if (patas == 4) {
			seq { true }
		} != >= <= < > + - * / %
	`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{CLASS, "class"},
		{IDENT, "Gato"},
		{LBRACE, "{"},
		{FUNC, "func"},
		{IDENT, "miar"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{TYPE_VOID, "void"},
		{LBRACE, "{"},
		{PRINT, "print"},
		{LPAREN, "("},
		{STRING, "miau"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{RBRACE, "}"},
		{IDENT, "patas"},
		{ASSIGN, "="},
		{NUMBER, "4"},
		{SEMICOLON, ";"},
		{IF, "if"},
		{LPAREN, "("},
		{IDENT, "patas"},
		{EQ, "=="},
		{NUMBER, "4"},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{SEQ, "seq"},
		{LBRACE, "{"},
		{TRUE, "true"},
		{RBRACE, "}"},
		{RBRACE, "}"},
		{NOT_EQ, "!="},
		{GTE, ">="},
		{LTE, "<="},
		{LT, "<"},
		{GT, ">"},
		{PLUS, "+"},
		{MINUS, "-"},
		{ASTERISK, "*"},
		{SLASH, "/"},
		{MOD, "%"},
		{EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("Testes[%d] - Tipo de token incorreto. Esperava=%q, Recebeu=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Testes[%d] - Literal incorreto. Esperava=%q, Recebeu=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
