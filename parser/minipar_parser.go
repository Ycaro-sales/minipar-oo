// Code generated from Minipar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Minipar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MiniparParser struct {
	*antlr.BaseParser
}

var MiniparParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func miniparParserInit() {
	staticData := &MiniparParserStaticData
	staticData.LiteralNames = []string{
		"", "'{'", "'}'", "'('", "')'", "'.'", "','", "'='", "';'", "'['", "']'",
		"':'", "'class'", "'extends'", "'new'", "'func'", "'if'", "'while'",
		"'do'", "'for'", "'in'", "'seq'", "'par'", "'print'", "'input'", "'return'",
		"'break'", "'continue'", "'pass'", "'s_channel'", "'c_channel'", "'send'",
		"'receive'", "'number'", "'string'", "'bool'", "'void'", "'list'", "'dict'",
		"'true'", "'false'", "'or'", "'and'", "'!'", "'=='", "'!='", "'<'",
		"'>'", "'<='", "'>='", "'+'", "'-'", "'*'", "'/'", "'%'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "CLASS", "EXTENDS",
		"NEW", "FUNC", "IF", "WHILE", "DO", "FOR", "IN", "SEQ", "PAR", "PRINT",
		"INPUT", "RETURN", "BREAK", "CONTINUE", "PASS", "S_CHANNEL", "C_CHANNEL",
		"SEND", "RECEIVE", "NUMBER_TYPE", "STRING_TYPE", "BOOL_TYPE", "VOID_TYPE",
		"LIST_TYPE", "DICT_TYPE", "TRUE", "FALSE", "OR", "AND", "NOT", "EQ",
		"NEQ", "LT", "GT", "LTE", "GTE", "PLUS", "MINUS", "MUL", "DIV", "MOD",
		"ID", "NUMERO", "STRING", "WS", "INLINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "declaration", "class_decl", "class_member", "method_decl",
		"method_call", "obj_creation", "func_decl", "params", "param", "func_call",
		"args", "var_decl", "assignment", "block", "stmt", "compound_stmt",
		"if_stmt", "while_stmt", "do_stmt", "for_stmt", "seq_stmt", "par_stmt",
		"print_call", "input_call", "simple_stmt", "expr_stmt", "channel_stmt",
		"s_channel_stmt", "c_channel_stmt", "send_stmt", "receive_stmt", "expr",
		"or_expr", "and_expr", "comparison_expr", "additive_expr", "multiplicative_expr",
		"unary_expr", "postfix_expr", "primary_expr", "list_literal", "dict_literal",
		"key_value_pair", "type", "id",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 60, 466, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 1, 0, 4, 0, 94, 8,
		0, 11, 0, 12, 0, 95, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 104, 8,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 110, 8, 2, 1, 2, 1, 2, 4, 2, 114, 8, 2,
		11, 2, 12, 2, 115, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 122, 8, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 128, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 3, 5, 140, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 148, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 156, 8, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 167, 8, 8, 10,
		8, 12, 8, 170, 9, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 3, 10, 178,
		8, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 5, 11, 185, 8, 11, 10, 11, 12,
		11, 188, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 194, 8, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 14, 5, 14, 201, 8, 14, 10, 14, 12, 14, 204, 9, 14,
		1, 15, 1, 15, 3, 15, 208, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 223, 8, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 3, 23, 274, 8, 23, 1,
		23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 3, 24, 282, 8, 24, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 294, 8,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 303, 8, 25,
		1, 26, 1, 26, 3, 26, 307, 8, 26, 1, 27, 1, 27, 3, 27, 311, 8, 27, 1, 28,
		1, 28, 1, 28, 1, 28, 3, 28, 317, 8, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30,
		333, 8, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3,
		31, 343, 8, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33,
		5, 33, 353, 8, 33, 10, 33, 12, 33, 356, 9, 33, 1, 34, 1, 34, 1, 34, 5,
		34, 361, 8, 34, 10, 34, 12, 34, 364, 9, 34, 1, 35, 1, 35, 1, 35, 5, 35,
		369, 8, 35, 10, 35, 12, 35, 372, 9, 35, 1, 36, 1, 36, 1, 36, 5, 36, 377,
		8, 36, 10, 36, 12, 36, 380, 9, 36, 1, 37, 1, 37, 1, 37, 5, 37, 385, 8,
		37, 10, 37, 12, 37, 388, 9, 37, 1, 38, 1, 38, 1, 38, 3, 38, 393, 8, 38,
		1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 399, 8, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 5, 39, 406, 8, 39, 10, 39, 12, 39, 409, 9, 39, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 3,
		40, 423, 8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 5, 41, 429, 8, 41, 10, 41,
		12, 41, 432, 9, 41, 3, 41, 434, 8, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42,
		1, 42, 5, 42, 442, 8, 42, 10, 42, 12, 42, 445, 9, 42, 3, 42, 447, 8, 42,
		1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1,
		44, 1, 44, 1, 44, 3, 44, 462, 8, 44, 1, 45, 1, 45, 1, 45, 0, 1, 78, 46,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
		38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
		74, 76, 78, 80, 82, 84, 86, 88, 90, 0, 4, 1, 0, 44, 49, 1, 0, 50, 51, 1,
		0, 52, 54, 2, 0, 43, 43, 51, 51, 487, 0, 93, 1, 0, 0, 0, 2, 103, 1, 0,
		0, 0, 4, 105, 1, 0, 0, 0, 6, 121, 1, 0, 0, 0, 8, 123, 1, 0, 0, 0, 10, 134,
		1, 0, 0, 0, 12, 143, 1, 0, 0, 0, 14, 151, 1, 0, 0, 0, 16, 163, 1, 0, 0,
		0, 18, 171, 1, 0, 0, 0, 20, 174, 1, 0, 0, 0, 22, 181, 1, 0, 0, 0, 24, 189,
		1, 0, 0, 0, 26, 195, 1, 0, 0, 0, 28, 202, 1, 0, 0, 0, 30, 207, 1, 0, 0,
		0, 32, 222, 1, 0, 0, 0, 34, 224, 1, 0, 0, 0, 36, 232, 1, 0, 0, 0, 38, 240,
		1, 0, 0, 0, 40, 249, 1, 0, 0, 0, 42, 260, 1, 0, 0, 0, 44, 265, 1, 0, 0,
		0, 46, 270, 1, 0, 0, 0, 48, 278, 1, 0, 0, 0, 50, 302, 1, 0, 0, 0, 52, 306,
		1, 0, 0, 0, 54, 310, 1, 0, 0, 0, 56, 312, 1, 0, 0, 0, 58, 321, 1, 0, 0,
		0, 60, 327, 1, 0, 0, 0, 62, 337, 1, 0, 0, 0, 64, 347, 1, 0, 0, 0, 66, 349,
		1, 0, 0, 0, 68, 357, 1, 0, 0, 0, 70, 365, 1, 0, 0, 0, 72, 373, 1, 0, 0,
		0, 74, 381, 1, 0, 0, 0, 76, 392, 1, 0, 0, 0, 78, 398, 1, 0, 0, 0, 80, 422,
		1, 0, 0, 0, 82, 424, 1, 0, 0, 0, 84, 437, 1, 0, 0, 0, 86, 450, 1, 0, 0,
		0, 88, 461, 1, 0, 0, 0, 90, 463, 1, 0, 0, 0, 92, 94, 3, 2, 1, 0, 93, 92,
		1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0,
		96, 97, 1, 0, 0, 0, 97, 98, 5, 0, 0, 1, 98, 1, 1, 0, 0, 0, 99, 104, 3,
		4, 2, 0, 100, 104, 3, 14, 7, 0, 101, 104, 3, 24, 12, 0, 102, 104, 3, 54,
		27, 0, 103, 99, 1, 0, 0, 0, 103, 100, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0,
		103, 102, 1, 0, 0, 0, 104, 3, 1, 0, 0, 0, 105, 106, 5, 12, 0, 0, 106, 109,
		3, 90, 45, 0, 107, 108, 5, 13, 0, 0, 108, 110, 3, 90, 45, 0, 109, 107,
		1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 113, 5, 1,
		0, 0, 112, 114, 3, 6, 3, 0, 113, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0,
		115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117,
		118, 5, 2, 0, 0, 118, 5, 1, 0, 0, 0, 119, 122, 3, 24, 12, 0, 120, 122,
		3, 8, 4, 0, 121, 119, 1, 0, 0, 0, 121, 120, 1, 0, 0, 0, 122, 7, 1, 0, 0,
		0, 123, 124, 3, 88, 44, 0, 124, 125, 3, 90, 45, 0, 125, 127, 5, 3, 0, 0,
		126, 128, 3, 16, 8, 0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128,
		129, 1, 0, 0, 0, 129, 130, 5, 4, 0, 0, 130, 131, 5, 1, 0, 0, 131, 132,
		3, 28, 14, 0, 132, 133, 5, 2, 0, 0, 133, 9, 1, 0, 0, 0, 134, 135, 3, 90,
		45, 0, 135, 136, 5, 5, 0, 0, 136, 137, 3, 90, 45, 0, 137, 139, 5, 3, 0,
		0, 138, 140, 3, 22, 11, 0, 139, 138, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0,
		140, 141, 1, 0, 0, 0, 141, 142, 5, 4, 0, 0, 142, 11, 1, 0, 0, 0, 143, 144,
		5, 14, 0, 0, 144, 145, 3, 90, 45, 0, 145, 147, 5, 3, 0, 0, 146, 148, 3,
		22, 11, 0, 147, 146, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 1, 0,
		0, 0, 149, 150, 5, 4, 0, 0, 150, 13, 1, 0, 0, 0, 151, 152, 5, 15, 0, 0,
		152, 153, 3, 90, 45, 0, 153, 155, 5, 3, 0, 0, 154, 156, 3, 16, 8, 0, 155,
		154, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158,
		5, 4, 0, 0, 158, 159, 3, 88, 44, 0, 159, 160, 5, 1, 0, 0, 160, 161, 3,
		28, 14, 0, 161, 162, 5, 2, 0, 0, 162, 15, 1, 0, 0, 0, 163, 168, 3, 18,
		9, 0, 164, 165, 5, 6, 0, 0, 165, 167, 3, 18, 9, 0, 166, 164, 1, 0, 0, 0,
		167, 170, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169,
		17, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 171, 172, 3, 88, 44, 0, 172, 173,
		3, 90, 45, 0, 173, 19, 1, 0, 0, 0, 174, 175, 3, 90, 45, 0, 175, 177, 5,
		3, 0, 0, 176, 178, 3, 22, 11, 0, 177, 176, 1, 0, 0, 0, 177, 178, 1, 0,
		0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 5, 4, 0, 0, 180, 21, 1, 0, 0, 0,
		181, 186, 3, 64, 32, 0, 182, 183, 5, 6, 0, 0, 183, 185, 3, 64, 32, 0, 184,
		182, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187,
		1, 0, 0, 0, 187, 23, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189, 190, 3, 88,
		44, 0, 190, 193, 3, 90, 45, 0, 191, 192, 5, 7, 0, 0, 192, 194, 3, 64, 32,
		0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 25, 1, 0, 0, 0, 195,
		196, 3, 90, 45, 0, 196, 197, 5, 7, 0, 0, 197, 198, 3, 64, 32, 0, 198, 27,
		1, 0, 0, 0, 199, 201, 3, 30, 15, 0, 200, 199, 1, 0, 0, 0, 201, 204, 1,
		0, 0, 0, 202, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 29, 1, 0, 0,
		0, 204, 202, 1, 0, 0, 0, 205, 208, 3, 32, 16, 0, 206, 208, 3, 50, 25, 0,
		207, 205, 1, 0, 0, 0, 207, 206, 1, 0, 0, 0, 208, 31, 1, 0, 0, 0, 209, 223,
		3, 34, 17, 0, 210, 223, 3, 36, 18, 0, 211, 223, 3, 38, 19, 0, 212, 223,
		3, 40, 20, 0, 213, 223, 3, 42, 21, 0, 214, 223, 3, 44, 22, 0, 215, 223,
		3, 54, 27, 0, 216, 223, 3, 46, 23, 0, 217, 223, 3, 48, 24, 0, 218, 219,
		3, 10, 5, 0, 219, 220, 5, 8, 0, 0, 220, 223, 1, 0, 0, 0, 221, 223, 3, 64,
		32, 0, 222, 209, 1, 0, 0, 0, 222, 210, 1, 0, 0, 0, 222, 211, 1, 0, 0, 0,
		222, 212, 1, 0, 0, 0, 222, 213, 1, 0, 0, 0, 222, 214, 1, 0, 0, 0, 222,
		215, 1, 0, 0, 0, 222, 216, 1, 0, 0, 0, 222, 217, 1, 0, 0, 0, 222, 218,
		1, 0, 0, 0, 222, 221, 1, 0, 0, 0, 223, 33, 1, 0, 0, 0, 224, 225, 5, 16,
		0, 0, 225, 226, 5, 3, 0, 0, 226, 227, 3, 64, 32, 0, 227, 228, 5, 4, 0,
		0, 228, 229, 5, 1, 0, 0, 229, 230, 3, 28, 14, 0, 230, 231, 5, 2, 0, 0,
		231, 35, 1, 0, 0, 0, 232, 233, 5, 17, 0, 0, 233, 234, 5, 3, 0, 0, 234,
		235, 3, 64, 32, 0, 235, 236, 5, 4, 0, 0, 236, 237, 5, 1, 0, 0, 237, 238,
		3, 28, 14, 0, 238, 239, 5, 2, 0, 0, 239, 37, 1, 0, 0, 0, 240, 241, 5, 18,
		0, 0, 241, 242, 5, 1, 0, 0, 242, 243, 3, 28, 14, 0, 243, 244, 5, 2, 0,
		0, 244, 245, 5, 17, 0, 0, 245, 246, 5, 3, 0, 0, 246, 247, 3, 64, 32, 0,
		247, 248, 5, 4, 0, 0, 248, 39, 1, 0, 0, 0, 249, 250, 5, 19, 0, 0, 250,
		251, 5, 3, 0, 0, 251, 252, 3, 88, 44, 0, 252, 253, 3, 90, 45, 0, 253, 254,
		5, 20, 0, 0, 254, 255, 3, 64, 32, 0, 255, 256, 5, 4, 0, 0, 256, 257, 5,
		1, 0, 0, 257, 258, 3, 28, 14, 0, 258, 259, 5, 2, 0, 0, 259, 41, 1, 0, 0,
		0, 260, 261, 5, 21, 0, 0, 261, 262, 5, 1, 0, 0, 262, 263, 3, 28, 14, 0,
		263, 264, 5, 2, 0, 0, 264, 43, 1, 0, 0, 0, 265, 266, 5, 22, 0, 0, 266,
		267, 5, 1, 0, 0, 267, 268, 3, 28, 14, 0, 268, 269, 5, 2, 0, 0, 269, 45,
		1, 0, 0, 0, 270, 271, 5, 23, 0, 0, 271, 273, 5, 3, 0, 0, 272, 274, 3, 22,
		11, 0, 273, 272, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0,
		275, 276, 5, 4, 0, 0, 276, 277, 5, 8, 0, 0, 277, 47, 1, 0, 0, 0, 278, 279,
		5, 24, 0, 0, 279, 281, 5, 3, 0, 0, 280, 282, 5, 57, 0, 0, 281, 280, 1,
		0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 284, 5, 4, 0,
		0, 284, 285, 5, 8, 0, 0, 285, 49, 1, 0, 0, 0, 286, 303, 3, 24, 12, 0, 287,
		303, 3, 26, 13, 0, 288, 289, 3, 52, 26, 0, 289, 290, 5, 8, 0, 0, 290, 303,
		1, 0, 0, 0, 291, 293, 5, 25, 0, 0, 292, 294, 3, 64, 32, 0, 293, 292, 1,
		0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 303, 5, 8, 0,
		0, 296, 297, 5, 26, 0, 0, 297, 303, 5, 8, 0, 0, 298, 299, 5, 27, 0, 0,
		299, 303, 5, 8, 0, 0, 300, 301, 5, 28, 0, 0, 301, 303, 5, 8, 0, 0, 302,
		286, 1, 0, 0, 0, 302, 287, 1, 0, 0, 0, 302, 288, 1, 0, 0, 0, 302, 291,
		1, 0, 0, 0, 302, 296, 1, 0, 0, 0, 302, 298, 1, 0, 0, 0, 302, 300, 1, 0,
		0, 0, 303, 51, 1, 0, 0, 0, 304, 307, 3, 26, 13, 0, 305, 307, 3, 20, 10,
		0, 306, 304, 1, 0, 0, 0, 306, 305, 1, 0, 0, 0, 307, 53, 1, 0, 0, 0, 308,
		311, 3, 56, 28, 0, 309, 311, 3, 58, 29, 0, 310, 308, 1, 0, 0, 0, 310, 309,
		1, 0, 0, 0, 311, 55, 1, 0, 0, 0, 312, 313, 5, 29, 0, 0, 313, 314, 3, 90,
		45, 0, 314, 316, 5, 3, 0, 0, 315, 317, 3, 22, 11, 0, 316, 315, 1, 0, 0,
		0, 316, 317, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0, 318, 319, 5, 4, 0, 0, 319,
		320, 5, 8, 0, 0, 320, 57, 1, 0, 0, 0, 321, 322, 5, 30, 0, 0, 322, 323,
		3, 90, 45, 0, 323, 324, 3, 90, 45, 0, 324, 325, 3, 90, 45, 0, 325, 326,
		5, 8, 0, 0, 326, 59, 1, 0, 0, 0, 327, 328, 3, 90, 45, 0, 328, 329, 5, 5,
		0, 0, 329, 330, 5, 31, 0, 0, 330, 332, 5, 3, 0, 0, 331, 333, 3, 22, 11,
		0, 332, 331, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334,
		335, 5, 4, 0, 0, 335, 336, 5, 8, 0, 0, 336, 61, 1, 0, 0, 0, 337, 338, 3,
		90, 45, 0, 338, 339, 5, 5, 0, 0, 339, 340, 5, 32, 0, 0, 340, 342, 5, 3,
		0, 0, 341, 343, 3, 22, 11, 0, 342, 341, 1, 0, 0, 0, 342, 343, 1, 0, 0,
		0, 343, 344, 1, 0, 0, 0, 344, 345, 5, 4, 0, 0, 345, 346, 5, 8, 0, 0, 346,
		63, 1, 0, 0, 0, 347, 348, 3, 66, 33, 0, 348, 65, 1, 0, 0, 0, 349, 354,
		3, 68, 34, 0, 350, 351, 5, 41, 0, 0, 351, 353, 3, 68, 34, 0, 352, 350,
		1, 0, 0, 0, 353, 356, 1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 354, 355, 1, 0,
		0, 0, 355, 67, 1, 0, 0, 0, 356, 354, 1, 0, 0, 0, 357, 362, 3, 70, 35, 0,
		358, 359, 5, 42, 0, 0, 359, 361, 3, 70, 35, 0, 360, 358, 1, 0, 0, 0, 361,
		364, 1, 0, 0, 0, 362, 360, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 69, 1,
		0, 0, 0, 364, 362, 1, 0, 0, 0, 365, 370, 3, 72, 36, 0, 366, 367, 7, 0,
		0, 0, 367, 369, 3, 72, 36, 0, 368, 366, 1, 0, 0, 0, 369, 372, 1, 0, 0,
		0, 370, 368, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 71, 1, 0, 0, 0, 372,
		370, 1, 0, 0, 0, 373, 378, 3, 74, 37, 0, 374, 375, 7, 1, 0, 0, 375, 377,
		3, 74, 37, 0, 376, 374, 1, 0, 0, 0, 377, 380, 1, 0, 0, 0, 378, 376, 1,
		0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 73, 1, 0, 0, 0, 380, 378, 1, 0, 0,
		0, 381, 386, 3, 76, 38, 0, 382, 383, 7, 2, 0, 0, 383, 385, 3, 76, 38, 0,
		384, 382, 1, 0, 0, 0, 385, 388, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 386,
		387, 1, 0, 0, 0, 387, 75, 1, 0, 0, 0, 388, 386, 1, 0, 0, 0, 389, 390, 7,
		3, 0, 0, 390, 393, 3, 76, 38, 0, 391, 393, 3, 78, 39, 0, 392, 389, 1, 0,
		0, 0, 392, 391, 1, 0, 0, 0, 393, 77, 1, 0, 0, 0, 394, 395, 6, 39, -1, 0,
		395, 399, 3, 80, 40, 0, 396, 399, 3, 10, 5, 0, 397, 399, 3, 20, 10, 0,
		398, 394, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 398, 397, 1, 0, 0, 0, 399,
		407, 1, 0, 0, 0, 400, 401, 10, 4, 0, 0, 401, 402, 5, 9, 0, 0, 402, 403,
		3, 64, 32, 0, 403, 404, 5, 10, 0, 0, 404, 406, 1, 0, 0, 0, 405, 400, 1,
		0, 0, 0, 406, 409, 1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 407, 408, 1, 0, 0,
		0, 408, 79, 1, 0, 0, 0, 409, 407, 1, 0, 0, 0, 410, 423, 5, 56, 0, 0, 411,
		423, 5, 57, 0, 0, 412, 423, 5, 39, 0, 0, 413, 423, 5, 40, 0, 0, 414, 423,
		3, 90, 45, 0, 415, 423, 3, 12, 6, 0, 416, 423, 3, 82, 41, 0, 417, 423,
		3, 84, 42, 0, 418, 419, 5, 3, 0, 0, 419, 420, 3, 64, 32, 0, 420, 421, 5,
		4, 0, 0, 421, 423, 1, 0, 0, 0, 422, 410, 1, 0, 0, 0, 422, 411, 1, 0, 0,
		0, 422, 412, 1, 0, 0, 0, 422, 413, 1, 0, 0, 0, 422, 414, 1, 0, 0, 0, 422,
		415, 1, 0, 0, 0, 422, 416, 1, 0, 0, 0, 422, 417, 1, 0, 0, 0, 422, 418,
		1, 0, 0, 0, 423, 81, 1, 0, 0, 0, 424, 433, 5, 9, 0, 0, 425, 430, 3, 64,
		32, 0, 426, 427, 5, 6, 0, 0, 427, 429, 3, 64, 32, 0, 428, 426, 1, 0, 0,
		0, 429, 432, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0, 430, 431, 1, 0, 0, 0, 431,
		434, 1, 0, 0, 0, 432, 430, 1, 0, 0, 0, 433, 425, 1, 0, 0, 0, 433, 434,
		1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 436, 5, 10, 0, 0, 436, 83, 1, 0,
		0, 0, 437, 446, 5, 1, 0, 0, 438, 443, 3, 86, 43, 0, 439, 440, 5, 6, 0,
		0, 440, 442, 3, 86, 43, 0, 441, 439, 1, 0, 0, 0, 442, 445, 1, 0, 0, 0,
		443, 441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444, 447, 1, 0, 0, 0, 445,
		443, 1, 0, 0, 0, 446, 438, 1, 0, 0, 0, 446, 447, 1, 0, 0, 0, 447, 448,
		1, 0, 0, 0, 448, 449, 5, 2, 0, 0, 449, 85, 1, 0, 0, 0, 450, 451, 5, 57,
		0, 0, 451, 452, 5, 11, 0, 0, 452, 453, 3, 64, 32, 0, 453, 87, 1, 0, 0,
		0, 454, 462, 5, 33, 0, 0, 455, 462, 5, 34, 0, 0, 456, 462, 5, 35, 0, 0,
		457, 462, 5, 36, 0, 0, 458, 462, 5, 37, 0, 0, 459, 462, 5, 38, 0, 0, 460,
		462, 3, 90, 45, 0, 461, 454, 1, 0, 0, 0, 461, 455, 1, 0, 0, 0, 461, 456,
		1, 0, 0, 0, 461, 457, 1, 0, 0, 0, 461, 458, 1, 0, 0, 0, 461, 459, 1, 0,
		0, 0, 461, 460, 1, 0, 0, 0, 462, 89, 1, 0, 0, 0, 463, 464, 5, 55, 0, 0,
		464, 91, 1, 0, 0, 0, 39, 95, 103, 109, 115, 121, 127, 139, 147, 155, 168,
		177, 186, 193, 202, 207, 222, 273, 281, 293, 302, 306, 310, 316, 332, 342,
		354, 362, 370, 378, 386, 392, 398, 407, 422, 430, 433, 443, 446, 461,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MiniparParserInit initializes any static state used to implement MiniparParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMiniparParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MiniparParserInit() {
	staticData := &MiniparParserStaticData
	staticData.once.Do(miniparParserInit)
}

// NewMiniparParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMiniparParser(input antlr.TokenStream) *MiniparParser {
	MiniparParserInit()
	this := new(MiniparParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MiniparParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Minipar.g4"

	return this
}

// MiniparParser tokens.
const (
	MiniparParserEOF            = antlr.TokenEOF
	MiniparParserT__0           = 1
	MiniparParserT__1           = 2
	MiniparParserT__2           = 3
	MiniparParserT__3           = 4
	MiniparParserT__4           = 5
	MiniparParserT__5           = 6
	MiniparParserT__6           = 7
	MiniparParserT__7           = 8
	MiniparParserT__8           = 9
	MiniparParserT__9           = 10
	MiniparParserT__10          = 11
	MiniparParserCLASS          = 12
	MiniparParserEXTENDS        = 13
	MiniparParserNEW            = 14
	MiniparParserFUNC           = 15
	MiniparParserIF             = 16
	MiniparParserWHILE          = 17
	MiniparParserDO             = 18
	MiniparParserFOR            = 19
	MiniparParserIN             = 20
	MiniparParserSEQ            = 21
	MiniparParserPAR            = 22
	MiniparParserPRINT          = 23
	MiniparParserINPUT          = 24
	MiniparParserRETURN         = 25
	MiniparParserBREAK          = 26
	MiniparParserCONTINUE       = 27
	MiniparParserPASS           = 28
	MiniparParserS_CHANNEL      = 29
	MiniparParserC_CHANNEL      = 30
	MiniparParserSEND           = 31
	MiniparParserRECEIVE        = 32
	MiniparParserNUMBER_TYPE    = 33
	MiniparParserSTRING_TYPE    = 34
	MiniparParserBOOL_TYPE      = 35
	MiniparParserVOID_TYPE      = 36
	MiniparParserLIST_TYPE      = 37
	MiniparParserDICT_TYPE      = 38
	MiniparParserTRUE           = 39
	MiniparParserFALSE          = 40
	MiniparParserOR             = 41
	MiniparParserAND            = 42
	MiniparParserNOT            = 43
	MiniparParserEQ             = 44
	MiniparParserNEQ            = 45
	MiniparParserLT             = 46
	MiniparParserGT             = 47
	MiniparParserLTE            = 48
	MiniparParserGTE            = 49
	MiniparParserPLUS           = 50
	MiniparParserMINUS          = 51
	MiniparParserMUL            = 52
	MiniparParserDIV            = 53
	MiniparParserMOD            = 54
	MiniparParserID             = 55
	MiniparParserNUMERO         = 56
	MiniparParserSTRING         = 57
	MiniparParserWS             = 58
	MiniparParserINLINE_COMMENT = 59
	MiniparParserBLOCK_COMMENT  = 60
)

// MiniparParser rules.
const (
	MiniparParserRULE_program             = 0
	MiniparParserRULE_declaration         = 1
	MiniparParserRULE_class_decl          = 2
	MiniparParserRULE_class_member        = 3
	MiniparParserRULE_method_decl         = 4
	MiniparParserRULE_method_call         = 5
	MiniparParserRULE_obj_creation        = 6
	MiniparParserRULE_func_decl           = 7
	MiniparParserRULE_params              = 8
	MiniparParserRULE_param               = 9
	MiniparParserRULE_func_call           = 10
	MiniparParserRULE_args                = 11
	MiniparParserRULE_var_decl            = 12
	MiniparParserRULE_assignment          = 13
	MiniparParserRULE_block               = 14
	MiniparParserRULE_stmt                = 15
	MiniparParserRULE_compound_stmt       = 16
	MiniparParserRULE_if_stmt             = 17
	MiniparParserRULE_while_stmt          = 18
	MiniparParserRULE_do_stmt             = 19
	MiniparParserRULE_for_stmt            = 20
	MiniparParserRULE_seq_stmt            = 21
	MiniparParserRULE_par_stmt            = 22
	MiniparParserRULE_print_call          = 23
	MiniparParserRULE_input_call          = 24
	MiniparParserRULE_simple_stmt         = 25
	MiniparParserRULE_expr_stmt           = 26
	MiniparParserRULE_channel_stmt        = 27
	MiniparParserRULE_s_channel_stmt      = 28
	MiniparParserRULE_c_channel_stmt      = 29
	MiniparParserRULE_send_stmt           = 30
	MiniparParserRULE_receive_stmt        = 31
	MiniparParserRULE_expr                = 32
	MiniparParserRULE_or_expr             = 33
	MiniparParserRULE_and_expr            = 34
	MiniparParserRULE_comparison_expr     = 35
	MiniparParserRULE_additive_expr       = 36
	MiniparParserRULE_multiplicative_expr = 37
	MiniparParserRULE_unary_expr          = 38
	MiniparParserRULE_postfix_expr        = 39
	MiniparParserRULE_primary_expr        = 40
	MiniparParserRULE_list_literal        = 41
	MiniparParserRULE_dict_literal        = 42
	MiniparParserRULE_key_value_pair      = 43
	MiniparParserRULE_type                = 44
	MiniparParserRULE_id                  = 45
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(MiniparParserEOF, 0)
}

func (s *ProgramContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *MiniparParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MiniparParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36029339795492864) != 0) {
		{
			p.SetState(92)
			p.Declaration()
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(97)
		p.Match(MiniparParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Class_decl() IClass_declContext
	Func_decl() IFunc_declContext
	Var_decl() IVar_declContext
	Channel_stmt() IChannel_stmtContext

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Class_decl() IClass_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClass_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClass_declContext)
}

func (s *DeclarationContext) Func_decl() IFunc_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_declContext)
}

