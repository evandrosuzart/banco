# Simulador de Banco em Go

Este é um projeto de estudo em linguagem Go que simula um banco com operações básicas para contas correntes e contas poupança.

## Funcionalidades

### Conta Poupança
- `Depositar(valor float64) bool`: Deposita um valor na conta poupança.
- `ObterSaldo() float64`: Consulta o saldo atual da conta poupança.
- `Sacar(valor float64) bool`: Realiza um saque na conta poupança.
- `PagarBoleto(valorDoBoleto float64)`: Realiza o pagamento de um boleto.

### Conta Corrente
- `Depositar(valor float64) bool`: Deposita um valor na conta corrente.
- `ObterSaldo() float64`: Consulta o saldo atual da conta corrente.
- `Sacar(valor float64) bool`: Realiza um saque na conta corrente.
- `PagarBoleto(valorDoBoleto float64)`: Realiza o pagamento de um boleto.
- `Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool`: Transfere um valor para outra conta corrente.

## Testes Unitários
Os testes unitários estão no pacote `test`. Eles verificam se as funções estão funcionando corretamente e se os resultados estão de acordo com o esperado.

Para executar os testes, use o comando:

```bash
go test ./tests
```


Para executar o projeto, use o comando:

```bash
go ./main.go
```
