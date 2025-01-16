package resource

import (
	"expense-manager/domains/bankaccount"
	"time"
)

type Interface interface {
	SaveValueReceived(value float32, date time.Time, description string, bank string)
	SaveBankAccount(bankAccount *bankaccount.BankAccount) error
	SelectBanksAccountsByUserID(id int) ([]*bankaccount.BankAccount, error)
}
