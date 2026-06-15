package lexer

// Tokenizer is the interface the parser depends on.
// *Lexer satisfies it; tests can substitute a mock.
type Tokenizer interface {
	NextToken() Token
	Line() int
}
