package cgen

import (
	"fmt"
	"strings"

	"minipar/internal/symtab"
	"minipar/semantic"
	"minipar/tac"
)

type CGenerator struct {
	instrs       []tac.Instruction
	tempTypes    map[string]string
	varTypes     map[string]string
	globalVars   map[string]bool
	localVars    map[string]bool
	paramBuf     []string
	currentClass string
	insideFunc   bool
	buf          strings.Builder
}

func New(instrs []tac.Instruction, tempTypes map[string]string, globalScope *symtab.Scope[*semantic.Type]) *CGenerator {
	g := &CGenerator{
		instrs:     instrs,
		tempTypes:  tempTypes,
		varTypes:   make(map[string]string),
		globalVars: make(map[string]bool),
		localVars:  make(map[string]bool),
	}
	g.collectVarTypes(globalScope)
	return g
}

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

func (g *CGenerator) resolveType(name string) string {
	if t, ok := g.tempTypes[name]; ok {
		return t
	}
	return g.varTypes[name]
}

func (g *CGenerator) isDeclared(name string) bool {
	return g.globalVars[name] || g.localVars[name]
}

func (g *CGenerator) markDeclared(name string) {
	if g.insideFunc {
		g.localVars[name] = true
	} else {
		g.globalVars[name] = true
	}
}

func (g *CGenerator) declPrefix(name string) string {
	if g.isDeclared(name) {
		return ""
	}
	g.markDeclared(name)
	return mapType(g.resolveType(name)) + " "
}

func (g *CGenerator) ind() string {
	if g.insideFunc {
		return "    "
	}
	return ""
}

func (g *CGenerator) emitBinary(result, arg1, op, arg2 string) {
	prefix := g.declPrefix(result)
	g.buf.WriteString(fmt.Sprintf("%s%s%s = %s %s %s;\n", g.ind(), prefix, result, arg1, op, arg2))
}

func (g *CGenerator) emitUnary(result, op, arg string) {
	prefix := g.declPrefix(result)
	g.buf.WriteString(fmt.Sprintf("%s%s%s = %s%s;\n", g.ind(), prefix, result, op, arg))
}

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

var builtinCName = map[string]string{
	"len":       "strlen",
	"to_number": "atoi",
	"isnum":     "isdigit",
}

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

func arrayElemType(arrType string) string {
	if len(arrType) >= 2 && arrType[0] == '[' && arrType[len(arrType)-1] == ']' {
		return arrType[1 : len(arrType)-1]
	}
	return ""
}

func tupleElemTypes(tupType string) []string {
	if len(tupType) >= 2 && tupType[0] == '(' && tupType[len(tupType)-1] == ')' {
		inner := tupType[1 : len(tupType)-1]
		return splitTypeList(inner)
	}
	return nil
}

func (g *CGenerator) emitCompositeTypedefs() {
	seen := make(map[string]bool)
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
			typedefName := mapType(t)
			g.buf.WriteString(fmt.Sprintf("typedef struct { %s* data; int len; } %s;\n", elemC, typedefName))
		} else if len(t) >= 2 && t[0] == '(' && t[len(t)-1] == ')' {
			seen[t] = true
			elems := tupleElemTypes(t)
			typedefName := mapType(t)
			g.buf.WriteString("typedef struct {\n")
			for i, e := range elems {
				g.buf.WriteString(fmt.Sprintf("    %s _%d;\n", mapType(strings.TrimSpace(e)), i))
			}
			g.buf.WriteString(fmt.Sprintf("} %s;\n", typedefName))
		}
	}
}

