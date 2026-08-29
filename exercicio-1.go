package main

import "fmt"

// somaSlice retorna a soma de todos os elementos da slice.
func somaSlice(numeros []int) int {
	soma := 0
	for _, n := range numeros {
		soma += n
	}
	return soma
}

func exercicio1() {
	// 1. Slice vazia com make e preenchida com append.
	numeros := make([]int, 0, 10)
	for i := 1; i <= 10; i++ {
		numeros = append(numeros, i)
	}
	fmt.Println("1) Slice original:", numeros)

	// 2. Sub-slice com os 5 primeiros elementos.
	primeiros := numeros[:5]
	fmt.Println("2) Primeiros 5:  ", primeiros)

	// 3. Cópia independente com copy.
	copia := make([]int, len(primeiros))
	copy(copia, primeiros)
	copia[0] = 99
	fmt.Println("3) Cópia alterada:", copia)
	fmt.Println("   Sub-slice:     ", primeiros)
	fmt.Println("   Original:      ", numeros)

	// 4. Soma dos elementos da slice original.
	fmt.Printf("4) Soma da slice original: %d\n", somaSlice(numeros))
}
