package main

import "fmt"

// buscarPreco procura um produto no map e retorna o preço junto de um
// booleano indicando se ele foi encontrado (comma ok idiom).
func buscarPreco(produtos map[string]float64, nome string) (float64, bool) {
	preco, ok := produtos[nome]
	return preco, ok
}

// mostrarBusca faz a busca e imprime a mensagem apropriada para cada caso.
func mostrarBusca(produtos map[string]float64, nome string) {
	if preco, ok := buscarPreco(produtos, nome); ok {
		fmt.Printf("   %-12s encontrado! Preço: R$ %.2f\n", nome, preco)
	} else {
		fmt.Printf("   %-12s não encontrado no estoque.\n", nome)
	}
}

func exercicio2() {
	// Cadastro dos produtos.
	produtos := map[string]float64{
		"café":   18.90,
		"arroz":  24.50,
		"feijão": 9.75,
		"açúcar": 4.29,
	}
	fmt.Println("1) Produtos cadastrados:")
	for nome, preco := range produtos {
		fmt.Printf("   %-12s R$ %.2f\n", nome, preco)
	}

	// Busca por um produto existente e por um inexistente.
	fmt.Println("2) Buscas:")
	mostrarBusca(produtos, "arroz")
	mostrarBusca(produtos, "chocolate")

	// Remoção com delete e nova busca.
	fmt.Println("3) Removendo \"feijão\" com delete():")
	mostrarBusca(produtos, "feijão")
	delete(produtos, "feijão")
	mostrarBusca(produtos, "feijão")
	fmt.Printf("   Restaram %d produtos.\n", len(produtos))
}
