package main

import (
	"encoding/json"
	"fmt"
	"minipar/parser"
)

func main() {
	// O TESTE DO ANALISADOR SINTÁTICO DA LINGUAGEM MINIPAR 2025.1
	codigoTeste := `
		# 1. Declarações Globais (Canais e Variáveis)
		c_channel canal_comunicacao (pc_cliente, pc_servidor);
		string nome_projeto = "Interpretador MiniPar"
		bool finalizado = true

		# 2. Funções Globais (Tipo de retorno DEPOIS dos parênteses, como na sua BNF)
		func inicializar(a, b) void {
			print("Sistema iniciado");
			return;
		}

		# 3. Orientação a Objetos (Classes e Herança)
		class Entidade {
			number id = 1
		}

		class Neuronio extends Entidade {
			number peso = 10
			number bias = 5
			
			func propagar() number {
				# 4. Matemática Extrema (Precedência: o parêntese e a divisão resolvem primeiro)
				number ativacao = peso * 2 + bias - (10 / 2)
				
				# 5. Controle de Fluxo (If sem else)
				if (ativacao >= 15) {
					print("Neuronio ativado!", ativacao);
					return ativacao;
				}
				
				return 0;
			}
		}

		# 6. Bloco Sequencial e Loops
		seq {
			number contador = 0
			while (contador < 3) {
				print("Processando...", contador);
				contador = contador + 1
			}
		}

		# 7. Bloco Paralelo e Lógica Booleana
		par {
			number x = 10 * (2 + 3)
			
			if (x == 50 and finalizado == true) {
				print("Executando thread paralela!");
			}
		}
	`
	astRaiz, erros := parser.ParseProgram(codigoTeste)

	if len(erros) > 0 {
		fmt.Println("ERROS SINTÁTICOS ENCONTRADOS:")
		for _, err := range erros {
			fmt.Printf("  - %s\n", err)
		}
		return
	}

	fmt.Println("AST GERADA COM SUCESSO!\n")

	astJSON, err := json.MarshalIndent(astRaiz, "", "  ")
	if err != nil {
		fmt.Println("Erro ao formatar a AST para visualização:", err)
		return
	}

	fmt.Println(string(astJSON))
	fmt.Println("Compilação finalizada.")
}
