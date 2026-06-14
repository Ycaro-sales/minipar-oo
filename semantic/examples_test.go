package semantic

import (
	"os"
	"path/filepath"
	"testing"

	"minipar/parser"
)

// TestExamplesAnalyze é a especificação executável da fase semântica: todo
// exemplo em tests/ deve ser um programa Minipar válido e, portanto, analisar
// com ZERO erros. Se um exemplo usa uma função builtin, um tipo ou uma operação
// que o analisador ainda não modela, a lacuna é do analisador — o teste deve
// ficar vermelho até que o componente seja implementado, nunca o contrário.
func TestExamplesAnalyze(t *testing.T) {
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
			prog, perrs := parser.ParseProgram(string(src))
			if len(perrs) > 0 {
				t.Fatalf("erros de parsing inesperados: %v", perrs)
			}
			if errs := New().Analyze(prog); len(errs) != 0 {
				t.Errorf("esperava 0 erros semânticos, recebeu: %v", errs)
			}
		})
	}
}
