grammar Minipar;

// ==========================================
// REGRAS SINTÁTICAS (Sempre em minúsculo)
// ==========================================

program: declaration+ EOF;

declaration: class_decl | func_decl | var_decl | channel_stmt;

/* Orientação a Objetos */
class_decl: CLASS id (EXTENDS id)? '{' class_member+ '}';
class_member: var_decl | method_decl;
method_decl: type id '(' params? ')' '{' block '}';
method_call: id '.' id '(' args? ')';
obj_creation: NEW id '(' args? ')';

/* Funções */
func_decl: FUNC id '(' params? ')' type '{' block '}';
params: param (',' param)*;
param: type id;
func_call: id '(' args? ')';
args: expr (',' expr)*;

/* Variáveis */
var_decl: type id ('=' expr)?;
assignment: id '=' expr;

/* Comandos e Blocos */
block: stmt*;
stmt: compound_stmt | simple_stmt;

compound_stmt: if_stmt
             | while_stmt
             | do_stmt
             | for_stmt
             | seq_stmt
             | par_stmt
             | channel_stmt
             | print_call
             | input_call
             | method_call ';'
             | expr ;

if_stmt: IF '(' expr ')' '{' block '}';
while_stmt: WHILE '(' expr ')' '{' block '}';
do_stmt: DO '{' block '}' WHILE '(' expr ')';
for_stmt: FOR '(' type id IN expr ')' '{' block '}';
seq_stmt: SEQ '{' block '}';
par_stmt: PAR '{' block '}';

print_call: PRINT '(' args? ')' ';';
input_call: INPUT '(' STRING? ')' ';';

simple_stmt: var_decl
           | assignment
           | expr_stmt ';'
           | RETURN expr? ';'
           | BREAK ';'
           | CONTINUE ';'
           | PASS ';';

expr_stmt: assignment | func_call;

/* Canais e Mensagens (Real Paralelismo) */
channel_stmt: s_channel_stmt | c_channel_stmt;
s_channel_stmt: S_CHANNEL id '(' args? ')' ';';
c_channel_stmt: C_CHANNEL id id id ';';
send_stmt: id '.' SEND '(' args? ')' ';';
receive_stmt: id '.' RECEIVE '(' args? ')' ';';

/* Expressões (Tratando Precedência) */
expr: or_expr;
or_expr: and_expr (OR and_expr)*;
and_expr: comparison_expr (AND comparison_expr)*;
comparison_expr: additive_expr ((EQ | NEQ | LT | GT | LTE | GTE) additive_expr)*;
additive_expr: multiplicative_expr ((PLUS | MINUS) multiplicative_expr)*;
multiplicative_expr: unary_expr ((MUL | DIV | MOD) unary_expr)*;
unary_expr: (MINUS | NOT) unary_expr | postfix_expr;
postfix_expr: postfix_expr '[' expr ']' | primary_expr | method_call | func_call;

primary_expr: NUMERO
            | STRING
            | TRUE
            | FALSE
            | id
            | obj_creation
            | list_literal
            | dict_literal
            | '(' expr ')';

list_literal: '[' (expr (',' expr)*)? ']';
dict_literal: '{' (key_value_pair (',' key_value_pair)*)? '}';
key_value_pair: STRING ':' expr;

/* Tipos */
type: NUMBER_TYPE | STRING_TYPE | BOOL_TYPE | VOID_TYPE | LIST_TYPE | DICT_TYPE | id;

// Regra auxiliar para o não-terminal <id>
id: ID;

// ==========================================
// REGRAS LÉXICAS (Tokens - Sempre em MAIÚSCULO)
// ==========================================

CLASS: 'class';
EXTENDS: 'extends';
NEW: 'new';
FUNC: 'func';
IF: 'if';
WHILE: 'while';
DO: 'do';
FOR: 'for';
IN: 'in';
SEQ: 'seq';
PAR: 'par';
PRINT: 'print';
INPUT: 'input';
RETURN: 'return';
BREAK: 'break';
CONTINUE: 'continue';
PASS: 'pass';
S_CHANNEL: 's_channel';
C_CHANNEL: 'c_channel';
SEND: 'send';
RECEIVE: 'receive';

NUMBER_TYPE: 'number';
STRING_TYPE: 'string';
BOOL_TYPE: 'bool';
VOID_TYPE: 'void';
LIST_TYPE: 'list';
DICT_TYPE: 'dict';

TRUE: 'true';
FALSE: 'false';
OR: 'or';
AND: 'and';
NOT: '!';

EQ: '==';
NEQ: '!=';
LT: '<';
GT: '>';
LTE: '<=';
GTE: '>=';
PLUS: '+';
MINUS: '-';
MUL: '*';
DIV: '/';
MOD: '%';

ID: [a-zA-Z_] [a-zA-Z0-9_]*;
NUMERO: ('0' | [1-9][0-9]*) ('.' [0-9]+)?;
STRING: '"' ~["]* '"'; 

WS: [ \t\r\n]+ -> skip;
INLINE_COMMENT: '#' ~[\r\n]* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;