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
	varTypes     map[string]string // nome → tipo minipar (da tabela de símbolos)
	declaredVars map[string]bool
	paramBuf     []string // argumentos acumulados antes de CALL
	currentClass  string // classe em definição, para name mangling de métodos
	insideFunc    bool  // controla indentação: global vs. dentro de função
	buf           strings.Builder
}

func New(instrs []tac.Instruction, tempTypes map[string]string, globalScope *symtab.Scope[*semantic.Type]) *CGenerator {
	g := &CGenerator{
		instrs:       instrs,
		tempTypes:    tempTypes,
		varTypes:     make(map[string]string),
		declaredVars: make(map[string]bool),
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

// declPrefix devolve o prefixo de declaração C se name ainda não foi declarado
// ("int32_t name" na primeira ocorrência, "" nas seguintes).
func (g *CGenerator) declPrefix(name string) string {
	if g.declaredVars[name] {
		return ""
	}
	g.declaredVars[name] = true
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

func (g *CGenerator) Generate() string {
	g.buf.WriteString("#include <stdint.h>\n")
	g.buf.WriteString("#include <stdbool.h>\n")
	g.buf.WriteString("#include <stdio.h>\n")
	g.buf.WriteString("#include <stdlib.h>\n")
	g.buf.WriteString("#include <string.h>\n")
	g.buf.WriteByte('\n')
	g.emitCompositeTypedefs()
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
			g.declaredVars = make(map[string]bool)
			// Consome os PARAM_DECL seguintes para montar a lista de parâmetros.
			var params []string
			for i+1 < len(g.instrs) && g.instrs[i+1].Op == "PARAM_DECL" {
				i++
				p := g.instrs[i]
				params = append(params, fmt.Sprintf("%s %s", mapType(p.Arg2), p.Arg1))
				g.declaredVars[p.Arg1] = true // param already declared in signature
			}
			g.buf.WriteString(fmt.Sprintf("%s %s(%s) {\n", retC, instr.Arg1, strings.Join(params, ", ")))
			g.insideFunc = true

		case "PARAM_DECL":
			// Já consumido pelo look-ahead de BEGIN_FUNC; ignorar se aparecer solto.

		case "PARAM":
			g.paramBuf = append(g.paramBuf, instr.Arg1)

		case "CALL":
			args := strings.Join(g.paramBuf, ", ")
			g.paramBuf = g.paramBuf[:0]
			retType := g.resolveType(instr.Result)
			if retType == "void" || retType == "" {
				g.buf.WriteString(fmt.Sprintf("%s%s(%s);\n", g.ind(), instr.Arg1, args))
			} else {
				prefix := g.declPrefix(instr.Result)
				g.buf.WriteString(fmt.Sprintf("%s%s%s = %s(%s);\n", g.ind(), prefix, instr.Result, instr.Arg1, args))
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
			g.declaredVars[arrTemp] = true

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
			g.declaredVars[tupTemp] = true

		case "TUPLE_SET":
			// Standalone TUPLE_SET: Arg1=tuple, Arg2=index, Result=value
			g.buf.WriteString(fmt.Sprintf("%s%s._%s = %s;\n",
				g.ind(), instr.Arg1, instr.Arg2, instr.Result))

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
				g.buf.WriteString("    return 0;\n")
			}
			g.buf.WriteString("}\n")
			g.insideFunc = false

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
			g.declaredVars = make(map[string]bool)
			g.declaredVars["self"] = true
			var params []string
			for i+1 < len(g.instrs) && g.instrs[i+1].Op == "PARAM_DECL" {
				i++
				p := g.instrs[i]
				params = append(params, fmt.Sprintf("%s %s", mapType(p.Arg2), p.Arg1))
				g.declaredVars[p.Arg1] = true
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
			g.buf.WriteString(fmt.Sprintf("    /* channel %s (tipo: %s) — não implementado */\n", instr.Arg2, instr.Arg1))
		}
	}

	return g.buf.String()
}
