package semantic

import "strings"

// Kind groups types into families so type rules don't switch on raw names.
type Kind int

const (
	KindInvalid Kind = iota // result of a type error; suppresses cascading errors
	KindAny                 // "any" — compatible with everything
	KindVoid                // functions with no return
	KindBool
	KindChar
	KindString
	KindInt   // i8..i64, u8..u64, "int"
	KindFloat // f16..f64, "float"
	KindClass
	KindInterface
	KindFunc
	KindArray
	KindTuple // immutable fixed-size sequence with per-position element types
	KindChan
)

// Type is the canonical, resolved type of an expression or symbol.
// AST nodes carry types as raw strings; the analyzer resolves them to *Type.
type Type struct {
	Name   string  // canonical name, used in error messages
	Kind   Kind
	Elem   *Type   // element type, for KindArray
	Elems  []*Type // per-position element types, for KindTuple
	Params []*Type // parameter types, for KindFunc
	Return *Type   // return type, for KindFunc
}

// Predeclared, shared instances.
var (
	tInvalid = &Type{Name: "<inválido>", Kind: KindInvalid}
	tAny     = &Type{Name: "any", Kind: KindAny}
	tVoid    = &Type{Name: "void", Kind: KindVoid}
	tBool    = &Type{Name: "bool", Kind: KindBool}
	tChar    = &Type{Name: "char", Kind: KindChar}
	tString  = &Type{Name: "string", Kind: KindString}
	tChan    = &Type{Name: "chan", Kind: KindChan}
)

// primitives maps every spellable primitive (including aliases) to its Type.
var primitives = map[string]*Type{
	"any":    tAny,
	"void":   tVoid,
	"bool":   tBool,
	"char":   tChar,
	"string": tString,
	"str":    tString,
	"int":    {Name: "int", Kind: KindInt},
	"i8":     {Name: "i8", Kind: KindInt},
	"i16":    {Name: "i16", Kind: KindInt},
	"i32":    {Name: "i32", Kind: KindInt},
	"i64":    {Name: "i64", Kind: KindInt},
	"u8":     {Name: "u8", Kind: KindInt},
	"u16":    {Name: "u16", Kind: KindInt},
	"u32":    {Name: "u32", Kind: KindInt},
	"u64":    {Name: "u64", Kind: KindInt},
	"float":  {Name: "float", Kind: KindFloat},
	"f16":    {Name: "f16", Kind: KindFloat},
	"f32":    {Name: "f32", Kind: KindFloat},
	"f64":    {Name: "f64", Kind: KindFloat},
}

func (t *Type) isNumeric() bool { return t.Kind == KindInt || t.Kind == KindFloat }

// assignable reports whether a value of type src may be stored where dst is expected.
// Invalid types are treated as compatible so one type error doesn't cascade.
func assignable(dst, src *Type) bool {
	if dst == nil || src == nil || dst.Kind == KindInvalid || src.Kind == KindInvalid {
		return true
	}
	if dst.Kind == KindAny || src.Kind == KindAny {
		return true
	}
	switch dst.Kind {
	case KindInt, KindFloat:
		// Numeric widening is permitted; an int literal fits a float slot.
		return src.isNumeric()
	case KindArray:
		return src.Kind == KindArray && assignable(dst.Elem, src.Elem)
	case KindTuple:
		if src.Kind != KindTuple || len(dst.Elems) != len(src.Elems) {
			return false
		}
		for i := range dst.Elems {
			if !assignable(dst.Elems[i], src.Elems[i]) {
				return false
			}
		}
		return true
	case KindInterface:
		// A class assigned to an interface must implement it; that structural
		// check happens in the analyzer (it owns the class table). Same-named
		// interface is trivially fine.
		return src.Name == dst.Name || src.Kind == KindClass
	default:
		return dst.Kind == src.Kind && dst.Name == src.Name
	}
}

// String renders a type for diagnostics.
func (t *Type) String() string {
	if t == nil {
		return "<nil>"
	}
	switch t.Kind {
	case KindArray:
		return "[" + t.Elem.String() + "]"
	case KindTuple:
		parts := make([]string, len(t.Elems))
		for i, e := range t.Elems {
			parts[i] = e.String()
		}
		return "(" + strings.Join(parts, ", ") + ")"
	case KindFunc:
		ps := make([]string, len(t.Params))
		for i, p := range t.Params {
			ps[i] = p.String()
		}
		return "func(" + strings.Join(ps, ", ") + ") -> " + t.Return.String()
	default:
		return t.Name
	}
}
