package cgen

import (
	"fmt"
	"strings"

	"minipar/internal/symtab"
	"minipar/semantic"
	"minipar/tac"
)

type CGenerator struct {
	instrs         []tac.Instruction
	tempTypes      map[string]string
	varTypes       map[string]string
	globalVars     map[string]bool
	localVars      map[string]bool
	paramBuf       []string
	currentClass   string
	insideFunc     bool
	currentFunc    string                     
	funcLocalTypes map[string]map[string]string 
	buf            strings.Builder
}

func New(instrs []tac.Instruction, tempTypes map[string]string, globalScope *symtab.Scope[*semantic.Type]) *CGenerator {
	g := &CGenerator{
		instrs:         instrs,
		tempTypes:      tempTypes,
		varTypes:       make(map[string]string),
		globalVars:     make(map[string]bool),
		localVars:      make(map[string]bool),
		funcLocalTypes: make(map[string]map[string]string), // Tem que inicializar aqui
	}
	g.collectVarTypes(globalScope)
	g.inferLocalTypes()
	return g
}

func (g *CGenerator) inferLocalTypes() {
	var curr string
	for _, instr := range g.instrs {
		switch instr.Op {
		case "BEGIN_FUNC", "BEGIN_METHOD":
			curr = instr.Arg1
			g.funcLocalTypes[curr] = make(map[string]string)
		case "PARAM_DECL":
			if curr != "" {
				g.funcLocalTypes[curr][instr.Arg1] = instr.Arg2
			}
		case "ASSIGN":
			if curr != "" {
				typ := ""
				if t, ok := g.tempTypes[instr.Arg1]; ok {
					typ = t
				} else if strings.HasPrefix(instr.Arg1, "\"") {
					typ = "string"
				} else if strings.Contains(instr.Arg1, ".") {
					typ = "f64"
				} else if instr.Arg1 == "true" || instr.Arg1 == "false" {
					typ = "bool"
				} else if len(instr.Arg1) > 0 && (instr.Arg1[0] >= '0' && instr.Arg1[0] <= '9' || instr.Arg1[0] == '-') {
					typ = "i32"
				} else if t, ok := g.funcLocalTypes[curr][instr.Arg1]; ok {
					typ = t
				} else if t, ok := g.varTypes[instr.Arg1]; ok {
					typ = t
				}
				// Guarda apenas a primeira atribuição (declaração)
				if _, exists := g.funcLocalTypes[curr][instr.Result]; !exists && typ != "" {
					g.funcLocalTypes[curr][instr.Result] = typ
				}
			}
		case "ARRAY_NEW", "TUPLE_NEW":
			if curr != "" {
				if t, ok := g.tempTypes[instr.Result]; ok {
					g.funcLocalTypes[curr][instr.Result] = t
				}
			}
		}
	}
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
	// NOVO: Olha no escopo local da função antes do global
	if g.insideFunc && g.currentFunc != "" {
		if t, ok := g.funcLocalTypes[g.currentFunc][name]; ok {
			return t
		}
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
	t := mapType(g.resolveType(name))
	if t == "void" {
		t = "int" // NOVO: Fallback seguro para evitar "variable declared void"
	}
	return t + " "
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

func mapTypeNameForC(t string) string {
	t = strings.ReplaceAll(t, "[", "arr_")
	t = strings.ReplaceAll(t, "]", "")
	t = strings.ReplaceAll(t, "(", "tup_")
	t = strings.ReplaceAll(t, ")", "")
	t = strings.ReplaceAll(t, " ", "")
	t = strings.ReplaceAll(t, ",", "_")
	return t
}

func mapType(miniparType string) string {
	if len(miniparType) >= 2 && miniparType[0] == '[' && miniparType[len(miniparType)-1] == ']' {
		return mapTypeNameForC(miniparType)
	}
	if len(miniparType) >= 2 && miniparType[0] == '(' && miniparType[len(miniparType)-1] == ')' {
		return mapTypeNameForC(miniparType)
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
	case "any":
		return "void*"
	default:
		if miniparType != "" {
			return miniparType
		}
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

func (g *CGenerator) emitStructDef(t string, seen map[string]bool) {
	if seen[t] {
		return
	}
	seen[t] = true

	if len(t) >= 2 && t[0] == '[' && t[len(t)-1] == ']' {
		elem := arrayElemType(t)
		if len(elem) >= 2 && (elem[0] == '[' || elem[0] == '(') {
			g.emitStructDef(elem, seen)
		}

		elemC := mapType(elem)
		typedefName := mapType(t)
		g.buf.WriteString(fmt.Sprintf("typedef struct { %s* data; int len; int cap; } %s_struct;\n", elemC, typedefName))
		g.buf.WriteString(fmt.Sprintf("typedef %s_struct* %s;\n", typedefName, typedefName))
		g.buf.WriteString(fmt.Sprintf("static void %s_append(%s arr, %s val) {\n", typedefName, typedefName, elemC))
		g.buf.WriteString("    if (arr->len >= arr->cap) {\n")
		g.buf.WriteString("        arr->cap = arr->cap == 0 ? 4 : arr->cap * 2;\n")
		g.buf.WriteString(fmt.Sprintf("        arr->data = realloc(arr->data, arr->cap * sizeof(%s));\n", elemC))
		g.buf.WriteString("    }\n")
		g.buf.WriteString("    arr->data[arr->len++] = val;\n")
		g.buf.WriteString("}\n")
	} else if len(t) >= 2 && t[0] == '(' && t[len(t)-1] == ')' {
		elems := tupleElemTypes(t)
		for _, e := range elems {
			e = strings.TrimSpace(e)
			if len(e) >= 2 && (e[0] == '[' || e[0] == '(') {
				g.emitStructDef(e, seen)
			}
		}
		typedefName := mapType(t)
		g.buf.WriteString("typedef struct {\n")
		for i, e := range elems {
			g.buf.WriteString(fmt.Sprintf("    %s _%d;\n", mapType(strings.TrimSpace(e)), i))
		}
		g.buf.WriteString(fmt.Sprintf("} %s;\n", typedefName))
	}
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
		g.emitStructDef(t, seen)
	}
}

type builtinUsage struct {
	toString  bool
	mpInput   bool
	ctype     bool
	openmp    bool
	network   bool
	strConcat bool
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
		case "ADD": // <- Agora está dentro do switch corretamente!
			if g.resolveType(instr.Result) == "string" {
				u.strConcat = true
			}
		}
	}
	return u
}

func (g *CGenerator) getFuncSig(funcName string) ([]string, string) {
	var params []string
	retType := "void"
	found := false
	for _, instr := range g.instrs {
		if instr.Op == "BEGIN_FUNC" && instr.Arg1 == funcName {
			retType = instr.Arg2
			found = true
			continue
		}
		if found {
			if instr.Op == "PARAM_DECL" {
				params = append(params, instr.Arg2)
			} else if instr.Op != "LABEL" {
				break
			}
		}
	}
	return params, retType
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
	if u.strConcat {
		g.buf.WriteString("static char* concat_string(const char* s1, const char* s2) {\n")
		g.buf.WriteString("    if (!s1) s1 = \"\";\n")
		g.buf.WriteString("    if (!s2) s2 = \"\";\n")
		g.buf.WriteString("    char* res = malloc(strlen(s1) + strlen(s2) + 1);\n")
		g.buf.WriteString("    strcpy(res, s1);\n    strcat(res, s2);\n    return res;\n")
		g.buf.WriteString("}\n")
	}
	if u.network {
		g.buf.WriteString("static int create_client(const char* ip, int port) {\n")
		g.buf.WriteString("    int sock = 0;\n")
		g.buf.WriteString("    if ((sock = socket(AF_INET, SOCK_STREAM, 0)) < 0) return -1;\n")
		g.buf.WriteString("    struct sockaddr_in serv_addr;\n")
		g.buf.WriteString("    serv_addr.sin_family = AF_INET; serv_addr.sin_port = htons(port);\n")
		g.buf.WriteString("    if (strcmp(ip, \"localhost\") == 0) ip = \"127.0.0.1\";\n")
		g.buf.WriteString("    inet_pton(AF_INET, ip, &serv_addr.sin_addr);\n")
		g.buf.WriteString("    if (connect(sock, (struct sockaddr *)&serv_addr, sizeof(serv_addr)) < 0) return -1;\n")
		g.buf.WriteString("    return sock;\n}\n\n")

		g.buf.WriteString("static void channel_send_str(int sock, const char* str) {\n")
		g.buf.WriteString("    if (sock < 0) return;\n")
		g.buf.WriteString("    send(sock, str, strlen(str), 0);\n")
		g.buf.WriteString("    char buffer[1024] = {0};\n")
		g.buf.WriteString("    int valread = recv(sock, buffer, 1023, 0);\n")
		g.buf.WriteString("    if (valread > 0) printf(\"  [CLIENTE] Recebeu a resposta TCP: %s\\n\", buffer);\n}\n\n")

		g.buf.WriteString("static void channel_close(int sock) {\n    if (sock >= 0) close(sock);\n}\n\n")

		for i, instr := range g.instrs {
			if instr.Op == "CHAN_DECL" && instr.Arg1 == "s_channel" {
				var pbuf []string
				for j := i - 1; j >= 0; j-- {
					if g.instrs[j].Op == "PARAM" {
						pbuf = append([]string{g.instrs[j].Arg1}, pbuf...)
					} else {
						break
					}
				}
				if len(pbuf) >= 4 {
					cb := pbuf[0]
					params, retType := g.getFuncSig(cb)

					// ===== CORREÇÃO: FORWARD DECLARATION INJETADA AQUI =====
					cRet := mapType(retType)
					if cRet == "" {
						cRet = "void"
					}
					var pTypes []string
					for _, p := range params {
						pTypes = append(pTypes, mapType(p))
					}
					sigParams := strings.Join(pTypes, ", ")
					if sigParams == "" {
						sigParams = "void"
					}
					g.buf.WriteString(fmt.Sprintf("%s %s(%s);\n", cRet, cb, sigParams))
					// ========================================================

					g.buf.WriteString(fmt.Sprintf("static void start_server_%s(const char* desc, const char* ip, int port) {\n", cb))
					g.buf.WriteString("    int server_fd, new_socket; struct sockaddr_in address;\n")
					g.buf.WriteString("    int opt = 1; int addrlen = sizeof(address);\n")
					g.buf.WriteString("    if ((server_fd = socket(AF_INET, SOCK_STREAM, 0)) == 0) exit(1);\n")
					g.buf.WriteString("    setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt));\n")
					g.buf.WriteString("    address.sin_family = AF_INET; address.sin_addr.s_addr = INADDR_ANY; address.sin_port = htons(port);\n")
					g.buf.WriteString("    bind(server_fd, (struct sockaddr *)&address, sizeof(address));\n")
					g.buf.WriteString("    listen(server_fd, 3);\n")
					g.buf.WriteString("    printf(\"Server [%s] started on %s:%d\\n\", desc, ip, port);\n")
					g.buf.WriteString("    while(1) {\n")
					g.buf.WriteString("        if ((new_socket = accept(server_fd, (struct sockaddr *)&address, (socklen_t*)&addrlen)) < 0) continue;\n")
					g.buf.WriteString("        char buffer[1024];\n")
					g.buf.WriteString("        while(1) {\n")
					g.buf.WriteString("            memset(buffer, 0, sizeof(buffer));\n")
					g.buf.WriteString("            if (recv(new_socket, buffer, 1023, 0) <= 0) break;\n")

					var formatParts, varDecls, varNames []string
					scanRefs := ""
					for pidx, pType := range params {
						vName := fmt.Sprintf("v%d", pidx)
						varNames = append(varNames, vName) // Devolvemos o varNames aqui!
						if pType == "f64" || pType == "f32" {
							formatParts = append(formatParts, "%lf")
							varDecls = append(varDecls, fmt.Sprintf("double %s;", vName))
							scanRefs += ", &" + vName
						} else if pType == "string" {
							formatParts = append(formatParts, "%s")
							varDecls = append(varDecls, fmt.Sprintf("char %s[256];", vName))
							scanRefs += ", " + vName // string não precisa de '&'
						} else {
							formatParts = append(formatParts, "%d")
							varDecls = append(varDecls, fmt.Sprintf("int %s;", vName))
							scanRefs += ", &" + vName
						}
					}
					for _, decl := range varDecls {
						g.buf.WriteString(fmt.Sprintf("            %s\n", decl))
					}
					scanFmt := strings.Join(formatParts, " ")

					g.buf.WriteString(fmt.Sprintf("            if (sscanf(buffer, \"%s\"%s) == %d) {\n", scanFmt, scanRefs, len(params)))
					callArgs := strings.Join(varNames, ", ")

					if retType != "void" && retType != "" {
						cRet := mapType(retType)
						g.buf.WriteString(fmt.Sprintf("                %s res = %s(%s);\n", cRet, cb, callArgs))
						if retType == "f64" || retType == "f32" {
							g.buf.WriteString("                snprintf(buffer, sizeof(buffer), \"%f\", res);\n")
						} else if retType == "string" { // CORREÇÃO AQUI
							g.buf.WriteString("                snprintf(buffer, sizeof(buffer), \"%s\", res);\n")
						} else {
							g.buf.WriteString("                snprintf(buffer, sizeof(buffer), \"%d\", res);\n")
						}
					} else {
						g.buf.WriteString(fmt.Sprintf("                %s(%s);\n", cb, callArgs))
						g.buf.WriteString("                snprintf(buffer, sizeof(buffer), \"OK\");\n")
					}
					g.buf.WriteString("                send(new_socket, buffer, strlen(buffer), 0);\n")
					g.buf.WriteString("            }\n        }\n")
					g.buf.WriteString("        close(new_socket);\n    }\n}\n")
				}
			}
		}
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
	g.buf.WriteString("#include <math.h>\n")
	if usage.ctype {
		g.buf.WriteString("#include <ctype.h>\n")
	}
	g.buf.WriteByte('\n')
	g.emitCompositeTypedefs()
	g.emitBuiltinHelpers(usage)
	g.buf.WriteByte('\n')
	// =========================================================================
	// Forward Declarations e Definições completas para Classes
	isClass := make(map[string]bool)
	classFields := make(map[string][]string)          // NOVO: Guarda os campos para os macros!
	classFieldDefaults := make(map[string]map[string]string) // NOVO: campo -> valor default literal

	for i := 0; i < len(g.instrs); i++ {
		if g.instrs[i].Op == "BEGIN_CLASS" {
			cls := g.instrs[i].Arg1
			isClass[cls] = true
			g.buf.WriteString(fmt.Sprintf("typedef struct %s_struct %s;\n", cls, cls))
			g.buf.WriteString(fmt.Sprintf("struct %s_struct {\n", cls))
			fieldCount := 0
			for j := i + 1; j < len(g.instrs); j++ {
				if g.instrs[j].Op == "END_CLASS" {
					break
				}
				if g.instrs[j].Op == "FIELD" {
					g.buf.WriteString(fmt.Sprintf("    %s %s;\n", mapType(g.instrs[j].Arg2), g.instrs[j].Arg1))
					classFields[cls] = append(classFields[cls], g.instrs[j].Arg1) // GUARDA O CAMPO!
					if def := g.instrs[j].Result; def != "" {                     // GUARDA O VALOR DEFAULT!
						if classFieldDefaults[cls] == nil {
							classFieldDefaults[cls] = make(map[string]string)
						}
						classFieldDefaults[cls][g.instrs[j].Arg1] = def
					}
					fieldCount++
				}
			}
			if fieldCount == 0 {
				g.buf.WriteString("    int _dummy;\n")
			}
			g.buf.WriteString("};\n\n")
		}
	}
	// Forward Declarations para Funções e Métodos
	clsName := ""
	for i := 0; i < len(g.instrs); i++ {
		if g.instrs[i].Op == "BEGIN_CLASS" {
			clsName = g.instrs[i].Arg1
		} else if g.instrs[i].Op == "END_CLASS" {
			clsName = ""
		} else if g.instrs[i].Op == "BEGIN_FUNC" {
			funcName := g.instrs[i].Arg1
			retC := mapType(g.instrs[i].Arg2)
			if funcName == "main" {
				retC = "int"
			}
			var pTypes []string
			for j := i + 1; j < len(g.instrs) && g.instrs[j].Op == "PARAM_DECL"; j++ {
				pTypes = append(pTypes, mapType(g.instrs[j].Arg2))
			}
			sigParams := strings.Join(pTypes, ", ")
			if sigParams == "" {
				sigParams = "void"
			}
			g.buf.WriteString(fmt.Sprintf("%s %s(%s);\n", retC, funcName, sigParams))
		} else if g.instrs[i].Op == "BEGIN_METHOD" {
			methodName := g.instrs[i].Arg1
			retC := mapType(g.instrs[i].Arg2)
			pTypes := []string{clsName + "*"} // self implícito
			for j := i + 1; j < len(g.instrs) && g.instrs[j].Op == "PARAM_DECL"; j++ {
				if g.instrs[j].Arg1 != "self" && g.instrs[j].Arg1 != "Self" {
					pTypes = append(pTypes, mapType(g.instrs[j].Arg2))
				}
			}
			g.buf.WriteString(fmt.Sprintf("%s %s_%s(%s);\n", retC, clsName, methodName, strings.Join(pTypes, ", ")))
		}
	}
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
			g.currentFunc = instr.Arg1 // NOVO: Avisa o gerador que entramos nesta função

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
			arg0 := ""
			if len(g.paramBuf) > 0 {
				arg0 = g.paramBuf[0]
			}
			args := strings.Join(g.paramBuf, ", ")
			g.paramBuf = g.paramBuf[:0]
			callName := instr.Arg1
			if cName, ok := builtinCName[callName]; ok {
				callName = cName
			}
			retType := g.resolveType(instr.Result)

			// 1. Instanciação de Objetos (Substitui func() por {0})
			if isClass[callName] {
				if g.insideFunc && g.currentFunc != "" {
					if g.funcLocalTypes[g.currentFunc] == nil {
						g.funcLocalTypes[g.currentFunc] = make(map[string]string)
					}
					g.funcLocalTypes[g.currentFunc][instr.Result] = callName // Registra a Classe localmente!
				}
				prefix := g.declPrefix(instr.Result)
				// Aplica os valores default dos campos da classe (designated initializers)
				// em vez de zerar tudo, preservando os defaults declarados na classe.
				init := "{0}"
				if defaults := classFieldDefaults[callName]; len(defaults) > 0 {
					var parts []string
					for _, f := range classFields[callName] {
						if v, ok := defaults[f]; ok {
							parts = append(parts, fmt.Sprintf(".%s = %s", f, v))
						}
					}
					if len(parts) > 0 {
						init = "{ " + strings.Join(parts, ", ") + " }"
					}
				}
				g.buf.WriteString(fmt.Sprintf("%s%s%s = (%s)%s;\n", g.ind(), prefix, instr.Result, callName, init))
				continue
			}

			// 2. Leitura nativa de tamanho de Arrays (Substitui strlen)
			if callName == "strlen" && arg0 != "" {
				t := g.resolveType(arg0)
				if strings.HasPrefix(t, "[") {
					prefix := g.declPrefix(instr.Result)
					g.buf.WriteString(fmt.Sprintf("%s%s%s = %s->len;\n", g.ind(), prefix, instr.Result, arg0))
					continue
				}
			}

			// 3. Chamada normal de funções
			if retType == "void" || retType == "" {
				g.buf.WriteString(fmt.Sprintf("%s%s(%s);\n", g.ind(), callName, args))
			} else {
				prefix := g.declPrefix(instr.Result)
				g.buf.WriteString(fmt.Sprintf("%s%s%s = %s(%s);\n", g.ind(), prefix, instr.Result, callName, args))
			}

		case "ASSIGN":
			if !g.isDeclared(instr.Result) && strings.HasPrefix(instr.Arg1, "lambda_") {
				g.markDeclared(instr.Result)
				g.buf.WriteString(fmt.Sprintf("%s__typeof__(%s)* %s = %s;\n", g.ind(), instr.Arg1, instr.Result, instr.Arg1))
			} else {
				prefix := g.declPrefix(instr.Result)
				retC := mapType(g.resolveType(instr.Result))
				rhsC := mapType(g.resolveType(instr.Arg1))
				
				cast := ""
				// Se a atribuição envolver qualquer array, convertemos para void* // Isso silencia 100% dos warnings de ponteiros incompativeis do GCC!
				if strings.HasPrefix(retC, "arr_") || strings.HasPrefix(rhsC, "arr_") {
					cast = "(void*)" 
				}
				
				g.buf.WriteString(fmt.Sprintf("%s%s%s = %s%s;\n", g.ind(), prefix, instr.Result, cast, instr.Arg1))
			}
		case "ADD":
			if g.resolveType(instr.Result) == "string" {
				prefix := g.declPrefix(instr.Result)
				g.buf.WriteString(fmt.Sprintf("%s%s%s = concat_string(%s, %s);\n", g.ind(), prefix, instr.Result, instr.Arg1, instr.Arg2))
			} else {
				g.emitBinary(instr.Result, instr.Arg1, "+", instr.Arg2)
			}
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

		case "EQ", "NEQ":
			isStr := strings.HasPrefix(instr.Arg1, "\"") || strings.HasPrefix(instr.Arg2, "\"") || 
					 g.resolveType(instr.Arg1) == "string" || g.resolveType(instr.Arg2) == "string"
			
			opC := "=="
			if instr.Op == "NEQ" { 
				opC = "!=" 
			}

			if isStr {
				prefix := g.declPrefix(instr.Result)
				g.buf.WriteString(fmt.Sprintf("%s%s%s = (strcmp(%s, %s) %s 0);\n", g.ind(), prefix, instr.Result, instr.Arg1, instr.Arg2, opC))
			} else {
				g.emitBinary(instr.Result, instr.Arg1, opC, instr.Arg2)
			}
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

			// NOVO: Malloc para criar a struct em tempo de execução
			g.buf.WriteString(fmt.Sprintf("%s%s %s = malloc(sizeof(%s_struct));\n", g.ind(), typedefName, arrTemp, typedefName))
			if len(values) > 0 {
				dataVar := "_" + arrTemp + "_data"
				g.buf.WriteString(fmt.Sprintf("%s%s* %s = malloc(%s * sizeof(%s));\n", g.ind(), elemC, dataVar, count, elemC))
				for idx, v := range values {
					g.buf.WriteString(fmt.Sprintf("%s%s[%d] = %s;\n", g.ind(), dataVar, idx, v))
				}
				g.buf.WriteString(fmt.Sprintf("%s%s->data = %s;\n", g.ind(), arrTemp, dataVar))
				g.buf.WriteString(fmt.Sprintf("%s%s->len = %s;\n", g.ind(), arrTemp, count))
				g.buf.WriteString(fmt.Sprintf("%s%s->cap = %s;\n", g.ind(), arrTemp, count))
			} else {
				g.buf.WriteString(fmt.Sprintf("%s%s->data = NULL;\n", g.ind(), arrTemp))
				g.buf.WriteString(fmt.Sprintf("%s%s->len = 0;\n", g.ind(), arrTemp))
				g.buf.WriteString(fmt.Sprintf("%s%s->cap = 0;\n", g.ind(), arrTemp))
			}
			g.markDeclared(arrTemp)

		case "ARRAY_SET":
			g.buf.WriteString(fmt.Sprintf("%s%s->data[%s] = %s;\n", g.ind(), instr.Arg1, instr.Arg2, instr.Result))

		case "ARRAY_GET":
			prefix := g.declPrefix(instr.Result)
			g.buf.WriteString(fmt.Sprintf("%s%s%s = %s->data[%s];\n", g.ind(), prefix, instr.Result, instr.Arg1, instr.Arg2))

		case "ARRAY_LEN":
			prefix := g.declPrefix(instr.Result)
			g.buf.WriteString(fmt.Sprintf("%s%s%s = %s->len;\n", g.ind(), prefix, instr.Result, instr.Arg1))
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
		case "FIELD":
		case "END_CLASS":
			g.currentClass = ""
		case "BEGIN_CTOR":
			g.buf.WriteString(fmt.Sprintf("void %s_init(%s* self) {\n", instr.Arg1, instr.Arg1))
		case "END_CTOR":
			g.buf.WriteString("}\n")
		case "BEGIN_METHOD":
			retC := mapType(instr.Arg2)
			g.localVars = make(map[string]bool)
			g.localVars["self"] = true
			g.currentFunc = instr.Arg1
			var params []string
			for i+1 < len(g.instrs) && g.instrs[i+1].Op == "PARAM_DECL" {
				i++
				p := g.instrs[i]
				g.localVars[p.Arg1] = true
				if p.Arg1 == "self" || p.Arg1 == "Self" {
					continue
				}
				params = append(params, fmt.Sprintf("%s %s", mapType(p.Arg2), p.Arg1))
			}
			paramList := strings.Join(params, ", ")
			if paramList != "" {
				paramList = ", " + paramList
			}
			g.buf.WriteString(fmt.Sprintf("%s %s_%s(%s* self%s) {\n", retC, g.currentClass, instr.Arg1, g.currentClass, paramList))
			g.insideFunc = true

			// NOVO: Magia dos Macros do C! Transforma variavel isolada em self->variavel
			for _, f := range classFields[g.currentClass] {
				g.localVars[f] = true 
				g.buf.WriteString(fmt.Sprintf("#define %s (self->%s)\n", f, f))
			}

		case "END_METHOD":
			// NOVO: Limpa a magia ao sair do método!
			for _, f := range classFields[g.currentClass] {
				g.buf.WriteString(fmt.Sprintf("#undef %s\n", f))
			}
			g.buf.WriteString("}\n")
			g.insideFunc = false
			g.currentFunc = ""
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
				// NOVO: Chama a start_server dinamicamente baseada no nome do callback
				if g.insideFunc {
					g.buf.WriteString(fmt.Sprintf("%sstart_server_%s(%s, %s, %s);\n", g.ind(), cb, desc, host, port))
				} else {
					g.buf.WriteString(fmt.Sprintf("__attribute__((constructor)) void __init_%s() {\n    start_server_%s(%s, %s, %s);\n}\n", instr.Arg2, cb, desc, host, port))
				}
			}
			g.paramBuf = g.paramBuf[:0]

		case "METHOD_CALL":
			parts := strings.Split(instr.Arg1, ".")
			obj := parts[0]
			method := parts[1]
			args := strings.Join(g.paramBuf, ", ")

			if method == "send" {
				g.buf.WriteString(fmt.Sprintf("%s{\n", g.ind()))
				g.buf.WriteString(fmt.Sprintf("%s    char send_buf[1024] = {0};\n", g.ind()))
				if len(g.paramBuf) > 0 {
					var formats []string
					for _, a := range g.paramBuf {
						t := g.resolveType(a)
						if t == "f64" || t == "f32" || strings.Contains(a, ".") {
							formats = append(formats, "%f")
						} else if t == "string" || strings.HasPrefix(a, "\"") {
							formats = append(formats, "%s")
						} else {
							formats = append(formats, "%d")
						}
					}
					fmtStr := strings.Join(formats, " ")
					g.buf.WriteString(fmt.Sprintf("%s    snprintf(send_buf, sizeof(send_buf), \"%s\", %s);\n", g.ind(), fmtStr, strings.Join(g.paramBuf, ", ")))
				}
				g.buf.WriteString(fmt.Sprintf("%s    channel_send_str(%s, send_buf);\n", g.ind(), obj))
				g.buf.WriteString(fmt.Sprintf("%s}\n", g.ind()))
			} else if method == "close" {
				g.buf.WriteString(fmt.Sprintf("%schannel_close(%s);\n", g.ind(), obj))
			} else {
				retC := g.resolveType(instr.Result)
				objCName := mapType(g.resolveType(obj))

				refOp := "&"
				if strings.HasPrefix(objCName, "arr_") {
					refOp = ""
				}

				cObj := obj
				// Injeta o ponteiro 'self' do C quando chamamos Self na linguagem
				if obj == "Self" || obj == "self" {
					objCName = g.currentClass
					cObj = "self"
					refOp = ""
				}

				callArgs := ""
				if args != "" {
					callArgs = ", " + args
				}

				if retC == "void" || retC == "" {
					g.buf.WriteString(fmt.Sprintf("%s%s_%s(%s%s%s);\n", g.ind(), objCName, method, refOp, cObj, callArgs))
				} else {
					prefix := g.declPrefix(instr.Result)
					g.buf.WriteString(fmt.Sprintf("%s%s%s = %s_%s(%s%s%s);\n", g.ind(), prefix, instr.Result, objCName, method, refOp, cObj, callArgs))
				}
			}
			g.paramBuf = g.paramBuf[:0]
		}
	}
	return g.buf.String()
}
