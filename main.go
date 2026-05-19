package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Token struct {
	Tipo    string
	Literal string
	Linha   int
	Coluna  int
}

func main() {

	input, erro := antlr.NewFileStream("input.txt")

	if erro != nil {
		fmt.Println("Erro: ", erro)
		return
	}

	lexer := NewLinguagem(input)

	var lista_tokens []Token

	for {

		token := lexer.NextToken()

		if token.GetTokenType() == antlr.TokenEOF {
			break
		}

		nome_tipo := lexer.SymbolicNames[token.GetTokenType()]

		novo_token := Token{
			Tipo:    nome_tipo,
			Literal: token.GetText(),
			Linha:   token.GetLine(),
			Coluna:  token.GetColumn(),
		}

		lista_tokens = append(lista_tokens, novo_token)
	}

	fmt.Println("---------------------------------------------------------")
	fmt.Printf("%-15s | %-15s | %-5s | %-6s\n", "TIPO", "LITERAL", "LINHA", "COLUNA")
	fmt.Println("---------------------------------------------------------")

	for _, t := range lista_tokens {
		fmt.Printf("%-15s | %-15s | %-5d | %-6d\n", t.Tipo, t.Literal, t.Linha, t.Coluna)
	}
	fmt.Println("---------------------------------------------------------")
}
