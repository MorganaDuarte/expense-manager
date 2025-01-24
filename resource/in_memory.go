package resource

import (
	"expense-manager/domains/bankaccount"
	"time"
)

type InMemoryResource struct {
	bankAccounts []*bankaccount.BankAccount
}

var inMemoryInstance *InMemoryResource

func GetInstance() *InMemoryResource {
	if inMemoryInstance != nil {
		return inMemoryInstance
	}

	inMemoryInstance = &InMemoryResource{}
	return inMemoryInstance
}

func (r *InMemoryResource) SaveValueReceived(value float32, date time.Time, description string, bank string) {
}

func (r *InMemoryResource) SaveBankAccount(bankAccount *bankaccount.BankAccount) error {
	b := &bankaccount.BankAccount{
		UserID:      bankAccount.UserID,
		Acronym:     bankAccount.Acronym,
		Description: bankAccount.Description,
	}
	r.bankAccounts = append(r.bankAccounts, b)
	return nil
}

func (r *InMemoryResource) SelectBanksAccountsByUserID(id int) ([]*bankaccount.BankAccount, error) {
	var bankAccounts []*bankaccount.BankAccount

	for _, bankAccount := range r.bankAccounts {
		if bankAccount.UserID == id {
			bankAccounts = append(bankAccounts, bankAccount)
		}
	}
	return bankAccounts, nil
}

func (r *InMemoryResource) CleanMemory() {
	r.bankAccounts = []*bankaccount.BankAccount{}
}
