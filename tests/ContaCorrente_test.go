package tests

import (
	"banco/contas"
	"testing"
)

func TestContaCorrente_Sacar(t *testing.T) {
	conta := contas.ContaCorrente{}
	conta.Depositar(100.0)

	// Teste para saque válido
	if !conta.Sacar(50.0) {
		t.Error("Falha no teste de saque válido")
	}

	// Teste para saque inválido (valor negativo)
	if conta.Sacar(-10.0) {
		t.Error("Falha no teste de saque inválido (valor negativo)")
	}

	// Teste para saque inválido (saldo insuficiente)
	if conta.Sacar(200.0) {
		t.Error("Falha no teste de saque inválido (saldo insuficiente)")
	}
}

func TestContaCorrente_Depositar(t *testing.T) {
	conta := contas.ContaCorrente{}
	conta.Depositar(100.0)

	if !conta.Depositar(50.0) {
		t.Error("Falha no teste de depósito válido")
	}

	if conta.Depositar(-10.0) {
		t.Error("Falha no teste de depósito inválido (valor negativo)")
	}
}

func TestContaCorrente_Transferir(t *testing.T) {
	contaSaldoOrigem := contas.ContaCorrente{}
	contaSaldoDestino := contas.ContaCorrente{}

	contaSaldoOrigem.Depositar(100.0)
	contaSaldoDestino.Depositar(100.0)

	if !contaSaldoOrigem.Transferir(100, &contaSaldoDestino) {
		t.Error("Falha no teste de transferirvalor valido")
	}

	if contaSaldoOrigem.Transferir(100, &contaSaldoDestino) {
		t.Error("Falha no teste de depósito inválido (valor negativo)")
	}
}
