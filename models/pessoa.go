package models

import (
	"github.com/google/uuid" // Importação necessária após o 'go get'
)

type Pessoa struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Nome           string    `json:"nome" gorm:"type:varchar(100);not null"`
	Cpf            string    `json:"cpf" gorm:"type:varchar(11);unique;not null"`
	DataNascimento string    `json:"data_nascimento" gorm:"type:date;not null"`
	Telefone       string    `json:"telefone" gorm:"type:varchar(15);not null"`
	Email          string    `json:"email" gorm:"type:varchar(100);unique;not null"`
}
