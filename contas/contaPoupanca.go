package contas

import (
	"banco/clientes"
	"fmt"
)

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	fmt.Println(mensagemOperacaoRealizadaComSucesso, "ObterSaldo")
	return c.saldo
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) bool {
	if valorDoSaque <= 0 {
		fmt.Println(mensagemValorMenorQueZero, "Sacar")
		return false
	}
	if valorDoSaque > c.ObterSaldo() {
		fmt.Println(mensagemsaldoInsuficiente, "Sacar")
		return false
	}
	c.saldo -= valorDoSaque
	fmt.Println(mensagemOperacaoRealizadaComSucesso, "Sacar")
	return true
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) bool {
	if valorDoDeposito <= 0 {
		fmt.Println(mensagemValorMenorQueZero, "Depositar")
		return false
	}
	c.saldo += valorDoDeposito
	fmt.Println(mensagemOperacaoRealizadaComSucesso, "Depositar")
	return true
}
