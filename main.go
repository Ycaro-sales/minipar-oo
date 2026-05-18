package main

import (
	"fmt"
	"minipar/ast"
	"minipar/parser"
)

func main() {
	// Um código MiniPar parrudo para testar todas as novas regras!
	codigoTeste := `
		c_channel rede_neural pc_local servidor_nuvem;

		class Neuronio extends Celula { 
			number peso = 0
			bool ativo = false
		}

		func treinarRede() void {
			number epocas = 0
			
			while (epocas < 100) {
				if (ativo == true) {
					epocas = 100;
				}
			}

			par {
				# Thread simulada pelo PAR
			}
		}
	`

	fmt.Println("Compilando código MiniPar...")
	astRaiz, erros := parser.ParseProgram(codigoTeste)

	if len(erros) > 0 {
		fmt.Println("Erros encontrados:")
		for _, err := range erros {
			fmt.Println(err)
		}
		return
	}

	fmt.Println("\n-- AST Gerada com Sucesso --")
	for i, node := range astRaiz.Declarations {
		fmt.Printf("\nNÓ RAÍZ [%d]: %+v\n", i, node)

		// Inspecionando o que tem dentro da Classe
		if classNode, ok := node.(*ast.ClassDecl); ok {
			fmt.Printf("  -> Membros da classe '%s':\n", classNode.Name)
			for j, membro := range classNode.Members {
				fmt.Printf("     [%d] %+v\n", j, membro)
			}
		}

		// Inspecionando o que tem dentro da Função
		if funcNode, ok := node.(*ast.FuncDecl); ok {
			fmt.Printf("  -> Comandos dentro da função '%s':\n", funcNode.Name)
			if funcNode.Body != nil {
				for j, stmt := range funcNode.Body.Statements {
					// Imprimindo cada instrução lida pelo seu parser
					fmt.Printf("     [%d] %+v\n", j, stmt)
				}
			}
		}
	}
}