type builtinUsage struct {
	toString bool
	mpInput  bool
	ctype    bool
	openmp   bool
	network  bool
}

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
		case "BEGIN_PAR":
			u.openmp = true
		case "CHAN_DECL", "METHOD_CALL":
			u.network = true
		}
	}
	return u
}

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
	if u.network {
		// ==========================================
		// CLIENTE TCP REAL (POSIX Sockets)
		// ==========================================
		g.buf.WriteString("static int create_client(const char* ip, int port) {\n")
		g.buf.WriteString("    int sock = 0;\n")
		g.buf.WriteString("    if ((sock = socket(AF_INET, SOCK_STREAM, 0)) < 0) { printf(\"Erro na criacao do socket\\n\"); return -1; }\n")
		g.buf.WriteString("    struct sockaddr_in serv_addr;\n")
		g.buf.WriteString("    serv_addr.sin_family = AF_INET;\n")
		g.buf.WriteString("    serv_addr.sin_port = htons(port);\n")
		g.buf.WriteString("    if (strcmp(ip, \"localhost\") == 0) ip = \"127.0.0.1\";\n")
		g.buf.WriteString("    inet_pton(AF_INET, ip, &serv_addr.sin_addr);\n")
		g.buf.WriteString("    if (connect(sock, (struct sockaddr *)&serv_addr, sizeof(serv_addr)) < 0) { printf(\"Erro de conexao (Servidor offline?)\\n\"); return -1; }\n")
		g.buf.WriteString("    return sock;\n")
		g.buf.WriteString("}\n\n")

		g.buf.WriteString("static void channel_send(int sock, const char* op, double v1, double v2) {\n")
		g.buf.WriteString("    if (sock < 0) return;\n")
		g.buf.WriteString("    char buffer[256] = {0};\n")
		g.buf.WriteString("    snprintf(buffer, sizeof(buffer), \"%s %f %f\", op, v1, v2);\n")
		g.buf.WriteString("    send(sock, buffer, strlen(buffer), 0); // Envia o pacote TCP real\n")

		g.buf.WriteString("    memset(buffer, 0, sizeof(buffer));\n")
		g.buf.WriteString("    int valread = recv(sock, buffer, 255, 0); // Fica travado aguardando o servidor devolver\n")
		g.buf.WriteString("    if (valread > 0) printf(\"  [CLIENTE] Recebeu a resposta TCP: %s\\n\", buffer);\n")
		g.buf.WriteString("}\n\n")

		g.buf.WriteString("static void channel_close(int sock) {\n")
		g.buf.WriteString("    if (sock >= 0) close(sock);\n")
		g.buf.WriteString("}\n\n")

		// ==========================================
		// SERVIDOR TCP REAL (POSIX Sockets)
		// ==========================================
		g.buf.WriteString("static void start_server(void* callback, const char* desc, const char* ip, int port) {\n")
		g.buf.WriteString("    int server_fd, new_socket;\n")
		g.buf.WriteString("    struct sockaddr_in address;\n")
		g.buf.WriteString("    int opt = 1;\n")
		g.buf.WriteString("    int addrlen = sizeof(address);\n")

		g.buf.WriteString("    if ((server_fd = socket(AF_INET, SOCK_STREAM, 0)) == 0) exit(EXIT_FAILURE);\n")
		g.buf.WriteString("    setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt)); // Permite reiniciar rapido\n")
		g.buf.WriteString("    address.sin_family = AF_INET;\n")
		g.buf.WriteString("    address.sin_addr.s_addr = INADDR_ANY;\n")
		g.buf.WriteString("    address.sin_port = htons(port);\n")

		g.buf.WriteString("    if (bind(server_fd, (struct sockaddr *)&address, sizeof(address)) < 0) { perror(\"bind\"); exit(EXIT_FAILURE); }\n")
		g.buf.WriteString("    if (listen(server_fd, 3) < 0) { perror(\"listen\"); exit(EXIT_FAILURE); }\n")

		g.buf.WriteString("    printf(\"Server [%s] started on %s:%d\\n\", desc, ip, port);\n")
		g.buf.WriteString("    printf(\"Aguardando conexoes REAIS...\\n\");\n")

		g.buf.WriteString("    while(1) {\n") // Loop principal do servidor (mantém ele vivo)
		g.buf.WriteString("        if ((new_socket = accept(server_fd, (struct sockaddr *)&address, (socklen_t*)&addrlen)) < 0) continue;\n")
		g.buf.WriteString("        printf(\"\\n  [SERVIDOR] *** Nova Conexao TCP Aceita ***\\n\");\n")

		g.buf.WriteString("        char buffer[256];\n")
		g.buf.WriteString("        while(1) {\n") // Loop de comunicação contínua com 1 cliente específico
		g.buf.WriteString("            memset(buffer, 0, sizeof(buffer));\n")
		g.buf.WriteString("            if (recv(new_socket, buffer, 255, 0) <= 0) break; // Sai se o cliente der close()\n")

		g.buf.WriteString("            printf(\"  [SERVIDOR] Mensagem recebida: %s\\n\", buffer);\n")
		g.buf.WriteString("            char op[16]; double v1, v2, res = 0;\n")

		g.buf.WriteString("            if (sscanf(buffer, \"%s %lf %lf\", op, &v1, &v2) == 3) {\n")
		g.buf.WriteString("                if (strcmp(op, \"+\") == 0) res = v1 + v2;\n")
		g.buf.WriteString("                else if (strcmp(op, \"-\") == 0) res = v1 - v2;\n")
		g.buf.WriteString("                else if (strcmp(op, \"*\") == 0) res = v1 * v2;\n")
		g.buf.WriteString("                else if (strcmp(op, \"/\") == 0) res = v1 / v2;\n")

		g.buf.WriteString("                snprintf(buffer, sizeof(buffer), \"%.2f\", res);\n")
		g.buf.WriteString("                send(new_socket, buffer, strlen(buffer), 0); // Envia resposta de volta!\n")
		g.buf.WriteString("                printf(\"  [SERVIDOR] Resposta enviada: %s\\n\", buffer);\n")
		g.buf.WriteString("            }\n")
		g.buf.WriteString("        }\n")
		g.buf.WriteString("        close(new_socket);\n")
		g.buf.WriteString("        printf(\"  [SERVIDOR] *** Cliente Desconectado ***\\n\");\n")
		g.buf.WriteString("    }\n")
		g.buf.WriteString("}\n\n")
	}
}

