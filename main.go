package main

import (
	"CRUD_Go/handlers"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {

	// Criando estilos reutilizáveis
	headerStyle := color.New(color.FgGreen, color.Bold)
	menuStyle := color.New(color.FgYellow, color.Bold)
	errorStyle := color.New(color.FgRed, color.Bold)
	titleStyle := color.New(color.FgCyan, color.Bold)
	messageStyle := color.New(color.FgHiMagenta, color.Bold)

	headerStyle.Println("\n\n**********////  SISTEMA CRUD_Go INICIADO  ////*********\n")
	for {

		titleStyle.Println("\n===========================================")
		titleStyle.Println("            GESTÃO DE PESSOAS     ")
		titleStyle.Println("===========================================")
		menuStyle.Println("	1 - Cadastrar Nova Pessoa")
		menuStyle.Println("	2 - Listar Todas as Pessoas")
		menuStyle.Println("	3 - Buscar Pessoa por ID")
		menuStyle.Println("	4 - Ver ID do Último Cadastro")
		menuStyle.Println("	5 - Atualizar Dados de uma Pessoa")
		menuStyle.Println("	6 - Deletar uma Pessoa")
		menuStyle.Println("	0 - Sair do Programa")
		titleStyle.Println("===========================================")
		fmt.Print("Escolha uma opção: ")

		var opcao string
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			handlers.CriarPessoaTerminal()
		case "2":
			handlers.ListarPessoasTerminal()
		case "3":
			handlers.BuscarPessoaPorIDTerminal()
		case "4":
			handlers.MostrarUltimoIDTerminal()
		case "5":
			handlers.AtualizarPessoaTerminal()
		case "6":
			handlers.DeletarPessoaTerminal()
		case "0":
			messageStyle.Println("\nEncerrando o sistema... Até logo!")
			os.Exit(0) // Finaliza o programa
		default:
			errorStyle.Println("\n[ERRO] Opção inválida! Tente novamente.")
		}
	}
}
