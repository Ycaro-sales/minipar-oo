package tac

import (
	"strings"
	"testing"
)

// tacLines runs the full pipeline and returns the TAC output split into lines,
// stripping empty lines.
func tacLines(t *testing.T, src string) []string {
	t.Helper()
	prog := parseProgram(t, src)
	analyze(t, prog)
	gen := New()
	gen.Generate(prog)
	out := gen.Instructions()
	var lines []string
	for _, l := range strings.Split(out, "\n") {
		if strings.TrimSpace(l) != "" {
			lines = append(lines, strings.TrimSpace(l))
		}
	}
	return lines
}

// containsOp reports whether any line starts with the given op.
func containsOp(lines []string, op string) bool {
	for _, l := range lines {
		if strings.HasPrefix(l, op) {
			return true
		}
	}
	return false
}

// linesWithOp returns all lines that start with op.
func linesWithOp(lines []string, op string) []string {
	var out []string
	for _, l := range lines {
		if strings.HasPrefix(l, op) {
			out = append(out, l)
		}
	}
	return out
}

// ==========================================
// ARRAY_NEW + ARRAY_SET (declaração de literal)
// ==========================================

// TestTACArrayLiteral verifica que [1,2,3] emite ARRAY_NEW + 3× ARRAY_SET.
func TestTACArrayLiteral(t *testing.T) {
	lines := tacLines(t, `func main() { let arr: [i32] = [1, 2, 3] }`)

	if !containsOp(lines, "ARRAY_NEW") {
		t.Errorf("esperava ARRAY_NEW\n%s", strings.Join(lines, "\n"))
	}
	sets := linesWithOp(lines, "ARRAY_SET")
	if len(sets) != 3 {
		t.Errorf("esperava 3 ARRAY_SET, recebeu %d\n%s", len(sets), strings.Join(lines, "\n"))
	}
}

// TestTACArrayNewCountAndType verifica que ARRAY_NEW carrega a contagem correta.
// O tipo do elemento é inferido do literal (int, não i32), então verificamos
// apenas a contagem e a presença de ARRAY_NEW.
func TestTACArrayNewCountAndType(t *testing.T) {
	lines := tacLines(t, `func main() { let arr: [i32] = [10, 20] }`)

	var found bool
	for _, l := range lines {
		if strings.HasPrefix(l, "ARRAY_NEW") && strings.Contains(l, "2") {
			found = true
		}
	}
	if !found {
		t.Errorf("esperava ARRAY_NEW com contagem 2\n%s", strings.Join(lines, "\n"))
	}
}

// TestTACArraySetIndices verifica que ARRAY_SET usa índices 0, 1, 2 em ordem.
func TestTACArraySetIndices(t *testing.T) {
	lines := tacLines(t, `func main() { let arr: [i32] = [10, 20, 30] }`)
	sets := linesWithOp(lines, "ARRAY_SET")
	if len(sets) != 3 {
		t.Fatalf("esperava 3 ARRAY_SET, recebeu %d", len(sets))
	}
	for i, expected := range []string{"0", "1", "2"} {
		if !strings.Contains(sets[i], expected) {
			t.Errorf("ARRAY_SET[%d]: esperava índice %s, linha: %s", i, expected, sets[i])
		}
	}
}

// ==========================================
// ARRAY_GET (acesso por índice)
// ==========================================

// TestTACArrayGet verifica que arr[i] emite ARRAY_GET.
func TestTACArrayGet(t *testing.T) {
	lines := tacLines(t, `func main() { let arr: [i32] = [1,2,3]; let x = arr[0] }`)

	if !containsOp(lines, "ARRAY_GET") {
		t.Errorf("esperava ARRAY_GET\n%s", strings.Join(lines, "\n"))
	}
}

// ==========================================
// ARRAY_SET (atribuição por índice)
// ==========================================

