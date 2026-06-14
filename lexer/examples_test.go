package lexer

import (
	"os"
	"path/filepath"
	"testing"
)

// TestExamplesTokenize lexes every sample under tests/ and asserts that the
// scanner consumes the whole file without producing an ILLEGAL token. This
// guards against the lexer choking on any real-world construct used in the
// example programs (comments, channels, char/string literals, operators...).
func TestExamplesTokenize(t *testing.T) {
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

			l := New(string(src))
			for {
				tok := l.NextToken()
				if tok.Type == ILLEGAL {
					t.Errorf("token ILLEGAL na linha %d: %q", tok.Line, tok.Literal)
				}
				if tok.Type == EOF {
					break
				}
			}
		})
	}
}
