package lexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Especiais
	ILLEGAL = "ILLEGAL" // Para caracteres que não pertencem à linguagem
	EOF     = "EOF"     // Fim do arquivo (End Of File)

	// Literais e identificadores
	IDENT  = "IDENT"  // Nomes criados pelo usuário
	NUMBER = "NUMBER" // Números inteiros ou decimais
	STRING = "STRING" // Texto entre aspas duplas

	// Operadores matemáticos e lógicos
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	MOD      = "%"

	// Operadores de comparação
	EQ     = "=="
	NOT_EQ = "!="
	LT     = "<"
	GT     = ">"
	LTE    = "<="
	GTE    = ">="

	// Pontuação e delimitadores
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	DOT       = "."
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Orientação a Objetos
	CLASS   = "CLASS"
	EXTENDS = "EXTENDS"
	NEW     = "NEW"

	// Funções e retornos
	FUNC     = "FUNC"
	RETURN   = "RETURN"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	PASS     = "PASS"

	// Controle de fluxo
	IF    = "IF"
	WHILE = "WHILE"
	DO    = "DO"
	FOR   = "FOR"
	IN    = "IN"

	// Concorrência e blocos
	SEQ = "SEQ"
	PAR = "PAR"

	// Input/Output e comunicação
	PRINT     = "PRINT"
	INPUT     = "INPUT"
	S_CHANNEL = "S_CHANNEL"
	C_CHANNEL = "C_CHANNEL"
	SEND      = "SEND"
	RECEIVE   = "RECEIVE"

	// Lógica e booleanos
	AND   = "AND"
	OR    = "OR"
	TRUE  = "TRUE"
	FALSE = "FALSE"

	// Tipos de dados
	TYPE_NUMBER = "TYPE_NUMBER"
	TYPE_STRING = "TYPE_STRING"
	TYPE_BOOL   = "TYPE_BOOL"
	TYPE_VOID   = "TYPE_VOID"
	TYPE_LIST   = "TYPE_LIST"
	TYPE_DICT   = "TYPE_DICT"
)

var keywords = map[string]TokenType{
	"class":     CLASS,
	"extends":   EXTENDS,
	"new":       NEW,
	"func":      FUNC,
	"return":    RETURN,
	"break":     BREAK,
	"continue":  CONTINUE,
	"pass":      PASS,
	"if":        IF,
	"while":     WHILE,
	"do":        DO,
	"for":       FOR,
	"in":        IN,
	"seq":       SEQ,
	"par":       PAR,
	"print":     PRINT,
	"input":     INPUT,
	"s_channel": S_CHANNEL,
	"c_channel": C_CHANNEL,
	"send":      SEND,
	"receive":   RECEIVE,
	"and":       AND,
	"or":        OR,
	"true":      TRUE,
	"false":     FALSE,
	"number":    TYPE_NUMBER,
	"string":    TYPE_STRING,
	"bool":      TYPE_BOOL,
	"void":      TYPE_VOID,
	"list":      TYPE_LIST,
	"dict":      TYPE_DICT,
}

func LookupIdent(ident string) TokenType {
	if token, ok := keywords[ident]; ok {
		return token
	}
	return IDENT
}
