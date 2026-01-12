package handlers

import (
	"CRUD_Go/services"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Estilos de cores para o terminal
var menuStyle = color.New(color.FgYellow)
var errorStyle = color.New(color.FgRed)

// Função auxiliar privada para ler e limpar qualquer entrada do teclado
// Centralizar isso resolve o problema do Scanln e limpa os espaços vazios
func lerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin) // Cria um scanner para ler a entrada padrão (teclado)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text()) // Retorna a linha lida, removendo espaços em branco extras
	}
	return "" // Retorna string vazia em caso de erro
}

// CriarPessoaTerminal, esta função simula a entrada de dados via console
func CriarPessoaTerminal() {

	menuStyle.Println("\n--------- CADASTRAR NOVA PESSOA ---------")

	fmt.Print("Nome: ")
	nome := lerEntrada()

	fmt.Print("CPF: ")
	cpf := lerEntrada()

	fmt.Print("Data Nascimento: ")
	dataNascimento := lerEntrada()

	fmt.Print("Telefone: ")
	tel := lerEntrada()

	fmt.Print("Email: ")
	email := lerEntrada()

	// Chama o service enviando as strings capturadas
	pessoa := services.CriarPessoaService(nome, cpf, dataNascimento, tel, email)

	fmt.Printf("\nSucesso! Pessoa cadastrada com ID: %s\n", pessoa.ID)
}

// ListarPessoasTerminal exibe todas as pessoas cadastradas
func ListarPessoasTerminal() {
	pessoas := services.ObterPessoasService()

	menuStyle.Println("\n--------- LISTAR TODAS AS PESSOAS ---------")
	if len(pessoas) == 0 {
		fmt.Println("\nNenhuma pessoa cadastrada.")
		return
	}

	for _, p := range pessoas {
		fmt.Printf("ID: %s | Nome: %s | CPF: %s | Data Nascimento: %s | Telefone: %s | Email: %s\n",
			p.ID, p.Nome, p.Cpf, p.DataNascimento, p.Telefone, p.Email)
	}
}

// BuscarPessoaPorIDTerminal pede um ID e exibe os detalhes da pessoa
func BuscarPessoaPorIDTerminal() {

	menuStyle.Println("\n--------- BUSCAR PESSOA POR ID ---------")
	fmt.Print("\nDigite o ID da pessoa que deseja buscar: ")
	id := lerEntrada()

	pessoa := services.ObterPessoaPorIDService(id)

	if pessoa == nil {
		errorStyle.Println("Erro: Pessoa não encontrada ou ID inválido.")
		return
	}

	fmt.Printf("\n--- Detalhes ---\nNome: %s\nCPF: %s\nData Nascimento: %s\nTelefone: %s\nEmail: %s\n",
		pessoa.Nome, pessoa.Cpf, pessoa.DataNascimento, pessoa.Telefone, pessoa.Email)
}

// MostrarUltimoIDTerminal exibe o ID da última pessoa criada
func MostrarUltimoIDTerminal() {
	ultimoID := services.ObterUltimoIDService()

	menuStyle.Println("\n--------- VER ID ÚLTIMO CADASTRO ---------")
	fmt.Printf("\nO ID da última pessoa criada é: %s\n", ultimoID)
}

// AtualizarPessoaTerminal solicita o ID e os novos dados
func AtualizarPessoaTerminal() {

	menuStyle.Println("\n--------- ATUALIZAR DADOS DE PESSOA ---------")
	fmt.Print("\nDigite o ID da pessoa que deseja ATUALIZAR: ")
	id := lerEntrada()

	fmt.Println("Digite os novos dados:")

	fmt.Print("Telefone: ")
	tel := lerEntrada()

	fmt.Print("Email: ")
	email := lerEntrada()

	pessoa := services.AtualizarPessoaService(id, tel, email) // Chama o service para atualizar

	if pessoa == nil {
		errorStyle.Println("\nErro: Não foi possível atualizar. Verifique o ID.")
	} else {
		fmt.Println("\nDados atualizados com sucesso!", pessoa.Nome)
	}
}

// DeletarPessoaTerminal remove uma pessoa pelo ID
func DeletarPessoaTerminal() {

	menuStyle.Println("\n--------- DELETAR PESSOA ---------")
	fmt.Print("\nDigite o ID da pessoa que deseja DELETAR: ")
	id := lerEntrada()

	sucesso := services.DeletarPessoaService(id)

	if sucesso {
		fmt.Println("\nPessoa removida com sucesso!")
	} else {
		errorStyle.Println("Erro: Pessoa não encontrada ou ID inválido.")
	}
}