func (s *DeclarationContext) Var_decl() IVar_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *DeclarationContext) Channel_stmt() IChannel_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChannel_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChannel_stmtContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *MiniparParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MiniparParserRULE_declaration)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniparParserCLASS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.Class_decl()
		}

	case MiniparParserFUNC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.Func_decl()
		}

	case MiniparParserNUMBER_TYPE, MiniparParserSTRING_TYPE, MiniparParserBOOL_TYPE, MiniparParserVOID_TYPE, MiniparParserLIST_TYPE, MiniparParserDICT_TYPE, MiniparParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(101)
			p.Var_decl()
		}

	case MiniparParserS_CHANNEL, MiniparParserC_CHANNEL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(102)
			p.Channel_stmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClass_declContext is an interface to support dynamic dispatch.
type IClass_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CLASS() antlr.TerminalNode
	AllId() []IIdContext
	Id(i int) IIdContext
	EXTENDS() antlr.TerminalNode
	AllClass_member() []IClass_memberContext
	Class_member(i int) IClass_memberContext

	// IsClass_declContext differentiates from other interfaces.
	IsClass_declContext()
}

type Class_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClass_declContext() *Class_declContext {
	var p = new(Class_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_class_decl
	return p
}

func InitEmptyClass_declContext(p *Class_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_class_decl
}

func (*Class_declContext) IsClass_declContext() {}

func NewClass_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Class_declContext {
	var p = new(Class_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_class_decl

	return p
}

func (s *Class_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Class_declContext) CLASS() antlr.TerminalNode {
	return s.GetToken(MiniparParserCLASS, 0)
}

func (s *Class_declContext) AllId() []IIdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdContext); ok {
			len++
		}
	}

	tst := make([]IIdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdContext); ok {
			tst[i] = t.(IIdContext)
			i++
		}
	}

	return tst
}