func (g *CGenerator) Generate() string {
	usage := g.scanBuiltins()

	if usage.openmp {
		g.buf.WriteString("#include <omp.h>\n")
	}
	if usage.network {
		g.buf.WriteString("#include <sys/socket.h>\n")
		g.buf.WriteString("#include <arpa/inet.h>\n")
		g.buf.WriteString("#include <unistd.h>\n")
	}

	g.buf.WriteString("#include <stdint.h>\n")
	g.buf.WriteString("#include <stdbool.h>\n")
	g.buf.WriteString("#include <stdio.h>\n")
	g.buf.WriteString("#include <stdlib.h>\n")
	g.buf.WriteString("#include <string.h>\n")
	if usage.ctype {
		g.buf.WriteString("#include <ctype.h>\n")
	}
	g.buf.WriteByte('\n')
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
			g.localVars = make(map[string]bool)
			var params []string
			for i+1 < len(g.instrs) && g.instrs[i+1].Op == "PARAM_DECL" {
				i++
				p := g.instrs[i]
				params = append(params, fmt.Sprintf("%s %s", mapType(p.Arg2), p.Arg1))
				g.localVars[p.Arg1] = true
			}
			g.buf.WriteString(fmt.Sprintf("%s %s(%s) {\n", retC, instr.Arg1, strings.Join(params, ", ")))
			g.insideFunc = true

		case "PARAM_DECL":
		case "PARAM":
			g.paramBuf = append(g.paramBuf, instr.Arg1)

		case "CALL":
			args := strings.Join(g.paramBuf, ", ")
			g.paramBuf = g.paramBuf[:0]
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

		case "AND":
			g.emitBinary(instr.Result, instr.Arg1, "&&", instr.Arg2)
		case "OR":
			g.emitBinary(instr.Result, instr.Arg1, "||", instr.Arg2)
		case "NOT":
			g.emitUnary(instr.Result, "!", instr.Arg1)

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

		case "ARRAY_NEW":
			arrTemp := instr.Result
			count := instr.Arg1
			arrType := g.resolveType(arrTemp)
			elemType := arrayElemType(arrType)
			if elemType == "" {
				elemType = instr.Arg2
			}
			elemC := mapType(elemType)
			typedefName := mapType(arrType)

			var values []string
			for i+1 < len(g.instrs) && g.instrs[i+1].Op == "ARRAY_SET" && g.instrs[i+1].Arg1 == arrTemp {
				i++
				values = append(values, g.instrs[i].Result)
			}

			dataVar := "_" + arrTemp + "_data"
			g.buf.WriteString(fmt.Sprintf("%s%s %s[] = {%s};\n", g.ind(), elemC, dataVar, strings.Join(values, ", ")))
			g.buf.WriteString(fmt.Sprintf("%s%s %s = { %s, %s };\n", g.ind(), typedefName, arrTemp, dataVar, count))
			g.markDeclared(arrTemp)

		case "ARRAY_SET":
			g.buf.WriteString(fmt.Sprintf("%s%s.data[%s] = %s;\n", g.ind(), instr.Arg1, instr.Arg2, instr.Result))

		case "ARRAY_GET":
			prefix := g.declPrefix(instr.Result)
			g.buf.WriteString(fmt.Sprintf("%s%s%s = %s.data[%s];\n", g.ind(), prefix, instr.Result, instr.Arg1, instr.Arg2))

		case "ARRAY_LEN":
			prefix := g.declPrefix(instr.Result)
			g.buf.WriteString(fmt.Sprintf("%s%s%s = %s.len;\n", g.ind(), prefix, instr.Result, instr.Arg1))

		case "TUPLE_NEW":
			tupTemp := instr.Result
			tupType := g.resolveType(tupTemp)
			typedefName := mapType(tupType)

			var values []string
			for i+1 < len(g.instrs) && g.instrs[i+1].Op == "TUPLE_SET" && g.instrs[i+1].Arg1 == tupTemp {
				i++
				values = append(values, g.instrs[i].Result)
			}

			g.buf.WriteString(fmt.Sprintf("%s%s %s = { %s };\n", g.ind(), typedefName, tupTemp, strings.Join(values, ", ")))
			g.markDeclared(tupTemp)

		case "TUPLE_SET":
			g.buf.WriteString(fmt.Sprintf("%s%s._%s = %s;\n", g.ind(), instr.Arg1, instr.Arg2, instr.Result))

		case "INPUT":
			if instr.Arg1 != "" {
				g.buf.WriteString(fmt.Sprintf("%sprintf(\"%%s\", %s);\n", g.ind(), instr.Arg1))
			}
			g.tempTypes[instr.Result] = "string"
			prefix := g.declPrefix(instr.Result)
			g.buf.WriteString(fmt.Sprintf("%s%s%s = mp_input();\n", g.ind(), prefix, instr.Result))

		case "PRINT":
			g.buf.WriteString(fmt.Sprintf("%sprintf(\"%s\\n\", %s);\n", g.ind(), g.printFmt(instr.Arg1), instr.Arg1))

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
				g.buf.WriteString("    return 0;\n")
			}
			g.buf.WriteString("}\n")
			g.insideFunc = false

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

		case "BEGIN_SEQ":
			g.buf.WriteString(fmt.Sprintf("%s/* Bloco Sequencial */\n", g.ind()))
		case "END_SEQ":
			g.buf.WriteString("\n")
		case "BEGIN_PAR":
			g.buf.WriteString(fmt.Sprintf("%s#pragma omp parallel sections\n%s{\n", g.ind(), g.ind()))
		case "END_PAR":
			g.buf.WriteString(fmt.Sprintf("%s}\n", g.ind()))
		case "PAR_TASK_START":
			g.buf.WriteString(fmt.Sprintf("%s    #pragma omp section\n%s    {\n", g.ind(), g.ind()))
		case "PAR_TASK_END":
			g.buf.WriteString(fmt.Sprintf("%s    }\n", g.ind()))

		case "CHAN_DECL":
			if instr.Arg1 == "c_channel" {
				host := g.paramBuf[0]
				port := g.paramBuf[1]
				if g.insideFunc {
					g.buf.WriteString(fmt.Sprintf("%sint %s = create_client(%s, %s);\n", g.ind(), instr.Arg2, host, port))
				} else {
					g.buf.WriteString(fmt.Sprintf("int %s;\n__attribute__((constructor)) void __init_%s() {\n    %s = create_client(%s, %s);\n}\n", instr.Arg2, instr.Arg2, instr.Arg2, host, port))
				}
			} else if instr.Arg1 == "s_channel" {
				cb := g.paramBuf[0]
				desc := g.paramBuf[1]
				host := g.paramBuf[2]
				port := g.paramBuf[3]
				if g.insideFunc {
					g.buf.WriteString(fmt.Sprintf("%sstart_server(%s, %s, %s, %s);\n", g.ind(), cb, desc, host, port))
				} else {
					g.buf.WriteString(fmt.Sprintf("__attribute__((constructor)) void __init_%s() {\n    start_server(%s, %s, %s, %s);\n}\n", instr.Arg2, cb, desc, host, port))
				}
			}
			g.paramBuf = g.paramBuf[:0]

		case "METHOD_CALL":
			parts := strings.Split(instr.Arg1, ".")
			obj := parts[0]
			method := parts[1]
			args := strings.Join(g.paramBuf, ", ")
			g.paramBuf = g.paramBuf[:0]

			if method == "send" {
				g.buf.WriteString(fmt.Sprintf("%schannel_send(%s, %s);\n", g.ind(), obj, args))
			} else if method == "close" {
				g.buf.WriteString(fmt.Sprintf("%schannel_close(%s);\n", g.ind(), obj))
			} else {
				retC := g.resolveType(instr.Result)
				if retC == "void" || retC == "" {
					g.buf.WriteString(fmt.Sprintf("%s%s_%s(&%s, %s);\n", g.ind(), g.resolveType(obj), method, obj, args))
				} else {
					prefix := g.declPrefix(instr.Result)
					g.buf.WriteString(fmt.Sprintf("%s%s%s = %s_%s(&%s, %s);\n", g.ind(), prefix, instr.Result, g.resolveType(obj), method, obj, args))
				}
			}
		}
	}

	return g.buf.String()
}
