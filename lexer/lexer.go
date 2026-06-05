package lexer

type Lexer struct {
	input        string // O código-fonte inteiro que o usuário digitou.
	position     int    // Índice do caractere atual.
	nextPosition int    // Índice do próximo caractere.
	character    byte   // Caractere atual.
}

// New é a função construtora do Lexer. Ela recebe o código-fonte como string e retorna um ponteiro para um Lexer inicializado.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readCharacter()
	return l
}

// readCharacter é a função responsável por avançar um caractere na entrada. Ela atualiza o campo 'character' com o próximo caractere e move os índices 'position' e 'nextPosition' para frente.
func (l *Lexer) readCharacter() {
	if l.nextPosition >= len(l.input) {
		l.character = 0
	} else {
		l.character = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

// peekCharacter é uma função auxiliar que permite olhar o próximo caractere sem avançar o índice. Isso é útil para identificar operadores de dois caracteres como '==' ou '>='.
func (l *Lexer) peekCharacter() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	}
	return l.input[l.nextPosition]
}

// newToken é uma função auxiliar para criar um novo token a partir de um tipo e um caractere. Ela converte o caractere em string para preencher o campo Literal do token.
func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

// skipWhiteSpace é uma função auxiliar que avança o índice do Lexer enquanto o caractere atual for um espaço, tabulação ou quebra de linha. Isso garante que o Lexer ignore os espaços em branco entre os tokens.
func (l *Lexer) skipWhiteSpace() {
	for l.character == ' ' || l.character == '\t' || l.character == '\n' || l.character == '\r' {
		l.readCharacter()
	}
}

// isLetter verifica se um caractere é uma letra. Isso é usado para identificar o início de identificadores e palavras-chave.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit verifica se um caractere é um dígito numérico. Isso é usado para identificar o início de números.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// readIdentifier lê um identificador ou palavra-chave completo. Ele começa na posição atual e continua avançando enquanto encontrar letras ou dígitos, permitindo que identificadores possam conter números após a primeira letra.
func (l *Lexer) readIdentifier() string {
	position := l.position                              // Guarda o índice onde a palavra começou.
	for isLetter(l.character) || isDigit(l.character) { // Permite letras e números dentro da palavra.
		l.readCharacter() // Continua avançando enquanto for letra ou número.
	}
	// Recorta a string do código-fonte, do início ao fim da palavra.
	return l.input[position:l.position]
}

// readNumber lê um número completo, incluindo dígitos e um possível ponto decimal. Ele começa na posição atual e continua avançando enquanto encontrar dígitos ou um ponto.
func (l *Lexer) readNumber() string {
	position := l.position                           // Guarda o índice onde o número começou.
	for isDigit(l.character) || l.character == '.' { // Permite dígitos e um ponto decimal.
		l.readCharacter()
	}
	// Recorta a string do código-fonte, do início ao fim do número.
	return l.input[position:l.position]
}

// readString lê uma string completa, começando e terminando com aspas duplas. Ele avança até encontrar a aspa de fechamento ou o fim do arquivo, e retorna o conteúdo entre as aspas.
func (l *Lexer) readString() string {
	position := l.position + 1 // Começa depois da aspa de abertura.
	for {
		l.readCharacter()
		if l.character == '"' || l.character == 0 { // Continua até encontrar a aspa de fechamento ou o fim do arquivo.
			break
		}
	}
	// Recorta a string do código-fonte, do início ao fim da string (sem as aspas).
	return l.input[position:l.position]
}

// skipSingleLineComment avança o índice do Lexer até encontrar uma quebra de linha ou o fim do arquivo, ignorando todo o texto do comentário.
func (l *Lexer) skipSingleLineComment() {
	for l.character != '\n' && l.character != 0 {
		l.readCharacter()
	}
	l.skipWhiteSpace() // Limpa os espaços que sobrarem após o comentário.
}

// skipMultiLineComment avança o índice do Lexer até encontrar o fechamento */, ignorando todo o texto do comentário.
func (l *Lexer) skipMultiLineComment() {
	for l.character != 0 {
		if l.character == '*' && l.peekCharacter() == '/' {
			l.readCharacter() // Consome o '*'
			l.readCharacter() // Consome o '/'
			break
		}
		l.readCharacter()
	}
	l.skipWhiteSpace() // Limpa os espaços que sobrarem após o comentário.
}

// NextToken é a função principal do Lexer. Ela analisa o caractere atual e decide qual token deve ser gerado, avançando o índice conforme necessário. Ela também lida com comentários e espaços em branco, garantindo que apenas os tokens relevantes sejam retornados.
func (l *Lexer) NextToken() Token {
	var token Token

	l.skipWhiteSpace()

	for l.character == '#' || (l.character == '/' && l.peekCharacter() == '*') {
		if l.character == '#' {
			l.skipSingleLineComment()
		} else if l.character == '/' && l.peekCharacter() == '*' {
			l.skipMultiLineComment()
		}
	}

	switch l.character {

	case '=':
		if l.peekCharacter() == '=' {
			ch := l.character
			l.readCharacter()
			token = Token{Type: EQ, Literal: string(ch) + string(l.character)}
		} else {
			token = newToken(ASSIGN, l.character)
		}

	case '!':
		if l.peekCharacter() == '=' {
			ch := l.character
			l.readCharacter()
			token = Token{Type: NOT_EQ, Literal: string(ch) + string(l.character)}
		} else {
			token = newToken(BANG, l.character)
		}

	case '<':
		if l.peekCharacter() == '=' {
			ch := l.character
			l.readCharacter()
			token = Token{Type: LTE, Literal: string(ch) + string(l.character)}
		} else {
			token = newToken(LT, l.character)
		}

	case '>':
		if l.peekCharacter() == '=' {
			ch := l.character
			l.readCharacter()
			token = Token{Type: GTE, Literal: string(ch) + string(l.character)}
		} else {
			token = newToken(GT, l.character)
		}

	case '+':
		token = newToken(PLUS, l.character)

	case '-':
		token = newToken(MINUS, l.character)

	case '*':
		token = newToken(ASTERISK, l.character)

	case '/':
		token = newToken(SLASH, l.character)

	case '%':
		token = newToken(MOD, l.character)

	case ',':
		token = newToken(COMMA, l.character)

	case ';':
		token = newToken(SEMICOLON, l.character)

	case ':':
		token = newToken(COLON, l.character)

	case '.':
		token = newToken(DOT, l.character)

	case '(':
		token = newToken(LPAREN, l.character)

	case ')':
		token = newToken(RPAREN, l.character)

	case '{':
		token = newToken(LBRACE, l.character)

	case '}':
		token = newToken(RBRACE, l.character)

	case '[':
		token = newToken(LBRACKET, l.character)

	case ']':
		token = newToken(RBRACKET, l.character)

	case '"':
		token.Type = STRING
		token.Literal = l.readString()

	case 0:
		token.Literal = ""
		token.Type = EOF

	default:
		if isLetter(l.character) {
			token.Literal = l.readIdentifier()
			token.Type = LookupIdent(token.Literal)
			return token

		} else if isDigit(l.character) {
			token.Type = NUMBER
			token.Literal = l.readNumber()
			return token

		} else {
			token = newToken(ILLEGAL, l.character)
		}
	}

	l.readCharacter()

	return token
}
