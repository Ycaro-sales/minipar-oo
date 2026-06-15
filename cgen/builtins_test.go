package cgen_test

// Testes das funções built-in: tradução de nome para C nativo e helpers próprios.
//
// Abordagem (decidida com o usuário):
//   - Builtins com equivalente C: traduzir o nome na hora do CALL (len→strlen,
//     to_number→atoi, isnum→isdigit). isalpha já tem nome idêntico na libc.
//   - Builtins sem equivalente: emitir helper próprio no preâmbulo (to_string, input).
//   - <ctype.h> incluído APENAS quando isalpha/isnum é usado.
//   - O handler CALL em cgen.go não muda estruturalmente, só o nome efetivo.

import "testing"

// ==========================================
// len → strlen
// ==========================================

func TestBuiltin_len_translates_to_strlen(t *testing.T) {
	src := `
func main() {
    let s: string = "abc"
    let n: i32 = len(s)
    print(n)
}`
	out := compileToC(t, src)
	assertContains(t, out, "strlen(")
	// O nome original "len" não deve aparecer como chamada de função C (apenas
	// como parte de "strlen").
	assertNotContains(t, out, "= len(")
}

// ==========================================
// to_number → atoi
// ==========================================

func TestBuiltin_toNumber_translates_to_atoi(t *testing.T) {
	src := `
func main() {
    let s: string = "5"
    let n: i32 = to_number(s)
    print(n)
}`
	out := compileToC(t, src)
	assertContains(t, out, "atoi(")
	assertNotContains(t, out, "to_number(")
}

// ==========================================
// isalpha → isalpha de <ctype.h>
// ==========================================

func TestBuiltin_isalpha_uses_ctype_h(t *testing.T) {
	src := `
func main() {
    let s: string = "a"
    let b: bool = isalpha(s)
    print(b)
}`
	out := compileToC(t, src)
	// Deve incluir <ctype.h> para que isalpha seja declarada.
	assertContains(t, out, "#include <ctype.h>")
	// A chamada deve usar o nome isalpha (sem renomear).
	assertContains(t, out, "isalpha(")
}

func TestBuiltin_isalpha_no_homonym_definition(t *testing.T) {
	// Não devemos definir um helper próprio chamado isalpha — usamos a da libc.
	src := `
func main() {
    let s: string = "x"
    let b: bool = isalpha(s)
    print(b)
}`
	out := compileToC(t, src)
	// Não pode aparecer "static" antes de "isalpha" (seria definição própria).
	assertNotContains(t, out, "static bool isalpha")
}

// ==========================================
// isnum → isdigit de <ctype.h>
// ==========================================

func TestBuiltin_isnum_translates_to_isdigit(t *testing.T) {
	src := `
func main() {
    let s: string = "1"
    let b: bool = isnum(s)
    print(b)
}`
	out := compileToC(t, src)
	assertContains(t, out, "isdigit(")
	assertNotContains(t, out, "= isnum(")
	// <ctype.h> também deve ser incluído para isdigit.
	assertContains(t, out, "#include <ctype.h>")
}

// ==========================================
// to_string → helper próprio (C não tem)
// ==========================================

func TestBuiltin_toString_emits_helper_definition(t *testing.T) {
	src := `
func main() {
    let n: i32 = 42
    let s: string = to_string(n)
    print(s)
}`
	out := compileToC(t, src)
	// O helper deve ser definido no preâmbulo.
	assertContains(t, out, "to_string(")
	assertContains(t, out, "snprintf(")
	// A definição deve ter assinatura long long → char*.
	assertContains(t, out, "long long")
}

func TestBuiltin_toString_helper_appears_before_main(t *testing.T) {
	// O helper deve aparecer no preâmbulo, antes da função main.
	src := `
func main() {
    let n: i32 = 1
    let s: string = to_string(n)
    print(s)
}`
	out := compileToC(t, src)
	helperIdx := indexOf(out, "static char* to_string")
	mainIdx := indexOf(out, "int main(")
	if helperIdx == -1 {
		t.Fatalf("helper 'to_string' não encontrado no output:\n%s", out)
	}
	if mainIdx == -1 {
		t.Fatalf("'int main(' não encontrado no output:\n%s", out)
	}
	if helperIdx > mainIdx {
		t.Errorf("helper to_string deve aparecer antes de main(); helper em %d, main em %d", helperIdx, mainIdx)
	}
}

// ==========================================
// input → helper mp_input próprio
// ==========================================

func TestBuiltin_input_emits_mp_input_helper(t *testing.T) {
	src := `
func main() {
    input("Enter: ")
}`
	out := compileToC(t, src)
	// O helper mp_input deve ser definido no preâmbulo.
	assertContains(t, out, "mp_input(")
	assertContains(t, out, "fgets(")
	// A chamada ao helper deve aparecer no corpo da função.
	assertContains(t, out, "= mp_input()")
}

func TestBuiltin_input_with_prompt_emits_printf(t *testing.T) {
	src := `
func main() {
    input("Digite algo: ")
}`
	out := compileToC(t, src)
	// O prompt deve ser emitido com printf antes de chamar mp_input.
	assertContains(t, out, `"Digite algo: "`)
}

func TestBuiltin_input_no_prompt(t *testing.T) {
	src := `
func main() {
    input()
}`
	out := compileToC(t, src)
	assertContains(t, out, "mp_input()")
}

// ==========================================
// Programa SEM builtins → sem helpers nem <ctype.h> desnecessários
// ==========================================

func TestBuiltin_no_helpers_when_unused(t *testing.T) {
	src := `
func main() {
    let x: i32 = 1 + 2
    print(x)
}`
	out := compileToC(t, src)
	// Nenhum helper desnecessário deve aparecer.
	assertNotContains(t, out,
		"static char* to_string",
		"static char* mp_input",
		"mp_input()",
	)
	// <ctype.h> não deve aparecer quando isalpha/isnum não é usado.
	assertNotContains(t, out, "#include <ctype.h>")
}

// ==========================================
// <ctype.h> ausente quando não há isalpha/isnum
// ==========================================

func TestBuiltin_no_ctype_h_when_no_isalpha_isnum(t *testing.T) {
	// to_string e len não precisam de <ctype.h>.
	src := `
func main() {
    let s: string = "hello"
    let n: i32 = len(s)
    let t: string = to_string(n)
    print(t)
}`
	out := compileToC(t, src)
	assertNotContains(t, out, "#include <ctype.h>")
}

// ==========================================
// Helpers de utilidade
// ==========================================

// indexOf retorna o índice byte da primeira ocorrência de sub em s, ou -1.
func indexOf(s, sub string) int {
	for i := range s {
		if len(s)-i >= len(sub) && s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