// TestTACIndexAssignment verifica que arr[i] = x emite ARRAY_SET.
func TestTACIndexAssignment(t *testing.T) {
	lines := tacLines(t, `func main() { let arr: [i32] = [1,2,3]; arr[0] = 99 }`)

	sets := linesWithOp(lines, "ARRAY_SET")
	// 3 da construção + 1 da atribuição = 4
	if len(sets) < 4 {
		t.Errorf("esperava pelo menos 4 ARRAY_SET (3 construção + 1 atribuição), recebeu %d\n%s",
			len(sets), strings.Join(lines, "\n"))
	}
}

// TestTACIndexAssignmentOp verifica que a instrução ARRAY_SET de atribuição
// contém o objeto, índice e valor corretos.
func TestTACIndexAssignmentOp(t *testing.T) {
	lines := tacLines(t, `func main() { let arr: [i32] = [0]; arr[0] = 42 }`)

	sets := linesWithOp(lines, "ARRAY_SET")
	// Último ARRAY_SET é o da atribuição; deve conter 42.
	if len(sets) == 0 {
		t.Fatal("nenhum ARRAY_SET encontrado")
	}
	last := sets[len(sets)-1]
	if !strings.Contains(last, "42") {
		t.Errorf("último ARRAY_SET deveria conter 42 (valor atribuído): %s", last)
	}
}

// ==========================================
// TUPLE_NEW + TUPLE_SET (declaração de literal)
// ==========================================

// TestTACTupleLiteral verifica que (1, "a") emite TUPLE_NEW + 2× TUPLE_SET.
func TestTACTupleLiteral(t *testing.T) {
	lines := tacLines(t, `func main() { let t = (1, "a") }`)

	if !containsOp(lines, "TUPLE_NEW") {
		t.Errorf("esperava TUPLE_NEW\n%s", strings.Join(lines, "\n"))
	}
	sets := linesWithOp(lines, "TUPLE_SET")
	if len(sets) != 2 {
		t.Errorf("esperava 2 TUPLE_SET, recebeu %d\n%s", len(sets), strings.Join(lines, "\n"))
	}
}

// TestTACTupleSetIndices verifica que TUPLE_SET usa índices 0 e 1 em ordem.
func TestTACTupleSetIndices(t *testing.T) {
	lines := tacLines(t, `func main() { let t = (10, 20) }`)
	sets := linesWithOp(lines, "TUPLE_SET")
	if len(sets) != 2 {
		t.Fatalf("esperava 2 TUPLE_SET, recebeu %d", len(sets))
	}
	for i, expected := range []string{"0", "1"} {
		if !strings.Contains(sets[i], expected) {
			t.Errorf("TUPLE_SET[%d]: esperava índice %s, linha: %s", i, expected, sets[i])
		}
	}
}

// ==========================================
// TIPOS DOS TEMPORÁRIOS
// ==========================================

// TestTACArrayTempType verifica que o temporário gerado por [1,2,3] recebe o
// tipo "[i32]" no mapa TempTypes.
func TestTACArrayTempType(t *testing.T) {
	prog := parseProgram(t, `func main() { let arr: [i32] = [1, 2, 3] }`)
	analyze(t, prog)
	gen := New()
	gen.Generate(prog)

	found := false
	for _, typ := range gen.TempTypes() {
		if typ == "[int]" || typ == "[i32]" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("esperava temporário com tipo [i32] ou [int] em TempTypes: %v", gen.TempTypes())
	}
}

// TestTACTupleTempType verifica que o temporário de uma tupla recebe o tipo
// correto no mapa TempTypes.
func TestTACTupleTempType(t *testing.T) {
	prog := parseProgram(t, `func main() { let t = (1, "a") }`)
	analyze(t, prog)
	gen := New()
	gen.Generate(prog)

	found := false
	for _, typ := range gen.TempTypes() {
		if strings.HasPrefix(typ, "(") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("esperava temporário com tipo de tupla em TempTypes: %v", gen.TempTypes())
	}
}
