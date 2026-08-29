package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	leitor := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
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
	}
}