func (s *Class_declContext) Id(i int) IIdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *Class_declContext) EXTENDS() antlr.TerminalNode {
	return s.GetToken(MiniparParserEXTENDS, 0)
}

func (s *Class_declContext) AllClass_member() []IClass_memberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IClass_memberContext); ok {
			len++
		}
	}

	tst := make([]IClass_memberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IClass_memberContext); ok {
			tst[i] = t.(IClass_memberContext)
			i++
		}
	}

	return tst
}

func (s *Class_declContext) Class_member(i int) IClass_memberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClass_memberContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClass_memberContext)
}

func (s *Class_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Class_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Class_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterClass_decl(s)
	}
}

func (s *Class_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitClass_decl(s)
	}
}

func (p *MiniparParser) Class_decl() (localctx IClass_declContext) {
	localctx = NewClass_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MiniparParserRULE_class_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(MiniparParserCLASS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Id()
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MiniparParserEXTENDS {
		{
			p.SetState(107)
			p.Match(MiniparParserEXTENDS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Id()
		}

	}
	{
		p.SetState(111)
		p.Match(MiniparParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36029338184843264) != 0) {
		{
			p.SetState(112)
			p.Class_member()
		}

		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
		p.Match(MiniparParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClass_memberContext is an interface to support dynamic dispatch.
type IClass_memberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_decl() IVar_declContext
	Method_decl() IMethod_declContext

	// IsClass_memberContext differentiates from other interfaces.
	IsClass_memberContext()
}

type Class_memberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClass_memberContext() *Class_memberContext {
	var p = new(Class_memberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_class_member
	return p
}

func InitEmptyClass_memberContext(p *Class_memberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_class_member
}

func (*Class_memberContext) IsClass_memberContext() {}

func NewClass_memberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Class_memberContext {
	var p = new(Class_memberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_class_member

	return p
}

func (s *Class_memberContext) GetParser() antlr.Parser { return s.parser }

func (s *Class_memberContext) Var_decl() IVar_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Class_memberContext) Method_decl() IMethod_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethod_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethod_declContext)
}

func (s *Class_memberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Class_memberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Class_memberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterClass_member(s)
	}
}

func (s *Class_memberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitClass_member(s)
	}
}

func (p *MiniparParser) Class_member() (localctx IClass_memberContext) {
	localctx = NewClass_memberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MiniparParserRULE_class_member)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Var_decl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Method_decl()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethod_declContext is an interface to support dynamic dispatch.
type IMethod_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Id() IIdContext
	Block() IBlockContext
	Params() IParamsContext

	// IsMethod_declContext differentiates from other interfaces.
	IsMethod_declContext()
}

