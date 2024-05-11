package contas

import (
	"banco/clientes"
	"fmt"
)

const (
	mensagemValorMenorQueZero           = "O valor deve ser maior que ZERO"
	mensagemOperacaoRealizadaComSucesso = "Operação realizada com sucesso!"
	mensagemsaldoInsuficiente           = "Saldo insuficiente!"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) ObterSaldo() float64 {
	fmt.Println(mensagemOperacaoRealizadaComSucesso, "ObterSaldo")
	return c.saldo
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) bool {
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

func (c *ContaCorrente) Depositar(valorDoDeposito float64) bool {
	if valorDoDeposito <= 0 {
		fmt.Println(mensagemValorMenorQueZero, "Depositar")
		return false
	}
	c.saldo += valorDoDeposito
	fmt.Println(mensagemOperacaoRealizadaComSucesso, "Depositar")
	return true
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia <= 0 {
		fmt.Println(mensagemValorMenorQueZero, "Transferir")
		return false
	}
	if valorDaTransferencia > c.ObterSaldo() {
		fmt.Println(mensagemsaldoInsuficiente, "Transferir")
		return false
	}
	c.Sacar(valorDaTransferencia)
	contaDestino.Depositar(valorDaTransferencia)
	fmt.Println(mensagemOperacaoRealizadaComSucesso, "Transferir")
	return true
}
