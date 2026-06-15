package lexer

import "testing"

// tok is a shorthand for building expected token entries.
type tok struct {
	wantType    TokenType
	wantLiteral string
}

func runTokens(t *testing.T, input string, expected []tok) {
	t.Helper()
	l := New(input)
	for i, tt := range expected {
		got := l.NextToken()
		if got.Type != tt.wantType {
			t.Fatalf("[%d] type: want %s, got %s (literal %q)", i, tt.wantType, got.Type, got.Literal)
		}
		if got.Literal != tt.wantLiteral {
			t.Fatalf("[%d] literal: want %q, got %q", i, tt.wantLiteral, got.Literal)
		}
	}
}

// ==========================================
// Keywords
// ==========================================

func TestKeywords(t *testing.T) {
	cases := []struct {
		name    string
		lexeme  string
		tokType TokenType
	}{
		{"let", "let", LET},
		{"func", "func", FUNC},
		{"class", "class", CLASS},
		{"interface", "interface", INTERFACE},
		{"implements", "implements", IMPLEMENTS},
		{"Self", "Self", SELF},
		{"return", "return", RETURN},
		{"break", "break", BREAK},
		{"continue", "continue", CONTINUE},
		{"pass", "pass", PASS},
		{"goto", "goto", GOTO},
		{"if", "if", IF},
		{"else", "else", ELSE},
		{"switch", "switch", SWITCH},
		{"while", "while", WHILE},
		{"do", "do", DO},
		{"for", "for", FOR},
		{"in", "in", IN},
		{"seq", "seq", SEQ},
		{"par", "par", PAR},
		{"print", "print", PRINT},
		{"input", "input", INPUT},
		{"and", "and", AND},
		{"or", "or", OR},
		{"true", "true", TRUE},
		{"false", "false", FALSE},
		{"s_channel", "s_channel", S_CHANNEL},
		{"c_channel", "c_channel", C_CHANNEL},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runTokens(t, c.lexeme, []tok{{c.tokType, c.lexeme}, {EOF, ""}})
		})
	}
}

// ==========================================
// Type keywords
// ==========================================

func TestTypeKeywords(t *testing.T) {
	cases := []struct {
		lexeme  string
		tokType TokenType
	}{
		{"i8", TYPE_I8}, {"i16", TYPE_I16}, {"i32", TYPE_I32}, {"i64", TYPE_I64},
		{"u8", TYPE_U8}, {"u16", TYPE_U16}, {"u32", TYPE_U32}, {"u64", TYPE_U64},
		{"f16", TYPE_F16}, {"f32", TYPE_F32}, {"f64", TYPE_F64},
		{"char", TYPE_CHAR}, {"string", TYPE_STRING}, {"bool", TYPE_BOOL},
		{"any", TYPE_ANY}, {"void", TYPE_VOID}, {"chan", TYPE_CHAN},
	}
	for _, c := range cases {
		t.Run(c.lexeme, func(t *testing.T) {
			runTokens(t, c.lexeme, []tok{{c.tokType, c.lexeme}, {EOF, ""}})
		})
	}
}

// ==========================================
// Operators
// ==========================================

func TestSingleCharOperators(t *testing.T) {
	runTokens(t, "+ - * / % ! = < > , ; : . ( ) { } [ ]", []tok{
		{PLUS, "+"}, {MINUS, "-"}, {ASTERISK, "*"}, {SLASH, "/"}, {MOD, "%"},
		{BANG, "!"}, {ASSIGN, "="}, {LT, "<"}, {GT, ">"},
		{COMMA, ","}, {SEMICOLON, ";"}, {COLON, ":"}, {DOT, "."},
		{LPAREN, "("}, {RPAREN, ")"}, {LBRACE, "{"}, {RBRACE, "}"},
		{LBRACKET, "["}, {RBRACKET, "]"},
		{EOF, ""},
	})
}

func TestTwoCharOperators(t *testing.T) {
	runTokens(t, "== != <= >= -> => += -= *= /=", []tok{
		{EQ, "=="}, {NOT_EQ, "!="}, {LTE, "<="}, {GTE, ">="},
		{ARROW, "->"}, {FAT_ARROW, "=>"},
		{PLUS_ASSIGN, "+="}, {MINUS_ASSIGN, "-="}, {STAR_ASSIGN, "*="}, {SLASH_ASSIGN, "/="},
		{EOF, ""},
	})
}

// ==========================================
// Literals
// ==========================================

func TestIntegerLiteral(t *testing.T) {
	runTokens(t, "0 42 100", []tok{
		{NUMBER, "0"}, {NUMBER, "42"}, {NUMBER, "100"}, {EOF, ""},
	})
}

func TestFloatLiteral(t *testing.T) {
	runTokens(t, "3.14 0.5", []tok{
		{NUMBER, "3.14"}, {NUMBER, "0.5"}, {EOF, ""},
	})
}

func TestStringLiteral(t *testing.T) {
	runTokens(t, `"hello" "world"`, []tok{
		{STRING, "hello"}, {STRING, "world"}, {EOF, ""},
	})
}

