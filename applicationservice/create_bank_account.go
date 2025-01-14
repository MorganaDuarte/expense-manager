package applicationservice

import (
	"expense-manager/resource"
	"fmt"
	"log"
)

type CreateBankAccountInput struct {
	Acronym     string `json:"acronym"`
	Description string `json:"description"`
}

func CreateBankAccount(input *CreateBankAccountInput, resource resource.Interface) error {
	err := validateSaveBankAccountInput(input)
	if err != nil {
		return err
	}

	err = resource.SaveBankAccount(input.Acronym, input.Description)
	if err != nil {
		return err
	}
	return nil
}

func validateSaveBankAccountInput(input *CreateBankAccountInput) error {
	if len(input.Acronym) == 0 {
		log.Println("AcronymValue can`t be empty")
		return fmt.Errorf("A sigla não pode ser vazia")
	}

	if len(input.Acronym) > 3 {
		log.Println("Invalid AcronymValue: must be up to 3 letters")
		return fmt.Errorf("A sigla deve ter no máximo 3 letras")
	}
	return nil
}
