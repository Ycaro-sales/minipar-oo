package lexer

type Lexer struct {
	input               string
	position            int
	nextPosition        int
	character           byte
	line                int
	unterminatedComment bool
}

func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1}
	l.readCharacter()
	return l
}

func (l *Lexer) Line() int { return l.line }

func (l *Lexer) readCharacter() {
	if l.nextPosition >= len(l.input) {
		l.character = 0
	} else {
		l.character = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition++
}

func (l *Lexer) peekCharacter() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	}
	return l.input[l.nextPosition]
}

func (l *Lexer) newToken(tt TokenType, literal string) Token {
	return Token{Type: tt, Literal: literal, Line: l.line}
}

func (l *Lexer) skipWhiteSpace() {
	for l.character == ' ' || l.character == '\t' || l.character == '\n' || l.character == '\r' {
		if l.character == '\n' {
			l.line++
		}
		l.readCharacter()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.character) || isDigit(l.character) {
		l.readCharacter()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readNumber() string {
	start := l.position
	hasDot := false
	for isDigit(l.character) || (l.character == '.' && !hasDot && isDigit(l.peekCharacter())) {
		if l.character == '.' {
			hasDot = true
		}
		l.readCharacter()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readString() string {
	start := l.position + 1
	for {
		l.readCharacter()
		if l.character == '"' || l.character == 0 {
			break
		}
		if l.character == '\n' {
			l.line++
		}
	}
	return l.input[start:l.position]
}

func (l *Lexer) skipSingleLineComment() {
	for l.character != '\n' && l.character != 0 {
		l.readCharacter()
	}
	l.skipWhiteSpace()
}

func (l *Lexer) skipMultiLineComment() {
	for l.character != 0 {
		if l.character == '\n' {
			l.line++
		}
		if l.character == '*' && l.peekCharacter() == '/' {
			l.readCharacter() // consume '*'
			l.readCharacter() // consume '/'
			l.skipWhiteSpace()
			return
		}
		l.readCharacter()
	}
	l.unterminatedComment = true
}

func (l *Lexer) NextToken() Token {
	if l.unterminatedComment {
		return l.newToken(ILLEGAL, "unterminated block comment")
	}

	l.skipWhiteSpace()

	for l.character == '#' || (l.character == '/' && l.peekCharacter() == '*') {
		if l.character == '#' {
			l.skipSingleLineComment()
		} else {
			l.skipMultiLineComment()
		}
		if l.unterminatedComment {
			return l.newToken(ILLEGAL, "unterminated block comment")
		}
	}

	switch l.character {
	case '=':
		if l.peekCharacter() == '=' {
			l.readCharacter()
			tok := l.newToken(EQ, "==")
			l.readCharacter()
			return tok
		} else if l.peekCharacter() == '>' {
			l.readCharacter()
			tok := l.newToken(FAT_ARROW, "=>")
			l.readCharacter()
			return tok
		}
		tok := l.newToken(ASSIGN, "=")
		l.readCharacter()
		return tok

	case '!':
		if l.peekCharacter() == '=' {
			l.readCharacter()
			tok := l.newToken(NOT_EQ, "!=")
			l.readCharacter()
			return tok
		}
		tok := l.newToken(BANG, "!")
		l.readCharacter()
		return tok

	case '<':
		if l.peekCharacter() == '=' {
			l.readCharacter()
			tok := l.newToken(LTE, "<=")
			l.readCharacter()
			return tok
		}
		tok := l.newToken(LT, "<")
		l.readCharacter()
		return tok

	case '>':
		if l.peekCharacter() == '=' {
			l.readCharacter()
			tok := l.newToken(GTE, ">=")
			l.readCharacter()
			return tok
		}
		tok := l.newToken(GT, ">")
		l.readCharacter()
		return tok

	case '+':
		if l.peekCharacter() == '=' {
			l.readCharacter()
			tok := l.newToken(PLUS_ASSIGN, "+=")
			l.readCharacter()
			return tok
		}
		tok := l.newToken(PLUS, "+")
		l.readCharacter()
		return tok

	case '-':
		if l.peekCharacter() == '>' {
			l.readCharacter()
			tok := l.newToken(ARROW, "->")
			l.readCharacter()
			return tok
		} else if l.peekCharacter() == '=' {
			l.readCharacter()
			tok := l.newToken(MINUS_ASSIGN, "-=")
			l.readCharacter()
			return tok
		}
		tok := l.newToken(MINUS, "-")
		l.readCharacter()
		return tok

	case '*':
		if l.peekCharacter() == '=' {
			l.readCharacter()
			tok := l.newToken(STAR_ASSIGN, "*=")
			l.readCharacter()
			return tok
		}
		tok := l.newToken(ASTERISK, "*")
		l.readCharacter()
		return tok

	case '/':
		if l.peekCharacter() == '=' {
			l.readCharacter()
			tok := l.newToken(SLASH_ASSIGN, "/=")
			l.readCharacter()
			return tok
		}
		tok := l.newToken(SLASH, "/")
		l.readCharacter()
		return tok

	case '%':
		tok := l.newToken(MOD, "%")
		l.readCharacter()
		return tok

	case ',':
		tok := l.newToken(COMMA, ",")
		l.readCharacter()
		return tok

	case ';':
		tok := l.newToken(SEMICOLON, ";")
		l.readCharacter()
		return tok

	case ':':
		tok := l.newToken(COLON, ":")
		l.readCharacter()
		return tok

	case '.':
		tok := l.newToken(DOT, ".")
		l.readCharacter()
		return tok

	case '(':
		tok := l.newToken(LPAREN, "(")
		l.readCharacter()
		return tok

	case ')':
		tok := l.newToken(RPAREN, ")")
		l.readCharacter()
		return tok

	case '{':
		tok := l.newToken(LBRACE, "{")
		l.readCharacter()
		return tok

	case '}':
		tok := l.newToken(RBRACE, "}")
		l.readCharacter()
		return tok

	case '[':
		tok := l.newToken(LBRACKET, "[")
		l.readCharacter()
		return tok

	case ']':
		tok := l.newToken(RBRACKET, "]")
		l.readCharacter()
		return tok

	case '"':
		literal := l.readString()
		if l.character == 0 {
			return l.newToken(ILLEGAL, literal)
		}
		tok := l.newToken(STRING, literal)
		l.readCharacter()
		return tok

	case '\'':
		l.readCharacter() // consume opening quote
		if l.character == 0 || l.character == '\'' {
			tok := l.newToken(ILLEGAL, "'")
			l.readCharacter()
			return tok
		}
		ch := l.character
		l.readCharacter() // consume the char itself
		if l.character != '\'' {
			return l.newToken(ILLEGAL, string(ch))
		}
		tok := l.newToken(CHAR, string(ch))
		l.readCharacter() // consume closing quote
		return tok

	case 0:
		return l.newToken(EOF, "")

	default:
		if isLetter(l.character) {
			literal := l.readIdentifier()
			return Token{Type: LookupIdent(literal), Literal: literal, Line: l.line}
		}
		if isDigit(l.character) {
			literal := l.readNumber()
			return Token{Type: NUMBER, Literal: literal, Line: l.line}
		}
		tok := l.newToken(ILLEGAL, string(l.character))
		l.readCharacter()
		return tok
	}
}
