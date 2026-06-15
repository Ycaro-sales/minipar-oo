package lexer

type TokenType int

func (t TokenType) String() string {
	if name, ok := tokenNames[t]; ok {
		return name
	}
	return "UNKNOWN"
}

var tokenNames = map[TokenType]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	IDENT:  "IDENT",
	NUMBER: "NUMBER",
	STRING: "STRING",
	CHAR:   "CHAR",

	ASSIGN:       "=",
	PLUS:         "+",
	MINUS:        "-",
	BANG:         "!",
	ASTERISK:     "*",
	SLASH:        "/",
	MOD:          "%",
	ARROW:        "->",
	FAT_ARROW:    "=>",
	PLUS_ASSIGN:  "+=",
	MINUS_ASSIGN: "-=",
	STAR_ASSIGN:  "*=",
	SLASH_ASSIGN: "/=",

	EQ:     "==",
	NOT_EQ: "!=",
	LT:     "<",
	GT:     ">",
	LTE:    "<=",
	GTE:    ">=",

	COMMA:     ",",
	SEMICOLON: ";",
	COLON:     ":",
	DOT:       ".",
	LPAREN:    "(",
	RPAREN:    ")",
	LBRACE:    "{",
	RBRACE:    "}",
	LBRACKET:  "[",
	RBRACKET:  "]",

	LET:        "LET",
	CLASS:      "CLASS",
	INTERFACE:  "INTERFACE",
	IMPLEMENTS: "IMPLEMENTS",
	SELF:       "SELF",
	FUNC:       "FUNC",
	RETURN:     "RETURN",
	BREAK:      "BREAK",
	CONTINUE:   "CONTINUE",
	PASS:       "PASS",
	GOTO:       "GOTO",

	IF:     "IF",
	ELSE:   "ELSE",
	SWITCH: "SWITCH",
	WHILE:  "WHILE",
	DO:     "DO",
	FOR:    "FOR",
	IN:     "IN",

	SEQ:       "SEQ",
	PAR:       "PAR",
	S_CHANNEL: "S_CHANNEL",
	C_CHANNEL: "C_CHANNEL",

	PRINT: "PRINT",
	INPUT: "INPUT",

	AND:   "AND",
	OR:    "OR",
	TRUE:  "TRUE",
	FALSE: "FALSE",

	TYPE_I8:     "i8",
	TYPE_I16:    "i16",
	TYPE_I32:    "i32",
	TYPE_I64:    "i64",
	TYPE_U8:     "u8",
	TYPE_U16:    "u16",
	TYPE_U32:    "u32",
	TYPE_U64:    "u64",
	TYPE_F16:    "f16",
	TYPE_F32:    "f32",
	TYPE_F64:    "f64",
	TYPE_CHAR:   "char",
	TYPE_STRING: "string",
	TYPE_BOOL:   "bool",
	TYPE_ANY:    "any",
	TYPE_VOID:   "void",
	TYPE_CHAN:   "chan",
}

const (
	ILLEGAL TokenType = iota
	EOF

	IDENT
	NUMBER
	STRING
	CHAR

	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH
	MOD
	ARROW
	FAT_ARROW
	PLUS_ASSIGN
	MINUS_ASSIGN
	STAR_ASSIGN
	SLASH_ASSIGN

	EQ
	NOT_EQ
	LT
	GT
	LTE
	GTE

	COMMA
	SEMICOLON
	COLON
	DOT
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET

	LET
	CLASS
	INTERFACE
	IMPLEMENTS
	SELF
	FUNC
	RETURN
	BREAK
	CONTINUE
	PASS
	GOTO

	IF
	ELSE
	SWITCH
	WHILE
	DO
	FOR
	IN

	SEQ
	PAR
	S_CHANNEL
	C_CHANNEL

	PRINT
	INPUT

	AND
	OR
	TRUE
	FALSE

	TYPE_I8
	TYPE_I16
	TYPE_I32
	TYPE_I64
	TYPE_U8
	TYPE_U16
	TYPE_U32
	TYPE_U64
	TYPE_F16
	TYPE_F32
	TYPE_F64
	TYPE_CHAR
	TYPE_STRING
	TYPE_BOOL
	TYPE_ANY
	TYPE_VOID
	TYPE_CHAN
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
}

var keywords = map[string]TokenType{
	"let":        LET,
	"class":      CLASS,
	"interface":  INTERFACE,
	"implements": IMPLEMENTS,
	"Self":       SELF,
	"func":       FUNC,
	"return":     RETURN,
	"break":      BREAK,
	"continue":   CONTINUE,
	"pass":       PASS,
	"goto":       GOTO,
	"if":         IF,
	"else":       ELSE,
	"switch":     SWITCH,
	"while":      WHILE,
	"do":         DO,
	"for":        FOR,
	"in":         IN,
	"seq":        SEQ,
	"par":        PAR,
	"s_channel":  S_CHANNEL,
	"c_channel":  C_CHANNEL,
	"print":      PRINT,
	"input":      INPUT,
	"and":        AND,
	"or":         OR,
	"true":       TRUE,
	"false":      FALSE,
	"i8":         TYPE_I8,
	"i16":        TYPE_I16,
	"i32":        TYPE_I32,
	"i64":        TYPE_I64,
	"u8":         TYPE_U8,
	"u16":        TYPE_U16,
	"u32":        TYPE_U32,
	"u64":        TYPE_U64,
	"f16":        TYPE_F16,
	"f32":        TYPE_F32,
	"f64":        TYPE_F64,
	"char":       TYPE_CHAR,
	"string":     TYPE_STRING,
	"bool":       TYPE_BOOL,
	"any":        TYPE_ANY,
	"void":       TYPE_VOID,
	"chan":        TYPE_CHAN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
