package main

import "fmt"

func somaSlice(numeros []int) int {
	soma := 0
	for _, n := range numeros {
		soma += n
	}
	return soma
}

func exercicio1() {
	numeros := make([]int, 0, 10)
	for i := 1; i <= 10; i++ {
		numeros = append(numeros, i)
	}
	fmt.Println("1) Slice original:", numeros)

	primeiros := numeros[:5]
	fmt.Println("2) Primeiros 5:  ", primeiros)

	copia := make([]int, len(primeiros))
	copy(copia, primeiros)
	copia[0] = 99
	fmt.Println("3) Cópia alterada:", copia)
	fmt.Println("   Sub-slice:     ", primeiros)
	fmt.Println("   Original:      ", numeros)

	fmt.Printf("4) Soma da slice original: %d\n", somaSlice(numeros))
}
