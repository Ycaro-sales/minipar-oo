package cgen

import (
	"fmt"
	"strconv"
	"strings"

	"minipar/internal/symtab"
	"minipar/semantic"
	"minipar/tac"
)

// chanInfo holds metadata about a declared channel (s_channel or c_channel).
type chanInfo struct {
	name     string   // Minipar channel variable name
	chanType string   // "s_channel" or "c_channel"
	args     []string // raw TAC PARAM values: (handler,desc,host,port) or (host,port)
}

// funcParam is one parameter of a Minipar function.
type funcParam struct{ name, typ string }

// funcSig holds the pre-scanned signature of a Minipar function.
type funcSig struct {
	params []funcParam
	ret    string
}

type CGenerator struct {
	instrs      []tac.Instruction
	tempTypes   map[string]string
	varTypes    map[string]string // nome → tipo minipar (da tabela de símbolos)
	globalVars  map[string]bool  // declaradas no escopo global; nunca resetadas
	localVars   map[string]bool  // declaradas na função atual; resetadas por função
	paramBuf    []string         // argumentos acumulados antes de CALL
	currentClass string          // classe em definição, para name mangling de métodos
	currentFunc  string          // função atual ("main", etc.)
	insideFunc  bool             // controla indentação: global vs. dentro de função
	channels    []chanInfo       // canais declarados (na ordem de declaração)
	chanSet     map[string]bool  // lookup rápido: nome → é canal?
	funcSigs    map[string]funcSig // assinaturas das funções para dispatch do servidor
	buf         strings.Builder
}

func New(instrs []tac.Instruction, tempTypes map[string]string, globalScope *symtab.Scope[*semantic.Type]) *CGenerator {
	g := &CGenerator{
		instrs:     instrs,
		tempTypes:  tempTypes,
		varTypes:   make(map[string]string),
		globalVars: make(map[string]bool),
		localVars:  make(map[string]bool),
		chanSet:    make(map[string]bool),
		funcSigs:   make(map[string]funcSig),
	}
	g.collectVarTypes(globalScope)
	return g
}

// collectVarTypes percorre a árvore de escopos e constrói o mapa nome → tipo.
func (g *CGenerator) collectVarTypes(scope *symtab.Scope[*semantic.Type]) {
	for _, sym := range scope.Symbols() {
		if sym.Type != nil {
			g.varTypes[sym.Name] = sym.Type.Name
		}
	}
	for _, child := range scope.Children() {
		g.collectVarTypes(child)
	}
}

// resolveType devolve o tipo minipar de uma variável ou temporário.
func (g *CGenerator) resolveType(name string) string {
	if t, ok := g.tempTypes[name]; ok {
		return t
	}
	return g.varTypes[name]
}

// isDeclared informa se name já foi declarado em qualquer escopo.
func (g *CGenerator) isDeclared(name string) bool {
	return g.globalVars[name] || g.localVars[name]
}

// markDeclared registra name no mapa adequado ao escopo atual.
func (g *CGenerator) markDeclared(name string) {
	if g.insideFunc {
		g.localVars[name] = true
	} else {
		g.globalVars[name] = true
	}
}

// declPrefix devolve o prefixo de declaração C se name ainda não foi declarado
// ("int32_t name" na primeira ocorrência, "" nas seguintes).
func (g *CGenerator) declPrefix(name string) string {
	if g.isDeclared(name) {
		return ""
	}
	g.markDeclared(name)
	return mapType(g.resolveType(name)) + " "
}

// ind devolve a indentação adequada ao contexto atual.
func (g *CGenerator) ind() string {
	if g.insideFunc {
		return "    "
	}
	return ""
}

// emitBinary emite uma instrução binária, declarando o resultado se necessário.
func (g *CGenerator) emitBinary(result, arg1, op, arg2 string) {
	prefix := g.declPrefix(result)
	g.buf.WriteString(fmt.Sprintf("%s%s%s = %s %s %s;\n", g.ind(), prefix, result, arg1, op, arg2))
}

// emitUnary emite uma instrução unária, declarando o resultado se necessário.
func (g *CGenerator) emitUnary(result, op, arg string) {
	prefix := g.declPrefix(result)
	g.buf.WriteString(fmt.Sprintf("%s%s%s = %s%s;\n", g.ind(), prefix, result, op, arg))
}

// printFmt devolve o especificador printf adequado para um valor.
// Resolução: (1) literal string → %s; (2) literal numérico com '.' → %f;
// (3) temporário → consulta tempTypes; (4) variável → consulta varTypes; (5) default %d.
func (g *CGenerator) printFmt(val string) string {
	if len(val) >= 2 && val[0] == '"' {
		return "%s"
	}
	hasDot := false
	for _, c := range val {
		if c == '.' {
			hasDot = true
			break
		}
	}
	if hasDot {
		return "%f"
	}
	return fmtForType(g.resolveType(val))
}

