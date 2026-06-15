package semantic

import (
	"testing"

	"minipar/parser"
)

// Estes testes especificam o contrato dos tipos de canal: declarações
// s_channel/c_channel carregam um tipo elemento (chan<T>) e expõem a interface
// tipada send(T) / recv() -> T / close(). send/recv são verificados contra T.

// TestChannelTypesValid exercita usos válidos de canais; devem analisar sem erros.
func TestChannelTypesValid(t *testing.T) {
	cases := map[string]string{
		"s_channel tipado": `s_channel srv<string>("localhost", 1234)
			func main() { srv.send("oi"); srv.close(); }`,
		"c_channel tipado": `c_channel cli<i32>("localhost", 1234)
			func main() { cli.send(7); cli.close(); }`,
		"recv retorna o tipo elemento": `c_channel cli<i32>("localhost", 1234)
			func main() { let n: i32 = cli.recv(); print(n); }`,
		"parametro tipado chan": `func f(c: chan<string>) { c.send("x"); c.close(); }`,
	}
	for name, src := range cases {
		t.Run(name, func(t *testing.T) {
			if errs := analyze(t, src); len(errs) != 0 {
				t.Errorf("esperava 0 erros, recebeu: %v", errs)
			}
		})
	}
}

// TestChannelTypeErrors cobre violações do contrato tipado.
func TestChannelTypeErrors(t *testing.T) {
	cases := map[string]string{
		"send com tipo incompatível": `c_channel cli<i32>("localhost", 1234)
			func main() { cli.send("texto"); }`,
		"send com aridade errada": `c_channel cli<i32>("localhost", 1234)
			func main() { cli.send(1, 2); }`,
		"recv não aceita argumentos": `c_channel cli<i32>("localhost", 1234)
			func main() { let n: i32 = cli.recv(1); print(n); }`,
		"método inexistente": `c_channel cli<i32>("localhost", 1234)
			func main() { cli.flush(); }`,
	}
	for name, src := range cases {
		t.Run(name, func(t *testing.T) {
			if errs := analyze(t, src); len(errs) == 0 {
				t.Errorf("esperava ao menos 1 erro, recebeu 0")
			}
		})
	}
}

// TestChannelTypeResolved confirma que uma declaração de canal define, no escopo
// global, uma variável cujo tipo resolvido é KindChan com o tipo elemento certo.
func TestChannelTypeResolved(t *testing.T) {
	src := `s_channel srv<string>("localhost", 1234) func main() { }`
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
	if sym.Type.Elem == nil || sym.Type.Elem.Kind != KindString {
		t.Errorf("tipo elemento de 'srv': esperava string, recebeu %v", sym.Type.Elem)
	}
}