type Method_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethod_declContext() *Method_declContext {
	var p = new(Method_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_method_decl
	return p
}

func InitEmptyMethod_declContext(p *Method_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_method_decl
}

func (*Method_declContext) IsMethod_declContext() {}

func NewMethod_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Method_declContext {
	var p = new(Method_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_method_decl

	return p
}

func (s *Method_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Method_declContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *Method_declContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *Method_declContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Method_declContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *Method_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Method_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Method_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterMethod_decl(s)
	}
}

func (s *Method_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitMethod_decl(s)
	}
}

func (p *MiniparParser) Method_decl() (localctx IMethod_declContext) {
	localctx = NewMethod_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MiniparParserRULE_method_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Type_()
	}
	{
		p.SetState(124)
		p.Id()
	}
	{
		p.SetState(125)
		p.Match(MiniparParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36029338184843264) != 0 {
		{
			p.SetState(126)
			p.Params()
		}

	}
	{
		p.SetState(129)
		p.Match(MiniparParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Match(MiniparParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Block()
	}
	{
		p.SetState(132)
		p.Match(MiniparParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethod_callContext is an interface to support dynamic dispatch.
type IMethod_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllId() []IIdContext
	Id(i int) IIdContext
	Args() IArgsContext

	// IsMethod_callContext differentiates from other interfaces.
	IsMethod_callContext()
}

type Method_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethod_callContext() *Method_callContext {
	var p = new(Method_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_method_call
	return p
}

func InitEmptyMethod_callContext(p *Method_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_method_call
}

func (*Method_callContext) IsMethod_callContext() {}

func NewMethod_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Method_callContext {
	var p = new(Method_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_method_call

	return p
}

func (s *Method_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Method_callContext) AllId() []IIdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdContext); ok {
			len++
		}
	}

	tst := make([]IIdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdContext); ok {
			tst[i] = t.(IIdContext)
			i++
		}
	}

	return tst
}

func (s *Method_callContext) Id(i int) IIdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *Method_callContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *Method_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Method_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Method_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterMethod_call(s)
	}
}

func (s *Method_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitMethod_call(s)
	}
}

func (p *MiniparParser) Method_call() (localctx IMethod_callContext) {
	localctx = NewMethod_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MiniparParserRULE_method_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Id()
	}
	{
		p.SetState(135)
		p.Match(MiniparParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.Id()
	}
	{
		p.SetState(137)
		p.Match(MiniparParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&254463824306913802) != 0 {
		{
			p.SetState(138)
			p.Args()
		}

	}
	{
		p.SetState(141)
		p.Match(MiniparParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObj_creationContext is an interface to support dynamic dispatch.
type IObj_creationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NEW() antlr.TerminalNode
	Id() IIdContext
	Args() IArgsContext

	// IsObj_creationContext differentiates from other interfaces.
	IsObj_creationContext()
}

type Obj_creationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObj_creationContext() *Obj_creationContext {
	var p = new(Obj_creationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_obj_creation
	return p
}

func InitEmptyObj_creationContext(p *Obj_creationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_obj_creation
}

func (*Obj_creationContext) IsObj_creationContext() {}

func NewObj_creationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Obj_creationContext {
	var p = new(Obj_creationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_obj_creation

	return p
}

func (s *Obj_creationContext) GetParser() antlr.Parser { return s.parser }

func (s *Obj_creationContext) NEW() antlr.TerminalNode {
	return s.GetToken(MiniparParserNEW, 0)
}

func (s *Obj_creationContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *Obj_creationContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *Obj_creationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Obj_creationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Obj_creationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterObj_creation(s)
	}
}

func (s *Obj_creationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitObj_creation(s)
	}
}

func (p *MiniparParser) Obj_creation() (localctx IObj_creationContext) {
	localctx = NewObj_creationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MiniparParserRULE_obj_creation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(MiniparParserNEW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Id()
	}
	{
		p.SetState(145)
		p.Match(MiniparParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&254463824306913802) != 0 {
		{
			p.SetState(146)
			p.Args()
		}

	}
	{
		p.SetState(149)
		p.Match(MiniparParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_declContext is an interface to support dynamic dispatch.
type IFunc_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	Id() IIdContext
	Type_() ITypeContext
	Block() IBlockContext
	Params() IParamsContext

	// IsFunc_declContext differentiates from other interfaces.
	IsFunc_declContext()
}

type Func_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_declContext() *Func_declContext {
	var p = new(Func_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_func_decl
	return p
}

func InitEmptyFunc_declContext(p *Func_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_func_decl
}

func (*Func_declContext) IsFunc_declContext() {}

func NewFunc_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_declContext {
	var p = new(Func_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_func_decl

	return p
}

func (s *Func_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_declContext) FUNC() antlr.TerminalNode {
	return s.GetToken(MiniparParserFUNC, 0)
}

func (s *Func_declContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *Func_declContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *Func_declContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Func_declContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *Func_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterFunc_decl(s)
	}
}

func (s *Func_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitFunc_decl(s)
	}
}

func (p *MiniparParser) Func_decl() (localctx IFunc_declContext) {
	localctx = NewFunc_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MiniparParserRULE_func_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(MiniparParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Id()
	}
	{
		p.SetState(153)
		p.Match(MiniparParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36029338184843264) != 0 {
		{
			p.SetState(154)
			p.Params()
		}

	}
	{
		p.SetState(157)
		p.Match(MiniparParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Type_()
	}
	{
		p.SetState(159)
		p.Match(MiniparParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.Block()
	}
	{
		p.SetState(161)
		p.Match(MiniparParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_params
	return p
}

func InitEmptyParamsContext(p *ParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_params
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamsContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitParams(s)
	}
}

func (p *MiniparParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MiniparParserRULE_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Param()
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniparParserT__5 {
		{
			p.SetState(164)
			p.Match(MiniparParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Param()
		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Id() IIdContext

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParamContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *MiniparParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MiniparParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Type_()
	}
	{
		p.SetState(172)
		p.Id()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_callContext is an interface to support dynamic dispatch.
type IFunc_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id() IIdContext
	Args() IArgsContext

	// IsFunc_callContext differentiates from other interfaces.
	IsFunc_callContext()
}

type Func_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_callContext() *Func_callContext {
	var p = new(Func_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_func_call
	return p
}

func InitEmptyFunc_callContext(p *Func_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_func_call
}

func (*Func_callContext) IsFunc_callContext() {}

func NewFunc_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_callContext {
	var p = new(Func_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_func_call

	return p
}

func (s *Func_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_callContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *Func_callContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *Func_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterFunc_call(s)
	}
}

func (s *Func_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitFunc_call(s)
	}
}

func (p *MiniparParser) Func_call() (localctx IFunc_callContext) {
	localctx = NewFunc_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MiniparParserRULE_func_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Id()
	}
	{
		p.SetState(175)
		p.Match(MiniparParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&254463824306913802) != 0 {
		{
			p.SetState(176)
			p.Args()
		}

	}
	{
		p.SetState(179)
		p.Match(MiniparParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_args
	return p
}

func InitEmptyArgsContext(p *ArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_args
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArgsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (p *MiniparParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MiniparParserRULE_args)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Expr()
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniparParserT__5 {
		{
			p.SetState(182)
			p.Match(MiniparParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Expr()
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_declContext is an interface to support dynamic dispatch.
type IVar_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Id() IIdContext
	Expr() IExprContext

	// IsVar_declContext differentiates from other interfaces.
	IsVar_declContext()
}

type Var_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_declContext() *Var_declContext {
	var p = new(Var_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_var_decl
	return p
}

func InitEmptyVar_declContext(p *Var_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_var_decl
}

func (*Var_declContext) IsVar_declContext() {}

func NewVar_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declContext {
	var p = new(Var_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_var_decl

	return p
}

func (s *Var_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *Var_declContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *Var_declContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Var_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterVar_decl(s)
	}
}

func (s *Var_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitVar_decl(s)
	}
}

func (p *MiniparParser) Var_decl() (localctx IVar_declContext) {
	localctx = NewVar_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MiniparParserRULE_var_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Type_()
	}
	{
		p.SetState(190)
		p.Id()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MiniparParserT__6 {
		{
			p.SetState(191)
			p.Match(MiniparParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.Expr()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id() IIdContext
	Expr() IExprContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *AssignmentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *MiniparParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MiniparParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Id()
	}
	{
		p.SetState(196)
		p.Match(MiniparParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Expr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *MiniparParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MiniparParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&254464367619162634) != 0 {
		{
			p.SetState(199)
			p.Stmt()
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Compound_stmt() ICompound_stmtContext
	Simple_stmt() ISimple_stmtContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Compound_stmt() ICompound_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompound_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompound_stmtContext)
}

func (s *StmtContext) Simple_stmt() ISimple_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_stmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *MiniparParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MiniparParserRULE_stmt)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Compound_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Simple_stmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompound_stmtContext is an interface to support dynamic dispatch.
type ICompound_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If_stmt() IIf_stmtContext
	While_stmt() IWhile_stmtContext
	Do_stmt() IDo_stmtContext
	For_stmt() IFor_stmtContext
	Seq_stmt() ISeq_stmtContext
	Par_stmt() IPar_stmtContext
	Channel_stmt() IChannel_stmtContext
	Print_call() IPrint_callContext
	Input_call() IInput_callContext
	Method_call() IMethod_callContext
	Expr() IExprContext

	// IsCompound_stmtContext differentiates from other interfaces.
	IsCompound_stmtContext()
}

type Compound_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompound_stmtContext() *Compound_stmtContext {
	var p = new(Compound_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_compound_stmt
	return p
}

func InitEmptyCompound_stmtContext(p *Compound_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_compound_stmt
}

func (*Compound_stmtContext) IsCompound_stmtContext() {}

func NewCompound_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Compound_stmtContext {
	var p = new(Compound_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_compound_stmt

	return p
}

func (s *Compound_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Compound_stmtContext) If_stmt() IIf_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *Compound_stmtContext) While_stmt() IWhile_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_stmtContext)
}

func (s *Compound_stmtContext) Do_stmt() IDo_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDo_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDo_stmtContext)
}

func (s *Compound_stmtContext) For_stmt() IFor_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_stmtContext)
}

func (s *Compound_stmtContext) Seq_stmt() ISeq_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISeq_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISeq_stmtContext)
}

func (s *Compound_stmtContext) Par_stmt() IPar_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPar_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPar_stmtContext)
}

func (s *Compound_stmtContext) Channel_stmt() IChannel_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChannel_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChannel_stmtContext)
}

func (s *Compound_stmtContext) Print_call() IPrint_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrint_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrint_callContext)
}

func (s *Compound_stmtContext) Input_call() IInput_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInput_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInput_callContext)
}

func (s *Compound_stmtContext) Method_call() IMethod_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethod_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethod_callContext)
}

func (s *Compound_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Compound_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Compound_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Compound_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterCompound_stmt(s)
	}
}

func (s *Compound_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitCompound_stmt(s)
	}
}

func (p *MiniparParser) Compound_stmt() (localctx ICompound_stmtContext) {
	localctx = NewCompound_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MiniparParserRULE_compound_stmt)
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.If_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(210)
			p.While_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(211)
			p.Do_stmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(212)
			p.For_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(213)
			p.Seq_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(214)
			p.Par_stmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(215)
			p.Channel_stmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(216)
			p.Print_call()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(217)
			p.Input_call()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(218)
			p.Method_call()
		}
		{
			p.SetState(219)
			p.Match(MiniparParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(221)
			p.Expr()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext

	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_if_stmt
	return p
}

func InitEmptyIf_stmtContext(p *If_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_if_stmt
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) IF() antlr.TerminalNode {
	return s.GetToken(MiniparParserIF, 0)
}

func (s *If_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *If_stmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterIf_stmt(s)
	}
}

