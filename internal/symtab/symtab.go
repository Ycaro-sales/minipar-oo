// Package symtab implements a lexically-nested symbol table for the Minipar
// compiler. Scopes form a retained tree (parent + children) so passes after
// semantic analysis — code generation, for instance — can look names up again
// instead of rebuilding the table.
//
// The table is generic over the payload type a symbol carries (typically the
// resolved type from the semantic analyzer). The only requirement is that the
// payload be printable, which keeps symtab free of any compiler-specific
// import and avoids an import cycle.
package symtab

import (
	"fmt"
	"strings"
)

// Printable is the constraint on a symbol's payload: it must render itself for
// the debug dump. The semantic analyzer's *Type satisfies it.
type Printable interface{ String() string }

// Kind distinguishes what a name refers to, so callers can reject nonsensical
// uses (e.g. calling a variable, indexing a function).
type Kind int

const (
	Var Kind = iota
	Param
	Func
	Class
	Interface
	Field
	Method
)

func (k Kind) String() string {
	switch k {
	case Var:
		return "var"
	case Param:
		return "param"
	case Func:
		return "func"
	case Class:
		return "class"
	case Interface:
		return "interface"
	case Field:
		return "field"
	case Method:
		return "method"
	default:
		return "unknown"
	}
}

// ScopeKind labels a scope's role, both for clearer rules and for diagnostics.
type ScopeKind int

const (
	GlobalScope ScopeKind = iota
	FunctionScope
	ClassScope
	BlockScope
)

func (k ScopeKind) String() string {
	switch k {
	case GlobalScope:
		return "global"
	case FunctionScope:
		return "function"
	case ClassScope:
		return "class"
	case BlockScope:
		return "block"
	default:
		return "unknown"
	}
}

// Symbol is a single named entry in a scope.
type Symbol[T Printable] struct {
	Name string
	Kind Kind
	Type T
	Line int
}

// Scope is one level of a lexically-nested symbol table. Scopes are retained in
// a tree: a parent keeps references to the children pushed from it.
type Scope[T Printable] struct {
	Kind     ScopeKind
	parent   *Scope[T]
	children []*Scope[T]
	symbols  map[string]*Symbol[T]
	order    []string // names in insertion order, for deterministic iteration
}

// NewGlobal returns a fresh root scope of kind GlobalScope.
func NewGlobal[T Printable]() *Scope[T] {
	return &Scope[T]{Kind: GlobalScope, symbols: map[string]*Symbol[T]{}}
}

// Push creates a child scope of the given kind, links it to this scope, and
// returns it. The child is retained in the tree; popping is just moving a
// cursor back to the parent, so the structure survives the analysis walk.
func (s *Scope[T]) Push(kind ScopeKind) *Scope[T] {
	child := &Scope[T]{Kind: kind, parent: s, symbols: map[string]*Symbol[T]{}}
	s.children = append(s.children, child)
	return child
}

// Define inserts sym into this scope. It returns the existing symbol (and
// false) when the name is already bound *in this same scope* — shadowing a name
// from an enclosing scope is allowed and succeeds.
func (s *Scope[T]) Define(sym *Symbol[T]) (*Symbol[T], bool) {
	if prev, ok := s.symbols[sym.Name]; ok {
		return prev, false
	}
	s.symbols[sym.Name] = sym
	s.order = append(s.order, sym.Name)
	return sym, true
}

// Resolve looks the name up in this scope and every enclosing scope.
func (s *Scope[T]) Resolve(name string) *Symbol[T] {
	for sc := s; sc != nil; sc = sc.parent {
		if sym, ok := sc.symbols[name]; ok {
			return sym
		}
	}
	return nil
}

// ResolveLocal looks the name up in this scope only.
func (s *Scope[T]) ResolveLocal(name string) *Symbol[T] {
	return s.symbols[name]
}

// Symbols returns this scope's symbols in insertion order.
func (s *Scope[T]) Symbols() []*Symbol[T] {
	out := make([]*Symbol[T], 0, len(s.order))
	for _, name := range s.order {
		out = append(out, s.symbols[name])
	}
	return out
}

// Children returns the scopes pushed from this one, in creation order.
func (s *Scope[T]) Children() []*Scope[T] { return s.children }

// Parent returns the enclosing scope, or nil for the global scope.
func (s *Scope[T]) Parent() *Scope[T] { return s.parent }

// String renders the scope and its subtree as an indented dump, useful for
// debugging and golden tests.
func (s *Scope[T]) String() string {
	var b strings.Builder
	s.dump(&b, 0)
	return b.String()
}

func (s *Scope[T]) dump(b *strings.Builder, depth int) {
	indent := strings.Repeat("  ", depth)
	fmt.Fprintf(b, "%s%s {\n", indent, s.Kind)
	for _, sym := range s.Symbols() {
		fmt.Fprintf(b, "%s  %s %s: %s (linha %d)\n",
			indent, sym.Kind, sym.Name, sym.Type.String(), sym.Line)
	}
	for _, child := range s.children {
		child.dump(b, depth+1)
	}
	fmt.Fprintf(b, "%s}\n", indent)
}