// fmtForType mapeia um tipo minipar para o especificador printf correspondente.
func fmtForType(miniparType string) string {
	switch miniparType {
	case "i64":
		return "%ld"
	case "u8", "u16", "u32":
		return "%u"
	case "u64":
		return "%lu"
	case "f32", "f64", "float", "f16":
		return "%f"
	case "char":
		return "%c"
	case "string":
		return "%s"
	default:
		return "%d"
	}
}

// mapType converte um tipo minipar para o equivalente em C.
// Para tipos compostos ([T] e (T0,T1,...)) devolve o nome do typedef gerado.
func mapType(miniparType string) string {
	if len(miniparType) >= 2 && miniparType[0] == '[' && miniparType[len(miniparType)-1] == ']' {
		inner := miniparType[1 : len(miniparType)-1]
		return "arr_" + inner
	}
	if len(miniparType) >= 2 && miniparType[0] == '(' && miniparType[len(miniparType)-1] == ')' {
		inner := miniparType[1 : len(miniparType)-1]
		parts := splitTypeList(inner)
		for i, p := range parts {
			parts[i] = strings.TrimSpace(p)
		}
		return "tup_" + strings.Join(parts, "_")
	}
	switch miniparType {
	case "i8":
		return "int8_t"
	case "i16":
		return "int16_t"
	case "i32":
		return "int32_t"
	case "i64":
		return "int64_t"
	case "u8":
		return "uint8_t"
	case "u16":
		return "uint16_t"
	case "u32":
		return "uint32_t"
	case "u64":
		return "uint64_t"
	case "f16", "f32", "float":
		return "float"
	case "f64":
		return "double"
	case "bool":
		return "bool"
	case "char":
		return "char"
	case "string":
		return "char*"
	case "void", "":
		return "void"
	default:
		return "int"
	}
}

// builtinCName mapeia nomes de builtins minipar para a função C equivalente.
// Apenas os casos que precisam de renomeação estão aqui; "isalpha" já tem o
// mesmo nome na libc e não precisa de entrada.
var builtinCName = map[string]string{
	"len":       "strlen",
	"to_number": "atoi",
	"isnum":     "isdigit",
}

// splitTypeList divide uma lista de tipos separados por vírgula, respeitando
// profundidade de colchetes/parênteses (para tipos compostos aninhados).
func splitTypeList(s string) []string {
	var parts []string
	depth := 0
	start := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[':
			depth++
		case ')', ']':
			depth--
		case ',':
			if depth == 0 {
				parts = append(parts, strings.TrimSpace(s[start:i]))
				start = i + 1
			}
		}
	}
	if start < len(s) {
		parts = append(parts, strings.TrimSpace(s[start:]))
	}
	return parts
}

// arrayElemType extrai o tipo do elemento de um tipo array ("[T]" → "T").
func arrayElemType(arrType string) string {
	if len(arrType) >= 2 && arrType[0] == '[' && arrType[len(arrType)-1] == ']' {
		return arrType[1 : len(arrType)-1]
	}
	return ""
}

// tupleElemTypes extrai os tipos dos elementos de um tipo tupla ("(T0, T1)" → ["T0","T1"]).
func tupleElemTypes(tupType string) []string {
	if len(tupType) >= 2 && tupType[0] == '(' && tupType[len(tupType)-1] == ')' {
		inner := tupType[1 : len(tupType)-1]
		return splitTypeList(inner)
	}
	return nil
}

