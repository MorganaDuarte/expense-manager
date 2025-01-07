package resource

import (
	"expense-manager/domains/bankaccounts"
	"time"
)

type InMemoryResource struct {
	bankAccounts []*bankaccounts.BankAccount
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

func (r *InMemoryResource) SaveBankAccount(acronym, description string) error {
	bankAccount := &bankaccounts.BankAccount{
		Acronym:     acronym,
		Description: description,
	}
	r.bankAccounts = append(r.bankAccounts, bankAccount)
	return nil
}

func (r *InMemoryResource) SelectBanksAccountsByUserID(id int64) ([]*bankaccounts.BankAccount, error) {
	return r.bankAccounts, nil
}

func (r *InMemoryResource) CleanMemory() {
	r.bankAccounts = []*bankaccounts.BankAccount{}
}
