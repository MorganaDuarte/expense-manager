package applicationservice

import (
	"expense-manager/domains/bankaccount"
	"expense-manager/resource"
	"fmt"
	"log"
)

type CreateBankAccountInput struct {
	Acronym     string `json:"acronym"`
	Description string `json:"description"`
}

func CreateBankAccount(input *CreateBankAccountInput, resource resource.Interface) error {
	fakeUserID := 1
	existingBankAccounts, err := resource.SelectBanksAccountsByUserID(fakeUserID)
	if err != nil {
		return err
	}

	for _, existingBankAccount := range existingBankAccounts {
		if existingBankAccount.Acronym == input.Acronym {
			log.Println("Invalid AcronymValue: existing acronym")
			return fmt.Errorf("sigla já existente")
		}
	}

	bankAccount, err := bankaccount.New(0, fakeUserID, input.Acronym, input.Description)
	if err != nil {
		return err
	}

	err = resource.SaveBankAccount(bankAccount)
	if err != nil {
		return err
	}
	return nil
}
