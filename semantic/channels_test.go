package semantic

import (
	"testing"

	"minipar/parser"
)

// Estes testes especificam o contrato dos tipos de canal: declarações
// s_channel/c_channel (e o tipo `chan` em anotações) resolvem para um tipo de
// canal, e qualquer método chamado sobre um canal (send/close/recv/...) é aceito
// e produz um valor `any` utilizável.

// TestChannelTypesValid exercita usos válidos de canais; devem analisar sem erros.
func TestChannelTypesValid(t *testing.T) {
	cases := map[string]string{
		"s_channel no topo": `s_channel srv("localhost", 1234)
			func main() { srv.send("oi"); srv.close(); }`,
		"c_channel no topo": `c_channel cli("localhost", 1234)
			func main() { cli.send("x", 1, 2); cli.close(); }`,
		"parametro tipado chan": `func f(c: chan) { c.send("x"); c.close(); }`,
	}
	for name, src := range cases {
		t.Run(name, func(t *testing.T) {
			if errs := analyze(t, src); len(errs) != 0 {
				t.Errorf("esperava 0 erros, recebeu: %v", errs)
			}
		})
	}
}

// TestChannelTypeResolved confirma que uma declaração de canal define, no escopo
// global, uma variável cujo tipo resolvido é KindChan.
func TestChannelTypeResolved(t *testing.T) {
	src := `s_channel srv("localhost", 1234) func main() { }`
	prog, perrs := parser.ParseProgram(src)
	if len(perrs) > 0 {
		t.Fatalf("erros de parsing inesperados: %v", perrs)
	}
	a := New()
	if errs := a.Analyze(prog); len(errs) != 0 {
		t.Fatalf("esperava 0 erros, recebeu: %v", errs)
	}
	sym := a.GlobalScope().Resolve("srv")
	if sym == nil {
		t.Fatal("canal 'srv' não foi definido no escopo global")
	}
	if sym.Type == nil || sym.Type.Kind != KindChan {
		t.Errorf("tipo de 'srv': esperava KindChan, recebeu %v", sym.Type)
	}
}

// TestChannelMethodReturnsAny garante que um método de canal produz um valor
// `any` utilizável em expressões (any + int -> any, atribuível a i32).
func TestChannelMethodReturnsAny(t *testing.T) {
	src := `c_channel cli("localhost", 1234)
		func main() { let n: i32 = cli.recv() + 1; print(n); }`
	if errs := analyze(t, src); len(errs) != 0 {
		t.Errorf("esperava 0 erros, recebeu: %v", errs)
	}
}
