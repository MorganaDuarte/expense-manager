package bankaccount

import (
	"fmt"
	"log"
)

type BankAccount struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Acronym     string `json:"acronym"`
	Description string `json:"description"`
}

func New(id, userID int, acronym, description string) (*BankAccount, error) {
	err := validateBankAccountAcronym(acronym)
	if err != nil {
		return nil, err
	}
	return &BankAccount{id, userID, acronym, description}, nil
}

func validateBankAccountAcronym(acronym string) error {
	if len(acronym) == 0 {
		log.Println("AcronymValue can`t be empty")
		return fmt.Errorf("A sigla não pode ser vazia")
	}

	if len(acronym) > 3 {
		log.Println("Invalid AcronymValue: must be up to 3 letters")
		return fmt.Errorf("A sigla deve ter no máximo 3 letras")
	}
	return nil
}
