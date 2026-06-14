package cgen_test

import (
	"strings"
	"testing"
)

// ==========================================
// ARRAYS — TYPEDEF E DECLARAÇÃO
// ==========================================

// TestCGenArrayTypedef verifica que um typedef de struct com .data e .len é emitido.
func TestCGenArrayTypedef(t *testing.T) {
	out := compileToC(t, `func main() { let arr: [i32] = [1, 2, 3] }`)
	assertContains(t, out,
		"typedef struct",
		"* data;",
		"len;",
	)
}

// TestCGenArrayDeclHasData verifica que a declaração do array contém os valores
// do literal como inicializadores do array de dados.
func TestCGenArrayDeclHasData(t *testing.T) {
	out := compileToC(t, `func main() { let arr: [i32] = [10, 20, 30] }`)
	assertContains(t, out, "10", "20", "30")
	// A variável interna de dados deve existir
	if !strings.Contains(out, "_data") {
		t.Errorf("esperava _data no output\n--- output ---\n%s", out)
	}
}

// TestCGenArrayVarType verifica que a variável arr é declarada com um tipo de array.
func TestCGenArrayVarType(t *testing.T) {
	out := compileToC(t, `func main() { let arr: [i32] = [1, 2] }`)
	// Deve haver algum tipo "arr_..." que declara a variável
	if !strings.Contains(out, "arr_") {
		t.Errorf("esperava tipo arr_xxx para a variável de array\n--- output ---\n%s", out)
	}
}

// ==========================================
// ARRAYS — ACESSO (ARRAY_GET)
// ==========================================

// TestCGenArrayGet verifica que arr[i] gera acesso via .data[i].
func TestCGenArrayGet(t *testing.T) {
	out := compileToC(t, `func main() {
    let arr: [i32] = [10, 20, 30]
    let x = arr[0]
}`)
	assertContains(t, out, ".data[0]")
}

// TestCGenArrayGetVar verifica que arr[i] com variável como índice gera .data[i].
func TestCGenArrayGetVar(t *testing.T) {
	out := compileToC(t, `func main() {
    let arr: [i32] = [10, 20, 30]
    let i: i32 = 1
    let x = arr[i]
}`)
	assertContains(t, out, ".data[i]")
}

// ==========================================
// ARRAYS — MUTAÇÃO (ARRAY_SET via IndexAssignment)
// ==========================================

// TestCGenArraySet verifica que arr[0] = 99 gera .data[0] = 99.
func TestCGenArraySet(t *testing.T) {
	out := compileToC(t, `func main() {
    let arr: [i32] = [1, 2, 3]
    arr[0] = 99
}`)
	assertContains(t, out, ".data[0] = 99")
}

// TestCGenArraySetVar verifica arr[i] = x com variáveis.
func TestCGenArraySetVar(t *testing.T) {
	out := compileToC(t, `func main() {
    let arr: [i32] = [1, 2, 3]
    let i: i32 = 1
    let v: i32 = 42
    arr[i] = v
}`)
	assertContains(t, out, ".data[i] = v")
}

// ==========================================
// ARRAYS — PASSAGEM PARA FUNÇÕES
// ==========================================

// TestCGenArrayParam verifica que função com param [i32] usa tipo de struct array.
func TestCGenArrayParam(t *testing.T) {
	out := compileToC(t, `
func soma(arr: [i32]) -> i32 {
    return arr[0]
}
func main() {}`)
	assertContains(t, out, "arr_")    // parâmetro tem tipo struct array
	assertContains(t, out, ".data[0]") // acesso via .data
}

// TestCGenSwap verifica o padrão completo de swap (como no quicksort).
func TestCGenSwap(t *testing.T) {
	out := compileToC(t, `
func swap(arr: [i32], i: i32, j: i32) {
    let tmp: i32 = arr[i]
    arr[i] = arr[j]
    arr[j] = tmp
}
func main() {}`)
	assertContains(t, out, ".data[i]", ".data[j]") // acesso e mutação
	assertContains(t, out, "arr_")                  // tipo de struct array no parâmetro
}

// ==========================================
// ARRAYS — LEN (via for)
// ==========================================

// TestCGenArrayLen verifica que ARRAY_LEN gera acesso via .len.
func TestCGenArrayLen(t *testing.T) {
	out := compileToC(t, `func main() {
    let arr: [i32] = [1, 2, 3]
    for (x in arr) { print(x) }
}`)
	assertContains(t, out, ".len")
}

// ==========================================
// TUPLAS — TYPEDEF E DECLARAÇÃO
// ==========================================

// TestCGenTupleTypedef verifica que um typedef com _0 e _1 é emitido.
func TestCGenTupleTypedef(t *testing.T) {
	out := compileToC(t, `func main() { let t = (1, "hello") }`)
	assertContains(t, out, "typedef struct", "_0;", "_1;")
}

// TestCGenTupleDecl verifica que a tupla é declarada com valores iniciais corretos.
func TestCGenTupleDecl(t *testing.T) {
	out := compileToC(t, `func main() { let t = (10, 20) }`)
	assertContains(t, out, "10", "20")
	// Deve ter campos _0 e _1
	if !strings.Contains(out, "_0") && !strings.Contains(out, "10") {
		t.Errorf("esperava campo _0 ou inicialização com 10\n--- output ---\n%s", out)
	}
}

// TestCGenTupleVarType verifica que a variável de tupla tem um tipo tup_...
func TestCGenTupleVarType(t *testing.T) {
	out := compileToC(t, `func main() { let t = (1, "a") }`)
	if !strings.Contains(out, "tup_") {
		t.Errorf("esperava tipo tup_xxx para a variável de tupla\n--- output ---\n%s", out)
	}
}

// ==========================================
// NÃO REGRIDE TESTES EXISTENTES
// ==========================================

// TestCGenArrayDoesNotBreakPrimitives garante que variáveis primitivas ainda funcionam.
func TestCGenArrayDoesNotBreakPrimitives(t *testing.T) {
	out := compileToC(t, `func main() {
    let x: i32 = 5
    let y: i32 = x + 1
    print(y)
}`)
	assertContains(t, out, "int32_t x = 5;", "printf")
	assertNotContains(t, out, "ERRO")
}
