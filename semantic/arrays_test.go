package semantic

import "testing"

// Estes testes especificam o contrato do método de array `append`: é aceito
// sobre qualquer array (vazio inferido ou tipado), e chamar `append` — ou um
// método desconhecido — sobre algo que não é array é um erro.

// TestArrayAppendValid exercita usos válidos de append; devem analisar sem erros.
func TestArrayAppendValid(t *testing.T) {
	cases := map[string]string{
		"append em lista vazia":    `func f() { let xs = []; xs.append(1); }`,
		"append em lista tipada":   `func f() { let xs = [1, 2, 3]; xs.append(4); }`,
		"appends sequenciais":      `func f() { let xs = []; xs.append(1); xs.append(2); xs.append(3); }`,
		"append de string em lista": `func f() { let xs = ["a"]; xs.append("b"); }`,
	}
	for name, src := range cases {
		t.Run(name, func(t *testing.T) {
			if errs := analyze(t, src); len(errs) != 0 {
				t.Errorf("esperava 0 erros, recebeu: %v", errs)
			}
		})
	}
}

// TestArrayAppendErrors garante que append (e métodos desconhecidos) só existem
// para arrays: usá-los em outros tipos gera "não possui método".
func TestArrayAppendErrors(t *testing.T) {
	cases := []struct {
		name string
		src  string
	}{
		{"metodo desconhecido em array", `func f() { let xs = []; xs.foo(); }`},
		{"append em string", `func f() { let s = "x"; s.append("y"); }`},
		{"append em inteiro", `func f() { let n = 1; n.append(1); }`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			errs := analyze(t, tc.src)
			if !containsAny(errs, "não possui método") {
				t.Errorf("esperava erro 'não possui método', recebeu: %v", errs)
			}
		})
	}
}
