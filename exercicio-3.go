package main

import "fmt"

type Livro struct {
	Titulo        string
	Autor         string
	AnoPublicacao int
	Disponivel    bool
}

func contarDisponiveis(livros []Livro) int {
	total := 0
	for _, l := range livros {
		if l.Disponivel {
			total++
		}
	}
	return total
}

func buscarPorAutor(livros []Livro, autor string) []Livro {
	var encontrados []Livro
	for _, l := range livros {
		if l.Autor == autor {
			encontrados = append(encontrados, l)
		}
	}
	return encontrados
}

func descrever(l Livro) string {
	status := "indisponível"
	if l.Disponivel {
		status = "disponível"
	}
	return fmt.Sprintf("%-28s %-22s %d  (%s)", l.Titulo, l.Autor, l.AnoPublicacao, status)
}

func exercicio3() {
	livros := []Livro{
		{Titulo: "Dom Casmurro", Autor: "Machado de Assis", AnoPublicacao: 1899, Disponivel: true},
		{Titulo: "Memórias Póstumas", Autor: "Machado de Assis", AnoPublicacao: 1881, Disponivel: false},
		{Titulo: "Grande Sertão: Veredas", Autor: "Guimarães Rosa", AnoPublicacao: 1956, Disponivel: true},
		{Titulo: "Vidas Secas", Autor: "Graciliano Ramos", AnoPublicacao: 1938, Disponivel: true},
		{Titulo: "São Bernardo", Autor: "Graciliano Ramos", AnoPublicacao: 1934, Disponivel: false},
	}

	fmt.Println("1) Acervo cadastrado:")
	for _, l := range livros {
		fmt.Println("   " + descrever(l))
	}

	fmt.Printf("2) Livros disponíveis: %d de %d\n", contarDisponiveis(livros), len(livros))

	autor := "Machado de Assis"
	encontrados := buscarPorAutor(livros, autor)
	fmt.Printf("3) Livros de %s (%d):\n", autor, len(encontrados))
	for _, l := range encontrados {
		fmt.Println("   " + descrever(l))
	}

	inexistente := buscarPorAutor(livros, "Clarice Lispector")
	fmt.Printf("   Livros de Clarice Lispector: %d\n", len(inexistente))
}