// stripQuotes remove as aspas de uma string literal TAC (ex.: `"localhost"` → `localhost`).
func stripQuotes(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

// hasChanInstrs verifica se há pelo menos uma instrução CHAN_DECL.
func hasChanInstrs(instrs []tac.Instruction) bool {
	for _, instr := range instrs {
		if instr.Op == "CHAN_DECL" {
			return true
		}
	}
	return false
}

// preScan faz uma passada única pelas instruções para coletar:
//   - Assinaturas de todas as funções (funcSigs) — usadas no dispatch do servidor.
//   - Informações de todos os canais (channels/chanSet) — usadas na geração de sockets.
func (g *CGenerator) preScan() {
	var curFunc string
	var sig funcSig

	for i, instr := range g.instrs {
		switch instr.Op {
		case "BEGIN_FUNC", "BEGIN_METHOD":
			curFunc = instr.Arg1
			sig = funcSig{ret: instr.Arg2}
		case "PARAM_DECL":
			if curFunc != "" {
				sig.params = append(sig.params, funcParam{instr.Arg1, instr.Arg2})
			}
		case "END_FUNC", "END_METHOD":
			if curFunc != "" {
				g.funcSigs[curFunc] = sig
			}
			curFunc = ""
			sig = funcSig{}
		case "CHAN_DECL":
			nArgs, _ := strconv.Atoi(instr.Result)
			var args []string
			start := i - nArgs
			if start < 0 {
				start = 0
			}
			for j := start; j < i; j++ {
				if g.instrs[j].Op == "PARAM" {
					args = append(args, g.instrs[j].Arg1)
				}
			}
			ci := chanInfo{name: instr.Arg2, chanType: instr.Arg1, args: args}
			g.channels = append(g.channels, ci)
			g.chanSet[instr.Arg2] = true
		}
	}
}

// emitChanInit emite o código C de inicialização de um canal dentro de main().
// Para c_channel: socket + connect + recv boas-vindas.
// Para s_channel: socket + setsockopt + bind + listen.
func (g *CGenerator) emitChanInit(ci chanInfo) string {
	var sb strings.Builder
	name := ci.name

	if ci.chanType == "c_channel" {
		// args: [host_expr, port_expr]
		host := "\"127.0.0.1\""
		port := "5000"
		if len(ci.args) >= 1 {
			host = ci.args[0]
		}
		if len(ci.args) >= 2 {
			port = ci.args[1]
		}
		hostStr := stripQuotes(host)

		sb.WriteString(fmt.Sprintf("    /* Connect client channel '%s' */\n", name))
		sb.WriteString(fmt.Sprintf("    %s_fd = socket(AF_INET, SOCK_STREAM, 0);\n", name))
		sb.WriteString(fmt.Sprintf("    if (%s_fd < 0) { perror(\"socket\"); exit(1); }\n", name))
		sb.WriteString("    {\n")
		sb.WriteString(fmt.Sprintf("        struct sockaddr_in _sa_%s;\n", name))
		sb.WriteString(fmt.Sprintf("        memset(&_sa_%s, 0, sizeof(_sa_%s));\n", name, name))
		sb.WriteString(fmt.Sprintf("        _sa_%s.sin_family = AF_INET;\n", name))
		sb.WriteString(fmt.Sprintf("        struct hostent* _he_%s = gethostbyname(\"%s\");\n", name, hostStr))
		sb.WriteString(fmt.Sprintf("        if (_he_%s) memcpy(&_sa_%s.sin_addr, _he_%s->h_addr, _he_%s->h_length);\n",
			name, name, name, name))
		sb.WriteString(fmt.Sprintf("        _sa_%s.sin_port = htons(%s);\n", name, port))
		sb.WriteString(fmt.Sprintf("        if (connect(%s_fd, (struct sockaddr*)&_sa_%s, sizeof(_sa_%s)) < 0) { perror(\"connect\"); exit(1); }\n",
			name, name, name))
		sb.WriteString("    }\n")
		sb.WriteString("    {\n")
		sb.WriteString(fmt.Sprintf("        char _wb_%s[4096];\n", name))
		sb.WriteString(fmt.Sprintf("        memset(_wb_%s, 0, sizeof(_wb_%s));\n", name, name))
		sb.WriteString(fmt.Sprintf("        recv(%s_fd, _wb_%s, sizeof(_wb_%s)-1, 0);\n", name, name, name))
		sb.WriteString(fmt.Sprintf("        printf(\"[%s] Server says: %%s\\n\", _wb_%s);\n", name, name))
		sb.WriteString("    }\n")
	} else {
		// s_channel: args: [handler, desc, host, port]
		port := "5000"
		if len(ci.args) >= 4 {
			port = ci.args[3]
		}

		sb.WriteString(fmt.Sprintf("    /* Setup server channel '%s' */\n", name))
		sb.WriteString(fmt.Sprintf("    %s_fd = socket(AF_INET, SOCK_STREAM, 0);\n", name))
		sb.WriteString(fmt.Sprintf("    if (%s_fd < 0) { perror(\"socket\"); exit(1); }\n", name))
		sb.WriteString("    {\n")
		sb.WriteString(fmt.Sprintf("        int _opt_%s = 1;\n", name))
		sb.WriteString(fmt.Sprintf("        setsockopt(%s_fd, SOL_SOCKET, SO_REUSEADDR, &_opt_%s, sizeof(_opt_%s));\n",
			name, name, name))
		sb.WriteString(fmt.Sprintf("        struct sockaddr_in _sa_%s;\n", name))
		sb.WriteString(fmt.Sprintf("        memset(&_sa_%s, 0, sizeof(_sa_%s));\n", name, name))
		sb.WriteString(fmt.Sprintf("        _sa_%s.sin_family = AF_INET;\n", name))
		sb.WriteString(fmt.Sprintf("        _sa_%s.sin_addr.s_addr = INADDR_ANY;\n", name))
		sb.WriteString(fmt.Sprintf("        _sa_%s.sin_port = htons(%s);\n", name, port))
		sb.WriteString(fmt.Sprintf("        if (bind(%s_fd, (struct sockaddr*)&_sa_%s, sizeof(_sa_%s)) < 0) { perror(\"bind\"); exit(1); }\n",
			name, name, name))
		sb.WriteString(fmt.Sprintf("        listen(%s_fd, 5);\n", name))
		sb.WriteString(fmt.Sprintf("        printf(\"[Server '%%s'] Listening on port %s\\n\", \"%s\");\n", port, name))
		sb.WriteString("    }\n")
	}
	return sb.String()
}

// emitServerLoop emite o loop de accept/recv/dispatch para um s_channel, injetado
// no final de main() antes do return 0.
func (g *CGenerator) emitServerLoop(ci chanInfo) string {
	var sb strings.Builder
	name := ci.name

	// s_channel args: [handler, desc, host, port]
	handlerName := ""
	descExpr := ""
	if len(ci.args) >= 1 {
		handlerName = ci.args[0]
	}
	if len(ci.args) >= 2 {
		descExpr = ci.args[1]
	}

	sig := g.funcSigs[handlerName]

	sb.WriteString(fmt.Sprintf("    /* Server accept-loop for channel '%s' (handler: %s) */\n", name, handlerName))
	sb.WriteString("    {\n")
	sb.WriteString("        while (1) {\n")
	sb.WriteString(fmt.Sprintf("            struct sockaddr_in _ca_%s;\n", name))
	sb.WriteString(fmt.Sprintf("            socklen_t _cl_%s = sizeof(_ca_%s);\n", name, name))
	sb.WriteString(fmt.Sprintf("            int _conn_%s = accept(%s_fd, (struct sockaddr*)&_ca_%s, &_cl_%s);\n",
		name, name, name, name))
	sb.WriteString(fmt.Sprintf("            if (_conn_%s < 0) break;\n", name))

	// Send description as welcome message
	if descExpr != "" {
		sb.WriteString(fmt.Sprintf("            send(_conn_%s, %s, strlen(%s), 0);\n", name, descExpr, descExpr))
	} else {
		sb.WriteString(fmt.Sprintf("            send(_conn_%s, \"%s\", %d, 0);\n", name, name, len(name)))
	}

	// Inner request-handling loop
	sb.WriteString(fmt.Sprintf("            char _buf_%s[4096];\n", name))
	sb.WriteString("            while (1) {\n")
	sb.WriteString(fmt.Sprintf("                memset(_buf_%s, 0, sizeof(_buf_%s));\n", name, name))
	sb.WriteString(fmt.Sprintf("                int _n_%s = (int)recv(_conn_%s, _buf_%s, sizeof(_buf_%s)-1, 0);\n",
		name, name, name, name))
	sb.WriteString(fmt.Sprintf("                if (_n_%s <= 0) break;\n", name))

	// Parse comma-separated tokens into typed handler params
	var callArgs []string
	if len(sig.params) > 0 {
		sb.WriteString("                char* _tok;\n")
		for idx, p := range sig.params {
			varName := fmt.Sprintf("_p%d_%s", idx, name)
			cType := mapType(p.typ)
			if idx == 0 {
				sb.WriteString(fmt.Sprintf("                _tok = strtok(_buf_%s, \",\");\n", name))
			} else {
				sb.WriteString("                _tok = strtok(NULL, \",\");\n")
			}
			switch p.typ {
			case "string":
				sb.WriteString(fmt.Sprintf("                %s %s = _tok ? _tok : \"\";\n", cType, varName))
			case "f64", "f32", "float", "f16":
				sb.WriteString(fmt.Sprintf("                %s %s = _tok ? atof(_tok) : 0;\n", cType, varName))
			case "i64":
				sb.WriteString(fmt.Sprintf("                %s %s = _tok ? atoll(_tok) : 0;\n", cType, varName))
			default:
				sb.WriteString(fmt.Sprintf("                %s %s = _tok ? atoi(_tok) : 0;\n", cType, varName))
			}
			callArgs = append(callArgs, varName)
		}
	}

	// Call the handler and send result back
	callStr := fmt.Sprintf("%s(%s)", handlerName, strings.Join(callArgs, ", "))
	retType := mapType(sig.ret)
	retFmt := fmtForType(sig.ret)
	if sig.ret == "" || sig.ret == "void" {
		sb.WriteString(fmt.Sprintf("                %s;\n", callStr))
		sb.WriteString(fmt.Sprintf("                send(_conn_%s, \"OK\", 2, 0);\n", name))
	} else {
		sb.WriteString(fmt.Sprintf("                %s _res_%s = %s;\n", retType, name, callStr))
		sb.WriteString("                char _resp[256];\n")
		sb.WriteString(fmt.Sprintf("                snprintf(_resp, sizeof(_resp), \"%s\", _res_%s);\n", retFmt, name))
		sb.WriteString(fmt.Sprintf("                send(_conn_%s, _resp, strlen(_resp), 0);\n", name))
	}

	sb.WriteString("            }\n") // end inner while
	sb.WriteString(fmt.Sprintf("            close(_conn_%s);\n", name))
	sb.WriteString("        }\n") // end outer while
	sb.WriteString("    }\n")
	return sb.String()
}

// emitChanSend emite o código C de um channel.send(args...) — serializa os args
// em CSV, envia via socket, recebe a resposta e imprime.
func (g *CGenerator) emitChanSend(chanName string, args []string) string {
	var sb strings.Builder
	var fmts []string
	for _, arg := range args {
		fmts = append(fmts, g.printFmt(arg))
	}
	fmtStr := strings.Join(fmts, ",")
	argsJoined := strings.Join(args, ", ")

	sb.WriteString("    {\n")
	sb.WriteString("        char _sbuf[4096];\n")
	if len(args) == 0 {
		sb.WriteString("        strcpy(_sbuf, \"\");\n")
	} else {
		sb.WriteString(fmt.Sprintf("        snprintf(_sbuf, sizeof(_sbuf), \"%s\", %s);\n", fmtStr, argsJoined))
	}
	sb.WriteString(fmt.Sprintf("        send(%s_fd, _sbuf, strlen(_sbuf), 0);\n", chanName))
	sb.WriteString("        char _rbuf[4096];\n")
	sb.WriteString("        memset(_rbuf, 0, sizeof(_rbuf));\n")
	sb.WriteString(fmt.Sprintf("        recv(%s_fd, _rbuf, sizeof(_rbuf)-1, 0);\n", chanName))
	sb.WriteString(fmt.Sprintf("        printf(\"[%s] Response: %%s\\n\", _rbuf);\n", chanName))
	sb.WriteString("    }\n")
	return sb.String()
}

// emitCompositeTypedefs escaneia todas as instruções e tipos conhecidos, e emite
// um typedef C para cada tipo composto único ([T] e (T0,T1,...)) encontrado.
func (g *CGenerator) emitCompositeTypedefs() {
	seen := make(map[string]bool)

	// coleta todos os tipos minipar conhecidos
	var allTypes []string
	for _, t := range g.tempTypes {
		allTypes = append(allTypes, t)
	}
	for _, t := range g.varTypes {
		allTypes = append(allTypes, t)
	}
	for _, instr := range g.instrs {
		if instr.Op == "PARAM_DECL" {
			allTypes = append(allTypes, instr.Arg2)
		}
	}

	for _, t := range allTypes {
		if seen[t] {
			continue
		}
		if len(t) >= 2 && t[0] == '[' && t[len(t)-1] == ']' {
			seen[t] = true
			elem := arrayElemType(t)
			elemC := mapType(elem)
			typedefName := mapType(t) // "arr_i32"
			g.buf.WriteString(fmt.Sprintf("typedef struct { %s* data; int len; } %s;\n", elemC, typedefName))
		} else if len(t) >= 2 && t[0] == '(' && t[len(t)-1] == ')' {
			seen[t] = true
			elems := tupleElemTypes(t)
			typedefName := mapType(t) // "tup_int_string"
			g.buf.WriteString("typedef struct {\n")
			for i, e := range elems {
				g.buf.WriteString(fmt.Sprintf("    %s _%d;\n", mapType(strings.TrimSpace(e)), i))
			}
			g.buf.WriteString(fmt.Sprintf("} %s;\n", typedefName))
		}
	}
}

// builtinUsage rastreia quais builtins/helpers são necessários no preâmbulo.
type builtinUsage struct {
	toString bool // helper próprio: to_string(long long)
	mpInput  bool // helper próprio: mp_input()
	ctype    bool // necessário #include <ctype.h> (isalpha ou isnum/isdigit)
}

// scanBuiltins faz uma única passada pelas instruções para detectar quais
// builtins são utilizados, evitando emitir helpers/includes desnecessários.
func (g *CGenerator) scanBuiltins() builtinUsage {
	var u builtinUsage
	for _, instr := range g.instrs {
		switch instr.Op {
		case "CALL":
			switch instr.Arg1 {
			case "to_string":
				u.toString = true
			case "isalpha", "isnum":
				u.ctype = true
			}
		case "INPUT":
			u.mpInput = true
		}
	}
	return u
}

// emitBuiltinHelpers escreve no preâmbulo as definições dos helpers próprios
// (aqueles sem equivalente direto na libc) que foram detectados no pré-scan.
func (g *CGenerator) emitBuiltinHelpers(u builtinUsage) {
	if u.toString {
		g.buf.WriteString("static char* to_string(long long v) {\n")
		g.buf.WriteString("    static char b[32];\n")
		g.buf.WriteString("    snprintf(b, sizeof b, \"%lld\", v);\n")
		g.buf.WriteString("    return b;\n")
		g.buf.WriteString("}\n")
	}
	if u.mpInput {
		g.buf.WriteString("static char* mp_input(void) {\n")
		g.buf.WriteString("    static char b[256];\n")
		g.buf.WriteString("    if (fgets(b, sizeof b, stdin)) b[strcspn(b, \"\\n\")] = '\\0';\n")
		g.buf.WriteString("    return b;\n")
		g.buf.WriteString("}\n")
	}
}

func (g *CGenerator) Generate() string {
	g.preScan()
	usage := g.scanBuiltins()

	g.buf.WriteString("#include <stdint.h>\n")
	g.buf.WriteString("#include <stdbool.h>\n")
	g.buf.WriteString("#include <stdio.h>\n")
	g.buf.WriteString("#include <stdlib.h>\n")
	g.buf.WriteString("#include <string.h>\n")
	if usage.ctype {
		g.buf.WriteString("#include <ctype.h>\n")
	}
	if hasChanInstrs(g.instrs) {
		g.buf.WriteString("#include <sys/socket.h>\n")
		g.buf.WriteString("#include <netinet/in.h>\n")
		g.buf.WriteString("#include <arpa/inet.h>\n")
		g.buf.WriteString("#include <netdb.h>\n")
		g.buf.WriteString("#include <unistd.h>\n")
	}
	g.buf.WriteByte('\n')
	// Emit global socket fd variables for every declared channel.
	for _, ci := range g.channels {
		g.buf.WriteString(fmt.Sprintf("static int %s_fd = -1;\n", ci.name))
	}
	g.emitCompositeTypedefs()
	g.emitBuiltinHelpers(usage)
	g.buf.WriteByte('\n')

	for i := 0; i < len(g.instrs); i++ {
		instr := g.instrs[i]
		switch instr.Op {
		case "BEGIN_FUNC":
			retC := mapType(instr.Arg2)
			if instr.Arg1 == "main" {
				retC = "int"
			}
			// Reset per-function declared vars so parameters of previous
			// functions do not shadow locals in this function.
			// Globals are kept separate in globalVars and are never reset.
			g.localVars = make(map[string]bool)
			// Consome os PARAM_DECL seguintes para montar a lista de parâmetros.
			var params []string
			for i+1 < len(g.instrs) && g.instrs[i+1].Op == "PARAM_DECL" {
				i++
				p := g.instrs[i]
				params = append(params, fmt.Sprintf("%s %s", mapType(p.Arg2), p.Arg1))
				g.localVars[p.Arg1] = true // param already declared in signature
			}
			g.currentFunc = instr.Arg1
			g.buf.WriteString(fmt.Sprintf("%s %s(%s) {\n", retC, instr.Arg1, strings.Join(params, ", ")))
			g.insideFunc = true
			// Inject channel initialization at the top of main().
			if instr.Arg1 == "main" {
				for _, ci := range g.channels {
					g.buf.WriteString(g.emitChanInit(ci))
				}
			}

		case "PARAM_DECL":
			// Já consumido pelo look-ahead de BEGIN_FUNC; ignorar se aparecer solto.

		case "PARAM":
			g.paramBuf = append(g.paramBuf, instr.Arg1)

		case "CALL":
			args := strings.Join(g.paramBuf, ", ")
			g.paramBuf = g.paramBuf[:0]
			// Traduzir nome do builtin para a função C equivalente (quando houver).
			callName := instr.Arg1
			if cName, ok := builtinCName[callName]; ok {
				callName = cName
			}
			retType := g.resolveType(instr.Result)
			if retType == "void" || retType == "" {
				g.buf.WriteString(fmt.Sprintf("%s%s(%s);\n", g.ind(), callName, args))
			} else {
				prefix := g.declPrefix(instr.Result)
				g.buf.WriteString(fmt.Sprintf("%s%s%s = %s(%s);\n", g.ind(), prefix, instr.Result, callName, args))
			}

		case "ASSIGN":
			prefix := g.declPrefix(instr.Result)
			g.buf.WriteString(fmt.Sprintf("%s%s%s = %s;\n", g.ind(), prefix, instr.Result, instr.Arg1))

		// Aritmética
		case "ADD":
			g.emitBinary(instr.Result, instr.Arg1, "+", instr.Arg2)
		case "SUB":
			g.emitBinary(instr.Result, instr.Arg1, "-", instr.Arg2)
		case "MUL":
			g.emitBinary(instr.Result, instr.Arg1, "*", instr.Arg2)
		case "DIV":
			g.emitBinary(instr.Result, instr.Arg1, "/", instr.Arg2)
		case "MOD":
			g.emitBinary(instr.Result, instr.Arg1, "%", instr.Arg2)
		case "NEG":
			g.emitUnary(instr.Result, "-", instr.Arg1)

		// Lógica
		case "AND":
			g.emitBinary(instr.Result, instr.Arg1, "&&", instr.Arg2)
		case "OR":
			g.emitBinary(instr.Result, instr.Arg1, "||", instr.Arg2)
		case "NOT":
			g.emitUnary(instr.Result, "!", instr.Arg1)

		// Comparação
		case "EQ":
			g.emitBinary(instr.Result, instr.Arg1, "==", instr.Arg2)
		case "NEQ":
			g.emitBinary(instr.Result, instr.Arg1, "!=", instr.Arg2)
		case "LT":
			g.emitBinary(instr.Result, instr.Arg1, "<", instr.Arg2)
		case "GT":
			g.emitBinary(instr.Result, instr.Arg1, ">", instr.Arg2)
		case "LEQ":
			g.emitBinary(instr.Result, instr.Arg1, "<=", instr.Arg2)
		case "GEQ":
			g.emitBinary(instr.Result, instr.Arg1, ">=", instr.Arg2)

		// Arrays
		case "ARRAY_NEW":
			arrTemp := instr.Result
			count := instr.Arg1
			arrType := g.resolveType(arrTemp) // "[i32]" after type propagation
			elemType := arrayElemType(arrType)
			if elemType == "" {
				elemType = instr.Arg2 // fallback: use Arg2 from TAC
			}
			elemC := mapType(elemType)
			typedefName := mapType(arrType)

			// look-ahead: collect consecutive ARRAY_SET for this temp
			var values []string
			for i+1 < len(g.instrs) && g.instrs[i+1].Op == "ARRAY_SET" && g.instrs[i+1].Arg1 == arrTemp {
				i++
				values = append(values, g.instrs[i].Result)
			}

			dataVar := "_" + arrTemp + "_data"
			g.buf.WriteString(fmt.Sprintf("%s%s %s[] = {%s};\n",
				g.ind(), elemC, dataVar, strings.Join(values, ", ")))
			g.buf.WriteString(fmt.Sprintf("%s%s %s = { %s, %s };\n",
				g.ind(), typedefName, arrTemp, dataVar, count))
			g.markDeclared(arrTemp)

		case "ARRAY_SET":
			// Standalone ARRAY_SET (mutation): Arg1=array, Arg2=index, Result=value
			g.buf.WriteString(fmt.Sprintf("%s%s.data[%s] = %s;\n",
				g.ind(), instr.Arg1, instr.Arg2, instr.Result))

		case "ARRAY_GET":
			// Arg1=array, Arg2=index, Result=temp
			prefix := g.declPrefix(instr.Result)
			g.buf.WriteString(fmt.Sprintf("%s%s%s = %s.data[%s];\n",
				g.ind(), prefix, instr.Result, instr.Arg1, instr.Arg2))

		case "ARRAY_LEN":
			// Arg1=array, Result=temp
			prefix := g.declPrefix(instr.Result)
			g.buf.WriteString(fmt.Sprintf("%s%s%s = %s.len;\n",
				g.ind(), prefix, instr.Result, instr.Arg1))

		// Tuplas
		case "TUPLE_NEW":
			tupTemp := instr.Result
			tupType := g.resolveType(tupTemp)
			typedefName := mapType(tupType)

			// look-ahead: collect consecutive TUPLE_SET for this temp
			var values []string
			for i+1 < len(g.instrs) && g.instrs[i+1].Op == "TUPLE_SET" && g.instrs[i+1].Arg1 == tupTemp {
				i++
				values = append(values, g.instrs[i].Result)
			}

			g.buf.WriteString(fmt.Sprintf("%s%s %s = { %s };\n",
				g.ind(), typedefName, tupTemp, strings.Join(values, ", ")))
			g.markDeclared(tupTemp)

		case "TUPLE_SET":
			// Standalone TUPLE_SET: Arg1=tuple, Arg2=index, Result=value
			g.buf.WriteString(fmt.Sprintf("%s%s._%s = %s;\n",
				g.ind(), instr.Arg1, instr.Arg2, instr.Result))

		case "INPUT":
			// Arg1 = prompt (pode ser ""), Result = temp que recebe o valor lido.
			if instr.Arg1 != "" {
				g.buf.WriteString(fmt.Sprintf("%sprintf(\"%%s\", %s);\n", g.ind(), instr.Arg1))
			}
			// Força o tipo do resultado para "string" (o TAC cria o temp sem tipo).
			g.tempTypes[instr.Result] = "string"
			prefix := g.declPrefix(instr.Result)
			g.buf.WriteString(fmt.Sprintf("%s%s%s = mp_input();\n", g.ind(), prefix, instr.Result))

		case "PRINT":
			g.buf.WriteString(fmt.Sprintf("%sprintf(\"%s\\n\", %s);\n", g.ind(), g.printFmt(instr.Arg1), instr.Arg1))

		// Controle de fluxo
		case "LABEL":
			g.buf.WriteString(fmt.Sprintf("%s:;\n", instr.Arg1))
		case "GOTO":
			g.buf.WriteString(fmt.Sprintf("%sgoto %s;\n", g.ind(), instr.Arg1))
		case "IF_FALSE":
			g.buf.WriteString(fmt.Sprintf("%sif (!%s) goto %s;\n", g.ind(), instr.Arg1, instr.Result))
		case "RETURN":
			if instr.Arg1 == "" {
				g.buf.WriteString(fmt.Sprintf("%sreturn;\n", g.ind()))
			} else {
				g.buf.WriteString(fmt.Sprintf("%sreturn %s;\n", g.ind(), instr.Arg1))
			}

		case "END_FUNC":
			if instr.Arg1 == "main" {
				// Inject accept-loops for any s_channel servers before return.
				for _, ci := range g.channels {
					if ci.chanType == "s_channel" {
						g.buf.WriteString(g.emitServerLoop(ci))
					}
				}
				g.buf.WriteString("    return 0;\n")
			}
			g.buf.WriteString("}\n")
			g.insideFunc = false
			g.currentFunc = ""

		// OOP
		case "BEGIN_CLASS":
			g.currentClass = instr.Arg1
			g.buf.WriteString("typedef struct {\n")
		case "FIELD":
			g.buf.WriteString(fmt.Sprintf("    %s %s;\n", mapType(instr.Arg2), instr.Arg1))
		case "END_CLASS":
			g.buf.WriteString(fmt.Sprintf("} %s;\n", instr.Arg1))
			g.currentClass = ""
		case "BEGIN_CTOR":
			g.buf.WriteString(fmt.Sprintf("void %s_init(%s* self) {\n", instr.Arg1, instr.Arg1))
		case "END_CTOR":
			g.buf.WriteString("}\n")
		case "BEGIN_METHOD":
			retC := mapType(instr.Arg2)
			g.localVars = make(map[string]bool)
			g.localVars["self"] = true
			var params []string
			for i+1 < len(g.instrs) && g.instrs[i+1].Op == "PARAM_DECL" {
				i++
				p := g.instrs[i]
				params = append(params, fmt.Sprintf("%s %s", mapType(p.Arg2), p.Arg1))
				g.localVars[p.Arg1] = true
			}
			paramList := strings.Join(params, ", ")
			if paramList != "" {
				paramList = ", " + paramList
			}
			g.buf.WriteString(fmt.Sprintf("%s %s_%s(%s* self%s) {\n", retC, g.currentClass, instr.Arg1, g.currentClass, paramList))
		case "END_METHOD":
			g.buf.WriteString("}\n")
		case "INTERFACE":
			g.buf.WriteString(fmt.Sprintf("/* interface %s (sem equivalente em C) */\n", instr.Arg1))

		// Concorrência (stubs)
		case "BEGIN_SEQ":
			g.buf.WriteString("    /* BEGIN_SEQ */\n")
		case "END_SEQ":
			g.buf.WriteString("    /* END_SEQ */\n")
		case "BEGIN_PAR":
			g.buf.WriteString("    /* BEGIN_PAR — pthreads: TODO */\n")
		case "END_PAR":
			g.buf.WriteString("    /* END_PAR */\n")
		case "CHAN_DECL":
			// Channel was already registered in preScan and its fd global was emitted
			// in the preâmbulo. Drain any PARAMs that preceded this instruction so
			// they are not mistakenly used by a subsequent CALL.
			g.paramBuf = g.paramBuf[:0]

		case "METHOD_CALL":
			// Arg1 = "<obj>.<method>", Result = arg count, paramBuf holds the args.
			nArgs, _ := strconv.Atoi(instr.Result)
			args := make([]string, len(g.paramBuf))
			copy(args, g.paramBuf)
			g.paramBuf = g.paramBuf[:0]
			_ = nArgs

			// Split "obj.method".
			dot := strings.LastIndex(instr.Arg1, ".")
			if dot < 0 {
				g.buf.WriteString(fmt.Sprintf("    /* METHOD_CALL %s — formato inválido */\n", instr.Arg1))
				break
			}
			obj := instr.Arg1[:dot]
			method := instr.Arg1[dot+1:]

			if g.chanSet[obj] {
				switch method {
				case "send":
					g.buf.WriteString(g.emitChanSend(obj, args))
				case "close":
					g.buf.WriteString(fmt.Sprintf("    close(%s_fd);\n", obj))
				default:
					g.buf.WriteString(fmt.Sprintf("    /* channel method '%s' não implementado */\n", method))
				}
			} else {
				// Regular object method — class support is separate; emit a plain call.
				callArgs := strings.Join(args, ", ")
				result := instr.Arg2
				if result == "" || result == "_" {
					g.buf.WriteString(fmt.Sprintf("    %s_%s(%s);\n", obj, method, callArgs))
				} else {
					prefix := g.declPrefix(result)
					g.buf.WriteString(fmt.Sprintf("    %s%s = %s_%s(%s);\n", prefix, result, obj, method, callArgs))
				}
			}
		}
	}

	return g.buf.String()
}
