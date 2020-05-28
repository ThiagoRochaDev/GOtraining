package main

import (
	"fmt"
	"os"
)

type ContaCorrente struct {
	titular string
	conta   int
	agencia int
	saldo   float64
}

func main() {
	cliente := ContaCorrente{}
	cliente.titular = "Thiago"
	cliente.saldo = 1000

	contaDestino := ContaCorrente{titular: "Cliente Fake da silva", saldo: 500}

	for {
		fmt.Println("------------BANCO THG------------")
		fmt.Println("Bem vindo", cliente.titular)
		menu()
		comando := insereMenu()
		switch comando {
		case 1:
			fmt.Println("--SELECIONADO SAQUE--")
			fmt.Println("INSIRA O VALOR DE SAQUE:")
			valorsaque := insereSaque()

			fmt.Println(cliente.calculaSaque(valorsaque))
			fmt.Println("Saldo final -> R$ ", cliente.saldo)
			os.Exit(1)

		case 2:
			fmt.Println("--SELECIONADO DEPÓSITO--")
			fmt.Println("INSIRA O VALOR DO DEPÓSITO:")
			deposito := insereDeposito()
			fmt.Println(cliente.calculaDepoisto(deposito), cliente.saldo)
			os.Exit(1)
		case 3:
			fmt.Println("--SELECIONADO SALDO--")
			fmt.Println("SEU SALDO É : R$ ", cliente.saldo)
			os.Exit(1)
		case 4:
			fmt.Println("--SELECIONADO TRANSFERENCIA--")
			fmt.Println("INSIRA O VALOR DA TRANSFERENCIA:")
			transferencia := insereTransferencia()
			cliente.calculaTransferencia(transferencia, &contaDestino)
			fmt.Println("TRANSFERIDO : R$ ", transferencia, " para:", contaDestino.titular, ", Seu Saldo final agora é:", cliente.saldo)
			os.Exit(1)
		case 5:
			fmt.Println("SAINDO..")
			os.Exit(1)
		default:
			fmt.Println("OPÇÃO INVÁLIDA")
			os.Exit(-1)
		}

	}

}

func menu() { // escreve menu
	fmt.Println("[1] - Para sacar")
	fmt.Println("[2] - Para depositar")
	fmt.Println("[3] - Para consultar saldo")
	fmt.Println("[4] - Para Transferir")
	fmt.Println("[5] - Para sair")
}

func insereMenu() int { // recebe valor do menu
	var valor int
	fmt.Scan(&valor)
	return valor
}

func insereSaque() float64 { // recebe valor para saque
	var saque float64
	fmt.Scan(&saque)
	return saque
}

func insereDeposito() float64 { // recebe valor para deposito
	var valor float64
	fmt.Scan(&valor)
	return valor
}

func insereTransferencia() float64 { // recebe valor para deposito
	var valor float64
	fmt.Scan(&valor)
	return valor
}

func (c *ContaCorrente) calculaSaque(valorsaque float64) string { // calcula meu saque - saldo
	saldofinal := valorsaque <= c.saldo && valorsaque > 0
	if saldofinal {
		c.saldo -= valorsaque
		return "Saque realizado com sucesso !"

	} else {
		return "Sem saldo para saque"
	}

}

func (c *ContaCorrente) calculaDepoisto(deposito float64) string { // calcula meu deposito + saldo
	novosaldo := deposito+c.saldo > 0

	if novosaldo {
		c.saldo += deposito
		return "Valor depositado com sucesso !"
	} else {
		return "erro ao depositar"
	}
}

func (c *ContaCorrente) calculaTransferencia(transferencia float64, contaDestino *ContaCorrente) bool {
	if transferencia < c.saldo && transferencia > 0 {
		c.saldo -= transferencia
		contaDestino.calculaDepoisto(transferencia)
		return true
	} else {
		return false
	}

}
