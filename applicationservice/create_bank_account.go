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
	err := verifyIfAlreadyExist(input.Acronym, resource, fakeUserID)
	if err != nil {
		return err
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

func verifyIfAlreadyExist(acronym string, resource resource.Interface, userID int) error {
	existingBankAccounts, err := resource.SelectBanksAccountsByUserID(userID)
	if err != nil {
		return err
	}

	for _, existingBankAccount := range existingBankAccounts {
		if existingBankAccount.Acronym == acronym {
			log.Println("Invalid AcronymValue: existing acronym")
			return fmt.Errorf("sigla já existente")
		}
	}
	return nil
}
