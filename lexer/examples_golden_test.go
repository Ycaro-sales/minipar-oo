package lexer

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// updateGolden regrava os arquivos golden de tokens em vez de compará-los.
// Uso: go test ./lexer -run TestExamplesTokenStream -update
var updateGolden = flag.Bool("update", false, "regrava os arquivos golden de tokens")

// dumpTokens serializa o stream completo de tokens de src (incluindo o EOF
// final), uma linha por token: número da linha, tipo e literal.
func dumpTokens(src string) string {
	var b strings.Builder
	l := New(src)
	for {
		tok := l.NextToken()
		fmt.Fprintf(&b, "%3d  %-10s %q\n", tok.Line, tok.Type, tok.Literal)
		if tok.Type == EOF {
			break
		}
	}
	return b.String()
}

// TestExamplesTokenStream é a especificação dos tokens produzidos para cada
// exemplo: o stream serializado deve bater exatamente com o golden em testdata/.
// Qualquer mudança na saída do lexer deixa o teste vermelho até o golden ser
// revisado e regravado com -update.
func TestExamplesTokenStream(t *testing.T) {
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
			got := dumpTokens(string(src))

			golden := filepath.Join("testdata", filepath.Base(file)+".tokens")
			if *updateGolden {
				if err := os.WriteFile(golden, []byte(got), 0o644); err != nil {
					t.Fatalf("não foi possível gravar golden %s: %v", golden, err)
				}
				return
			}

			want, err := os.ReadFile(golden)
			if err != nil {
				t.Fatalf("golden ausente (%v); gere com 'go test ./lexer -run TestExamplesTokenStream -update'", err)
			}
			if got != string(want) {
				t.Errorf("stream de tokens diferente do golden %s:\n%s", golden, firstDiff(string(want), got))
			}
		})
	}
}

// firstDiff retorna uma mensagem apontando a primeira linha divergente entre
// want e got, para diagnósticos mais legíveis que um dump completo.
func firstDiff(want, got string) string {
	wl := strings.Split(want, "\n")
	gl := strings.Split(got, "\n")
	for i := 0; i < len(wl) || i < len(gl); i++ {
		var w, g string
		if i < len(wl) {
			w = wl[i]
		}
		if i < len(gl) {
			g = gl[i]
		}
		if w != g {
			return fmt.Sprintf("linha %d:\n  golden: %q\n  obtido: %q", i+1, w, g)
		}
	}
	return "(sem diferença de linha encontrada)"
}