func TestCharLiteral(t *testing.T) {
	runTokens(t, `'a' 'Z' '0'`, []tok{
		{CHAR, "a"}, {CHAR, "Z"}, {CHAR, "0"}, {EOF, ""},
	})
}

func TestBooleanLiterals(t *testing.T) {
	runTokens(t, "true false", []tok{
		{TRUE, "true"}, {FALSE, "false"}, {EOF, ""},
	})
}

// ==========================================
// Identifiers
// ==========================================

func TestIdentifier(t *testing.T) {
	runTokens(t, "foo bar_baz myVar x1", []tok{
		{IDENT, "foo"}, {IDENT, "bar_baz"}, {IDENT, "myVar"}, {IDENT, "x1"}, {EOF, ""},
	})
}

// ==========================================
// Comments (skipped)
// ==========================================

func TestLineComment(t *testing.T) {
	runTokens(t, "# this is a comment\n42", []tok{
		{NUMBER, "42"}, {EOF, ""},
	})
}

func TestBlockComment(t *testing.T) {
	runTokens(t, "/* block\ncomment */42", []tok{
		{NUMBER, "42"}, {EOF, ""},
	})
}

func TestCommentAtEOF(t *testing.T) {
	runTokens(t, "# comment at EOF", []tok{{EOF, ""}})
}

// ==========================================
// Line tracking
// ==========================================

func TestLineTracking(t *testing.T) {
	input := "foo\nbar\nbaz"
	l := New(input)

	tok1 := l.NextToken()
	if tok1.Line != 1 {
		t.Errorf("foo: want line 1, got %d", tok1.Line)
	}
	tok2 := l.NextToken()
	if tok2.Line != 2 {
		t.Errorf("bar: want line 2, got %d", tok2.Line)
	}
	tok3 := l.NextToken()
	if tok3.Line != 3 {
		t.Errorf("baz: want line 3, got %d", tok3.Line)
	}
}

// ==========================================
// Edge cases
// ==========================================

func TestEmptyInput(t *testing.T) {
	runTokens(t, "", []tok{{EOF, ""}})
}

func TestWhitespaceOnly(t *testing.T) {
	runTokens(t, "   \t\n  ", []tok{{EOF, ""}})
}

func TestUnknownCharacter(t *testing.T) {
	runTokens(t, "@", []tok{{ILLEGAL, "@"}, {EOF, ""}})
}

func TestUnterminatedString(t *testing.T) {
	l := New(`"unterminated`)
	tok := l.NextToken()
	if tok.Type != ILLEGAL {
		t.Errorf("want ILLEGAL for unterminated string, got %s", tok.Type)
	}
}

func TestUnterminatedBlockComment(t *testing.T) {
	l := New("/* no end")
	tok := l.NextToken()
	if tok.Type != ILLEGAL {
		t.Errorf("want ILLEGAL for unterminated block comment, got %s", tok.Type)
	}
}

func TestMultipleDotsInNumber(t *testing.T) {
	// 1.2.3 should lex as NUMBER("1.2"), DOT, NUMBER("3")
	runTokens(t, "1.2.3", []tok{
		{NUMBER, "1.2"}, {DOT, "."}, {NUMBER, "3"}, {EOF, ""},
	})
}

func TestInvalidCharLiteral(t *testing.T) {
	// 'ab' is not a valid char literal
	l := New("'ab'")
	tok := l.NextToken()
	if tok.Type != ILLEGAL {
		t.Errorf("want ILLEGAL for multi-char literal, got %s", tok.Type)
	}
}

// ==========================================
// Full token stream (integration)
// ==========================================

func TestFullTokenStream(t *testing.T) {
	input := `
		# line comment
		class Gato {
			func miar() -> void {
				print("miau");
			}
		}
		let patas: i32 = 4;
		if (patas == 4) {
			seq { true }
		}
	`
	expected := []tok{
		{CLASS, "class"}, {IDENT, "Gato"}, {LBRACE, "{"},
		{FUNC, "func"}, {IDENT, "miar"}, {LPAREN, "("}, {RPAREN, ")"},
		{ARROW, "->"}, {TYPE_VOID, "void"}, {LBRACE, "{"},
		{PRINT, "print"}, {LPAREN, "("}, {STRING, "miau"}, {RPAREN, ")"}, {SEMICOLON, ";"},
		{RBRACE, "}"}, {RBRACE, "}"},
		{LET, "let"}, {IDENT, "patas"}, {COLON, ":"}, {TYPE_I32, "i32"}, {ASSIGN, "="}, {NUMBER, "4"}, {SEMICOLON, ";"},
		{IF, "if"}, {LPAREN, "("}, {IDENT, "patas"}, {EQ, "=="}, {NUMBER, "4"}, {RPAREN, ")"},
		{LBRACE, "{"}, {SEQ, "seq"}, {LBRACE, "{"}, {TRUE, "true"}, {RBRACE, "}"}, {RBRACE, "}"},
		{EOF, ""},
	}
	runTokens(t, input, expected)
}
