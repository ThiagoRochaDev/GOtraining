package contas

import (
	"fmt"

	"github.com/alura/orientacao_objeto/clientes"
)

type ContaCorrente struct { // cria estrutura do objeto , (maiusculo permite encontrar em outros pacotes do projeto)
	Titular clientes.Cliente
	Conta   int
	Agencia int
	Saldo   float64
}

func Menu() { // escreve menu
	fmt.Println("[1] - Para sacar")
	fmt.Println("[2] - Para depositar")
	fmt.Println("[3] - Para consultar saldo")
	fmt.Println("[4] - Para Transferir")
	fmt.Println("[5] - Para sair")
}

func InsereMenu() int { // recebe valor do menu
	var valor int
	fmt.Scan(&valor)
	return valor
}

func InsereSaque() float64 { // recebe valor para saque
	var saque float64
	fmt.Scan(&saque)
	return saque
}

func InsereDeposito() float64 { // recebe valor para deposito
	var valor float64
	fmt.Scan(&valor)
	return valor
}

func InsereTransferencia() float64 { // recebe valor para deposito
	var valor float64
	fmt.Scan(&valor)
	return valor
}

func (c *ContaCorrente) CalculaSaque(valorsaque float64) string { // calcula meu saque - saldo
	saldofinal := valorsaque <= c.Saldo && valorsaque > 0
	if saldofinal {
		c.Saldo -= valorsaque
		return "Saque realizado com sucesso !"

	} else {
		return "Sem saldo para saque"
	}

}

func (c *ContaCorrente) CalculaDeposito(deposito float64) string { // calcula meu deposito + saldo
	novosaldo := deposito+c.Saldo > 0

	if novosaldo {
		c.Saldo += deposito
		return "Valor depositado com sucesso !"
	} else {
		return "erro ao depositar"
	}
}

func (c *ContaCorrente) CalculaTransferencia(transferencia float64, contaDestino *ContaCorrente) bool {
	if transferencia < c.Saldo && transferencia > 0 {
		c.Saldo -= transferencia
		contaDestino.CalculaDeposito(transferencia)
		return true
	} else {
		return false
	}

}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.Saldo
}
