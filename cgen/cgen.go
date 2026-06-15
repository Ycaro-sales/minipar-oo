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
	memChans     map[string]bool // names bound to an in-memory channel (chan<T>())
	netChans     map[string]bool // names bound to a network channel (s_/c_channel)
	buf          strings.Builder
}

func New(instrs []tac.Instruction, tempTypes map[string]string, globalScope *symtab.Scope[*semantic.Type]) *CGenerator {
	g := &CGenerator{
		instrs:     instrs,
		tempTypes:  tempTypes,
		varTypes:   make(map[string]string),
		globalVars: make(map[string]bool),
		localVars:  make(map[string]bool),
		memChans:   make(map[string]bool),
		netChans:   make(map[string]bool),
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

// isStringOperand reports whether a TAC operand is a string: a string literal
// (starts with a quote) or a value whose resolved type is string.
func (g *CGenerator) isStringOperand(name string) bool {
	if len(name) > 0 && name[0] == '"' {
		return true
	}
	return g.resolveType(name) == "string"
}

// emitStringCompare emits an equality/inequality on strings via strcmp.
func (g *CGenerator) emitStringCompare(result, arg1, arg2, op string) {
	prefix := g.declPrefix(result)
	g.buf.WriteString(fmt.Sprintf("%s%s%s = strcmp(%s, %s) %s 0;\n", g.ind(), prefix, result, arg1, arg2, op))
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
	if strings.HasPrefix(miniparType, "chan<") {
		return "mp_chan*"
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

// serName mangles a minipar type into a C identifier suffix for its
// mp_enc_*/mp_dec_* serializer functions (e.g. "(string, f64)" -> "tup_string_f64").
func serName(miniparType string) string {
	if strings.HasPrefix(miniparType, "(") {
		return mapType(miniparType)
	}
	return strings.ReplaceAll(miniparType, " ", "")
}

// emitSerializers writes mp_enc_<T>/mp_dec_<T> for each element type (and its
// nested field types), in dependency order, deduplicated.
func (g *CGenerator) emitSerializers(types []string) {
	seen := make(map[string]bool)
	for _, t := range types {
		g.emitSerializer(t, seen)
	}
}

func (g *CGenerator) emitSerializer(t string, seen map[string]bool) {
	if t == "" || seen[t] {
		return
	}
	seen[t] = true
	name := serName(t)
	cType := mapType(t)

	if elems := tupleElemTypes(t); elems != nil {
		// Emit field serializers first so they're declared before use.
		for _, e := range elems {
			g.emitSerializer(strings.TrimSpace(e), seen)
		}
		fmt.Fprintf(&g.buf, "static void mp_enc_%s(int s, %s *v) {\n", name, cType)
		for i, e := range elems {
			fmt.Fprintf(&g.buf, "    mp_enc_%s(s, &v->_%d);\n", serName(strings.TrimSpace(e)), i)
		}
		g.buf.WriteString("}\n")
		fmt.Fprintf(&g.buf, "static void mp_dec_%s(int s, %s *v) {\n", name, cType)
		for i, e := range elems {
			fmt.Fprintf(&g.buf, "    mp_dec_%s(s, &v->_%d);\n", serName(strings.TrimSpace(e)), i)
		}
		g.buf.WriteString("}\n\n")
		return
	}

	if t == "string" {
		fmt.Fprintf(&g.buf, "static void mp_enc_%s(int s, %s *v) {\n", name, cType)
		g.buf.WriteString("    uint32_t n = (uint32_t)strlen(*v);\n")
		g.buf.WriteString("    mp_write_all(s, &n, sizeof(n));\n")
		g.buf.WriteString("    mp_write_all(s, *v, n);\n")
		g.buf.WriteString("}\n")
		fmt.Fprintf(&g.buf, "static void mp_dec_%s(int s, %s *v) {\n", name, cType)
		g.buf.WriteString("    uint32_t n = 0; mp_read_all(s, &n, sizeof(n));\n")
		g.buf.WriteString("    char *b = malloc(n + 1); mp_read_all(s, b, n); b[n] = '\\0'; *v = b;\n")
		g.buf.WriteString("}\n\n")
		return
	}

	// Fixed-width primitive: raw bytes (assumes matching architecture on both ends).
	fmt.Fprintf(&g.buf, "static void mp_enc_%s(int s, %s *v) { mp_write_all(s, v, sizeof(%s)); }\n", name, cType, cType)
	fmt.Fprintf(&g.buf, "static void mp_dec_%s(int s, %s *v) { mp_read_all(s, v, sizeof(%s)); }\n\n", name, cType, cType)
}

// chanElem extracts the minipar element type from a "chan<T>" string.
func chanElem(chanType string) string {
	if strings.HasPrefix(chanType, "chan<") && strings.HasSuffix(chanType, ">") {
		return chanType[len("chan<") : len(chanType)-1]
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
		// Network channel element types need their struct typedef even when no
		// literal of that exact type appears in the program.
		if instr.Op == "CHAN_DECL" {
			allTypes = append(allTypes, instr.Result)
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
	ctype        bool
	openmp       bool
	network      bool
	memChan      bool
	netElemTypes []string // element types carried by network channels (for serializers)
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
		case "CHAN_DECL":
			// Network channels are always introduced by a CHAN_DECL; a bare
			// METHOD_CALL (class methods, in-memory channel ops) does not imply
			// the TCP runtime.
			u.network = true
			u.netElemTypes = append(u.netElemTypes, instr.Result) // element type
		case "CHAN_NEW":
			u.memChan = true
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
	if u.memChan || u.network {
		// ==========================================
		// CANAL UNIFICADO (estilo Go): MEM (fila limitada entre threads) e
		// NET (socket TCP). Ambos expõem send/recv/close tipados; o NET
		// (de)serializa o valor com mp_enc_*/mp_dec_* gerados por tipo.
		// ==========================================
		g.buf.WriteString(`typedef struct {
    int kind; /* 0 = memoria, 1 = rede */
    int sock; /* socket, para kind=1 */
    pthread_mutex_t mu;
    pthread_cond_t  not_empty;
    pthread_cond_t  not_full;
    size_t elem_size, cap, count, head;
    int closed;
    unsigned char *buf;
} mp_chan;

static void mp_chan_close(mp_chan *c) {
    if (c->kind == 1) { if (c->sock >= 0) close(c->sock); return; }
    pthread_mutex_lock(&c->mu);
    c->closed = 1;
    pthread_cond_broadcast(&c->not_empty);
    pthread_cond_broadcast(&c->not_full);
    pthread_mutex_unlock(&c->mu);
}

`)
	}
	if u.memChan {
		g.buf.WriteString(`static mp_chan* mp_chan_new(size_t elem_size, size_t cap) {
    if (cap == 0) cap = 1; /* canal sem buffer aproximado como capacidade 1 */
    mp_chan *c = calloc(1, sizeof(mp_chan));
    c->kind = 0; c->sock = -1;
    pthread_mutex_init(&c->mu, NULL);
    pthread_cond_init(&c->not_empty, NULL);
    pthread_cond_init(&c->not_full, NULL);
    c->elem_size = elem_size; c->cap = cap; c->count = 0; c->head = 0; c->closed = 0;
    c->buf = malloc(elem_size * cap);
    return c;
}

static void mp_chan_send(mp_chan *c, const void *v) {
    pthread_mutex_lock(&c->mu);
    while (c->count == c->cap && !c->closed) pthread_cond_wait(&c->not_full, &c->mu);
    if (c->closed) { pthread_mutex_unlock(&c->mu); return; }
    size_t tail = (c->head + c->count) % c->cap;
    memcpy(c->buf + tail * c->elem_size, v, c->elem_size);
    c->count++;
    pthread_cond_signal(&c->not_empty);
    pthread_mutex_unlock(&c->mu);
}

static int mp_chan_recv(mp_chan *c, void *out) {
    pthread_mutex_lock(&c->mu);
    while (c->count == 0 && !c->closed) pthread_cond_wait(&c->not_empty, &c->mu);
    if (c->count == 0 && c->closed) { pthread_mutex_unlock(&c->mu); return 0; }
    memcpy(out, c->buf + c->head * c->elem_size, c->elem_size);
    c->head = (c->head + 1) % c->cap;
    c->count--;
    pthread_cond_signal(&c->not_full);
    pthread_mutex_unlock(&c->mu);
    return 1;
}

`)
	}
	if u.network {
		// Setup TCP + framing helpers, then the per-type serializers.
		g.buf.WriteString(`static void mp_write_all(int s, const void *p, size_t n) {
    const unsigned char *b = p; size_t off = 0;
    while (off < n) { ssize_t w = send(s, b + off, n - off, 0); if (w <= 0) return; off += (size_t)w; }
}

static int mp_read_all(int s, void *p, size_t n) {
    unsigned char *b = p; size_t off = 0;
    while (off < n) { ssize_t r = recv(s, b + off, n - off, 0); if (r <= 0) return 0; off += (size_t)r; }
    return 1;
}

static int mp_create_client(const char *ip, int port) {
    int sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock < 0) { perror("socket"); return -1; }
    struct sockaddr_in addr; addr.sin_family = AF_INET; addr.sin_port = htons(port);
    if (strcmp(ip, "localhost") == 0) ip = "127.0.0.1";
    inet_pton(AF_INET, ip, &addr.sin_addr);
    if (connect(sock, (struct sockaddr *)&addr, sizeof(addr)) < 0) { perror("connect"); return -1; }
    return sock;
}

static int mp_accept_one(const char *ip, int port) {
    (void)ip;
    int server_fd = socket(AF_INET, SOCK_STREAM, 0);
    if (server_fd < 0) { perror("socket"); exit(EXIT_FAILURE); }
    int opt = 1; setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt));
    struct sockaddr_in addr; addr.sin_family = AF_INET; addr.sin_addr.s_addr = INADDR_ANY; addr.sin_port = htons(port);
    if (bind(server_fd, (struct sockaddr *)&addr, sizeof(addr)) < 0) { perror("bind"); exit(EXIT_FAILURE); }
    if (listen(server_fd, 1) < 0) { perror("listen"); exit(EXIT_FAILURE); }
    socklen_t len = sizeof(addr);
    int c = accept(server_fd, (struct sockaddr *)&addr, &len);
    if (c < 0) { perror("accept"); exit(EXIT_FAILURE); }
    close(server_fd);
    return c;
}

static mp_chan* mp_chan_wrap(int sock) {
    mp_chan *c = calloc(1, sizeof(mp_chan));
    c->kind = 1; c->sock = sock;
    return c;
}
static mp_chan* mp_chan_client(const char *ip, int port) { return mp_chan_wrap(mp_create_client(ip, port)); }
static mp_chan* mp_chan_server(const char *ip, int port) { return mp_chan_wrap(mp_accept_one(ip, port)); }

`)
		g.emitSerializers(u.netElemTypes)
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
	if usage.memChan || usage.network {
		// The unified mp_chan struct embeds pthread sync primitives.
		g.buf.WriteString("#include <pthread.h>\n")
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
			// An in-memory channel handle is just a pointer; copying it keeps
			// both names referring to the same channel.
			if g.memChans[instr.Arg1] {
				g.memChans[instr.Result] = true
			}

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
			if g.isStringOperand(instr.Arg1) || g.isStringOperand(instr.Arg2) {
				g.emitStringCompare(instr.Result, instr.Arg1, instr.Arg2, "==")
			} else {
				g.emitBinary(instr.Result, instr.Arg1, "==", instr.Arg2)
			}
		case "NEQ":
			if g.isStringOperand(instr.Arg1) || g.isStringOperand(instr.Arg2) {
				g.emitStringCompare(instr.Result, instr.Arg1, instr.Arg2, "!=")
			} else {
				g.emitBinary(instr.Result, instr.Arg1, "!=", instr.Arg2)
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

		case "TUPLE_GET":
			prefix := g.declPrefix(instr.Result)
			g.buf.WriteString(fmt.Sprintf("%s%s%s = %s._%s;\n", g.ind(), prefix, instr.Result, instr.Arg1, instr.Arg2))

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

		case "CHAN_NEW":
			// In-memory channel: Arg1 = element type, Arg2 = capacity, Result = temp.
			elemC := mapType(instr.Arg1)
			prefix := g.declPrefix(instr.Result)
			g.buf.WriteString(fmt.Sprintf("%s%s%s = mp_chan_new(sizeof(%s), %s);\n",
				g.ind(), prefix, instr.Result, elemC, instr.Arg2))
			g.memChans[instr.Result] = true

		case "CHAN_DECL":
			ctor := "mp_chan_client"
			if instr.Arg1 == "s_channel" {
				ctor = "mp_chan_server"
			}
			// Streaming endpoint form is <name><T>(host, port). host and port are
			// the last two args, so any legacy leading args are simply ignored.
			name := instr.Arg2
			host, port := "\"127.0.0.1\"", "0"
			if n := len(g.paramBuf); n >= 2 {
				host = g.paramBuf[n-2]
				port = g.paramBuf[n-1]
			}
			if g.insideFunc {
				g.buf.WriteString(fmt.Sprintf("%smp_chan* %s = %s(%s, %s);\n", g.ind(), name, ctor, host, port))
			} else {
				g.buf.WriteString(fmt.Sprintf("mp_chan* %s;\n__attribute__((constructor)) void __init_%s() {\n    %s = %s(%s, %s);\n}\n", name, name, name, ctor, host, port))
			}
			g.netChans[name] = true
			g.markDeclared(name)
			g.paramBuf = g.paramBuf[:0]

		case "METHOD_CALL":
			parts := strings.Split(instr.Arg1, ".")
			obj := parts[0]
			method := parts[1]
			args := strings.Join(g.paramBuf, ", ")
			g.paramBuf = g.paramBuf[:0]

			if g.memChans[obj] {
				// In-memory channel: send/recv exchange elem_size bytes by address.
				elemC := mapType(chanElem(g.resolveType(obj)))
				switch method {
				case "send":
					g.buf.WriteString(fmt.Sprintf("%s{ %s _mpv = %s; mp_chan_send(%s, &_mpv); }\n",
						g.ind(), elemC, args, obj))
				case "recv":
					prefix := g.declPrefix(instr.Result)
					if prefix != "" {
						g.buf.WriteString(fmt.Sprintf("%s%s%s;\n", g.ind(), prefix, instr.Result))
					}
					g.buf.WriteString(fmt.Sprintf("%smp_chan_recv(%s, &%s);\n", g.ind(), obj, instr.Result))
				case "close":
					g.buf.WriteString(fmt.Sprintf("%smp_chan_close(%s);\n", g.ind(), obj))
				}
			} else if g.netChans[obj] {
				// Network channel: (de)serialize the typed value over the socket.
				elem := chanElem(g.resolveType(obj))
				elemC := mapType(elem)
				ser := serName(elem)
				switch method {
				case "send":
					g.buf.WriteString(fmt.Sprintf("%s{ %s _mpv = %s; mp_enc_%s(%s->sock, &_mpv); }\n",
						g.ind(), elemC, args, ser, obj))
				case "recv":
					prefix := g.declPrefix(instr.Result)
					if prefix != "" {
						g.buf.WriteString(fmt.Sprintf("%s%s%s;\n", g.ind(), prefix, instr.Result))
					}
					g.buf.WriteString(fmt.Sprintf("%smp_dec_%s(%s->sock, &%s);\n", g.ind(), ser, obj, instr.Result))
				case "close":
					g.buf.WriteString(fmt.Sprintf("%smp_chan_close(%s);\n", g.ind(), obj))
				}
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
