package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

type verificarConta interface {
	Sacar(valor float64) bool
	Depositar(valor float64) bool
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func Depositar(conta verificarConta, valorDoBoleto float64) bool {
	return conta.Depositar(valorDoBoleto)
}

func main() {
	giovanni := clientes.Titular{
		Nome: "Giovanni Augusto Fogaça",
		CPF:  "269.343.220-02", Profissao: "Vendedor(a)"}

	contaPoupancaGiovanni := contas.ContaPoupanca{Titular: giovanni}
	Depositar(&contaPoupancaGiovanni, 500)
	PagarBoleto(&contaPoupancaGiovanni, 50)
	fmt.Println(contaPoupancaGiovanni)
}