func (s *If_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitIf_stmt(s)
	}
}

func (p *MiniparParser) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MiniparParserRULE_if_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(MiniparParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Match(MiniparParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.Expr()
	}
	{
		p.SetState(227)
		p.Match(MiniparParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.Match(MiniparParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Block()
	}
	{
		p.SetState(230)
		p.Match(MiniparParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhile_stmtContext is an interface to support dynamic dispatch.
type IWhile_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext

	// IsWhile_stmtContext differentiates from other interfaces.
	IsWhile_stmtContext()
}

type While_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_stmtContext() *While_stmtContext {
	var p = new(While_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_while_stmt
	return p
}

func InitEmptyWhile_stmtContext(p *While_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_while_stmt
}

func (*While_stmtContext) IsWhile_stmtContext() {}

func NewWhile_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmtContext {
	var p = new(While_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_while_stmt

	return p
}

func (s *While_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(MiniparParserWHILE, 0)
}

func (s *While_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *While_stmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *While_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterWhile_stmt(s)
	}
}

func (s *While_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitWhile_stmt(s)
	}
}

func (p *MiniparParser) While_stmt() (localctx IWhile_stmtContext) {
	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MiniparParserRULE_while_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(MiniparParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(233)
		p.Match(MiniparParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.Expr()
	}
	{
		p.SetState(235)
		p.Match(MiniparParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.Match(MiniparParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Block()
	}
	{
		p.SetState(238)
		p.Match(MiniparParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDo_stmtContext is an interface to support dynamic dispatch.
type IDo_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DO() antlr.TerminalNode
	Block() IBlockContext
	WHILE() antlr.TerminalNode
	Expr() IExprContext

	// IsDo_stmtContext differentiates from other interfaces.
	IsDo_stmtContext()
}

type Do_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDo_stmtContext() *Do_stmtContext {
	var p = new(Do_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_do_stmt
	return p
}

func InitEmptyDo_stmtContext(p *Do_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_do_stmt
}

func (*Do_stmtContext) IsDo_stmtContext() {}

func NewDo_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Do_stmtContext {
	var p = new(Do_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_do_stmt

	return p
}

func (s *Do_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Do_stmtContext) DO() antlr.TerminalNode {
	return s.GetToken(MiniparParserDO, 0)
}

func (s *Do_stmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Do_stmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(MiniparParserWHILE, 0)
}

func (s *Do_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Do_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Do_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Do_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterDo_stmt(s)
	}
}

func (s *Do_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitDo_stmt(s)
	}
}

func (p *MiniparParser) Do_stmt() (localctx IDo_stmtContext) {
	localctx = NewDo_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MiniparParserRULE_do_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(MiniparParserDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Match(MiniparParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.Block()
	}
	{
		p.SetState(243)
		p.Match(MiniparParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(MiniparParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Match(MiniparParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.Expr()
	}
	{
		p.SetState(247)
		p.Match(MiniparParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_stmtContext is an interface to support dynamic dispatch.
type IFor_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	Type_() ITypeContext
	Id() IIdContext
	IN() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext

	// IsFor_stmtContext differentiates from other interfaces.
	IsFor_stmtContext()
}

type For_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmtContext() *For_stmtContext {
	var p = new(For_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_for_stmt
	return p
}

func InitEmptyFor_stmtContext(p *For_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_for_stmt
}

func (*For_stmtContext) IsFor_stmtContext() {}

func NewFor_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmtContext {
	var p = new(For_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_for_stmt

	return p
}

func (s *For_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(MiniparParserFOR, 0)
}

func (s *For_stmtContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *For_stmtContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *For_stmtContext) IN() antlr.TerminalNode {
	return s.GetToken(MiniparParserIN, 0)
}

func (s *For_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *For_stmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *For_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterFor_stmt(s)
	}
}

func (s *For_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitFor_stmt(s)
	}
}

func (p *MiniparParser) For_stmt() (localctx IFor_stmtContext) {
	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MiniparParserRULE_for_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(MiniparParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.Match(MiniparParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(251)
		p.Type_()
	}
	{
		p.SetState(252)
		p.Id()
	}
	{
		p.SetState(253)
		p.Match(MiniparParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
		p.Expr()
	}
	{
		p.SetState(255)
		p.Match(MiniparParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(256)
		p.Match(MiniparParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.Block()
	}
	{
		p.SetState(258)
		p.Match(MiniparParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISeq_stmtContext is an interface to support dynamic dispatch.
type ISeq_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEQ() antlr.TerminalNode
	Block() IBlockContext

	// IsSeq_stmtContext differentiates from other interfaces.
	IsSeq_stmtContext()
}

type Seq_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeq_stmtContext() *Seq_stmtContext {
	var p = new(Seq_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_seq_stmt
	return p
}

func InitEmptySeq_stmtContext(p *Seq_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_seq_stmt
}

func (*Seq_stmtContext) IsSeq_stmtContext() {}

func NewSeq_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Seq_stmtContext {
	var p = new(Seq_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_seq_stmt

	return p
}

func (s *Seq_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Seq_stmtContext) SEQ() antlr.TerminalNode {
	return s.GetToken(MiniparParserSEQ, 0)
}

func (s *Seq_stmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Seq_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Seq_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Seq_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterSeq_stmt(s)
	}
}

func (s *Seq_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitSeq_stmt(s)
	}
}

func (p *MiniparParser) Seq_stmt() (localctx ISeq_stmtContext) {
	localctx = NewSeq_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MiniparParserRULE_seq_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.Match(MiniparParserSEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Match(MiniparParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.Block()
	}
	{
		p.SetState(263)
		p.Match(MiniparParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPar_stmtContext is an interface to support dynamic dispatch.
type IPar_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PAR() antlr.TerminalNode
	Block() IBlockContext

	// IsPar_stmtContext differentiates from other interfaces.
	IsPar_stmtContext()
}

type Par_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPar_stmtContext() *Par_stmtContext {
	var p = new(Par_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_par_stmt
	return p
}

func InitEmptyPar_stmtContext(p *Par_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_par_stmt
}

func (*Par_stmtContext) IsPar_stmtContext() {}

func NewPar_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Par_stmtContext {
	var p = new(Par_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_par_stmt

	return p
}

func (s *Par_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Par_stmtContext) PAR() antlr.TerminalNode {
	return s.GetToken(MiniparParserPAR, 0)
}

func (s *Par_stmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Par_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Par_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Par_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterPar_stmt(s)
	}
}

func (s *Par_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitPar_stmt(s)
	}
}

func (p *MiniparParser) Par_stmt() (localctx IPar_stmtContext) {
	localctx = NewPar_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MiniparParserRULE_par_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Match(MiniparParserPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(266)
		p.Match(MiniparParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Block()
	}
	{
		p.SetState(268)
		p.Match(MiniparParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrint_callContext is an interface to support dynamic dispatch.
type IPrint_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	Args() IArgsContext

	// IsPrint_callContext differentiates from other interfaces.
	IsPrint_callContext()
}

type Print_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrint_callContext() *Print_callContext {
	var p = new(Print_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_print_call
	return p
}

func InitEmptyPrint_callContext(p *Print_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_print_call
}

func (*Print_callContext) IsPrint_callContext() {}

func NewPrint_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Print_callContext {
	var p = new(Print_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_print_call

	return p
}

func (s *Print_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Print_callContext) PRINT() antlr.TerminalNode {
	return s.GetToken(MiniparParserPRINT, 0)
}

func (s *Print_callContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *Print_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Print_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Print_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterPrint_call(s)
	}
}

func (s *Print_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitPrint_call(s)
	}
}

func (p *MiniparParser) Print_call() (localctx IPrint_callContext) {
	localctx = NewPrint_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MiniparParserRULE_print_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Match(MiniparParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.Match(MiniparParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&254463824306913802) != 0 {
		{
			p.SetState(272)
			p.Args()
		}

	}
	{
		p.SetState(275)
		p.Match(MiniparParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(276)
		p.Match(MiniparParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInput_callContext is an interface to support dynamic dispatch.
type IInput_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INPUT() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsInput_callContext differentiates from other interfaces.
	IsInput_callContext()
}

type Input_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInput_callContext() *Input_callContext {
	var p = new(Input_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_input_call
	return p
}

func InitEmptyInput_callContext(p *Input_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_input_call
}

func (*Input_callContext) IsInput_callContext() {}

func NewInput_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Input_callContext {
	var p = new(Input_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_input_call

	return p
}

func (s *Input_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Input_callContext) INPUT() antlr.TerminalNode {
	return s.GetToken(MiniparParserINPUT, 0)
}

func (s *Input_callContext) STRING() antlr.TerminalNode {
	return s.GetToken(MiniparParserSTRING, 0)
}

func (s *Input_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Input_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Input_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterInput_call(s)
	}
}

func (s *Input_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitInput_call(s)
	}
}

func (p *MiniparParser) Input_call() (localctx IInput_callContext) {
	localctx = NewInput_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MiniparParserRULE_input_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Match(MiniparParserINPUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.Match(MiniparParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MiniparParserSTRING {
		{
			p.SetState(280)
			p.Match(MiniparParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(283)
		p.Match(MiniparParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)
		p.Match(MiniparParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimple_stmtContext is an interface to support dynamic dispatch.
type ISimple_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_decl() IVar_declContext
	Assignment() IAssignmentContext
	Expr_stmt() IExpr_stmtContext
	RETURN() antlr.TerminalNode
	Expr() IExprContext
	BREAK() antlr.TerminalNode
	CONTINUE() antlr.TerminalNode
	PASS() antlr.TerminalNode

	// IsSimple_stmtContext differentiates from other interfaces.
	IsSimple_stmtContext()
}

type Simple_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_stmtContext() *Simple_stmtContext {
	var p = new(Simple_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_simple_stmt
	return p
}

func InitEmptySimple_stmtContext(p *Simple_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_simple_stmt
}

func (*Simple_stmtContext) IsSimple_stmtContext() {}

func NewSimple_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_stmtContext {
	var p = new(Simple_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_simple_stmt

	return p
}

func (s *Simple_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_stmtContext) Var_decl() IVar_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Simple_stmtContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *Simple_stmtContext) Expr_stmt() IExpr_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_stmtContext)
}

func (s *Simple_stmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(MiniparParserRETURN, 0)
}

func (s *Simple_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Simple_stmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(MiniparParserBREAK, 0)
}

func (s *Simple_stmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(MiniparParserCONTINUE, 0)
}

func (s *Simple_stmtContext) PASS() antlr.TerminalNode {
	return s.GetToken(MiniparParserPASS, 0)
}

func (s *Simple_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterSimple_stmt(s)
	}
}

func (s *Simple_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitSimple_stmt(s)
	}
}

func (p *MiniparParser) Simple_stmt() (localctx ISimple_stmtContext) {
	localctx = NewSimple_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MiniparParserRULE_simple_stmt)
	var _la int

	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(286)
			p.Var_decl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(287)
			p.Assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(288)
			p.Expr_stmt()
		}
		{
			p.SetState(289)
			p.Match(MiniparParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(291)
			p.Match(MiniparParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&254463824306913802) != 0 {
			{
				p.SetState(292)
				p.Expr()
			}

		}
		{
			p.SetState(295)
			p.Match(MiniparParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(296)
			p.Match(MiniparParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)
			p.Match(MiniparParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(298)
			p.Match(MiniparParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.Match(MiniparParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(300)
			p.Match(MiniparParserPASS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)
			p.Match(MiniparParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpr_stmtContext is an interface to support dynamic dispatch.
type IExpr_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	Func_call() IFunc_callContext

	// IsExpr_stmtContext differentiates from other interfaces.
	IsExpr_stmtContext()
}

type Expr_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_stmtContext() *Expr_stmtContext {
	var p = new(Expr_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_expr_stmt
	return p
}

func InitEmptyExpr_stmtContext(p *Expr_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_expr_stmt
}

func (*Expr_stmtContext) IsExpr_stmtContext() {}

func NewExpr_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_stmtContext {
	var p = new(Expr_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_expr_stmt

	return p
}

func (s *Expr_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_stmtContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *Expr_stmtContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *Expr_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterExpr_stmt(s)
	}
}

func (s *Expr_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitExpr_stmt(s)
	}
}

func (p *MiniparParser) Expr_stmt() (localctx IExpr_stmtContext) {
	localctx = NewExpr_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MiniparParserRULE_expr_stmt)
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(304)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(305)
			p.Func_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IChannel_stmtContext is an interface to support dynamic dispatch.
type IChannel_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	S_channel_stmt() IS_channel_stmtContext
	C_channel_stmt() IC_channel_stmtContext

	// IsChannel_stmtContext differentiates from other interfaces.
	IsChannel_stmtContext()
}

type Channel_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChannel_stmtContext() *Channel_stmtContext {
	var p = new(Channel_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_channel_stmt
	return p
}

func InitEmptyChannel_stmtContext(p *Channel_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_channel_stmt
}

func (*Channel_stmtContext) IsChannel_stmtContext() {}

func NewChannel_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Channel_stmtContext {
	var p = new(Channel_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_channel_stmt

	return p
}

func (s *Channel_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Channel_stmtContext) S_channel_stmt() IS_channel_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IS_channel_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IS_channel_stmtContext)
}

func (s *Channel_stmtContext) C_channel_stmt() IC_channel_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IC_channel_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IC_channel_stmtContext)
}

func (s *Channel_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Channel_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Channel_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterChannel_stmt(s)
	}
}

func (s *Channel_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitChannel_stmt(s)
	}
}

func (p *MiniparParser) Channel_stmt() (localctx IChannel_stmtContext) {
	localctx = NewChannel_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MiniparParserRULE_channel_stmt)
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniparParserS_CHANNEL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(308)
			p.S_channel_stmt()
		}

	case MiniparParserC_CHANNEL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(309)
			p.C_channel_stmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IS_channel_stmtContext is an interface to support dynamic dispatch.
type IS_channel_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	S_CHANNEL() antlr.TerminalNode
	Id() IIdContext
	Args() IArgsContext

	// IsS_channel_stmtContext differentiates from other interfaces.
	IsS_channel_stmtContext()
}

type S_channel_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyS_channel_stmtContext() *S_channel_stmtContext {
	var p = new(S_channel_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_s_channel_stmt
	return p
}

func InitEmptyS_channel_stmtContext(p *S_channel_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_s_channel_stmt
}

func (*S_channel_stmtContext) IsS_channel_stmtContext() {}

func NewS_channel_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *S_channel_stmtContext {
	var p = new(S_channel_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_s_channel_stmt

	return p
}

func (s *S_channel_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *S_channel_stmtContext) S_CHANNEL() antlr.TerminalNode {
	return s.GetToken(MiniparParserS_CHANNEL, 0)
}

func (s *S_channel_stmtContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *S_channel_stmtContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *S_channel_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_channel_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *S_channel_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterS_channel_stmt(s)
	}
}

func (s *S_channel_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitS_channel_stmt(s)
	}
}

func (p *MiniparParser) S_channel_stmt() (localctx IS_channel_stmtContext) {
	localctx = NewS_channel_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, MiniparParserRULE_s_channel_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(MiniparParserS_CHANNEL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
		p.Id()
	}
	{
		p.SetState(314)
		p.Match(MiniparParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&254463824306913802) != 0 {
		{
			p.SetState(315)
			p.Args()
		}

	}
	{
		p.SetState(318)
		p.Match(MiniparParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(319)
		p.Match(MiniparParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IC_channel_stmtContext is an interface to support dynamic dispatch.
type IC_channel_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	C_CHANNEL() antlr.TerminalNode
	AllId() []IIdContext
	Id(i int) IIdContext

	// IsC_channel_stmtContext differentiates from other interfaces.
	IsC_channel_stmtContext()
}

type C_channel_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyC_channel_stmtContext() *C_channel_stmtContext {
	var p = new(C_channel_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_c_channel_stmt
	return p
}

func InitEmptyC_channel_stmtContext(p *C_channel_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_c_channel_stmt
}

func (*C_channel_stmtContext) IsC_channel_stmtContext() {}

func NewC_channel_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *C_channel_stmtContext {
	var p = new(C_channel_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_c_channel_stmt

	return p
}

func (s *C_channel_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *C_channel_stmtContext) C_CHANNEL() antlr.TerminalNode {
	return s.GetToken(MiniparParserC_CHANNEL, 0)
}

func (s *C_channel_stmtContext) AllId() []IIdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdContext); ok {
			len++
		}
	}

	tst := make([]IIdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdContext); ok {
			tst[i] = t.(IIdContext)
			i++
		}
	}

	return tst
}

func (s *C_channel_stmtContext) Id(i int) IIdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *C_channel_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *C_channel_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *C_channel_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterC_channel_stmt(s)
	}
}

func (s *C_channel_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitC_channel_stmt(s)
	}
}

func (p *MiniparParser) C_channel_stmt() (localctx IC_channel_stmtContext) {
	localctx = NewC_channel_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, MiniparParserRULE_c_channel_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Match(MiniparParserC_CHANNEL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Id()
	}
	{
		p.SetState(323)
		p.Id()
	}
	{
		p.SetState(324)
		p.Id()
	}
	{
		p.SetState(325)
		p.Match(MiniparParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISend_stmtContext is an interface to support dynamic dispatch.
type ISend_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id() IIdContext
	SEND() antlr.TerminalNode
	Args() IArgsContext

	// IsSend_stmtContext differentiates from other interfaces.
	IsSend_stmtContext()
}

type Send_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySend_stmtContext() *Send_stmtContext {
	var p = new(Send_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_send_stmt
	return p
}

func InitEmptySend_stmtContext(p *Send_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_send_stmt
}

func (*Send_stmtContext) IsSend_stmtContext() {}

func NewSend_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Send_stmtContext {
	var p = new(Send_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_send_stmt

	return p
}

func (s *Send_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Send_stmtContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *Send_stmtContext) SEND() antlr.TerminalNode {
	return s.GetToken(MiniparParserSEND, 0)
}

func (s *Send_stmtContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *Send_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Send_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Send_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterSend_stmt(s)
	}
}

func (s *Send_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitSend_stmt(s)
	}
}

func (p *MiniparParser) Send_stmt() (localctx ISend_stmtContext) {
	localctx = NewSend_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, MiniparParserRULE_send_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Id()
	}
	{
		p.SetState(328)
		p.Match(MiniparParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(329)
		p.Match(MiniparParserSEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(330)
		p.Match(MiniparParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&254463824306913802) != 0 {
		{
			p.SetState(331)
			p.Args()
		}

	}
	{
		p.SetState(334)
		p.Match(MiniparParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(335)
		p.Match(MiniparParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReceive_stmtContext is an interface to support dynamic dispatch.
type IReceive_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id() IIdContext
	RECEIVE() antlr.TerminalNode
	Args() IArgsContext

	// IsReceive_stmtContext differentiates from other interfaces.
	IsReceive_stmtContext()
}

type Receive_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReceive_stmtContext() *Receive_stmtContext {
	var p = new(Receive_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_receive_stmt
	return p
}

func InitEmptyReceive_stmtContext(p *Receive_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_receive_stmt
}

func (*Receive_stmtContext) IsReceive_stmtContext() {}

func NewReceive_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Receive_stmtContext {
	var p = new(Receive_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_receive_stmt

	return p
}

func (s *Receive_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Receive_stmtContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *Receive_stmtContext) RECEIVE() antlr.TerminalNode {
	return s.GetToken(MiniparParserRECEIVE, 0)
}

func (s *Receive_stmtContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *Receive_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Receive_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Receive_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterReceive_stmt(s)
	}
}

func (s *Receive_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitReceive_stmt(s)
	}
}

func (p *MiniparParser) Receive_stmt() (localctx IReceive_stmtContext) {
	localctx = NewReceive_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, MiniparParserRULE_receive_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Id()
	}
	{
		p.SetState(338)
		p.Match(MiniparParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(339)
		p.Match(MiniparParserRECEIVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)
		p.Match(MiniparParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&254463824306913802) != 0 {
		{
			p.SetState(341)
			p.Args()
		}

	}
	{
		p.SetState(344)
		p.Match(MiniparParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)
		p.Match(MiniparParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Or_expr() IOr_exprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Or_expr() IOr_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOr_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOr_exprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *MiniparParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, MiniparParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Or_expr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOr_exprContext is an interface to support dynamic dispatch.
type IOr_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAnd_expr() []IAnd_exprContext
	And_expr(i int) IAnd_exprContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsOr_exprContext differentiates from other interfaces.
	IsOr_exprContext()
}

type Or_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOr_exprContext() *Or_exprContext {
	var p = new(Or_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_or_expr
	return p
}

func InitEmptyOr_exprContext(p *Or_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_or_expr
}

func (*Or_exprContext) IsOr_exprContext() {}

func NewOr_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Or_exprContext {
	var p = new(Or_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_or_expr

	return p
}

func (s *Or_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Or_exprContext) AllAnd_expr() []IAnd_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnd_exprContext); ok {
			len++
		}
	}

	tst := make([]IAnd_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnd_exprContext); ok {
			tst[i] = t.(IAnd_exprContext)
			i++
		}
	}

	return tst
}

func (s *Or_exprContext) And_expr(i int) IAnd_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnd_exprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnd_exprContext)
}

func (s *Or_exprContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(MiniparParserOR)
}

func (s *Or_exprContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(MiniparParserOR, i)
}

func (s *Or_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Or_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterOr_expr(s)
	}
}

func (s *Or_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitOr_expr(s)
	}
}

func (p *MiniparParser) Or_expr() (localctx IOr_exprContext) {
	localctx = NewOr_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, MiniparParserRULE_or_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.And_expr()
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniparParserOR {
		{
			p.SetState(350)
			p.Match(MiniparParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.And_expr()
		}

		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnd_exprContext is an interface to support dynamic dispatch.
type IAnd_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllComparison_expr() []IComparison_exprContext
	Comparison_expr(i int) IComparison_exprContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsAnd_exprContext differentiates from other interfaces.
	IsAnd_exprContext()
}

type And_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnd_exprContext() *And_exprContext {
	var p = new(And_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_and_expr
	return p
}

func InitEmptyAnd_exprContext(p *And_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_and_expr
}

func (*And_exprContext) IsAnd_exprContext() {}

func NewAnd_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *And_exprContext {
	var p = new(And_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_and_expr

	return p
}

func (s *And_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *And_exprContext) AllComparison_expr() []IComparison_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparison_exprContext); ok {
			len++
		}
	}

	tst := make([]IComparison_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparison_exprContext); ok {
			tst[i] = t.(IComparison_exprContext)
			i++
		}
	}

	return tst
}

func (s *And_exprContext) Comparison_expr(i int) IComparison_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparison_exprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparison_exprContext)
}

func (s *And_exprContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(MiniparParserAND)
}

func (s *And_exprContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(MiniparParserAND, i)
}

func (s *And_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *And_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *And_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterAnd_expr(s)
	}
}

func (s *And_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitAnd_expr(s)
	}
}

func (p *MiniparParser) And_expr() (localctx IAnd_exprContext) {
	localctx = NewAnd_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, MiniparParserRULE_and_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Comparison_expr()
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MiniparParserAND {
		{
			p.SetState(358)
			p.Match(MiniparParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(359)
			p.Comparison_expr()
		}

		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparison_exprContext is an interface to support dynamic dispatch.
type IComparison_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAdditive_expr() []IAdditive_exprContext
	Additive_expr(i int) IAdditive_exprContext
	AllEQ() []antlr.TerminalNode
	EQ(i int) antlr.TerminalNode
	AllNEQ() []antlr.TerminalNode
	NEQ(i int) antlr.TerminalNode
	AllLT() []antlr.TerminalNode
	LT(i int) antlr.TerminalNode
	AllGT() []antlr.TerminalNode
	GT(i int) antlr.TerminalNode
	AllLTE() []antlr.TerminalNode
	LTE(i int) antlr.TerminalNode
	AllGTE() []antlr.TerminalNode
	GTE(i int) antlr.TerminalNode

	// IsComparison_exprContext differentiates from other interfaces.
	IsComparison_exprContext()
}

type Comparison_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparison_exprContext() *Comparison_exprContext {
	var p = new(Comparison_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_comparison_expr
	return p
}

func InitEmptyComparison_exprContext(p *Comparison_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_comparison_expr
}

func (*Comparison_exprContext) IsComparison_exprContext() {}

func NewComparison_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comparison_exprContext {
	var p = new(Comparison_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_comparison_expr

	return p
}

func (s *Comparison_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Comparison_exprContext) AllAdditive_expr() []IAdditive_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditive_exprContext); ok {
			len++
		}
	}

	tst := make([]IAdditive_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditive_exprContext); ok {
			tst[i] = t.(IAdditive_exprContext)
			i++
		}
	}

	return tst
}

func (s *Comparison_exprContext) Additive_expr(i int) IAdditive_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditive_exprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditive_exprContext)
}

func (s *Comparison_exprContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(MiniparParserEQ)
}

func (s *Comparison_exprContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(MiniparParserEQ, i)
}

func (s *Comparison_exprContext) AllNEQ() []antlr.TerminalNode {
	return s.GetTokens(MiniparParserNEQ)
}

func (s *Comparison_exprContext) NEQ(i int) antlr.TerminalNode {
	return s.GetToken(MiniparParserNEQ, i)
}

func (s *Comparison_exprContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(MiniparParserLT)
}

func (s *Comparison_exprContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(MiniparParserLT, i)
}

func (s *Comparison_exprContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(MiniparParserGT)
}

func (s *Comparison_exprContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(MiniparParserGT, i)
}

func (s *Comparison_exprContext) AllLTE() []antlr.TerminalNode {
	return s.GetTokens(MiniparParserLTE)
}

func (s *Comparison_exprContext) LTE(i int) antlr.TerminalNode {
	return s.GetToken(MiniparParserLTE, i)
}

func (s *Comparison_exprContext) AllGTE() []antlr.TerminalNode {
	return s.GetTokens(MiniparParserGTE)
}

func (s *Comparison_exprContext) GTE(i int) antlr.TerminalNode {
	return s.GetToken(MiniparParserGTE, i)
}

func (s *Comparison_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparison_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comparison_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterComparison_expr(s)
	}
}

func (s *Comparison_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitComparison_expr(s)
	}
}

func (p *MiniparParser) Comparison_expr() (localctx IComparison_exprContext) {
	localctx = NewComparison_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, MiniparParserRULE_comparison_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Additive_expr()
	}
	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1108307720798208) != 0 {
		{
			p.SetState(366)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1108307720798208) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(367)
			p.Additive_expr()
		}

		p.SetState(372)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdditive_exprContext is an interface to support dynamic dispatch.
type IAdditive_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMultiplicative_expr() []IMultiplicative_exprContext
	Multiplicative_expr(i int) IMultiplicative_exprContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsAdditive_exprContext differentiates from other interfaces.
	IsAdditive_exprContext()
}

type Additive_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditive_exprContext() *Additive_exprContext {
	var p = new(Additive_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_additive_expr
	return p
}

func InitEmptyAdditive_exprContext(p *Additive_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_additive_expr
}

func (*Additive_exprContext) IsAdditive_exprContext() {}

func NewAdditive_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Additive_exprContext {
	var p = new(Additive_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_additive_expr

	return p
}

func (s *Additive_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Additive_exprContext) AllMultiplicative_expr() []IMultiplicative_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicative_exprContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicative_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicative_exprContext); ok {
			tst[i] = t.(IMultiplicative_exprContext)
			i++
		}
	}

	return tst
}

func (s *Additive_exprContext) Multiplicative_expr(i int) IMultiplicative_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicative_exprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicative_exprContext)
}

func (s *Additive_exprContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(MiniparParserPLUS)
}

func (s *Additive_exprContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(MiniparParserPLUS, i)
}

func (s *Additive_exprContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(MiniparParserMINUS)
}

func (s *Additive_exprContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(MiniparParserMINUS, i)
}

func (s *Additive_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Additive_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Additive_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterAdditive_expr(s)
	}
}

func (s *Additive_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitAdditive_expr(s)
	}
}

func (p *MiniparParser) Additive_expr() (localctx IAdditive_exprContext) {
	localctx = NewAdditive_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, MiniparParserRULE_additive_expr)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)
		p.Multiplicative_expr()
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(374)
				_la = p.GetTokenStream().LA(1)

				if !(_la == MiniparParserPLUS || _la == MiniparParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(375)
				p.Multiplicative_expr()
			}

		}
		p.SetState(380)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplicative_exprContext is an interface to support dynamic dispatch.
type IMultiplicative_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnary_expr() []IUnary_exprContext
	Unary_expr(i int) IUnary_exprContext
	AllMUL() []antlr.TerminalNode
	MUL(i int) antlr.TerminalNode
	AllDIV() []antlr.TerminalNode
	DIV(i int) antlr.TerminalNode
	AllMOD() []antlr.TerminalNode
	MOD(i int) antlr.TerminalNode

	// IsMultiplicative_exprContext differentiates from other interfaces.
	IsMultiplicative_exprContext()
}

type Multiplicative_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicative_exprContext() *Multiplicative_exprContext {
	var p = new(Multiplicative_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_multiplicative_expr
	return p
}

func InitEmptyMultiplicative_exprContext(p *Multiplicative_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_multiplicative_expr
}

func (*Multiplicative_exprContext) IsMultiplicative_exprContext() {}

func NewMultiplicative_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Multiplicative_exprContext {
	var p = new(Multiplicative_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_multiplicative_expr

	return p
}

func (s *Multiplicative_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Multiplicative_exprContext) AllUnary_expr() []IUnary_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnary_exprContext); ok {
			len++
		}
	}

	tst := make([]IUnary_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnary_exprContext); ok {
			tst[i] = t.(IUnary_exprContext)
			i++
		}
	}

	return tst
}

func (s *Multiplicative_exprContext) Unary_expr(i int) IUnary_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnary_exprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnary_exprContext)
}

func (s *Multiplicative_exprContext) AllMUL() []antlr.TerminalNode {
	return s.GetTokens(MiniparParserMUL)
}

func (s *Multiplicative_exprContext) MUL(i int) antlr.TerminalNode {
	return s.GetToken(MiniparParserMUL, i)
}

func (s *Multiplicative_exprContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(MiniparParserDIV)
}

func (s *Multiplicative_exprContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(MiniparParserDIV, i)
}

func (s *Multiplicative_exprContext) AllMOD() []antlr.TerminalNode {
	return s.GetTokens(MiniparParserMOD)
}

func (s *Multiplicative_exprContext) MOD(i int) antlr.TerminalNode {
	return s.GetToken(MiniparParserMOD, i)
}

func (s *Multiplicative_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Multiplicative_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Multiplicative_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterMultiplicative_expr(s)
	}
}

func (s *Multiplicative_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitMultiplicative_expr(s)
	}
}

func (p *MiniparParser) Multiplicative_expr() (localctx IMultiplicative_exprContext) {
	localctx = NewMultiplicative_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, MiniparParserRULE_multiplicative_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.Unary_expr()
	}
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31525197391593472) != 0 {
		{
			p.SetState(382)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31525197391593472) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(383)
			p.Unary_expr()
		}

		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnary_exprContext is an interface to support dynamic dispatch.
type IUnary_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Unary_expr() IUnary_exprContext
	MINUS() antlr.TerminalNode
	NOT() antlr.TerminalNode
	Postfix_expr() IPostfix_exprContext

	// IsUnary_exprContext differentiates from other interfaces.
	IsUnary_exprContext()
}

type Unary_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnary_exprContext() *Unary_exprContext {
	var p = new(Unary_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_unary_expr
	return p
}

func InitEmptyUnary_exprContext(p *Unary_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_unary_expr
}

func (*Unary_exprContext) IsUnary_exprContext() {}

func NewUnary_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unary_exprContext {
	var p = new(Unary_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_unary_expr

	return p
}

func (s *Unary_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Unary_exprContext) Unary_expr() IUnary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnary_exprContext)
}

func (s *Unary_exprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MiniparParserMINUS, 0)
}

func (s *Unary_exprContext) NOT() antlr.TerminalNode {
	return s.GetToken(MiniparParserNOT, 0)
}

func (s *Unary_exprContext) Postfix_expr() IPostfix_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfix_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfix_exprContext)
}

func (s *Unary_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unary_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterUnary_expr(s)
	}
}

func (s *Unary_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitUnary_expr(s)
	}
}

func (p *MiniparParser) Unary_expr() (localctx IUnary_exprContext) {
	localctx = NewUnary_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, MiniparParserRULE_unary_expr)
	var _la int

	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniparParserNOT, MiniparParserMINUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(389)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MiniparParserNOT || _la == MiniparParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(390)
			p.Unary_expr()
		}

	case MiniparParserT__0, MiniparParserT__2, MiniparParserT__8, MiniparParserNEW, MiniparParserTRUE, MiniparParserFALSE, MiniparParserID, MiniparParserNUMERO, MiniparParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(391)
			p.postfix_expr(0)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPostfix_exprContext is an interface to support dynamic dispatch.
type IPostfix_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Primary_expr() IPrimary_exprContext
	Method_call() IMethod_callContext
	Func_call() IFunc_callContext
	Postfix_expr() IPostfix_exprContext
	Expr() IExprContext

	// IsPostfix_exprContext differentiates from other interfaces.
	IsPostfix_exprContext()
}

type Postfix_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfix_exprContext() *Postfix_exprContext {
	var p = new(Postfix_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_postfix_expr
	return p
}

func InitEmptyPostfix_exprContext(p *Postfix_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_postfix_expr
}

func (*Postfix_exprContext) IsPostfix_exprContext() {}

func NewPostfix_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Postfix_exprContext {
	var p = new(Postfix_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_postfix_expr

	return p
}

func (s *Postfix_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Postfix_exprContext) Primary_expr() IPrimary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *Postfix_exprContext) Method_call() IMethod_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethod_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethod_callContext)
}

func (s *Postfix_exprContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *Postfix_exprContext) Postfix_expr() IPostfix_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfix_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfix_exprContext)
}

