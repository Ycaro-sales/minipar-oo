package symtab

import "testing"

// stub is a minimal Printable payload so the table can be tested without
// depending on the semantic analyzer's type system.
type stub string

func (s stub) String() string { return string(s) }

func sym(name string, line int) *Symbol[stub] {
	return &Symbol[stub]{Name: name, Kind: Var, Type: stub(name + "Type"), Line: line}
}

func TestDefineAndResolveLocal(t *testing.T) {
	g := NewGlobal[stub]()
	if _, ok := g.Define(sym("x", 1)); !ok {
		t.Fatal("primeira definição de x deveria ter sucesso")
	}
	if prev, ok := g.Define(sym("x", 9)); ok {
		t.Error("redefinição no mesmo escopo deveria falhar")
	} else if prev.Line != 1 {
		t.Errorf("redefinição deveria retornar o símbolo original (linha 1), recebeu linha %d", prev.Line)
	}
	if got := g.ResolveLocal("x"); got == nil || got.Line != 1 {
		t.Errorf("ResolveLocal(x) deveria achar o símbolo original")
	}
	if g.ResolveLocal("y") != nil {
		t.Error("ResolveLocal(y) deveria ser nil")
	}
}

func TestResolveWalksParents(t *testing.T) {
	g := NewGlobal[stub]()
	g.Define(sym("g", 1))
	fn := g.Push(FunctionScope)
	fn.Define(sym("p", 2))

	if fn.Resolve("g") == nil {
		t.Error("escopo filho deveria resolver nome do escopo global")
	}
	if g.Resolve("p") != nil {
		t.Error("escopo global não deveria enxergar nome do filho")
	}
}

func TestShadowing(t *testing.T) {
	g := NewGlobal[stub]()
	g.Define(&Symbol[stub]{Name: "x", Type: stub("outer"), Line: 1})
	blk := g.Push(BlockScope)
	if _, ok := blk.Define(&Symbol[stub]{Name: "x", Type: stub("inner"), Line: 2}); !ok {
		t.Fatal("sombrear nome de escopo externo deveria ser permitido")
	}
	if got := blk.Resolve("x"); got.Type != "inner" {
		t.Errorf("escopo interno deveria resolver para 'inner', recebeu %q", got.Type)
	}
	if got := g.Resolve("x"); got.Type != "outer" {
		t.Errorf("escopo externo deveria continuar 'outer', recebeu %q", got.Type)
	}
}

func TestSymbolsOrdered(t *testing.T) {
	g := NewGlobal[stub]()
	for _, name := range []string{"c", "a", "b"} {
		g.Define(sym(name, 0))
	}
	want := []string{"c", "a", "b"} // insertion order, not sorted
	got := g.Symbols()
	if len(got) != len(want) {
		t.Fatalf("esperava %d símbolos, recebeu %d", len(want), len(got))
	}
	for i, s := range got {
		if s.Name != want[i] {
			t.Errorf("posição %d: esperava %q, recebeu %q", i, want[i], s.Name)
		}
	}
}

func TestRetainedTree(t *testing.T) {
	g := NewGlobal[stub]()
	fn := g.Push(FunctionScope)
	blk := fn.Push(BlockScope)

	// The tree must survive without holding the intermediate references.
	if len(g.Children()) != 1 || g.Children()[0] != fn {
		t.Fatal("global deveria reter o escopo de função como filho")
	}
	if len(fn.Children()) != 1 || fn.Children()[0] != blk {
		t.Fatal("função deveria reter o bloco como filho")
	}
	if blk.Parent() != fn || fn.Parent() != g || g.Parent() != nil {
		t.Error("ponteiros de parent estão incorretos")
	}
}

func TestStringDump(t *testing.T) {
	g := NewGlobal[stub]()
	g.Define(&Symbol[stub]{Name: "answer", Kind: Var, Type: stub("i32"), Line: 1})
	fn := g.Push(FunctionScope)
	fn.Define(&Symbol[stub]{Name: "p", Kind: Param, Type: stub("string"), Line: 2})

	want := "global {\n" +
		"  var answer: i32 (linha 1)\n" +
		"  function {\n" +
		"    param p: string (linha 2)\n" +
		"  }\n" +
		"}\n"
	if got := g.String(); got != want {
		t.Errorf("dump incorreto:\n--- esperado ---\n%s\n--- recebido ---\n%s", want, got)
	}
}
