package semantic

import "testing"

// Estes testes especificam o contrato das funções builtin registradas por
// registerBuiltins: nome, aridade (1 argumento) e tipo de retorno. Eles definem
// o comportamento esperado — se um builtin for alterado, o teste deve quebrar.

// TestBuiltinsValid exercita cada builtin no seu uso correto, atribuindo o
// retorno a uma variável do tipo esperado. Devem analisar sem erros.
func TestBuiltinsValid(t *testing.T) {
	cases := map[string]string{
		"len retorna int":        `func f() { let n: i32 = len("abc"); }`,
		"to_string retorna str":  `func f() { let s: string = to_string(42); }`,
		"to_number retorna int":  `func f() { let n: i32 = to_number("5"); }`,
		"isalpha retorna bool":   `func f() { let b: bool = isalpha("a"); }`,
		"isnum retorna bool":     `func f() { let b: bool = isnum("1"); }`,
		"isalpha em condicao if": `func f() { if (isalpha("a")) { let x = 1; } }`,
		"isnum com and":          `func f() { if (isnum("1") and true) { let x = 1; } }`,
		"len em aritmetica":      `func f() { let n: i32 = len("ab") + 1; }`,
		"input retorna str":      `func f() { let s: string = input("p"); }`,
		"input sem argumentos":   `func f() { let s: string = input(); }`,
		"print sem argumentos":   `func f() { print(); }`,
		"print um argumento":     `func f() { print("a"); }`,
		"print variadico":        `func f() { print("a", 1, true); }`,
	}
	for name, src := range cases {
		t.Run(name, func(t *testing.T) {
			if errs := analyze(t, src); len(errs) != 0 {
				t.Errorf("esperava 0 erros, recebeu: %v", errs)
			}
		})
	}
}

// TestBuiltinsReturnTypeMismatch garante que o tipo de retorno é realmente
// imposto: atribuir o resultado a um tipo incompatível deve gerar erro.
func TestBuiltinsReturnTypeMismatch(t *testing.T) {
	cases := []struct {
		name string
		src  string
	}{
		{"len nao e string", `func f() { let s: string = len("abc"); }`},
		{"to_string nao e int", `func f() { let n: i32 = to_string(1); }`},
		{"to_number nao e bool", `func f() { let b: bool = to_number("1"); }`},
		{"isalpha nao e int", `func f() { let n: i32 = isalpha("a"); }`},
		{"isnum nao e string", `func f() { let s: string = isnum("1"); }`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			errs := analyze(t, tc.src)
			if !containsAny(errs, "não é possível atribuir") {
				t.Errorf("esperava erro de atribuição incompatível, recebeu: %v", errs)
			}
		})
	}
}

// TestBuiltinsArity garante que os builtins têm aridade 1: chamá-los com número
// errado de argumentos deve gerar erro de aridade.
func TestBuiltinsArity(t *testing.T) {
	cases := []struct {
		name string
		src  string
	}{
		{"len sem argumentos", `func f() { let n = len(); }`},
		{"len com dois argumentos", `func f() { let n = len("a", "b"); }`},
		{"to_string com dois argumentos", `func f() { let s = to_string(1, 2); }`},
		{"isnum sem argumentos", `func f() { let b = isnum(); }`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			errs := analyze(t, tc.src)
			if !containsAny(errs, "argumento(s)") {
				t.Errorf("esperava erro de aridade, recebeu: %v", errs)
			}
		})
	}
}

// TestBuiltinsAreFunctions confirma que os nomes builtin estão definidos no
// escopo global como funções utilizáveis (resolvidos como chamada, não como
// identificador livre).
func TestBuiltinsAreFunctions(t *testing.T) {
	for _, name := range []string{"len", "to_string", "to_number", "isalpha", "isnum", "print", "input"} {
		t.Run(name, func(t *testing.T) {
			a := New()
			sym := a.GlobalScope().Resolve(name)
			if sym == nil {
				t.Fatalf("builtin '%s' não está registrado no escopo global", name)
			}
			if sym.Type == nil || sym.Type.Kind != KindFunc {
				t.Errorf("builtin '%s' deveria ser uma função, tipo: %v", name, sym.Type)
			}
		})
	}
}
