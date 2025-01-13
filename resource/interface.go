package resource

import (
	"expense-manager/domains/bankaccounts"
	"time"
)

type Interface interface {
	SaveValueReceived(value float32, date time.Time, description string, bank string)
	SaveBankAccount(acronym, description string) error
	SelectBanksAccountsByUserID(id int64) ([]*bankaccounts.BankAccount, error)
}
