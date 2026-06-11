lexer grammar Linguagem;

/* PALAVRAS-CHAVE (Keywords) */
CLASS      : 'class' ;
EXTENDS    : 'extends' ;
FUNC       : 'func' ;
IF         : 'if' ;
WHILE      : 'while' ;
DO         : 'do' ;
FOR        : 'for' ;
IN         : 'in' ;
SEQ        : 'seq' ;
PAR        : 'par' ;
PRINT      : 'print' ;
INPUT      : 'input' ;
RETURN     : 'return' ;
BREAK      : 'break' ;
CONTINUE   : 'continue' ;
PASS       : 'pass' ;
S_CHANNEL  : 's_channel' ;
C_CHANNEL  : 'c_channel' ;
SEND       : 'send' ;
RECEIVE    : 'receive' ;
OR         : 'or' ;
AND        : 'and' ;
TRUE       : 'true' ;
FALSE      : 'false' ;
NEW        : 'new' ;

/* TIPOS */
T_NUMBER   : 'number' ;
T_STRING   : 'string' ;
T_BOOL     : 'bool' ;
T_VOID     : 'void' ;
T_LIST     : 'list' ;
T_DICT     : 'dict' ;

/* SÍMBOLOS E OPERADORES */
SOMA          : '+' ;
SUBTRACAO     : '-' ;
MULTIPLICACAO : '*' ;
DIVISAO       : '/' ;
MODULO        : '%' ;

IGUAL_IGUAL   : '==' ;
DIFERENTE     : '!=' ;
MENOR_IGUAL   : '<=' ;
MAIOR_IGUAL   : '>=' ;
MENOR         : '<' ;
MAIOR         : '>' ;
ATRIBUICAO    : '=' ;
NEGACAO       : '!' ;

ABRE_PAR      : '(' ;
FECHA_PAR     : ')' ;
ABRE_CHAV     : '{' ;
FECHA_CHAV    : '}' ;
ABRE_COL      : '[' ;
FECHA_COL     : ']' ;

PONTO         : '.' ;
VIRGULA       : ',' ;
DOIS_PONTOS   : ':' ;
PONTO_VIRG    : ';' ;

/* REGRAS COMPLEXAS (Terminais da BNF) */
// Número inteiro ou decimal
NUMERO : ('0' | [1-9][0-9]*) ('.' [0-9]+)? ;

// Identificadores (Variáveis, nomes de funções, etc.)
ID : [a-z] [a-zA-Z0-9]* ;

// Strings (Qualquer coisa entre aspas duplas)
STRING : '"' .*? '"' ;

/* COMENTÁRIOS E ESPAÇOS EM BRANCO (Ignorados pelo Lexer) */
COMENTARIO_LINHA : '#' ~[\r\n]* -> skip ;
COMENTARIO_BLOCO : '/*' .*? '*/' -> skip ;
WS               : [ \t\r\n]+ -> skip ;