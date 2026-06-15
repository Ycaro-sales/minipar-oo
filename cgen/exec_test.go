package cgen_test

// Testes end-to-end com gcc: compilam o C gerado, executam e comparam stdout.
// Todos os testes são pulados quando gcc não está disponível.

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// runC compila cSrc com gcc (lendo de stdin) e executa o binário, retornando stdout.
// Pula o teste se gcc não estiver disponível.
func runC(t *testing.T, cSrc string) string {
	t.Helper()
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("gcc não encontrado; pulando teste end-to-end")
	}

	tmp, err := os.CreateTemp("", "minipar_exec_test_*")
	if err != nil {
		t.Fatalf("CreateTemp: %v", err)
	}
	binPath := tmp.Name()
	tmp.Close()
	defer os.Remove(binPath)

	cmd := exec.Command("gcc", "-x", "c", "-", "-o", binPath)
	cmd.Stdin = strings.NewReader(cSrc)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("gcc falhou: %v\n--- C source ---\n%s", err, cSrc)
	}

	out, err := exec.Command(binPath).Output()
	if err != nil {
		t.Fatalf("execução falhou: %v\n--- C source ---\n%s", err, cSrc)
	}
	return strings.TrimRight(string(out), "\n")
}

// ==========================================
// to_string
// ==========================================

func TestExec_toString_integer(t *testing.T) {
	src := `
func main() {
    let n: i32 = 21
    let s: string = to_string(n)
    print(s)
}`
	cSrc := compileToC(t, src)
	got := runC(t, cSrc)
	if got != "21" {
		t.Errorf("to_string(21): want %q, got %q", "21", got)
	}
}

func TestExec_toString_arithmetic(t *testing.T) {
	src := `
func main() {
    let s: string = to_string(20 + 1)
    print(s)
}`
	cSrc := compileToC(t, src)
	got := runC(t, cSrc)
	if got != "21" {
		t.Errorf("to_string(20+1): want %q, got %q", "21", got)
	}
}

// ==========================================
// to_number
// ==========================================

func TestExec_toNumber(t *testing.T) {
	src := `
func main() {
    let s: string = "42"
    let n: i32 = to_number(s)
    print(n)
}`
	cSrc := compileToC(t, src)
	got := runC(t, cSrc)
	if got != "42" {
		t.Errorf("to_number: want %q, got %q", "42", got)
	}
}

// ==========================================
// len
// ==========================================

func TestExec_len_string(t *testing.T) {
	src := `
func main() {
    let s: string = "hello"
    let n: i32 = len(s)
    print(n)
}`
	cSrc := compileToC(t, src)
	got := runC(t, cSrc)
	if got != "5" {
		t.Errorf("len(\"hello\"): want %q, got %q", "5", got)
	}
}

// Nota: testes exec para isalpha/isnum requerem indexação de string (char),
// funcionalidade ainda não implementada em cgen (ARRAY_GET usa .data[i] para
// arrays com struct, não para char* strings). Esses casos são cobertos apenas
// pelos testes de asserção de string em builtins_test.go.
