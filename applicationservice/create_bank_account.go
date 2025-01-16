package applicationservice

import (
	"expense-manager/domains/bankaccount"
	"expense-manager/resource"
)

type CreateBankAccountInput struct {
	Acronym     string `json:"acronym"`
	Description string `json:"description"`
}

func CreateBankAccount(input *CreateBankAccountInput, resource resource.Interface) error {
	fakeUserID := 1
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