func (s *Postfix_exprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Postfix_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Postfix_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Postfix_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterPostfix_expr(s)
	}
}

func (s *Postfix_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitPostfix_expr(s)
	}
}

func (p *MiniparParser) Postfix_expr() (localctx IPostfix_exprContext) {
	return p.postfix_expr(0)
}

func (p *MiniparParser) postfix_expr(_p int) (localctx IPostfix_exprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPostfix_exprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPostfix_exprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 78
	p.EnterRecursionRule(localctx, 78, MiniparParserRULE_postfix_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(395)
			p.Primary_expr()
		}

	case 2:
		{
			p.SetState(396)
			p.Method_call()
		}

	case 3:
		{
			p.SetState(397)
			p.Func_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPostfix_exprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, MiniparParserRULE_postfix_expr)
			p.SetState(400)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				goto errorExit
			}
			{
				p.SetState(401)
				p.Match(MiniparParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(402)
				p.Expr()
			}
			{
				p.SetState(403)
				p.Match(MiniparParserT__9)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimary_exprContext is an interface to support dynamic dispatch.
type IPrimary_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMERO() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	Id() IIdContext
	Obj_creation() IObj_creationContext
	List_literal() IList_literalContext
	Dict_literal() IDict_literalContext
	Expr() IExprContext

	// IsPrimary_exprContext differentiates from other interfaces.
	IsPrimary_exprContext()
}

type Primary_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimary_exprContext() *Primary_exprContext {
	var p = new(Primary_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_primary_expr
	return p
}

func InitEmptyPrimary_exprContext(p *Primary_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_primary_expr
}

func (*Primary_exprContext) IsPrimary_exprContext() {}

func NewPrimary_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primary_exprContext {
	var p = new(Primary_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_primary_expr

	return p
}

func (s *Primary_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Primary_exprContext) NUMERO() antlr.TerminalNode {
	return s.GetToken(MiniparParserNUMERO, 0)
}

func (s *Primary_exprContext) STRING() antlr.TerminalNode {
	return s.GetToken(MiniparParserSTRING, 0)
}

func (s *Primary_exprContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MiniparParserTRUE, 0)
}

func (s *Primary_exprContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MiniparParserFALSE, 0)
}

func (s *Primary_exprContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *Primary_exprContext) Obj_creation() IObj_creationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObj_creationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObj_creationContext)
}

