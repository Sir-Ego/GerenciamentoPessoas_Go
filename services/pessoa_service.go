package services

import (
	"CRUD_Go/models"     // Importa a definição da estrutura Pessoa
	"CRUD_Go/repository" // Importa o pacote repository para acessar as funções de manipulação de dados
	"fmt"

	"github.com/google/uuid" // Importa a biblioteca para manipular IDs únicos (UUID)
)

// CriarPessoaService é a função de serviço que cria uma nova pessoa com validação.
// Ela recebe os dados como strings e retorna a pessoa criada ou um erro de validação.
func CriarPessoaService(nome, cpf, dataNascimento, telefone, email string) (models.Pessoa, error) {

	// Cria uma nova instância da struct Pessoa com os dados fornecidos.
	pessoa := models.Pessoa{
		Nome:           nome,
		Cpf:            cpf,
		DataNascimento: dataNascimento,
		Telefone:       telefone,
		Email:          email,
	}

	// Validação final de pessoa antes de criar
	if err := pessoa.Validate(); err != nil {
		return models.Pessoa{}, err
	}

	// Chama a função do repositório para criar a pessoa e obter o objeto completo.
	pessoaCriada := repository.CriarPessoa(nome, cpf, dataNascimento, telefone, email)

	// Retorna o objeto da pessoa criada para quem chamou o serviço.
	return pessoaCriada, nil
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
func AtualizarPessoaService(id string, telefone, email string) (*models.Pessoa, error) {
	// Converte a string ID para o tipo uuid.UUID
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err // Retorna nil se a conversão falhar
	}

	// Busca a pessoa existente
	pessoaExistente := repository.ObterPessoaPorID(uid)
	if pessoaExistente == nil {
		return nil, fmt.Errorf("Pessoa não encontrada")
	}

	// Valida apenas os campos que podem ser atualizados
	if err := models.ValidarTelefone(telefone); err != nil {
		return nil, err
	}

	if err := models.ValidarEmail(email); err != nil {
		return nil, err
	}

	// Atualiza os os novos campos na struct existente
	pessoaExistente.Telefone = telefone
	pessoaExistente.Email = email

	// Chama a função do repositório para atualizar a pessoa.
	pessoaAtualizada := repository.AtualizarPessoa(uid, telefone, email)

	return pessoaAtualizada, nil // Retorna a pessoa atualizada para quem chamou o serviço.
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
