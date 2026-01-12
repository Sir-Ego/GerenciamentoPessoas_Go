package repository

import (
	"CRUD_Go/models" // Importa a definição da estrutura Pessoa

	"github.com/google/uuid" // Importa a biblioteca para gerar IDs únicos (UUID)
)

// pessoas é um "slice" (lista dinâmica) que armazena todas as pessoas na memória.
// Funciona como um banco de dados temporário enquanto o programa estiver rodando.
var pessoas []models.Pessoa

// ultimoID armazena o UUID da última pessoa criada.
// Isso permite acesso instantâneo ao último registro sem precisar buscar na lista.
var ultimoID uuid.UUID

// CriarPessoa recebe os dados, gera um ID, salva na lista e retorna a pessoa criada.
func CriarPessoa(nome, cpf, dataNascimento, telefone, email string) models.Pessoa {

	novoID := uuid.New() // Gera um identificador único universal (v4) para a nova pessoa.

	// Cria uma nova instância da struct Pessoa (definida na pasta models).
	novaPessoa := models.Pessoa{
		ID:             novoID,
		Nome:           nome,
		Cpf:            cpf,
		DataNascimento: dataNascimento,
		Telefone:       telefone,
		Email:          email,
	}

	pessoas = append(pessoas, novaPessoa) // Adiciona a nova pessoa ao final da lista 'pessoas'.
	// O comando append cria um novo slice com o elemento adicionado.

	// Atualiza a variável global com o ID que acabamos de gerar.
	// Agora, qualquer outra função pode consultar 'ultimoID' para saber quem entrou por último.
	ultimoID = novoID

	return novaPessoa // Retorna o objeto completo da pessoa para quem chamou a função (ex: camada de Service).
}

// ObterUltimoID é uma função auxiliar (opcional) para retornar o valor da variável privada.
func ObterUltimoID() uuid.UUID {
	return ultimoID
}

// ObterPessoas retorna a lista completa de pessoas armazenadas na memória.
func ObterPessoas() []models.Pessoa {
	return pessoas
}

// ObterPessoaPorID busca uma pessoa na lista pelo seu ID e retorna a pessoa encontrada.
func ObterPessoaPorID(id uuid.UUID) *models.Pessoa {
	for i, pessoa := range pessoas { // Percorre a lista de pessoas
		if pessoa.ID == id { // Verifica se o ID corresponde ao da pessoa buscada
			return &pessoas[i] // Retorna o ponteiro para a pessoa encontrada
		}
	}
	return nil // Retorna nil se a pessoa com o ID fornecido não for encontrada
}

// AtualizarPessoa atualiza os dados de uma pessoa existente na lista.
func AtualizarPessoa(id uuid.UUID, telefone, email string) *models.Pessoa {
	for i := range pessoas { // Percorre a lista de pessoas
		if pessoas[i].ID == id {
			pessoas[i].Telefone = telefone
			pessoas[i].Email = email

			return &pessoas[i] // Retorna o ponteiro para a pessoa atualizada
		}
	}
	return nil // Retorna nil se a pessoa com o ID fornecido não for encontrada
}

// DeletarPessoa remove uma pessoa da lista pelo seu ID.
func DeletarPessoa(id uuid.UUID) bool {
	for i, pessoa := range pessoas { // Percorre a lista de pessoas
		if pessoa.ID == id { // Verifica se o ID corresponde ao da pessoa a ser deletada
			pessoas = append(pessoas[:i], pessoas[i+1:]...) // Remove a pessoa da lista usando slicing
			return true                                     // Retorna true indicando que a pessoa foi deletada com sucesso
		}
	}
	return false // Retorna false se a pessoa com o ID fornecido não for encontrada
}
