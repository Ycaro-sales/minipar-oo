package semantic

import "minipar/internal/symtab"

// The analyzer uses the generic symbol table instantiated with the resolved
// type (*Type) as the payload. These aliases keep call sites terse and hide the
// generic parameter from the rest of the package.
type (
	scope  = symtab.Scope[*Type]
	symbol = symtab.Symbol[*Type]
)
