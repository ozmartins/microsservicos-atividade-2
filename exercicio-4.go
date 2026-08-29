package main

import "fmt"

// ContaBancaria representa a conta de um titular.
type ContaBancaria struct {
	Titular string
	Saldo   float64
}

// Depositar adiciona valor ao saldo da conta.
func (c *ContaBancaria) Depositar(valor float64) {
	c.Saldo += valor
}

// Sacar subtrai valor do saldo. Se o valor for maior que o saldo disponível,
// retorna um erro e o saldo permanece inalterado.
func (c *ContaBancaria) Sacar(valor float64) error {
	if valor > c.Saldo {
		return fmt.Errorf("saldo insuficiente: tentativa de sacar R$ %.2f com saldo de R$ %.2f", valor, c.Saldo)
	}
	c.Saldo -= valor
	return nil
}

// tentarSaque executa o saque e imprime o resultado, tratando o erro.
func tentarSaque(c *ContaBancaria, valor float64) {
	if err := c.Sacar(valor); err != nil {
		fmt.Printf("   Saque de R$ %8.2f -> falhou: %v\n", valor, err)
		return
	}
	fmt.Printf("   Saque de R$ %8.2f -> ok. Saldo: R$ %.2f\n", valor, c.Saldo)
}

func exercicio4() {
	conta := ContaBancaria{Titular: "Oseias", Saldo: 0}
	fmt.Printf("1) Conta criada para %s. Saldo inicial: R$ %.2f\n", conta.Titular, conta.Saldo)

	// 2. Depósitos.
	fmt.Println("2) Depósitos:")
	for _, valor := range []float64{500.00, 250.50} {
		conta.Depositar(valor)
		fmt.Printf("   Depósito de R$ %8.2f -> Saldo: R$ %.2f\n", valor, conta.Saldo)
	}

	// 3. Saques, incluindo um que deve falhar.
	fmt.Println("3) Saques:")
	tentarSaque(&conta, 200.00)
	tentarSaque(&conta, 1000.00) // maior que o saldo: deve falhar
	tentarSaque(&conta, 50.50)

	fmt.Printf("4) Saldo final de %s: R$ %.2f\n", conta.Titular, conta.Saldo)
}
