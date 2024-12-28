package resource

import (
	"time"
)

type Interface interface {
	SaveValueReceived(value float32, date time.Time, description string, bank string)
	SaveBankAccount(bank, account, acronym string) (int64, error)
	SelectBanksAccountsByUserID(id int64) ([]Account, error)
}

type Account struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Bank    string `json:"bank"`
	Account string `json:"account"`
	Acronym string `json:"acronym"`
}
