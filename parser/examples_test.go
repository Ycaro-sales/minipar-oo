package parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestExamplesParse ensures every sample under tests/ is syntactically valid in
// the current grammar. ParseProgram returns only parser (syntactic) errors, so
// undeclared builtins (len, to_string, .append, ...) do not affect this check.
func TestExamplesParse(t *testing.T) {
	files, err := filepath.Glob("../tests/*.minipar")
	if err != nil {
		t.Fatalf("glob falhou: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("nenhum exemplo .minipar encontrado em ../tests")
	}

	for _, file := range files {
		t.Run(filepath.Base(file), func(t *testing.T) {
			src, err := os.ReadFile(file)
			if err != nil {
				t.Fatalf("não foi possível ler %s: %v", file, err)
			}
			// Arquivos marcados com "# ALVO:" são código-alvo que depende de
			// features ainda não implementadas; ignorados até estarem prontos.
			if strings.HasPrefix(strings.TrimSpace(string(src)), "# ALVO:") {
				t.Skip("código-alvo: depende de features pendentes no parser/semântico")
			}
			_, errs := ParseProgram(string(src))
			if len(errs) != 0 {
				t.Errorf("erros sintáticos em %s:\n  %v", filepath.Base(file), errs)
			}
		})
	}
}
