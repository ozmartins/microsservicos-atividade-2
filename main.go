package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// limparTela move o cursor para o topo e apaga o conteúdo do terminal
// usando sequências de escape ANSI.
func limparTela() {
	fmt.Print("\033[H\033[2J")
}

// pausar espera o ENTER para que a saída do exercício possa ser lida
// antes de o menu limpar a tela novamente.
func pausar(leitor *bufio.Scanner) {
	fmt.Print("\nPressione ENTER para voltar ao menu...")
	leitor.Scan()
}

func main() {
	leitor := bufio.NewScanner(os.Stdin)

	for {
		limparTela()
		fmt.Println("=== Atividade 2 ===")
		fmt.Println("1) Exercício 1 - Slices")
		fmt.Println("2) Exercício 2 - Map de produtos")
		fmt.Println("3) Exercício 3 - Structs de livros")
		fmt.Println("4) Exercício 4 - Conta bancária")
		fmt.Println("0) Sair")
		fmt.Print("Escolha uma opção: ")

		if !leitor.Scan() {
			// Fim da entrada (Ctrl+D ou pipe encerrado).
			fmt.Println()
			return
		}

		fmt.Println()
		switch strings.TrimSpace(leitor.Text()) {
		case "1":
			exercicio1()
		case "2":
			exercicio2()
		case "3":
			exercicio3()
		case "4":
			exercicio4()
		case "0":
			fmt.Println("Até mais!")
			return
		default:
			fmt.Println("Opção inválida.")
		}

		pausar(leitor)
	}
}