func (s *Primary_exprContext) List_literal() IList_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_literalContext)
}

func (s *Primary_exprContext) Dict_literal() IDict_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDict_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDict_literalContext)
}

func (s *Primary_exprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Primary_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Primary_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterPrimary_expr(s)
	}
}

func (s *Primary_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitPrimary_expr(s)
	}
}

func (p *MiniparParser) Primary_expr() (localctx IPrimary_exprContext) {
	localctx = NewPrimary_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, MiniparParserRULE_primary_expr)
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniparParserNUMERO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(410)
			p.Match(MiniparParserNUMERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniparParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(411)
			p.Match(MiniparParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniparParserTRUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(412)
			p.Match(MiniparParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniparParserFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(413)
			p.Match(MiniparParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniparParserID:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(414)
			p.Id()
		}

	case MiniparParserNEW:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(415)
			p.Obj_creation()
		}

	case MiniparParserT__8:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(416)
			p.List_literal()
		}

	case MiniparParserT__0:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(417)
			p.Dict_literal()
		}

	case MiniparParserT__2:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(418)
			p.Match(MiniparParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.Expr()
		}
		{
			p.SetState(420)
			p.Match(MiniparParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_literalContext is an interface to support dynamic dispatch.
type IList_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsList_literalContext differentiates from other interfaces.
	IsList_literalContext()
}

type List_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_literalContext() *List_literalContext {
	var p = new(List_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_list_literal
	return p
}

func InitEmptyList_literalContext(p *List_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_list_literal
}

func (*List_literalContext) IsList_literalContext() {}

func NewList_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_literalContext {
	var p = new(List_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_list_literal

	return p
}

func (s *List_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *List_literalContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *List_literalContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *List_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterList_literal(s)
	}
}

func (s *List_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitList_literal(s)
	}
}

func (p *MiniparParser) List_literal() (localctx IList_literalContext) {
	localctx = NewList_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, MiniparParserRULE_list_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(424)
		p.Match(MiniparParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&254463824306913802) != 0 {
		{
			p.SetState(425)
			p.Expr()
		}
		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MiniparParserT__5 {
			{
				p.SetState(426)
				p.Match(MiniparParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(427)
				p.Expr()
			}

			p.SetState(432)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(435)
		p.Match(MiniparParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDict_literalContext is an interface to support dynamic dispatch.
type IDict_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllKey_value_pair() []IKey_value_pairContext
	Key_value_pair(i int) IKey_value_pairContext

	// IsDict_literalContext differentiates from other interfaces.
	IsDict_literalContext()
}

type Dict_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDict_literalContext() *Dict_literalContext {
	var p = new(Dict_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_dict_literal
	return p
}

func InitEmptyDict_literalContext(p *Dict_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_dict_literal
}

func (*Dict_literalContext) IsDict_literalContext() {}

func NewDict_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dict_literalContext {
	var p = new(Dict_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_dict_literal

	return p
}

func (s *Dict_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Dict_literalContext) AllKey_value_pair() []IKey_value_pairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKey_value_pairContext); ok {
			len++
		}
	}

	tst := make([]IKey_value_pairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKey_value_pairContext); ok {
			tst[i] = t.(IKey_value_pairContext)
			i++
		}
	}

	return tst
}

func (s *Dict_literalContext) Key_value_pair(i int) IKey_value_pairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKey_value_pairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKey_value_pairContext)
}

func (s *Dict_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dict_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dict_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterDict_literal(s)
	}
}

