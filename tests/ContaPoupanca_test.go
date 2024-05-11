package tests

import (
	"banco/contas"
	"testing"
)

func TestContaPoupanca_Sacar(t *testing.T) {
	conta := contas.ContaPoupanca{}
	conta.Depositar(100.0)

	if !conta.Sacar(50.0) {
		t.Error("Falha no teste de saque válido")
	}

	if conta.Sacar(-10.0) {
		t.Error("Falha no teste de saque inválido (valor negativo)")
	}

	if conta.Sacar(200.0) {
		t.Error("Falha no teste de saque inválido (saldo insuficiente)")
	}
}

func TestContaPoupanca_Depositar(t *testing.T) {
	conta := contas.ContaPoupanca{}
	conta.Depositar(100.0)

	if !conta.Depositar(50.0) {
		t.Error("Falha no teste de depósito válido")
	}

	if conta.Depositar(-10.0) {
		t.Error("Falha no teste de depósito inválido (valor negativo)")
	}
}
