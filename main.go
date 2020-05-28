package main

import (
	"fmt"
	"os"

	"github.com/alura/orientacao_objeto/contas"
)

func main() {
	cliente := contas.ContaCorrente{}
	cliente.Titular = "Thiago"
	cliente.Saldo = 1000

	contaDestino := contas.ContaCorrente{Titular: "Cliente Fake da silva", Saldo: 500}

	for {
		fmt.Println("------------BANCO THG------------")
		fmt.Println("Bem vindo", cliente.Titular)
		contas.Menu()
		comando := contas.InsereMenu()
		switch comando {
		case 1:
			fmt.Println("--SELECIONADO SAQUE--")
			fmt.Println("INSIRA O VALOR DE SAQUE:")
			valorsaque := contas.InsereSaque()

			fmt.Println(cliente.CalculaSaque(valorsaque))
			fmt.Println("Saldo final -> R$ ", cliente.Saldo)
			os.Exit(1)

		case 2:
			fmt.Println("--SELECIONADO DEPÓSITO--")
			fmt.Println("INSIRA O VALOR DO DEPÓSITO:")
			deposito := contas.InsereDeposito()
			fmt.Println(cliente.CalculaDeposito(deposito), cliente.Saldo)
			os.Exit(1)
		case 3:
			fmt.Println("--SELECIONADO SALDO--")
			fmt.Println("SEU SALDO É : R$ ", cliente.Saldo)
			os.Exit(1)
		case 4:
			fmt.Println("--SELECIONADO TRANSFERENCIA--")
			fmt.Println("INSIRA O VALOR DA TRANSFERENCIA:")
			transferencia := contas.InsereTransferencia()
			cliente.CalculaTransferencia(transferencia, &contaDestino)
			fmt.Println("TRANSFERIDO : R$ ", transferencia, " para:", contaDestino.Titular, ", Seu Saldo final agora é:", cliente.Saldo)
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
