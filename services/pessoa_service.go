package services

import (
	"CRUD_Go/models"     // Importa a definição da estrutura Pessoa
	"CRUD_Go/repository" // Importa o pacote repository para acessar as funções de manipulação de dados

	"github.com/google/uuid" // Importa a biblioteca para manipular IDs únicos (UUID)
)

// CriarPessoaService é a função de serviço que orquestra a criação de uma nova pessoa.
// Ela recebe os dados necessários, chama o repositório para criar a pessoa e retorna o resultado.
func CriarPessoaService(nome, cpf, dataNascimento, telefone, email string) models.Pessoa {
	// Chama a função do repositório para criar a pessoa e obter o objeto completo.
	pessoaCriada := repository.CriarPessoa(nome, cpf, dataNascimento, telefone, email)

	// Retorna o objeto da pessoa criada para quem chamou o serviço.
	return pessoaCriada
}

// ObterPessoasService é a função de serviço que recupera todas as pessoas armazenadas.
func ObterPessoasService() []models.Pessoa {
	// Chama a função do repositório para obter a lista completa de pessoas.
	pessoas := repository.ObterPessoas()

	// Retorna a lista de pessoas para quem chamou o serviço.
	return pessoas
}

// ObterPessoaPorIDService é a função de serviço que recupera uma pessoa pelo seu ID.
func ObterPessoaPorIDService(id string) *models.Pessoa {
	// Converte a string ID para o tipo uuid.UUID
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil // Retorna nil se a conversão falhar
	}

	// Chama a função do repositório para obter a pessoa pelo ID.
	pessoa := repository.ObterPessoaPorID(uid)

	// Retorna a pessoa encontrada ou nil para quem chamou o serviço.
	return pessoa
}

// ObterUltimoIDService é a função de serviço que recupera o ID da última pessoa criada.
func ObterUltimoIDService() uuid.UUID {
	// Chama a função do repositório para obter o último ID.
	ultimoID := repository.ObterUltimoID()

	// Retorna o último ID para quem chamou o serviço.
	return ultimoID
}

// AtualizarPessoaService é a função de serviço que atualiza os dados de uma pessoa existente.
func AtualizarPessoaService(id string, telefone, email string) *models.Pessoa {
	// Converte a string ID para o tipo uuid.UUID
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil // Retorna nil se a conversão falhar
	}

	// Chama a função do repositório para atualizar a pessoa.
	pessoaAtualizada := repository.AtualizarPessoa(uid, telefone, email)

	return pessoaAtualizada // Retorna a pessoa atualizada para quem chamou o serviço.
}

// DeletarPessoaService é a função de serviço que remove uma pessoa pelo seu ID.
func DeletarPessoaService(id string) bool {
	// Converte a string ID para o tipo uuid.UUID
	uid, err := uuid.Parse(id)
	if err != nil {
		return false // Retorna false se a conversão falhar
	}

	// Chama a função do repositório para deletar a pessoa.
	return repository.DeletarPessoa(uid)
}
