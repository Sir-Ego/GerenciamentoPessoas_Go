package models

import (
	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Definição de Regex para validações específicas
var (
	// Permite apenas letras e espaços, min 3 max 25
	regexNome = regexp.MustCompile(`^[a-zA-ZÀ-ÿ\s]{3,25}$`)

	// Apenas 11 números
	regexCPF = regexp.MustCompile(`^\d{11}$`)

	// Formato DD/MM/YYYY (apenas números e barras)
	regexData = regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])/(0[1-9]|1[0-2])/\d{4}$`)

	// Permite + opcional no início e números (formato internacional)
	regexTelefone = regexp.MustCompile(`^\+?[0-9]{10,15}$`)
)

func ValidarNome(nome string) error { 
	return validation.Validate(nome, 
		validation.Required.Error("O nome é obrigatório"), 
		validation.Match(regexNome).Error("O nome deve conter apenas letras e ter entre 3 e 25 caracteres"), 
	) 
} 
		
func ValidarCpf(cpf string) error { 
	return validation.Validate(cpf, 
		validation.Required.Error("O CPF é obrigatório"), 
		validation.Match(regexCPF).Error("O CPF deve conter exatamente 11 números"), 
	) 
} 

func ValidarData(data string) error { 
	return validation.Validate(data, 
		validation.Required.Error("A data de nascimento é obrigatória"), 
		validation.Match(regexData).Error("A data deve estar no formato DD/MM/YYYY"), 
	) 
} 

func ValidarTelefone(tel string) error { 
	return validation.Validate(tel, 
		validation.Required.Error("O telefone é obrigatório"), 
		validation.Match(regexTelefone).Error("Telefone inválido. Use o formato: +DDI DDD N°TELEFONE ou apenas DDD N°TELEFONE "), 
	) 
} 

func ValidarEmail(email string) error { 
	return validation.Validate(email, 
		validation.Required.Error("O e-mail é obrigatório"), 
		is.Email.Error("O formato do e-mail é inválido"), 
	)
}

// Validate aplica as regras fluentes na struct Pessoa
func (p Pessoa) Validate() error {
	return validation.ValidateStruct(&p,

		// Nome: Letras e espaços, 3-25 caracteres
		validation.Field(&p.Nome,
			validation.Required.Error("O nome é obrigatório"),
			validation.Match(regexNome).Error("O nome deve conter apenas letras e ter entre 3 e 25 caracteres"),
		),

		// CPF: Obrigatório e tem de ter 11 dígitos
		validation.Field(&p.Cpf,
			validation.Required.Error("O CPF é obrigatório"),
			validation.Match(regexCPF).Error("O CPF deve conter exatamente 11 números"),
		),

		// Data: Obrigatória e formato DD/MM/YYYY
		validation.Field(&p.DataNascimento,
			validation.Required.Error("A data de nascimento é obrigatória"),
			validation.Match(regexData).Error("A data deve estar no formato DD/MM/YYYY"),
		),

		// Telefone: Aceita + e números
		validation.Field(&p.Telefone,
			validation.Required.Error("O telefone é obrigatório"),
			validation.Match(regexTelefone).Error("Telefone inválido. Use o formato: +5511999999999 ou apenas números"),
		),

		// Email: Obrigatório e formato válido
		validation.Field(&p.Email,
			validation.Required.Error("O e-mail é obrigatório"),
			is.Email.Error("O formato do e-mail é inválido"),
		),
	)
}
