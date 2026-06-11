package main

import (
	"encoding/json"
	"fmt"
	"minipar/parser"
)

func main() {
	// Código MiniPar de teste englobando as regras da sua BNF!
	codigoTeste := `
		# 1. Testando Canal (Obrigatório ponto e vírgula)
		c_channel rede_neural (pc_local, servidor_nuvem);

		# 2. Testando Orientação a Objetos
		class Neuronio extends Celula {
			number peso = 10
			
			func treinar() void {
				# 3. Testando precedência matemática: a multiplicação ocorre antes da soma
				peso = peso + 1 * 5
				print(peso);
			}
		}

		# 4. Testando bloco sequencial, loop e condicional
		seq {
			number epocas = 0
			while (epocas < 100) {
				if (peso == 0) {
					peso = 1
				}
				epocas = epocas + 1
			}
		}

		# 5. Testando paralelismo e chamadas simples
		par {
			print("Treinando em paralelo");
			return;
		}
	`

	fmt.Println("========================================")
	fmt.Println(" Iniciando o Compilador MiniPar...")
	fmt.Println("========================================\n")

	// Chama a nossa Facade que agora é 100% manual
	astRaiz, erros := parser.ParseProgram(codigoTeste)

	// Se o Parser encontrou erros de sintaxe (ex: faltou ponto e vírgula), aborta.
	if len(erros) > 0 {
		fmt.Println("❌ ERROS SINTÁTICOS ENCONTRADOS:")
		for _, err := range erros {
			fmt.Printf("  - %s\n", err)
		}
		return
	}

	fmt.Println(" AST GERADA COM SUCESSO!\n")

	// Para visualizar a árvore complexa no terminal sem ser apenas endereços de memória,
	// o jeito mais profisional e limpo no Go é converter a struct para JSON formatado.
	astJSON, err := json.MarshalIndent(astRaiz, "", "  ")
	if err != nil {
		fmt.Println("Erro ao formatar a AST para visualização:", err)
		return
	}

	fmt.Println(string(astJSON))
	fmt.Println("\n========================================")
	fmt.Println("Compilação finalizada.")
}