func (s *Dict_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitDict_literal(s)
	}
}

func (p *MiniparParser) Dict_literal() (localctx IDict_literalContext) {
	localctx = NewDict_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, MiniparParserRULE_dict_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(437)
		p.Match(MiniparParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MiniparParserSTRING {
		{
			p.SetState(438)
			p.Key_value_pair()
		}
		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MiniparParserT__5 {
			{
				p.SetState(439)
				p.Match(MiniparParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(440)
				p.Key_value_pair()
			}

			p.SetState(445)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(448)
		p.Match(MiniparParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKey_value_pairContext is an interface to support dynamic dispatch.
type IKey_value_pairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	Expr() IExprContext

	// IsKey_value_pairContext differentiates from other interfaces.
	IsKey_value_pairContext()
}

type Key_value_pairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKey_value_pairContext() *Key_value_pairContext {
	var p = new(Key_value_pairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_key_value_pair
	return p
}

func InitEmptyKey_value_pairContext(p *Key_value_pairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_key_value_pair
}

func (*Key_value_pairContext) IsKey_value_pairContext() {}

func NewKey_value_pairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Key_value_pairContext {
	var p = new(Key_value_pairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_key_value_pair

	return p
}

func (s *Key_value_pairContext) GetParser() antlr.Parser { return s.parser }

func (s *Key_value_pairContext) STRING() antlr.TerminalNode {
	return s.GetToken(MiniparParserSTRING, 0)
}

func (s *Key_value_pairContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Key_value_pairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Key_value_pairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Key_value_pairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterKey_value_pair(s)
	}
}

func (s *Key_value_pairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitKey_value_pair(s)
	}
}

func (p *MiniparParser) Key_value_pair() (localctx IKey_value_pairContext) {
	localctx = NewKey_value_pairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, MiniparParserRULE_key_value_pair)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(450)
		p.Match(MiniparParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(451)
		p.Match(MiniparParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(452)
		p.Expr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER_TYPE() antlr.TerminalNode
	STRING_TYPE() antlr.TerminalNode
	BOOL_TYPE() antlr.TerminalNode
	VOID_TYPE() antlr.TerminalNode
	LIST_TYPE() antlr.TerminalNode
	DICT_TYPE() antlr.TerminalNode
	Id() IIdContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(MiniparParserNUMBER_TYPE, 0)
}

func (s *TypeContext) STRING_TYPE() antlr.TerminalNode {
	return s.GetToken(MiniparParserSTRING_TYPE, 0)
}

func (s *TypeContext) BOOL_TYPE() antlr.TerminalNode {
	return s.GetToken(MiniparParserBOOL_TYPE, 0)
}

func (s *TypeContext) VOID_TYPE() antlr.TerminalNode {
	return s.GetToken(MiniparParserVOID_TYPE, 0)
}

func (s *TypeContext) LIST_TYPE() antlr.TerminalNode {
	return s.GetToken(MiniparParserLIST_TYPE, 0)
}

func (s *TypeContext) DICT_TYPE() antlr.TerminalNode {
	return s.GetToken(MiniparParserDICT_TYPE, 0)
}

func (s *TypeContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *MiniparParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, MiniparParserRULE_type)
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MiniparParserNUMBER_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(454)
			p.Match(MiniparParserNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniparParserSTRING_TYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(455)
			p.Match(MiniparParserSTRING_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniparParserBOOL_TYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(456)
			p.Match(MiniparParserBOOL_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniparParserVOID_TYPE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(457)
			p.Match(MiniparParserVOID_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniparParserLIST_TYPE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(458)
			p.Match(MiniparParserLIST_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniparParserDICT_TYPE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(459)
			p.Match(MiniparParserDICT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MiniparParserID:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(460)
			p.Id()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdContext is an interface to support dynamic dispatch.
type IIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsIdContext differentiates from other interfaces.
	IsIdContext()
}

type IdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdContext() *IdContext {
	var p = new(IdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_id
	return p
}

func InitEmptyIdContext(p *IdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniparParserRULE_id
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniparParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniparParserID, 0)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniparListener); ok {
		listenerT.ExitId(s)
	}
}

func (p *MiniparParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, MiniparParserRULE_id)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(463)
		p.Match(MiniparParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *MiniparParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 39:
		var t *Postfix_exprContext = nil
		if localctx != nil {
			t = localctx.(*Postfix_exprContext)
		}
		return p.Postfix_expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MiniparParser) Postfix_expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
